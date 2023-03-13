package solimpl

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/OffchainLabs/challenge-protocol-v2/protocol"
	"github.com/OffchainLabs/challenge-protocol-v2/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	_ = protocol.AssertionChain(&AssertionChain{})
	_ = protocol.ChainReadWriter(&AssertionChain{})
	_ = protocol.Assertion(&Assertion{})
	_ = protocol.ActiveTx(&activeTx{})
)

func TestAssertionStateHash(t *testing.T) {
	ctx := context.Background()
	chain, _, _, _, _ := setupAssertionChainWithChallengeManager(t)
	tx := &activeTx{readWriteTx: true}
	assertion, err := chain.LatestConfirmed(ctx, tx)
	require.NoError(t, err)

	execState := &protocol.ExecutionState{
		GlobalState: protocol.GoGlobalState{
			BlockHash: common.Hash{},
		},
		MachineStatus: protocol.MachineStatusFinished,
	}
	computed := protocol.ComputeStateHash(execState, big.NewInt(1))
	require.Equal(t, computed, assertion.StateHash())
}

func TestCreateAssertion(t *testing.T) {
	ctx := context.Background()
	chain, accs, addresses, backend, headerReader := setupAssertionChainWithChallengeManager(t)
	tx := &activeTx{readWriteTx: true}

	t.Run("OK", func(t *testing.T) {
		height := uint64(1)
		prev := uint64(0)
		minAssertionPeriod, err := chain.userLogic.MinimumAssertionPeriod(chain.callOpts)
		require.NoError(t, err)

		latestBlockHash := common.Hash{}
		for i := uint64(0); i < minAssertionPeriod.Uint64(); i++ {
			latestBlockHash = backend.Commit()
		}

		prevState := &protocol.ExecutionState{
			GlobalState:   protocol.GoGlobalState{},
			MachineStatus: protocol.MachineStatusFinished,
		}
		postState := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash:  latestBlockHash,
				SendRoot:   common.Hash{},
				Batch:      1,
				PosInBatch: 0,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		prevInboxMaxCount := big.NewInt(1)
		created, err := chain.CreateAssertion(
			ctx,
			tx,
			height,
			protocol.AssertionSequenceNumber(prev),
			prevState,
			postState,
			prevInboxMaxCount,
		)
		require.NoError(t, err)
		computed := protocol.ComputeStateHash(postState, big.NewInt(2))
		require.Equal(t, computed, created.StateHash(), "Unequal computed hash")

		_, err = chain.CreateAssertion(
			ctx,
			tx,
			height,
			protocol.AssertionSequenceNumber(prev),
			prevState,
			postState,
			prevInboxMaxCount,
		)
		require.ErrorContains(t, err, "ALREADY_STAKED")
	})
	t.Run("can create fork", func(t *testing.T) {
		chain, err := NewAssertionChain(
			ctx,
			addresses.Rollup,
			accs[2].txOpts,
			&bind.CallOpts{},
			accs[2].accountAddr,
			backend,
			headerReader,
		)
		require.NoError(t, err)
		height := uint64(1)
		prev := uint64(0)
		minAssertionPeriod, err := chain.userLogic.MinimumAssertionPeriod(chain.callOpts)
		require.NoError(t, err)

		for i := uint64(0); i < minAssertionPeriod.Uint64(); i++ {
			backend.Commit()
		}

		prevState := &protocol.ExecutionState{
			GlobalState:   protocol.GoGlobalState{},
			MachineStatus: protocol.MachineStatusFinished,
		}
		postState := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash:  common.BytesToHash([]byte("evil hash")),
				SendRoot:   common.Hash{},
				Batch:      1,
				PosInBatch: 0,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		prevInboxMaxCount := big.NewInt(1)
		chain.txOpts.From = accs[2].accountAddr
		forked, err := chain.CreateAssertion(
			ctx,
			tx,
			height,
			protocol.AssertionSequenceNumber(prev),
			prevState,
			postState,
			prevInboxMaxCount,
		)
		require.NoError(t, err)
		computed := protocol.ComputeStateHash(postState, big.NewInt(2))
		require.Equal(t, computed, forked.StateHash(), "Unequal computed hash")
	})
}

func TestAssertionBySequenceNum(t *testing.T) {
	ctx := context.Background()
	chain, _, _, _, _ := setupAssertionChainWithChallengeManager(t)
	tx := &activeTx{readWriteTx: true}

	resp, err := chain.AssertionBySequenceNum(ctx, tx, 0)
	require.NoError(t, err)

	require.Equal(t, true, resp.StateHash() != [32]byte{})

	_, err = chain.AssertionBySequenceNum(ctx, tx, 1)
	require.ErrorIs(t, err, ErrNotFound)
}

func TestAssertion_Confirm(t *testing.T) {
	ctx := context.Background()
	tx := &activeTx{readWriteTx: true}
	t.Run("OK", func(t *testing.T) {
		chain, _, _, backend, _ := setupAssertionChainWithChallengeManager(t)

		height := uint64(1)
		prev := uint64(0)
		minAssertionPeriod, err := chain.userLogic.MinimumAssertionPeriod(chain.callOpts)
		require.NoError(t, err)

		assertionBlockHash := common.Hash{}
		for i := uint64(0); i < minAssertionPeriod.Uint64(); i++ {
			assertionBlockHash = backend.Commit()
		}

		prevState := &protocol.ExecutionState{
			GlobalState:   protocol.GoGlobalState{},
			MachineStatus: protocol.MachineStatusFinished,
		}
		postState := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash:  assertionBlockHash,
				SendRoot:   common.Hash{},
				Batch:      1,
				PosInBatch: 0,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		prevInboxMaxCount := big.NewInt(1)
		_, err = chain.CreateAssertion(
			ctx,
			tx,
			height,
			protocol.AssertionSequenceNumber(prev),
			prevState,
			postState,
			prevInboxMaxCount,
		)
		require.NoError(t, err)

		err = chain.Confirm(ctx, tx, assertionBlockHash, common.Hash{})
		require.ErrorIs(t, err, ErrTooSoon)

		for i := uint64(0); i < minAssertionPeriod.Uint64(); i++ {
			backend.Commit()
		}
		require.NoError(t, chain.Confirm(ctx, tx, assertionBlockHash, common.Hash{}))
		require.ErrorIs(t, ErrNoUnresolved, chain.Confirm(ctx, tx, assertionBlockHash, common.Hash{}))
	})
}

func TestAssertion_Reject(t *testing.T) {
	ctx := context.Background()
	tx := &activeTx{readWriteTx: true}

	t.Run("Can reject assertion", func(t *testing.T) {
		t.Skip("TODO: Can't reject assertion. Blocked by one step proof")
	})

	t.Run("Already confirmed assertion", func(t *testing.T) {
		chain, _, _, backend, _ := setupAssertionChainWithChallengeManager(t)

		height := uint64(1)
		prev := uint64(0)
		minAssertionPeriod, err := chain.userLogic.MinimumAssertionPeriod(chain.callOpts)
		require.NoError(t, err)

		assertionBlockHash := common.Hash{}
		for i := uint64(0); i < minAssertionPeriod.Uint64(); i++ {
			assertionBlockHash = backend.Commit()
		}

		prevState := &protocol.ExecutionState{
			GlobalState:   protocol.GoGlobalState{},
			MachineStatus: protocol.MachineStatusFinished,
		}
		postState := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash:  assertionBlockHash,
				SendRoot:   common.Hash{},
				Batch:      1,
				PosInBatch: 0,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		prevInboxMaxCount := big.NewInt(1)
		_, err = chain.CreateAssertion(
			ctx,
			tx,
			height,
			protocol.AssertionSequenceNumber(prev),
			prevState,
			postState,
			prevInboxMaxCount,
		)
		require.NoError(t, err)

		for i := uint64(0); i < minAssertionPeriod.Uint64(); i++ {
			backend.Commit()
		}
		require.NoError(t, chain.Confirm(ctx, tx, assertionBlockHash, common.Hash{}))
		require.ErrorIs(t, ErrNoUnresolved, chain.Reject(ctx, tx, chain.stakerAddr))
	})
}

func TestChallengePeriodSeconds(t *testing.T) {
	ctx := context.Background()
	chain, _, _, _, _ := setupAssertionChainWithChallengeManager(t)
	tx := &activeTx{readWriteTx: true}
	manager, err := chain.CurrentChallengeManager(ctx, tx)
	require.NoError(t, err)

	chalPeriod, err := manager.ChallengePeriodSeconds(ctx, tx)
	require.NoError(t, err)
	require.Equal(t, time.Second*100, chalPeriod)
}

func TestCreateSuccessionChallenge(t *testing.T) {
	ctx := context.Background()
	tx := &activeTx{readWriteTx: true}
	t.Run("assertion does not exist", func(t *testing.T) {
		chain, _, _, _, _ := setupAssertionChainWithChallengeManager(t)
		_, err := chain.CreateSuccessionChallenge(ctx, tx, 2)
		require.ErrorIs(t, err, ErrInvalidChildren)
	})
	t.Run("at least two children required", func(t *testing.T) {
		chain, _, _, backend, _ := setupAssertionChainWithChallengeManager(t)
		height := uint64(1)
		prev := uint64(0)
		minAssertionPeriod, err := chain.userLogic.MinimumAssertionPeriod(chain.callOpts)
		require.NoError(t, err)

		latestBlockHash := common.Hash{}
		for i := uint64(0); i < minAssertionPeriod.Uint64(); i++ {
			latestBlockHash = backend.Commit()
		}

		prevState := &protocol.ExecutionState{
			GlobalState:   protocol.GoGlobalState{},
			MachineStatus: protocol.MachineStatusFinished,
		}
		postState := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash:  latestBlockHash,
				SendRoot:   common.Hash{},
				Batch:      1,
				PosInBatch: 0,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		prevInboxMaxCount := big.NewInt(1)
		_, err = chain.CreateAssertion(
			ctx,
			tx,
			height,
			protocol.AssertionSequenceNumber(prev),
			prevState,
			postState,
			prevInboxMaxCount,
		)
		require.NoError(t, err)

		_, err = chain.CreateSuccessionChallenge(ctx, tx, 0)
		require.ErrorIs(t, err, ErrInvalidChildren)
	})
	t.Run("assertion already rejected", func(t *testing.T) {
		t.Skip(
			"Needs a challenge manager to provide a winning claim first",
		)
	})
	t.Run("OK", func(t *testing.T) {
		chain, accs, addresses, backend, headerReader := setupAssertionChainWithChallengeManager(t)
		height := uint64(1)
		prev := uint64(0)
		minAssertionPeriod, err := chain.userLogic.MinimumAssertionPeriod(chain.callOpts)
		require.NoError(t, err)

		latestBlockHash := common.Hash{}
		for i := uint64(0); i < minAssertionPeriod.Uint64(); i++ {
			latestBlockHash = backend.Commit()
		}

		prevState := &protocol.ExecutionState{
			GlobalState:   protocol.GoGlobalState{},
			MachineStatus: protocol.MachineStatusFinished,
		}
		postState := &protocol.ExecutionState{
			GlobalState: protocol.GoGlobalState{
				BlockHash:  latestBlockHash,
				SendRoot:   common.Hash{},
				Batch:      1,
				PosInBatch: 0,
			},
			MachineStatus: protocol.MachineStatusFinished,
		}
		prevInboxMaxCount := big.NewInt(1)
		_, err = chain.CreateAssertion(
			ctx,
			tx,
			height,
			protocol.AssertionSequenceNumber(prev),
			prevState,
			postState,
			prevInboxMaxCount,
		)
		require.NoError(t, err)

		chain, err = NewAssertionChain(
			ctx,
			addresses.Rollup,
			accs[2].txOpts,
			&bind.CallOpts{},
			accs[2].accountAddr,
			backend,
			headerReader,
		)
		require.NoError(t, err)

		postState.GlobalState.BlockHash = common.BytesToHash([]byte("evil"))
		_, err = chain.CreateAssertion(
			ctx,
			tx,
			height,
			protocol.AssertionSequenceNumber(prev),
			prevState,
			postState,
			prevInboxMaxCount,
		)
		require.NoError(t, err)

		_, err = chain.CreateSuccessionChallenge(ctx, tx, 0)
		require.NoError(t, err)

		_, err = chain.CreateSuccessionChallenge(ctx, tx, 0)
		require.ErrorIs(t, err, ErrAlreadyExists)
	})
}

func setupAssertionChainWithChallengeManager(t *testing.T) (*AssertionChain, []*testAccount, *rollupAddresses, *backends.SimulatedBackend, *HeaderReader) {
	t.Helper()
	ctx := context.Background()
	accs, backend := setupAccounts(t, 3)
	prod := false
	wasmModuleRoot := common.Hash{}
	rollupOwner := accs[0].accountAddr
	chainId := big.NewInt(1337)
	loserStakeEscrow := common.Address{}
	challengePeriodSeconds := big.NewInt(100)
	miniStake := big.NewInt(1)
	cfg := generateRollupConfig(prod, wasmModuleRoot, rollupOwner, chainId, loserStakeEscrow, challengePeriodSeconds, miniStake)
	addresses := deployFullRollupStack(
		t,
		ctx,
		backend,
		accs[0].txOpts,
		common.Address{}, // Sequencer addr.
		cfg,
	)
	headerReader := NewHeaderReader(util.SimulatedBackendWrapper{SimulatedBackend: backend}, func() *Config { return &TestConfig })
	headerReader.Start(ctx)
	chain, err := NewAssertionChain(
		ctx,
		addresses.Rollup,
		accs[1].txOpts,
		&bind.CallOpts{},
		accs[1].accountAddr,
		backend,
		headerReader,
	)
	require.NoError(t, err)
	return chain, accs, addresses, backend, headerReader
}

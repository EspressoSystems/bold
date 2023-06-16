package solimpl_test

import (
	"context"
	"math/big"
	"testing"

	protocol "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction"
	solimpl "github.com/OffchainLabs/challenge-protocol-v2/chain-abstraction/sol-implementation"
	"github.com/OffchainLabs/challenge-protocol-v2/testing/setup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestCreateAssertion(t *testing.T) {
	ctx := context.Background()
	cfg, err := setup.ChainsWithEdgeChallengeManager()
	require.NoError(t, err)
	chain := cfg.Chains[0]
	backend := cfg.Backend

	t.Run("OK", func(t *testing.T) {
		latestBlockHash := common.Hash{}
		for i := uint64(0); i < 100; i++ {
			latestBlockHash = backend.Commit()
		}

		createdInfo := &protocol.AssertionCreatedInfo{
			AfterState: (&protocol.ExecutionState{
				GlobalState:   protocol.GoGlobalState{},
				MachineStatus: protocol.MachineStatusFinished,
			}).AsSolidityStruct(),
			InboxMaxCount: big.NewInt(1),
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
		_, err := chain.CreateAssertion(ctx, createdInfo, postState)
		require.NoError(t, err)

		_, err = chain.CreateAssertion(ctx, createdInfo, postState)
		require.ErrorContains(t, err, "ALREADY_STAKED")
	})
	t.Run("can create fork", func(t *testing.T) {
		t.Skip()
		assertionChain := cfg.Chains[1]

		for i := uint64(0); i < 100; i++ {
			backend.Commit()
		}

		creationInfo := &protocol.AssertionCreatedInfo{
			AfterState: (&protocol.ExecutionState{
				GlobalState:   protocol.GoGlobalState{},
				MachineStatus: protocol.MachineStatusFinished,
			}).AsSolidityStruct(),
			InboxMaxCount: big.NewInt(1),
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
		_, err := assertionChain.CreateAssertion(ctx, creationInfo, postState)
		require.NoError(t, err)
	})
}

func TestAssertionBySequenceNum(t *testing.T) {
	ctx := context.Background()
	cfg, err := setup.ChainsWithEdgeChallengeManager()
	require.NoError(t, err)
	chain := cfg.Chains[0]
	latestConfirmed, err := chain.LatestConfirmed(ctx)
	require.NoError(t, err)
	_, err = chain.GetAssertion(ctx, latestConfirmed.Id())
	require.NoError(t, err)

	_, err = chain.GetAssertion(ctx, protocol.AssertionId(common.BytesToHash([]byte("foo"))))
	require.ErrorIs(t, err, solimpl.ErrNotFound)
}

func TestChallengePeriodBlocks(t *testing.T) {
	ctx := context.Background()
	cfg, err := setup.ChainsWithEdgeChallengeManager()
	require.NoError(t, err)
	chain := cfg.Chains[0]

	manager, err := chain.SpecChallengeManager(ctx)
	require.NoError(t, err)

	chalPeriod, err := manager.ChallengePeriodBlocks(ctx)
	require.NoError(t, err)
	require.Equal(t, cfg.RollupConfig.ConfirmPeriodBlocks, chalPeriod)
}

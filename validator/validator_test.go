package validator

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/OffchainLabs/new-rollup-exploration/protocol"
	"github.com/OffchainLabs/new-rollup-exploration/testing/mocks"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
)

func Test_processChallengeStart(t *testing.T) {
	ctx := context.Background()

	seq := uint64(1)
	wantErr := errors.New("not found")

	t.Run("reading assertion fails", func(t *testing.T) {
		p := &mocks.MockProtocol{}
		s := &mocks.MockStateManager{}
		v, err := New(ctx, p, s)
		require.NoError(t, err)

		p.On("AssertionBySequenceNumber", ctx, seq).Return(&protocol.Assertion{}, wantErr)
		err = v.processChallengeStart(ctx, &protocol.StartChallengeEvent{
			ParentSeqNum: seq,
		})
		require.ErrorIs(t, err, wantErr)
	})
	t.Run("challenge does not concern us", func(t *testing.T) {
		p := &mocks.MockProtocol{}
		s := &mocks.MockStateManager{}
		v, err := New(ctx, p, s)
		require.NoError(t, err)

		p.On("AssertionBySequenceNumber", ctx, seq).Return(&protocol.Assertion{
			StateCommitment: protocol.StateCommitment{
				Height:    0,
				StateRoot: common.BytesToHash([]byte("foo")),
			},
		}, nil)
		err = v.processChallengeStart(ctx, &protocol.StartChallengeEvent{
			ParentSeqNum: seq,
		})
		require.NoError(t, err)
	})

}

type assertionLoggerFn func(string, ...interface{})

// AssertLogsContain checks that the desired string is a subset of the current log output.
func AssertLogsContain(tb testing.TB, hook *test.Hook, want string, msg ...interface{}) {
	checkLogs(tb, hook, want, true, msg...)
}

// AssertLogsDoNotContain is the inverse check of LogsContain.
func AssertLogsDoNotContain(tb testing.TB, hook *test.Hook, want string, msg ...interface{}) {
	checkLogs(tb, hook, want, false, msg...)
}

// LogsContain checks whether a given substring is a part of logs. If flag=false, inverse is checked.
func checkLogs(tb testing.TB, hook *test.Hook, want string, flag bool, msg ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	entries := hook.AllEntries()
	logs := make([]string, 0, len(entries))
	match := false
	for _, e := range entries {
		msg, err := e.String()
		if err != nil {
			tb.Errorf("%s:%d Failed to format log entry to string: %v", filepath.Base(file), line, err)
			return
		}
		if strings.Contains(msg, want) {
			match = true
		}
		for _, field := range e.Data {
			fieldStr, ok := field.(string)
			if !ok {
				continue
			}
			if strings.Contains(fieldStr, want) {
				match = true
			}
		}
		logs = append(logs, msg)
	}
	var errMsg string
	if flag && !match {
		errMsg = parseMsg("Expected log not found", msg...)
	} else if !flag && match {
		errMsg = parseMsg("Unexpected log found", msg...)
	}
	if errMsg != "" {
		tb.Errorf("%s:%d %s: %v\nSearched logs:\n%v", filepath.Base(file), line, errMsg, want, logs)
	}
}

func parseMsg(defaultMsg string, msg ...interface{}) string {
	if len(msg) >= 1 {
		msgFormat, ok := msg[0].(string)
		if !ok {
			return defaultMsg
		}
		return fmt.Sprintf(msgFormat, msg[1:]...)
	}
	return defaultMsg
}

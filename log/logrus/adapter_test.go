package logrus_test

import (
	"testing"

	"github.com/go-kit/kit/log/level"
	_logrus "github.com/goph/fxt/log/logrus"
	"github.com/sirupsen/logrus"
	logrus_test "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogrusAdapter_Log(t *testing.T) {
	l, hook := logrus_test.NewNullLogger()
	logger := _logrus.New(_logrus.Logger(l))

	logger.Log("msg", "message", "key", "value")

	require.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "message", hook.LastEntry().Message)
	assert.Equal(t, logrus.Fields(map[string]interface{}{"key": "value"}), hook.LastEntry().Data)
}

func TestMessageKey(t *testing.T) {
	l, hook := logrus_test.NewNullLogger()
	logger := _logrus.New(_logrus.Logger(l), _logrus.MessageKey("message"))

	logger.Log("message", "message")

	require.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "message", hook.LastEntry().Message)
}

func TestLevelKey(t *testing.T) {
	l, hook := logrus_test.NewNullLogger()
	logger := _logrus.New(_logrus.Logger(l), _logrus.LevelKey("lvl"))

	logger.Log("msg", "message", "lvl", "error")

	require.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "message", hook.LastEntry().Message)
}

func TestDefaultLevel(t *testing.T) {
	l, hook := logrus_test.NewNullLogger()
	l.Level = logrus.DebugLevel

	logger := _logrus.New(_logrus.Logger(l), _logrus.DefaultLevel(level.DebugValue()))

	logger.Log("msg", "message")

	require.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.DebugLevel, hook.LastEntry().Level)
	assert.Equal(t, "message", hook.LastEntry().Message)
}

func TestLevels(t *testing.T) {
	tests := map[string]struct {
		actual   interface{}
		expected logrus.Level
	}{
		"debug (string)": {"debug", logrus.DebugLevel},
		"info (string)":  {"info", logrus.InfoLevel},
		"warn (string)":  {"warn", logrus.WarnLevel},
		"error (string)": {"error", logrus.ErrorLevel},

		"debug (value)": {level.DebugValue(), logrus.DebugLevel},
		"info (value)":  {level.InfoValue(), logrus.InfoLevel},
		"warn (value)":  {level.WarnValue(), logrus.WarnLevel},
		"error (value)": {level.ErrorValue(), logrus.ErrorLevel},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			l, hook := logrus_test.NewNullLogger()
			l.Level = logrus.DebugLevel

			logger := _logrus.New(_logrus.Logger(l))

			logger.Log("msg", "message", "level", test.actual)

			require.Equal(t, 1, len(hook.Entries))
			assert.Equal(t, test.expected, hook.LastEntry().Level)
		})
	}
}

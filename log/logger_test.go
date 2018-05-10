package fxlog_test

import (
	"testing"

	"bytes"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/fxt/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	buf := new(bytes.Buffer)

	setUp := func(config *fxlog.Config) kitlog.Logger {
		buf.Reset()
		config.Output = buf

		logger, err := fxlog.NewLogger(config)

		require.NoError(t, err)

		return logger
	}

	t.Run("Defaults", func(t *testing.T) {
		config := fxlog.NewConfig()

		logger := setUp(config)

		logger.Log()

		assert.Equal(t, "{\"level\":\"info\"}\n", buf.String())
	})

	t.Run("LogfmtFormat", func(t *testing.T) {
		config := fxlog.NewConfig()
		config.Format = fxlog.LogfmtFormat

		logger := setUp(config)

		logger.Log()

		assert.Equal(t, "level=info\n", buf.String())
	})

	t.Run("FallbackLevel", func(t *testing.T) {
		config := fxlog.NewConfig()
		config.FallbackLevel = level.WarnValue()

		logger := setUp(config)

		logger.Log()

		assert.Equal(t, "{\"level\":\"warn\"}\n", buf.String())
	})

	t.Run("Context", func(t *testing.T) {
		config := fxlog.NewConfig()
		config.Context = []interface{}{"key", "value"}

		logger := setUp(config)

		logger.Log()

		assert.Equal(t, "{\"key\":\"value\",\"level\":\"info\"}\n", buf.String())
	})

	t.Run("Debug", func(t *testing.T) {
		t.Run("Off", func(t *testing.T) {
			config := fxlog.NewConfig()

			logger := setUp(config)

			logger.Log("level", level.DebugValue())

			assert.Equal(t, "", buf.String())
		})

		t.Run("On", func(t *testing.T) {
			config := fxlog.NewConfig()
			config.Debug = true

			logger := setUp(config)

			logger.Log("level", level.DebugValue())

			assert.Equal(t, "{\"level\":\"debug\"}\n", buf.String())
		})
	})
}

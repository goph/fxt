package errors_test

import (
	"testing"

	"bytes"

	"github.com/goph/emperror"
	"github.com/goph/fxt/errors"
	"github.com/goph/fxt/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	t.Run("Logger", func(t *testing.T) {
		buf := new(bytes.Buffer)
		config := log.NewConfig()
		config.Output = buf
		config.Format = log.LogfmtFormat
		logger, err := log.NewLogger(config)

		require.NoError(t, err)

		handler := errors.NewHandler(errors.HandlerParams{
			Logger: logger,
		})

		err = emperror.New("error")

		handler.Handle(err)

		assert.Equal(t, "level=error msg=error\n", buf.String())
	})
}

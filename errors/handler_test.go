package fxerrors_test

import (
	"testing"

	"bytes"
	stderrors "errors"

	"github.com/goph/fxt/errors"
	"github.com/goph/fxt/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	t.Run("Logger", func(t *testing.T) {
		buf := new(bytes.Buffer)
		config := fxlog.NewConfig()
		config.Output = buf
		config.Format = fxlog.LogfmtFormat
		logger, err := fxlog.NewLogger(config)

		require.NoError(t, err)

		handler := fxerrors.NewHandler(fxerrors.HandlerParams{
			Logger: logger,
		})

		err = stderrors.New("error")

		handler.Handle(err)

		assert.Equal(t, "level=error msg=error\n", buf.String())
	})
}

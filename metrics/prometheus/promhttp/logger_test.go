package fxpromhttp_test

import (
	"testing"

	"bytes"

	"github.com/go-kit/kit/log"
	"github.com/goph/fxt/metrics/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	logger := log.NewLogfmtLogger(buf)

	params := fxpromhttp.LoggerParams{
		Logger: logger,
	}

	l := fxpromhttp.NewLogger(params)

	l.Println("something")

	assert.Equal(t, "level=error component=prometheus msg=\"something\\n\"\n", buf.String())
}

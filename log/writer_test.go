package fxlog_test

import (
	"io"
	"testing"

	"bytes"

	kitlog "github.com/go-kit/kit/log"
	"github.com/goph/fxt/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriterAdapter(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := kitlog.NewLogfmtLogger(buf)
	writer := fxlog.NewWriterAdapter(logger)

	buf.Reset()

	n, err := writer.Write([]byte("Hello, World!"))
	require.NoError(t, err)
	assert.Equal(t, 13, n)
	assert.Equal(t, "msg=\"Hello, World!\"\n", buf.String())
}

func TestWriterAdapter_ImplementsWriter(t *testing.T) {
	assert.Implements(t, (*io.Writer)(nil), new(fxlog.WriterAdapter))
}

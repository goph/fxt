package log

import (
	"io"

	"github.com/go-kit/kit/log"
)

// WriterAdapter implements similar logic to StdlibAdapter in go-kit/log package,
// but instead of parsing incoming entries as standard go log entries WriterAdapter
// simply writes the whole entry as a log message.
type WriterAdapter struct {
	log.Logger

	messageKey string
}

// WriterAdapterOption sets a parameter for the WriterAdapter.
type WriterAdapterOption func(*WriterAdapter)

// MessageKey sets the key for the actual log message. By default, it's "msg".
func MessageKey(key string) WriterAdapterOption {
	return func(a *WriterAdapter) { a.messageKey = key }
}

// NewWriterAdapter returns a new WriterAdapter wrapper around the passed logger.
func NewWriterAdapter(logger log.Logger, options ...WriterAdapterOption) io.Writer {
	a := WriterAdapter{
		Logger:     logger,
		messageKey: "msg",
	}

	for _, option := range options {
		option(&a)
	}

	return a
}

func (a WriterAdapter) Write(p []byte) (int, error) {
	if err := a.Logger.Log(a.messageKey, string(p)); err != nil {
		return 0, err
	}

	return len(p), nil
}

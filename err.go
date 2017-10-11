package fxt

// ErrChan can be used as an arbitrary error acceptor channel
// and use it to stop the application.
type ErrChan chan error

// NewErrChan creates a new ErrChan channel.
func NewErrChan() ErrChan {
	return make(ErrChan, 1)
}

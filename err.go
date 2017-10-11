package fxt

// ErrChan can be used to stop the application as a result of an error.
type ErrChan chan error

// NewErrChan creates a new ErrChan channel.
func NewErrChan() ErrChan {
	return make(ErrChan, 1)
}

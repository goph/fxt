package fxt

// ErrIn accepts an error which causes the application to stop.
type ErrIn chan<- error

// ErrIn is the output counterpart of ErrIn.
type ErrOut <-chan error

// NewErr creates a new ErrIn channel.
func NewErr() (ErrIn, ErrOut) {
	ch := make(chan error, 1)

	return ch, ch
}

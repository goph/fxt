package fxt

import (
	"context"
)

// WithTimeout makes a function susceptible to deadlines.
func WithTimeout(ctx context.Context, f func(context.Context) error) error {
	ch := make(chan error, 1)

	go func() { ch <- f(ctx) }()

	select {
	case <-ctx.Done():
		return ctx.Err()

	case err := <-ch:
		return err
	}
}

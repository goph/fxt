package fxdebug

import "net/http"

// Err accepts an error which causes the application to stop.
type Err <-chan error

// ErrServerClosed is returned by the underlying server when it stops listening
// after graceful or forceful shutdown is initiated.
var ErrServerClosed = http.ErrServerClosed

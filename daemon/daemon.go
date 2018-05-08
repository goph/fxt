package fxdaemon

// Daemon is a long-running process in the background.
type Daemon interface {
	// Run has a loop inside and quits when it gets a signal from the quit channel.
	// It can handle errors internally or return an error if the daemon can't continue running.
	// Returning nil means the daemon somehow finished working (eg. received an external signal).
	Run(quit <-chan struct{}) error
}

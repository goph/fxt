package fxdev

import (
	"path"
	"runtime"

	"github.com/pkg/errors"
)

// GetCurrentDir returns the current directory, relative to the caller file.
// Should be used for development purposes only.
func GetCurrentDir() (string, error) {
	return getCurrentDir(1)
}

// getCurrentDir returns the current directory, relative to the caller depth.
// Should be used for development purposes only.
func getCurrentDir(depth int) (string, error) {
	_, filename, _, ok := runtime.Caller(depth + 1)
	if !ok {
		return "", errors.New("cannot get current dir: no caller information")
	}

	return path.Clean(path.Dir(filename)), nil
}

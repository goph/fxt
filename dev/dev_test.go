package fxdev_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/goph/fxt/dev"
	"github.com/stretchr/testify/require"
)

func TestGetCurrentDir(t *testing.T) {
	gopaths := strings.Split(os.Getenv("GOPATH"), ":")

	expected, err := fxdev.GetCurrentDir()
	require.NoError(t, err)

	for _, gopath := range gopaths {
		if  path.Clean(path.Join(gopath, "src/github.com/goph/fxt/dev")) == expected {
			return
		}
	}

	t.Fatal("failed asserting that the current dir is inside a GOPATH")
}

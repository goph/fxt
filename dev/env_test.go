package fxdev_test

import (
	"os"
	"testing"

	"github.com/goph/fxt/dev"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadEnvFromFile(t *testing.T) {
	err := fxdev.LoadEnvFromFile(".env")
	require.NoError(t, err)
	assert.Equal(t, "test_value", os.Getenv("TEST_VARIABLE"))
}

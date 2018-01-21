// +build !integration

package is_test

import (
	"testing"

	"github.com/goph/fxt/test/is"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	assert.False(t, is.Integration)
}

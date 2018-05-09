// +build !unit

package is_test

import (
	"testing"

	"github.com/goph/fxt/testing/is"
	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	assert.False(t, is.Unit)
}

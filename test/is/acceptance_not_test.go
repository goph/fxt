// +build !acceptance

package is_test

import (
	"testing"

	"github.com/goph/fxt/test/is"
	"github.com/stretchr/testify/assert"
)

func TestAcceptance(t *testing.T) {
	assert.False(t, is.Acceptance)
}

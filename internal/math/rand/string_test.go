package rand_test

import (
	"math/rand"
	"testing"

	_rand "github.com/goph/fxt/internal/math/rand"
	"github.com/stretchr/testify/assert"
)

func TestStringRand_Read(t *testing.T) {
	r := _rand.NewStringRand(rand.NewSource(1234))

	p := make([]byte, 4)

	n, err := r.Read(p)

	assert.Equal(t, 4, n)
	assert.NoError(t, err)
	assert.Equal(t, "Nxqa", string(p))
}

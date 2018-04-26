package correlationid_test

import (
	"testing"

	"github.com/goph/fxt/internal/correlationid"
	"github.com/stretchr/testify/assert"
)

func TestRandGenerator_Generate(t *testing.T) {
	generator := correlationid.NewRandGenerator(20)

	id := generator.Generate()

	assert.Len(t, id, 20)
}

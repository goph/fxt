package correlationid_test

import (
	"context"
	"testing"

	"github.com/goph/fxt/grpc/middleware/correlationid"
	"github.com/stretchr/testify/assert"
)

func TestGeneratorSource_ExtractCorrelationID(t *testing.T) {
	generator := new(correlationIdGenerator)
	generator.On("Generate").Return("cid")

	source := correlationid.NewGeneratorSource(generator)
	ctx := context.Background()

	cid := source.ExtractCorrelationID(ctx)

	assert.Equal(t, "cid", cid)
}

func TestDefaultGeneratorSource(t *testing.T) {
	generator := correlationid.DefaultGeneratorSource()

	ctx := context.Background()

	cid := generator.ExtractCorrelationID(ctx)

	assert.Len(t, cid, 32)
}

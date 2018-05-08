package fxcontext

import (
	"testing"

	"context"

	"github.com/stretchr/testify/assert"
)

func TestCorrleationId(t *testing.T) {
	ctx := context.WithValue(context.Background(), contextKeyCorrelationId, "cid")

	cid, ok := CorrleationId(ctx)

	assert.True(t, ok)
	assert.Equal(t, "cid", cid)
}

func TestWithCorrelationId(t *testing.T) {
	ctx := WithCorrelationId(context.Background(), "cid")

	cid, ok := ctx.Value(contextKeyCorrelationId).(string)

	assert.True(t, ok)
	assert.Equal(t, "cid", cid)
}

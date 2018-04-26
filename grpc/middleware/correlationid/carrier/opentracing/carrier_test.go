package opentracing_test

import (
	"testing"

	"context"

	otcarrier "github.com/goph/fxt/grpc/middleware/correlationid/carrier/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
)

func TestCarrier_GetCorrelationID(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	span.SetBaggageItem("correlation_id", "1234")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	c := otcarrier.New()

	correlationID, ok := c.GetCorrelationID(ctx)

	assert.Equal(t, "1234", correlationID)
	assert.True(t, ok)
}

func TestCarrier_GetCorrelationID_Empty(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	span.SetBaggageItem("correlation_id", "")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	c := otcarrier.New()

	correlationID, ok := c.GetCorrelationID(ctx)

	assert.Equal(t, "", correlationID)
	assert.False(t, ok)
}

func TestCarrier_SetCorrelationID(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	c := otcarrier.New()

	newCtx := c.SetCorrelationID(ctx, "1234")

	assert.Equal(t, "1234", span.BaggageItem("correlation_id"))
	assert.Equal(t, ctx, newCtx)
}

func TestWithRestrictedKey(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	span.SetBaggageItem("cid", "1234")
	ctx := context.Background()

	ctx = opentracing.ContextWithSpan(ctx, span)

	c := otcarrier.New(otcarrier.WithRestrictedKey("cid"))

	correlationID, ok := c.GetCorrelationID(ctx)

	assert.Equal(t, "1234", correlationID)
	assert.True(t, ok)

	span = tracer.StartSpan("test")
	ctx = context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	newCtx := c.SetCorrelationID(ctx, "1234")

	assert.Equal(t, "1234", span.BaggageItem("cid"))
	assert.Equal(t, ctx, newCtx)
}

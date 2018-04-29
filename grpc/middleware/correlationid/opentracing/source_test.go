package otcorrelationid_test

import (
	"testing"

	"context"

	"github.com/goph/fxt/grpc/middleware/correlationid/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
)

func TestOpentracingSource_ExtractCorrelationID(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	span.SetBaggageItem("correlation_id", "cid")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	source := otcorrelationid.NewOpentracingSource()

	correlationID := source.ExtractCorrelationID(ctx)

	assert.Equal(t, "cid", correlationID)
}

func TestOpentracingSource_ExtractCorrelationID_RestrictedKey(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	span.SetBaggageItem("correlationid", "cid")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	source := otcorrelationid.NewOpentracingSource(otcorrelationid.RestrictedKey("correlationid"))

	correlationID := source.ExtractCorrelationID(ctx)

	assert.Equal(t, "cid", correlationID)
}

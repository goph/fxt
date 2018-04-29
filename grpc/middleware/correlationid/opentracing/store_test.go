package otcorrelationid_test

import (
	"testing"

	"context"

	"github.com/goph/fxt/grpc/middleware/correlationid/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
)

func TestOpentracingStore_StoreCorrelationID(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	store := otcorrelationid.NewOpentracingStore()

	store.StoreCorrelationID(ctx, "cid")

	assert.Equal(t, "cid", span.BaggageItem("correlation_id"))
}

func TestOpentracingStore_StoreCorrelationID_RestrictedKey(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	store := otcorrelationid.NewOpentracingStore(otcorrelationid.RestrictedKey("correlationid"))

	store.StoreCorrelationID(ctx, "cid")

	assert.Equal(t, "cid", span.BaggageItem("correlationid"))
}

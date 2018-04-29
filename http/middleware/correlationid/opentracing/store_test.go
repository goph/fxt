package otcorrelationid_test

import (
	"testing"

	"context"
	"net/http"

	"github.com/goph/fxt/http/middleware/correlationid/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpentracingStore_StoreCorrelationID(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	req, err := http.NewRequest("GET", "", nil)
	require.NoError(t, err)

	req = req.WithContext(ctx)

	store := otcorrelationid.NewOpentracingStore()

	store.StoreCorrelationID(req, "cid")

	assert.Equal(t, "cid", span.BaggageItem("correlation_id"))
}

func TestOpentracingStore_StoreCorrelationID_RestrictedKey(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	req, err := http.NewRequest("GET", "", nil)
	require.NoError(t, err)

	req = req.WithContext(ctx)

	store := otcorrelationid.NewOpentracingStore(otcorrelationid.RestrictedKey("correlationid"))

	store.StoreCorrelationID(req, "cid")

	assert.Equal(t, "cid", span.BaggageItem("correlationid"))
}

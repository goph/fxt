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

func TestOpentracingPlugin_ExtractCorrelationID(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	span.SetBaggageItem("correlation_id", "cid")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	req, err := http.NewRequest("GET", "", nil)
	require.NoError(t, err)

	req = req.WithContext(ctx)

	source := otcorrelationid.NewOpentracingSource()

	correlationID := source.ExtractCorrelationID(req)

	assert.Equal(t, "cid", correlationID)
}

func TestOpentracingPlugin_ExtractCorrelationID_RestrictedKey(t *testing.T) {
	tracer := mocktracer.New()

	span := tracer.StartSpan("test")
	span.SetBaggageItem("correlationid", "cid")
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)

	req, err := http.NewRequest("GET", "", nil)
	require.NoError(t, err)

	req = req.WithContext(ctx)

	source := otcorrelationid.NewOpentracingSource(otcorrelationid.RestrictedKey("correlationid"))

	correlationID := source.ExtractCorrelationID(req)

	assert.Equal(t, "cid", correlationID)
}

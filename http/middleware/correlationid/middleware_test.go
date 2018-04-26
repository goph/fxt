package correlationid_test

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/goph/fxt/context"
	"github.com/goph/fxt/http/middleware/correlationid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMiddleware_Handler(t *testing.T) {
	generator := new(MockGenerator)
	generator.On("Generate").Return("cid")

	var cid string
	var ok bool
	m := correlationid.New(correlationid.Generator(generator))
	ts := httptest.NewServer(m.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid, ok = context.CorrleationId(r.Context())
	})))
	defer ts.Close()

	http.Get(ts.URL)

	assert.True(t, ok)
	assert.Equal(t, "cid", cid)
	generator.AssertExpectations(t)
}

func TestMiddleware_Handler_Headers(t *testing.T) {
	var cid string
	var ok bool
	m := correlationid.New(correlationid.Headers("Correlation-ID", "X-Correlation-ID"))
	ts := httptest.NewServer(m.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid, ok = context.CorrleationId(r.Context())
	})))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	require.NoError(t, err)

	req.Header.Set("X-Correlation-ID", "cid")

	http.DefaultClient.Do(req)

	assert.True(t, ok)
	assert.Equal(t, "cid", cid)
}

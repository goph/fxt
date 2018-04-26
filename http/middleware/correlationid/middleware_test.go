package correlationid_test

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/goph/fxt/context"
	"github.com/goph/fxt/http/middleware/correlationid"
	"github.com/goph/fxt/internal/correlationid/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMiddleware_Handler(t *testing.T) {
	generator := new(mocks.Generator)

	var cid string
	var ok bool
	middleware := correlationid.New(generator)
	ts := httptest.NewServer(middleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid, ok = context.CorrleationId(r.Context())
	})))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	require.NoError(t, err)

	req.Header.Set("Correlation-ID", "cid")

	http.DefaultClient.Do(req)

	assert.True(t, ok)
	assert.Equal(t, "cid", cid)
	generator.AssertNotCalled(t, "Generate")
}

func TestMiddleware_Handler_Generate(t *testing.T) {
	generator := new(mocks.Generator)
	generator.On("Generate").Return("cid")

	var cid string
	var ok bool
	middleware := correlationid.New(generator)
	ts := httptest.NewServer(middleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid, ok = context.CorrleationId(r.Context())
	})))
	defer ts.Close()

	http.Get(ts.URL)

	assert.True(t, ok)
	assert.Equal(t, "cid", cid)
	generator.AssertExpectations(t)
}

func TestHeader(t *testing.T) {
	generator := new(mocks.Generator)

	var cid string
	var ok bool
	middleware := correlationid.New(generator, correlationid.Header("correlationid"))
	ts := httptest.NewServer(middleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid, ok = context.CorrleationId(r.Context())
	})))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	require.NoError(t, err)

	req.Header.Set("correlationid", "cid")

	http.DefaultClient.Do(req)

	assert.True(t, ok)
	assert.Equal(t, "cid", cid)
	generator.AssertNotCalled(t, "Generate")
}

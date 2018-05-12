package correlationid_test

import (
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/goph/fxt/context"
	"github.com/goph/fxt/http/middleware/correlationid"
	"github.com/goph/fxt/http/middleware/correlationid/opentracing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMiddleware_Handler(t *testing.T) {
	var ok bool

	m := correlationid.New()
	ts := httptest.NewServer(m.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok = fxcontext.CorrleationId(r.Context())
	})))
	defer ts.Close()

	http.Get(ts.URL)

	assert.False(t, ok)
}

func TestMiddleware_Handler_Source(t *testing.T) {
	var cid string
	var ok bool

	s1 := new(correlationIdSource)
	s1.On("ExtractCorrelationID", mock.Anything).Return("")
	s2 := new(correlationIdSource)
	s2.On("ExtractCorrelationID", mock.Anything).Return("cid")

	m := correlationid.New(correlationid.Source(s1, s2))
	ts := httptest.NewServer(m.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid, ok = fxcontext.CorrleationId(r.Context())
	})))
	defer ts.Close()

	http.Get(ts.URL)

	assert.True(t, ok)
	assert.Equal(t, "cid", cid)
	mock.AssertExpectationsForObjects(t, s1, s2)
}

func TestMiddleware_Handler_WithSource(t *testing.T) {
	var cid string
	var ok bool

	s1 := new(correlationIdSource)
	s1.On("ExtractCorrelationID", mock.Anything).Return("")
	s2 := new(correlationIdSource)
	s2.On("ExtractCorrelationID", mock.Anything).Return("cid")

	m := correlationid.New(correlationid.Source(s1), correlationid.WithSource(s2))
	ts := httptest.NewServer(m.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid, ok = fxcontext.CorrleationId(r.Context())
	})))
	defer ts.Close()

	http.Get(ts.URL)

	assert.True(t, ok)
	assert.Equal(t, "cid", cid)
	mock.AssertExpectationsForObjects(t, s1, s2)
}

func TestMiddleware_Handler_Store(t *testing.T) {
	s1 := new(correlationIdSource)
	s1.On("ExtractCorrelationID", mock.Anything).Return("cid")

	st1 := new(correlationIdStore)
	st1.On("StoreCorrelationID", mock.Anything, "cid")

	m := correlationid.New(correlationid.Source(s1), correlationid.Store(st1))
	ts := httptest.NewServer(m.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// noop
	})))
	defer ts.Close()

	http.Get(ts.URL)

	mock.AssertExpectationsForObjects(t, s1, st1)
}

func ExampleCorrelationID() {
	m := correlationid.New(
		otcorrelationid.NewOption(), // Configure Opentracing
		correlationid.WithSource(
			correlationid.NewHeaderSource("Correlation-ID"), // Find correlation ID in headers
			correlationid.DefaultGeneratorSource(),          // Generate a new one if none is found
		),
	)
	ts := httptest.NewServer(m.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid, ok := fxcontext.CorrleationId(r.Context())
		if ok {
			fmt.Println(cid)
		}
	})))
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL, nil)
	req.Header.Set("Correlation-ID", "cid")

	http.DefaultClient.Do(req)

	// Output: cid
}

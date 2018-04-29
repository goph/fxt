package correlationid_test

import (
	"net/http"
	"testing"

	"github.com/goph/fxt/http/middleware/correlationid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderSource_ExtractCorrelationID(t *testing.T) {
	source := correlationid.NewHeaderSource("Correlation-ID")

	req, err := http.NewRequest("", "", nil)
	require.NoError(t, err)

	req.Header.Set("Correlation-ID", "cid")

	cid := source.ExtractCorrelationID(req)

	assert.Equal(t, "cid", cid)
}

func TestDefaultHeaderSource(t *testing.T) {
	source := correlationid.DefaultHeaderSource()

	req, err := http.NewRequest("", "", nil)
	require.NoError(t, err)

	req.Header.Set("X-Correlation-ID", "cid")

	cid := source.ExtractCorrelationID(req)

	assert.Equal(t, "cid", cid)
}

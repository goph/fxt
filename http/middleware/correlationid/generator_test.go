package correlationid_test

import (
	"net/http"
	"testing"

	"github.com/goph/fxt/http/middleware/correlationid"
	"github.com/stretchr/testify/assert"
)

func TestGeneratorSource_ExtractCorrelationID(t *testing.T) {
	generator := new(correlationIdGenerator)
	generator.On("Generate").Return("cid")

	source := correlationid.NewGeneratorSource(generator)
	req := &http.Request{}

	cid := source.ExtractCorrelationID(req)

	assert.Equal(t, "cid", cid)
}

func TestDefaultGeneratorSource(t *testing.T) {
	generator := correlationid.DefaultGeneratorSource()

	req := &http.Request{}

	cid := generator.ExtractCorrelationID(req)

	assert.Len(t, cid, 32)
}

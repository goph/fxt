package correlationid

import (
	"net/http"

	"github.com/goph/fxt/internal/correlationid"
)

//go:generate sh -c "CGO_ENABLED=0 mockery -name=correlationIdGenerator -output . -outpkg correlationid_test -testonly"

// correlationIdGenerator generates a correlation ID.
type correlationIdGenerator interface {
	// Generate generates a unique string which can be used as a correlation ID.
	Generate() string
}

// generatorSource generates a correlation ID.
type generatorSource struct {
	generator correlationIdGenerator
}

// NewGeneratorSource returns a new correlation ID source which generates a new correlation ID.
func NewGeneratorSource(generator correlationIdGenerator) *generatorSource {
	return &generatorSource{generator}
}

// DefaultGeneratorSource returns a new correlation ID source which generates a new correlation ID.
// The underlying generator is a default random string generator.
func DefaultGeneratorSource() *generatorSource {
	return &generatorSource{correlationid.NewRandGenerator(32)}
}

func (s *generatorSource) ExtractCorrelationID(r *http.Request) string {
	return s.generator.Generate()
}

package correlationid

const (
	defaultHeader = "Correlation-ID"
)

// Option is used to define middleware configuration.
type Option interface {
	apply(*middleware)
}

// Headers sets the headers to be checked in the middleware.
func Headers(h ...string) Option {
	return headersOption(h)
}

type headersOption []string

func (o headersOption) apply(m *middleware) {
	m.headers = o
}

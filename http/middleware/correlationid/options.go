package correlationid

const (
	defaultHeader = "Correlation-ID"
)

// Option is used to define middleware configuration.
type Option interface {
	apply(*middleware)
}

type optionFunc func(*middleware)

func (f optionFunc) apply(m *middleware) {
	f(m)
}

// Headers sets the headers to be checked in the middleware.
func Headers(h ...string) Option {
	return headersOption(h)
}

type headersOption []string

func (o headersOption) apply(m *middleware) {
	m.headers = o
}

// Generator sets a generator instance in the middleware.
func Generator(g correlationIdGenerator) Option {
	return optionFunc(func(m *middleware) {
		m.generator = g
	})
}

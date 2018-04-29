package correlationid

// Option is used to define middleware configuration.
type Option interface {
	apply(*middleware)
}

// Options is a set of middleware options together.
type Options []Option

func (o Options) apply(m *middleware) {
	for _, op := range o {
		op.apply(m)
	}
}

// Source sets a set of correlation ID sources in the middleware.
func Source(s ...correlationIdSource) Option {
	return sourceOption(s)
}

type sourceOption []correlationIdSource

func (o sourceOption) apply(m *middleware) {
	m.sources = o
}

// WithSource appends a set of correlation ID sources in the middleware.
func WithSource(s ...correlationIdSource) Option {
	return withSourceOption(s)
}

type withSourceOption []correlationIdSource

func (o withSourceOption) apply(m *middleware) {
	m.sources = append(m.sources, o...)
}

// Store sets a set of correlation ID stores in the middleware.
func Store(s ...correlationIdStore) Option {
	return storeOption(s)
}

type storeOption []correlationIdStore

func (o storeOption) apply(m *middleware) {
	m.stores = o
}

// WithStore appends a set of correlation ID stores in the middleware.
func WithStore(s ...correlationIdStore) Option {
	return withStoreOption(s)
}

type withStoreOption []correlationIdStore

func (o withStoreOption) apply(m *middleware) {
	m.stores = append(m.stores, o...)
}

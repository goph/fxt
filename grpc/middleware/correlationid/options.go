package correlationid

type options struct {
	sources []correlationIdSource
	stores  []correlationIdStore
}

func newOptions(opts ...Option) *options {
	op := new(options)

	for _, o := range opts {
		o.apply(op)
	}

	return op
}

// Option is used to define middleware configuration.
type Option interface {
	apply(*options)
}

// Options is a set of middleware options together.
type Options []Option

func (o Options) apply(m *options) {
	for _, op := range o {
		op.apply(m)
	}
}

// Source sets a set of correlation ID sources in the middleware.
func Source(s ...correlationIdSource) Option {
	return sourceOption(s)
}

type sourceOption []correlationIdSource

func (o sourceOption) apply(op *options) {
	op.sources = o
}

// WithSource appends a set of correlation ID sources in the middleware.
func WithSource(s ...correlationIdSource) Option {
	return withSourceOption(s)
}

type withSourceOption []correlationIdSource

func (o withSourceOption) apply(op *options) {
	op.sources = append(op.sources, o...)
}

// Store sets a set of correlation ID stores in the middleware.
func Store(s ...correlationIdStore) Option {
	return storeOption(s)
}

type storeOption []correlationIdStore

func (o storeOption) apply(op *options) {
	op.stores = o
}

// WithStore appends a set of correlation ID stores in the middleware.
func WithStore(s ...correlationIdStore) Option {
	return withStoreOption(s)
}

type withStoreOption []correlationIdStore

func (o withStoreOption) apply(op *options) {
	op.stores = append(op.stores, o...)
}

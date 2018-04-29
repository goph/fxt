package otcorrelationid

const defaultRestrictedKey = "correlation_id"

type options struct {
	restrictedKey string
}

func newOptions(opts ...Option) *options {
	op := &options{
		restrictedKey: defaultRestrictedKey,
	}

	for _, o := range opts {
		o.apply(op)
	}

	return op
}

// Option is used to define middleware configuration.
type Option interface {
	apply(*options)
}

// RestrictedKey customizes the baggage item restricted key for the correlation ID
type RestrictedKey string

func (k RestrictedKey) apply(p *options) {
	p.restrictedKey = string(k)
}

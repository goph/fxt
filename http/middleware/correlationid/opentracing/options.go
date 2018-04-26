package opentracing

const defaultRestrictedKey = "correlation_id"

// Option is used to define middleware configuration.
type Option interface {
	apply(*middleware)
}

// RestrictedKey customizes the baggage item restricted key for the correlation ID
type RestrictedKey string

func (k RestrictedKey) apply(m *middleware) {
	m.restrictedKey = string(k)
}

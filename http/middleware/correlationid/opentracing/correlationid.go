package otcorrelationid

import "github.com/goph/fxt/http/middleware/correlationid"

// NewOption returns correlation ID options configuring opentracing.
func NewOption(options ...Option) correlationid.Option {
	return correlationid.Options{
		correlationid.WithSource(NewOpentracingSource(options...)),
		correlationid.WithStore(NewOpentracingStore(options...)),
	}
}

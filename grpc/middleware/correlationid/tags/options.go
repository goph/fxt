package tagscorrelationid

const TagCorrelationID = "correlationid"

// Option is used to define middleware configuration.
type Option interface {
	apply(*tagStore)
}

// Tag customizes the tag for the correlation ID
type Tag string

func (t Tag) apply(s *tagStore) {
	s.tag = string(t)
}

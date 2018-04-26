package correlationid

// correlationIdGenerator generates a correlation ID.
type correlationIdGenerator interface {
	// Generate generates a unique string which can be used as a correlation ID.
	Generate() string
}

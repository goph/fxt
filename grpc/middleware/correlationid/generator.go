package correlationid

type generator interface {
	// Generate generates a unique string which can be used as a correlation ID.
	Generate() string
}

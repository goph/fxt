package correlationid

type IdGenerator interface {
	// Generate generates a unique string which can be used as a correlation ID.
	Generate() string
}

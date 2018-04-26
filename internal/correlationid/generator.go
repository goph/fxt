package correlationid

//go:generate mockery -name=Generator

// Generator generates a correlation ID.
type Generator interface {
	// Generate generates a unique string which can be used as a correlation ID.
	Generate() string
}

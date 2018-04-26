package correlationid

//go:generate sh -c "CGO_ENABLED=0 mockery -name=correlationIdGenerator -output . -outpkg correlationid_test -testonly"

// correlationIdGenerator generates a correlation ID.
type correlationIdGenerator interface {
	// Generate generates a unique string which can be used as a correlation ID.
	Generate() string
}

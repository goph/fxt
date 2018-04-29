package correlationid

import "net/http"

//go:generate sh -c "CGO_ENABLED=0 mockery -name=correlationIdStore -output . -outpkg correlationid_test -testonly"

// correlationIdStore stores a correlation ID in an arbitrary storage.
type correlationIdStore interface {
	// StoreCorrelationID stores the correlation ID in an arbitrary storage.
	StoreCorrelationID(r *http.Request, cid string)
}

package correlationid_test

import (
	"testing"

	"context"

	"github.com/goph/fxt/grpc/middleware/correlationid"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/stretchr/testify/assert"
)

func TestMetadataSource_ExtractCorrelationID(t *testing.T) {
	md := metautils.NiceMD{}
	md.Set("correlation_id", "cid")
	ctx := md.ToIncoming(context.Background())

	source := correlationid.NewMetadataSource("correlation_id")

	cid := source.ExtractCorrelationID(ctx)

	assert.Equal(t, "cid", cid)
}

func TestDefaultMetadataSource(t *testing.T) {
	md := metautils.NiceMD{}
	md.Set("correlationid", "cid")
	ctx := md.ToIncoming(context.Background())

	source := correlationid.DefaultMetadataSource()

	cid := source.ExtractCorrelationID(ctx)

	assert.Equal(t, "cid", cid)
}

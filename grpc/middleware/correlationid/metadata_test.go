package correlationid_test

import (
	"testing"

	"context"

	"github.com/goph/fxt/grpc/middleware/correlationid"
	"github.com/goph/fxt/internal/correlationid/mocks"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/stretchr/testify/assert"
)

func TestMetadataCarrier_GetCorrelationID(t *testing.T) {
	md := metautils.NiceMD{}
	md.Set("correlationid", "1234")
	ctx := md.ToIncoming(context.Background())

	c := correlationid.NewMetadataCarrier()

	correlationID, ok := c.GetCorrelationID(ctx)

	assert.Equal(t, "1234", correlationID)
	assert.True(t, ok)
}

func TestMetadataCarrier_GetCorrelationID_Empty(t *testing.T) {
	md := metautils.NiceMD{}
	md.Set("correlationid", "")
	ctx := md.ToIncoming(context.Background())

	c := correlationid.NewMetadataCarrier()

	correlationID, ok := c.GetCorrelationID(ctx)

	assert.Equal(t, "", correlationID)
	assert.False(t, ok)
}

func TestMetadataCarrier_SetCorrelationID(t *testing.T) {
	c := correlationid.NewMetadataCarrier()

	ctx := c.SetCorrelationID(context.Background(), "1234")

	md := metautils.ExtractIncoming(ctx)

	assert.Equal(t, "1234", md.Get("correlationid"))
}

func TestWithHeader(t *testing.T) {
	md := metautils.NiceMD{}
	md.Set("cid", "1234")
	ctx := md.ToIncoming(context.Background())

	c := correlationid.NewMetadataCarrier(correlationid.WithHeader("cid"))

	correlationID, ok := c.GetCorrelationID(ctx)

	assert.Equal(t, "1234", correlationID)
	assert.True(t, ok)

	ctx = c.SetCorrelationID(context.Background(), "1234")

	md = metautils.ExtractIncoming(ctx)

	assert.Equal(t, "1234", md.Get("cid"))
}

func TestMetadataSourceCarrier_GetCorrelationID(t *testing.T) {
	md := metautils.NiceMD{}
	md.Set("correlationid", "1234")
	ctx := md.ToIncoming(context.Background())

	cm := new(mocks.Carrier)

	cm.On("GetCorrelationID", ctx).Return("", false)

	c := correlationid.NewMetadataSourceCarrier(cm)

	correlationID, ok := c.GetCorrelationID(ctx)

	assert.Equal(t, "1234", correlationID)
	assert.False(t, ok)

	cm.AssertExpectations(t)
}

func TestMetadataSourceCarrier_SetCorrelationID(t *testing.T) {
	md := metautils.NiceMD{}
	md.Set("correlationid", "1234")
	ctx := md.ToIncoming(context.Background())

	cm := new(mocks.Carrier)

	cm.On("SetCorrelationID", ctx, "1234").Return(ctx)

	c := correlationid.NewMetadataSourceCarrier(cm)

	ctx = c.SetCorrelationID(ctx, "1234")

	md = metautils.ExtractIncoming(ctx)

	assert.Equal(t, "1234", md.Get("correlationid"))

	cm.AssertExpectations(t)
}

package nettest_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/goph/fxt/testing/nettest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFreePort(t *testing.T) {
	port, err := nettest.GetFreePort()
	assert.NoError(t, err)
	assert.NotZero(t, port)

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	require.NoError(t, err)
	lis.Close()
}

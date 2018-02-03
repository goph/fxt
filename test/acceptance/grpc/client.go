package grpc

import (
	"github.com/DATA-DOG/godog"
	fxgrpc "github.com/goph/fxt/grpc"
	"google.golang.org/grpc"
)

// ClientContext creates a new client connection before every scenario.
type ClientContext struct {
	config fxgrpc.AppClientConfig

	conn *grpc.ClientConn

	frozen bool
}

// NewClientContext returns a new ClientContext which creates a new client connection before every scenario.
func NewClientContext(config fxgrpc.AppClientConfig) *ClientContext {
	return &ClientContext{
		config: config,
	}
}

func (c *ClientContext) FeatureContext(s *godog.Suite) {
	if c.frozen {
		panic("trying to use a frozen feature context")
	}
	c.frozen = true

	s.BeforeScenario(func(interface{}) {
		conn, err := fxgrpc.Dial(c.config)
		if err != nil {
			panic(err)
		}

		c.conn = conn
	})

	s.AfterScenario(func(interface{}, error) {
		if c.conn != nil {
			c.conn.Close()
		}
	})
}

func (c *ClientContext) Conn() *grpc.ClientConn {
	return c.conn
}

package acceptance

import (
	"github.com/DATA-DOG/godog"
	fxgrpc "github.com/goph/fxt/grpc"
	"google.golang.org/grpc"
)

// GrpcClientContext creates a new client connection before every scenario.
type GrpcClientContext struct {
	config fxgrpc.AppClientConfig

	conn *grpc.ClientConn

	frozen bool
}

// NewGrpcClientContext returns a new GrpcClientContext which creates a new client connection before every scenario.
func NewGrpcClientContext(config fxgrpc.AppClientConfig) *GrpcClientContext {
	return &GrpcClientContext{
		config: config,
	}
}

func (c *GrpcClientContext) FeatureContext(s *godog.Suite) {
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

func (c *GrpcClientContext) Conn() *grpc.ClientConn {
	return c.conn
}

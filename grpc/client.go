package fxgrpc

import "google.golang.org/grpc"

// Dial creates a client connection.
func Dial(config AppClientConfig, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	if config.Insecure {
		options = append(options, grpc.WithInsecure())

		if config.Host != "" {
			options = append(options, grpc.WithAuthority(config.Host))
		}
	}

	return grpc.Dial(config.Addr, options...)
}

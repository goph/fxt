package nettest

import (
	"net"
	"strconv"
)

// GetFreePort returns a free TCP port.
func GetFreePort() (int, error) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer lis.Close()

	_, port, err := net.SplitHostPort(lis.Addr().String())
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(port)
}

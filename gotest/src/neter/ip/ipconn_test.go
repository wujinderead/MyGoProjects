package ip

import (
	"net"
	"testing"
)

func TestIpConn(t *testing.T) {
	ipconn := net.DialIP()
	net.ListenIP("ip")
}

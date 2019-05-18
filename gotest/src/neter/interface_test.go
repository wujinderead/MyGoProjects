package neter

import (
	"fmt"
	"net"
	"testing"
)

func TestInterfaces(t *testing.T) {
	faces, err := net.Interfaces()
	if err != nil {
		fmt.Println("get interfaces err:", err)
		t.FailNow()
	}
	for i := range faces {
		displayInterface(&faces[i])
	}
}

func TestInterfaceAddrs(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get interface addrs err:", err)
		t.FailNow()
	}
	for i := range addrs {
		displayAddr(&addrs[i])
	}
}

// IpNet is the ip plus mask, like 10.0.0.1/24
func TestIpNet(t *testing.T) {
	ip := net.IPv4(10, 19, 138, 135)
	mask := net.IPv4Mask(255, 255, 0, 0)
	ipnet := net.IPNet{ip, mask}
	fmt.Println("ipnet network:", ipnet.Network())
	fmt.Println("ipnet:", ipnet.String())
	for _, v := range []net.IP{
		net.IPv4(10, 19, 138, 135),
		net.IPv4(10, 19, 0, 34),
		net.IPv4(10, 13, 45, 22),
	} {
		fmt.Println(ipnet, "contains:", v, ipnet.Contains(v))
	}
}

func TestLookUp(t *testing.T) {
	// todo finish net test
}

func displayInterface(face *net.Interface) {
	fmt.Println("interface name:", face.Name)
	fmt.Println("interface mtu:", face.MTU)
	fmt.Println("interface index:", face.Index)
	fmt.Println("interface flags:", face.Flags)
	fmt.Println("interface hardware addr:", face.HardwareAddr)
	addrs, err := face.Addrs()
	if err != nil {
		fmt.Println("get address err:", err)
	}
	for i := range addrs {
		displayAddr(&addrs[i])
	}
	multiAddrs, err := face.MulticastAddrs()
	if err != nil {
		fmt.Println("get multicast address err:", err)
	}
	for i := range multiAddrs {
		displayAddr(&multiAddrs[i])
	}
	fmt.Println()
}

func displayAddr(addr *net.Addr) {
	fmt.Println("addr:", (*addr).Network(), "=", (*addr).String())
}

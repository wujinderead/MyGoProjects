package neter

import (
	"fmt"
	"net"
	"os"
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
	hostname, err := os.Hostname()
	if err != nil {
		t.Fatal("get hostname failed")
	}
	// use ip to get hostname
	addrs := []string{"::1", "10.19.138.22", "127.0.0.2"}
	for i := range addrs {
		names, err := net.LookupAddr(addrs[i])
		fmt.Println(names, err)
	}
	fmt.Println()

	// get canonical hostname
	hosts := []string{"localhost", hostname, "www.baidu.com", "hub.docker.com"}
	for i := range hosts {
		cname, err := net.LookupCNAME(hosts[i])
		fmt.Println(cname, err)
	}
	fmt.Println()

	// get addresses from hostname
	for i := range hosts {
		addrs, err := net.LookupHost(hosts[i])
		fmt.Println(addrs, err)
	}
	fmt.Println()

	// get ip from hostname
	for i := range hosts {
		ips, err := net.LookupIP(hosts[i])
		fmt.Println(ips, err)
	}
	fmt.Println()

	// get port from service name
	services := []string{"ssh", "ftp", "http", "kerberos"}
	for i := range hosts {
		port, err := net.LookupPort("tcp", services[i])
		fmt.Println(port, err)
		port, err = net.LookupPort("udp", services[i])
		fmt.Println(port, err)
	}
	fmt.Println()

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

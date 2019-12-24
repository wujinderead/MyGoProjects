package main

import (
	"fmt"
	"net"
	"syscall"
	"time"
)

func main() {
	//udpWriteConnected()
	//udpWriteUnconnected()
	sendUdpSyscall()
}

func udpWriteConnected() {
	raddr := &net.UDPAddr{IP: net.IPv4(10, 20, 38, 188), Port: 18888}
	var laddr *net.UDPAddr = nil
	// these addresses are also worked
	// var laddr *net.UDPAddr = &net.UDPAddr{IP: net.IPv4(10,20,38,188), Port: 12345}
	// var laddr *net.UDPAddr = &net.UDPAddr{IP: net.IPv4(10,20,38,188)}
	// var laddr *net.UDPAddr = &net.UDPAddr{IP: net.IPv4(127,0,0,1)}
	// var laddr *net.UDPAddr = &net.UDPAddr{}
	conn, err := net.DialUDP("udp4", laddr, raddr)
	if err != nil {
		fmt.Println("dial udp err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 3; i++ {
		n, err := conn.Write([]byte("abcdefg")) // since connected, no need to specify write destination
		if err != nil {
			fmt.Println("write err:", err)
			break
		}
		fmt.Println("write:", n)
		buf := make([]byte, 20)
		n, err = conn.Read(buf) // since connected, no need to specify read source
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println("read:", string(buf[:n]))
		time.Sleep(1000)
	}
}

func udpWriteUnconnected() {
	laddr := &net.UDPAddr{IP: net.IPv4(10, 20, 38, 188)}
	raddr := &net.UDPAddr{IP: net.IPv4(10, 20, 38, 188), Port: 18888}
	conn, err := net.ListenUDP("udp4", laddr) // listen udp on random port
	if err != nil {
		fmt.Println("dial udp err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 3; i++ {
		n, err := conn.WriteToUDP([]byte("abcdefg"), raddr) // listener and ReadFrom and WriteTo
		if err != nil {
			fmt.Println("write err:", err)
			break
		}
		fmt.Println("write:", n)
		buf := make([]byte, 20)
		n, addr, err := conn.ReadFromUDP(buf) // since connected, no need to specify read source
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println("read:", string(buf[:n]), ", from", addr.String())
		time.Sleep(1000)
	}
}

// need only "socket" and "bind', no need to connect.
func sendUdpSyscall() {
	raddr := &syscall.SockaddrInet4{
		Port: 18888,
		Addr: [4]byte{10, 20, 38, 188},
	}
	sockfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM|syscall.SOCK_CLOEXEC, 0)
	if err != nil {
		fmt.Println("socket err:", err)
		return
	}
	err = syscall.Bind(sockfd, &syscall.SockaddrInet4{}) // listen on rand port
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}
	lsa, err := syscall.Getsockname(sockfd) // getsockname syscall to get the local addr
	if err != nil {
		fmt.Println("getsockname err:", err)
		return
	}
	fmt.Println("laddr:", lsa.(*syscall.SockaddrInet4))
	buf := make([]byte, 128)
	for i := 0; i < 3; i++ {
		err := syscall.Sendto(sockfd, []byte("my name is van"), 0, raddr)
		if err != nil {
			fmt.Println("sendto err:", err)
			break
		}
		n, rraddr, err := syscall.Recvfrom(sockfd, buf, 0)
		if err != nil {
			fmt.Println("recv from err:", err)
			break
		}
		fmt.Println("recv:", string(buf[:n]), ", from", rraddr.(*syscall.SockaddrInet4))
		time.Sleep(1 * time.Second)
	}
	defer closeFd(sockfd)

}

func closeFd(fd int) {
	err := syscall.Close(fd)
	if err != nil {
		fmt.Println("close fd err:", err)
		return
	}
}

package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//listenUdp()
	listenUdpSyscall()
}

func listenUdp() {
	lis, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.IPv4(10, 20, 38, 188), Port: 18888})
	if err != nil {
		fmt.Println("listen udp err:", err)
		return
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigc
		fmt.Printf("got signal %d, close conn and exit.\n", sig.(syscall.Signal))
		err := lis.Close()
		fmt.Println("listener close:", err)
	}()

	buf := make([]byte, 128)
	for {
		// udp is no connection, so udp listener can't accept; it just read packet and get sender address
		n, addr, err := lis.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		// write back
		n, err = lis.WriteToUDP([]byte("got"), addr)
		if err != nil {
			fmt.Println("write err:", err)
			break
		}
		fmt.Println("read from:", addr.String(), ", content:", string(buf[:n]), ", write:", n)
	}
}

// udp socket need only "socket" and "bind", no need listen
func listenUdpSyscall() {
	laddr := &syscall.SockaddrInet4{
		Port: 18888,
		Addr: [4]byte{10, 20, 38, 188},
	}
	sockfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM|syscall.SOCK_CLOEXEC, 0)
	if err != nil {
		fmt.Println("socket err:", err)
		return
	}
	err = syscall.Bind(sockfd, syscall.Sockaddr(laddr))
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}
	buf := make([]byte, 128)
	// use Recvfrom and Sendto instead of Read or Write to get the from addr and set the to addr
	for i := 0; i < 3; i++ {
		n, raddr, err := syscall.Recvfrom(sockfd, buf, 0)
		if err != nil {
			fmt.Println("recv from err:", err)
			break
		}
		fmt.Println("recv:", string(buf[:n]), ", from", raddr.(*syscall.SockaddrInet4))
		err = syscall.Sendto(sockfd, []byte("i got it"), 0, raddr)
		if err != nil {
			fmt.Println("send to err:", err)
			break
		}
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

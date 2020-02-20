package main

import (
	"fmt"
	"net"
	"syscall"
)

func main() {
	dialUnix()
	//dialUnixStreamSyscall()
	//dialUnixgramConnected()
	//dialUnixgramUnconnected()
}

func dialUnix() {
	conn, err := net.Dial("unix", "neter/unix/listen.sock")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("local:", conn.LocalAddr().String(), ", remote:", conn.RemoteAddr().String())
	buf := make([]byte, 32)
	for i:=0; i<3; i++ {
		n, err := conn.Write([]byte("my name is van."))
		if err != nil {
			fmt.Println("write err:", err)
			break
		}
		n, err = conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println("read:", string(buf[:n]))
	}
}

func dialUnixStreamSyscall() {
	raddr := &syscall.SockaddrUnix{Name: "neter/unix/listen.sock"}
	laddr := &syscall.SockaddrUnix{Name: "neter/unix/dial.sock"}
	// open socket
	sockfd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("socket err:", err)
		return
	}
	defer syscall.Unlink(laddr.Name)
	defer syscall.Close(sockfd)
	// bind to addr
	err = syscall.Bind(sockfd, laddr)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}
	// connect to dial
	err = syscall.Connect(sockfd, raddr)
	if err != nil {
		fmt.Println("connect err:", err)
		return
	}
	buf := make([]byte, 32)
	for i:=0; i<3; i++ {
		n, err := syscall.Write(sockfd, []byte("my name is van."))
		if err != nil {
			fmt.Println("write err:", err)
			break
		}
		n, err = syscall.Read(sockfd, buf)
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println("read:", string(buf[:n]))
	}
}

// unix datagram, no connection, like udp
func dialUnixgramConnected() {
	// create dialer conn
	raddr, _ := net.ResolveUnixAddr("unix", "neter/unix/listen.sock")
	laddr, _ := net.ResolveUnixAddr("unix", "neter/unix/dial.sock")
	conn, err := net.DialUnix("unixgram", laddr, raddr)  // after dial, no need to specify address
	if err != nil {
		fmt.Println("listen unix sock error: ", err)
		return
	}
	defer syscall.Unlink(laddr.Name)
	defer conn.Close()

	buf := make([]byte, 32)
	for i:=0; i<3; i++ {
		n, err := conn.Write([]byte("my name is van."))
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		n, err = conn.Read(buf)
		if err != nil {
			fmt.Println("write err:", err)
			break
		}
		fmt.Println("read:", string(buf[:n]))
	}
}

func dialUnixgramUnconnected() {
	// create listener
	raddr, _ := net.ResolveUnixAddr("unixgram", "neter/unix/listen.sock")  // raddr must be unixgram
	laddr, _ := net.ResolveUnixAddr("unix", "neter/unix/dial.sock")        // laddr can be unix
	conn, err := net.ListenUnixgram("unixgram", laddr)
	if err != nil {
		fmt.Println("listen unix sock error: ", err)
		return
	}
	defer syscall.Unlink(laddr.Name)
	defer conn.Close()

	buf := make([]byte, 32)
	for i:=0; i<3; i++ {
		n, err := conn.WriteToUnix([]byte("my name is van."), raddr)
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		n, addr, err := conn.ReadFromUnix(buf)
		if err != nil {
			fmt.Println("write err:", err)
			break
		}
		fmt.Println("read:", string(buf[:n]), "from:", addr.Network(), addr.String())
	}
}
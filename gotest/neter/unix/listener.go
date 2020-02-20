package main

import (
	"fmt"
	"net"
	"syscall"
)

func main() {
	//listenUnix()
	listenUnixStreamSyscall()
	//listenUnixgram()
}

func listenUnix() {
	// create listener
	listener, err := net.Listen("unix", "neter/unix/listen.sock")
	if err != nil {
		fmt.Println("listen unix sock error: ", err)
		return
	}
	defer listener.Close()
	lis := listener.(*net.UnixListener)
	lis.SetUnlinkOnClose(true)   // delete the socket file after close

	conn, err := lis.AcceptUnix()
	if err != nil {
		fmt.Println("accept error: ", err)
		return
	}
	buf := make([]byte, 32)
	fmt.Println("local:", conn.LocalAddr().String(), ", remote:", conn.RemoteAddr().String())
	for i:=0; i<3; i++ {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println("read:", string(buf[:n]))
		n, err = conn.Write([]byte("i got it"))
		if err != nil {
			fmt.Println("write err:", err)
			break
		}
	}
}

func listenUnixStreamSyscall() {
	skfd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("open socket err:", err)
		return
	}
	defer syscall.Close(skfd)
	laddr := &syscall.SockaddrUnix{Name: "neter/unix/listen.sock"}
	err = syscall.Bind(skfd, laddr)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}
	defer syscall.Unlink(laddr.Name)
	err = syscall.Listen(skfd, syscall.SOMAXCONN)
	if err != nil {
		fmt.Println("listen err:")
		return
	}

	connfd, raddr, err := syscall.Accept(skfd)
	if err != nil {
		fmt.Println("accept err:", err)
		return
	}
	fmt.Println("sockfd:", skfd, "connfd:", connfd, "raddr:", raddr.(*syscall.SockaddrUnix).Name)

	buf := make([]byte, 32)
	for i:=0; i<3; i++ {
		n, err := syscall.Read(connfd, buf)
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println("read:", string(buf[:n]))
		n, err = syscall.Write(connfd, []byte("i got it"))
		if err != nil {
			fmt.Println("write err:", err)
			break
		}
	}
}

func listenUnixgram() {
	// create listener
	laddr, _ := net.ResolveUnixAddr("unix", "neter/unix/listen.sock")
	conn, err := net.ListenUnixgram("unixgram", laddr)
	if err != nil {
		fmt.Println("listen unix sock error: ", err)
		return
	}
	defer syscall.Unlink(laddr.Name)
	defer conn.Close()

	buf := make([]byte, 32)
	for i:=0; i<3; i++ {
		n, addr, err := conn.ReadFromUnix(buf)
		if err != nil {
			fmt.Println("read err:", err)
			break
		}
		fmt.Println("read:", string(buf[:n]), ", from:", addr.Network(), addr.String())
		n, err = conn.WriteToUnix([]byte("i got it"), addr)
		if err != nil {
			fmt.Println("write err:", err)
			break
		}
	}
}

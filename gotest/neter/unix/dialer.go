package main

import (
	"fmt"
	"net"
)

func main() {
	Send("my name is van")
}

func Send(content string) {
	conn, err := net.Dial("unix", "neter/unix/test.sock")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	n, err := conn.Write([]byte(content))
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
	fmt.Println("wrote", n)
	buf := make([]byte, 512)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	println("client got:", string(buf[:n]))
}

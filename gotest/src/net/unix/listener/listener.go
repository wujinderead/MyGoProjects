package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	Listen()
}

func Listen() {
	// create address
	addr, err := net.ResolveUnixAddr("unix", "src/net/unix/test.sock")
	if err != nil {
		fmt.Println("get unix sock addr error: ", err)
		return
	}

	// create listener
	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		fmt.Println("listen unix sock error: ", err)
		return
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			fmt.Println("close listener err:", err)
		} else {
			fmt.Println("listener closed.")
		}
	}()

	// ctrl-c to stop process
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		os.Exit(0)
	}(sigc)

	// accept conn and use goroutine to process incoming conn
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error: ", err)
		}

		go echo(conn)
	}
}

// the process func
func echo(conn net.Conn) {
	buf := make([]byte, 512)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:n]
		println("Server got:", string(data))

		_, err = conn.Write([]byte("i got it!"))
		if err != nil {
			log.Fatal("Writing client error: ", err)
		}
	}
}

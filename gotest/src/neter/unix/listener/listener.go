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

	// ctrl-c to stop process
	// process: ctrl-c, notify sigc, get from sigc, close listener, accept error, main loop out,
	// wait hoop exit, main exit
	waithook := make(chan int)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(c chan os.Signal) {
		sig := <-c
		fmt.Printf("Caught signal %s: shutting down.\n", sig)
		err := listener.Close() // if not close, the socket file can't be used again
		if err != nil {
			fmt.Println("close listener err:", err)
		} else {
			fmt.Println("listener closed.")
		}
		fmt.Println("hook exit")
		waithook <- 1
	}(sigc)

	for {
		conn, err := listener.AcceptUnix()
		if err != nil {
			fmt.Println("accept error: ", err)
			break // if error, break
		}
		// accept conn and use goroutine to process incoming conn
		go echo(conn)
	}
	<-waithook
	fmt.Println("main exit")
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

		// response to client
		_, err = conn.Write([]byte("i got it!"))
		if err != nil {
			log.Fatal("Writing client error: ", err)
		}
	}
}

package util

import (
	"net"
	"fmt"
	"log"
	"os/signal"
	"os"
	"syscall"
	"time"
	"io"
	"strconv"
)

var sock_file_path = "/home/xzy/test.sock"

func reader(r io.Reader) {
	buf := make([]byte, 512)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("Client got:", string(buf[0:n]))
	}
}

func Send() {
	conn, err := net.Dial("unix", sock_file_path)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer conn.Close()


	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
		sig := <-sigc
		log.Printf("Caught signal %s: shutting down.", sig)
		os.Exit(0)
	}()

	go reader(conn)
	i := 0
	for {
		msg := "my name is van!" + strconv.Itoa(i)
		i++
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}
		println("Client sent:", msg)
		time.Sleep(3 * time.Second)
	}
}

func Listen() {
	addr, err := net.ResolveUnixAddr("unix", sock_file_path)
	if err != nil {
		fmt.Println("get unix sock addr error: ", err)
	}
	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		fmt.Println("listen unix sock error: ", err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(ln net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		ln.Close()
		os.Exit(0)
	}(listener, sigc)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error: ", err)
		}

		go echo(conn)
	}
}

func echo(conn net.Conn) {
	buf := make([]byte, 512)
	for {
		nr, err := conn.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println("Server got:", string(data))

		_, err = conn.Write(data)
		if err != nil {
			log.Fatal("Writing client error: ", err)
		}
	}
}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	Send("lgq", 3)
}

func Send(name string, sleep int64) {
	conn, err := net.Dial("unix", "src/net/unix/test.sock")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("close conn err:", err)
		} else {
			fmt.Println("conn closed.")
		}
	}()


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
		msg := fmt.Sprintf("my name is %s-%d", name, i)
		i++
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}
		println("Client sent:", msg)
		time.Sleep(time.Duration(sleep * int64(time.Second)))
	}
}

// a separate goroutine to read response from server
func reader(r io.Reader) {
	buf := make([]byte, 512)
	for {
		n, err := r.Read(buf)
		if err != nil {
			return
		}
		println("Client got:", string(buf[:n]))
	}
}

package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	Send("xzy", 3)
}

func Send(name string, sleep int64) {
	conn, err := net.Dial("unix", "src/net/unix/test.sock")
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

	go reader(conn)
looper:
	for i := 0; i >= 0; i++ {
		// since there is no block in loop, we can we shutdown hook in main goroutine by
		// 'select' to get interrupt signal and break the loop.
		select {
		case sig := <-sigc:
			fmt.Printf("Caught signal %s: shutting down.\n", sig)
			break looper
		default:
		}

		// write to server
		msg := fmt.Sprintf("my name is %s-%d", name, i)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}
		println("Client sent:", msg)
		time.Sleep(time.Duration(sleep * int64(time.Second)))
	}

	// close the conn when out of loop
	err = conn.Close()
	if err != nil {
		fmt.Println("close conn err:", err)
	} else {
		fmt.Println("conn closed.")
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

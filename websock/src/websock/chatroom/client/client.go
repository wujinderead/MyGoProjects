package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
	"websock/src/websock/util"
)

var addr = flag.String("addr", "localhost:8080", "server address")

type message struct {
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()

	// dial websocket server
	conn, response, err := websocket.DefaultDialer.Dial("ws://"+*addr, nil)
	if err != nil {
		log.Fatal("dial err:", err)
		return
	}
	defer util.Close(fmt.Sprintf("connection '%s'", util.ToString(conn.LocalAddr())), conn)
	log.Println("dial respond header:", response.Header)
	log.Printf("dial success, local addr: %s, remote addr: %s\n",
		util.ToString(conn.LocalAddr()), util.ToString(conn.RemoteAddr()))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// send message to other and get message from other
	id := new(string)
	fmt.Println("input your id:")
	_, _ = fmt.Scanf("%s\n", id)
	err = conn.WriteMessage(websocket.TextMessage, []byte(*id))
	if err != nil {
		log.Fatal("send id err:", err)
		return
	}

	// read message from console and send to server
	go func() {
		for {
			to := new(string)
			msg := new(string)
			_, err := fmt.Scanf("%s %s\n", to, msg)
			if err != nil {
				fmt.Println("input err:", err)
				continue
			}
			err = conn.WriteJSON(&message{*id, *to, *msg})
			if err != nil {
				log.Println("write message err:", err)
				break
			}
			fmt.Printf("to '%s': %s\n", *to, *msg)
		}
	}()

	done := make(chan struct{})
	// read message from conn and print
	go func() {
		defer close(done)
		for {
			msg := new(message)
			err := conn.ReadJSON(msg) // it will stuck here if no thing read
			if err != nil {           // err can be normal close or abnormal close
				log.Println("read message err:", err)
				break
			}
			fmt.Printf("from '%s': %s\n", msg.From, msg.Msg)
		}
	}()

	// print stack trace for all goroutine on certain time interval
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	trace := make([]byte, 10240)
	for {
		select {
		case <-done: // unexpected read err cause close(done), trigger <-done to return
			return
		case <-ticker.C:
			fmt.Println("start stack trace:")
			n := runtime.Stack(trace, false)
			fmt.Println(string(trace[:n]))
		case <-interrupt:
			fmt.Println("get interrupt signal.")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("write close error:", err)
				return
			}
			select { // no default, it will block until a case triggered
			case <-done:
				fmt.Println("done")
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
			return
		}
	}
}

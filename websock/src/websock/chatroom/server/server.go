package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"websock/src/websock/util"
)

var addr = flag.String("addr", "localhost:8080", "server address")

var upgrader = websocket.Upgrader{}

var connMap sync.Map

var messageQueue = make(chan *message, 100)

type message struct {
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// get websocket conn
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade err:", err)
		return
	}
	defer util.Close(fmt.Sprintf("connection to '%s'", util.ToString(conn.RemoteAddr())), conn)
	log.Printf("accept conn, local: %s, remote: %s\n",
		util.ToString(conn.LocalAddr()), util.ToString(conn.RemoteAddr()))

	// get client id and store its id-conn map
	_, id, err := conn.ReadMessage()
	if err != nil {
		log.Println("get name err:", err)
		return
	}
	log.Printf("register id: '%s' with conn: '%p'\n", string(id), conn)
	connMap.Store(string(id), conn)

	// read the message and dispatch to corresponding receiver
	for {
		msg := new(message)
		err := conn.ReadJSON(msg)
		if err != nil {
			log.Println("read msg err:", err)
			connMap.Delete(string(id))
			break
		}
		messageQueue <- msg
	}
}

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sig
		connMap.Range(func(key, value interface{}) bool {
			err := value.(*websocket.Conn).WriteMessage(
				websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("send close msg err:", err)
			}
			return true
		})
		os.Exit(int(syscall.SIGTERM))
	}()
	go func() {
		for msg := range messageQueue {
			receiver := msg.To
			receiveConn, ok := connMap.Load(receiver)
			if !ok {
				continue // no registered receiver
			}
			err := receiveConn.(*websocket.Conn).WriteJSON(msg)
			if err != nil {
				log.Println("writer receiver err:", err)
			}
		}
	}()
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

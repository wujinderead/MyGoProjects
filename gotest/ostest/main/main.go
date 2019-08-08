package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

// use signal to end a loop
func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	a := new(int32)
	*a = 1
	go func(){  // shutdown hook
		s := <-sig
		fmt.Println("got signal: ", s)
		time.Sleep(1*time.Second)
		atomic.StoreInt32(a, 0)
	}()

	index := 0
	for atomic.LoadInt32(a)==1 {
		time.Sleep(200*time.Millisecond)
		fmt.Println("working", index)
		index++
	}

	fmt.Println("main complete")
}

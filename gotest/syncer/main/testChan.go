package main

import (
	"fmt"
	"time"
)

func main() {
	//testCloseChan1()
	//testCloseChan2()
	//testNilChan()
	testChanClosed()
}

func testCloseChan1() {
	ch := make(chan int)
	// goroutine sender
	go func() {
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
		time.Sleep(time.Second)
		close(ch)
		/* close(ch)    // cause panic: close of closed channel  */
		/* ch <- 3      // cause panic: send on closed channel   */
	}()

	// main receiver
	for i := range ch { // when chan closed, the receiver will unblock immediately
		fmt.Println(i, time.Now())
	}
	fmt.Println(time.Now())
}

func testCloseChan2() {
	ch := make(chan int)
	// goroutine sender
	go func() {
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
		time.Sleep(time.Second)
		close(ch)
	}()

	// main receiver
	for i := 0; i < 4; i++ {
		i, ok := <-ch // when chan closed, the receiver will receive zero value and false immediately
		fmt.Println(i, ok, time.Now())
	}
}

func testNilChan() {
	// send to nil channel cause
	// 'fatal err:'fatal error: all goroutines are asleep - deadlock!'
	var ch chan int
	go func() {
		ch <- 1
	}()
	<-ch
}

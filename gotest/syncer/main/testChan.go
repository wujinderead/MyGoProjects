package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func main() {
	//testCloseChan1()
	//testCloseChan2()
	//testNilChan()
	testHchan()
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
		i, ok := <-ch // when unbuffered chan closed, the receiver will unblock and receive zero value and false immediately
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

// hchan defined in runtime/chan.go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters
	lock     mutex
}

func testHchan() {
	// create a buffered chan
	ch := make(chan int64, 11) // chan is actually *hchan
	// send 4 ints, receive 1 int
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	a, ok := <-ch
	fmt.Println(a, ok)

	hc := *(**hchan)(unsafe.Pointer(&ch)) // convert chan to *hchan
	buf := hc.buf
	fmt.Println("buf addr :", buf)
	etype := hc.elemtype
	fmt.Println("elem kind:", etype.kind, reflect.Kind(etype.kind))
	fmt.Println("elem size:", hc.elemsize, etype.size)
	fmt.Println("dataqsiz :", hc.dataqsiz) // the buffer of chan is a circular queue, the size is as specified when made
	fmt.Println()

	// since we have sent 4 data anf receive 1 data
	// so qcount=3 (3 ints in buffer), sendx=4 (next send index is 4), recvx=1 (next receive index is 1)
	fmt.Println("closed:", hc.closed) // 0 for open chan
	fmt.Println("qcount:", hc.qcount)
	fmt.Println("sendx :", hc.sendx)
	fmt.Println("recvx :", hc.recvx)
	fmt.Println()

	ch <- 5
	ch <- 6
	a, ok = <-ch
	fmt.Println(a, ok)
	close(ch)
	fmt.Println("closed:", hc.closed) // 1 for closed chan
	fmt.Println("qcount:", hc.qcount)
	fmt.Println("sendx :", hc.sendx)
	fmt.Println("recvx :", hc.recvx)
	fmt.Println()

	// check buffer
	// when data are received, the buffer is set to zero value
	fmt.Println("buffer:")
	for i := 0; i < 11; i++ {
		data := *(*int64)(unsafe.Pointer(uintptr(buf) + uintptr(i)*uintptr(hc.elemsize)))
		fmt.Print(data, " ") // 0 0 3 4 5 6 0 0 0 0 0
	}
	fmt.Println()

	// although chan have been closed, data are still buffered in queue, so we can receive data until queue empty
	for i := 0; i < 5; i++ {
		// 3 true
		// 4 true
		// 5 true
		// 6 true
		// 0 false, when queue empty, 'ok' become 'false' to indicate empty queue
		a, ok := <-ch
		fmt.Println(a, ok)
	}
}

type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        *typeAlg
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}
type tflag uint8
type mutex struct {
	key uintptr
}
type typeAlg struct {
	// function for hashing objects of this type
	// (ptr to object, seed) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}
type nameOff int32
type typeOff int32
type waitq struct {
	first uintptr
	last  uintptr
}

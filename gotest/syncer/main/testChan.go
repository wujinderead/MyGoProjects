package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

func main() {
	//testCloseChan1()
	//testCloseChan2()
	//testNilChan()
	//testHchan()
	//testLenCapChan()
	//testSelect()
	testChanClosedWhenSending()
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
	var ch chan int // uninitialized chan, i.e., nil chan
	go func() {
		// sending to nil channel will call gopark() to block current goroutine infinitely,
		// because no one would call gounpark() to wake it up.
		// since this goroutine never ends, its resources never gets released, thus this goroutine leaks.
		ch <- 1
		fmt.Println("aaa") // never get executed
	}()
	go func() {
		// receiving from nil channel also call gopark() to block infinitely
		<-ch
		fmt.Println("aaa") // never get executed
	}()
	time.Sleep(time.Second)
}

func testSelect() {
	var wg sync.WaitGroup
	wg.Add(3)
	achan := make(chan int)
	bchan := make(chan int)
	cchan := make(chan int)
	go func() {
		select {
		case a := <-achan:
			fmt.Println("a:", a)
		case <-time.NewTimer(time.Second).C:
			fmt.Println("a timeout")
		}
		wg.Done()
	}()
	go func() {
		select { // if we don't want to block unexpectedly infinitely, use select and a timer
		case a := <-bchan:
			fmt.Println("b:", a)
		case <-time.NewTimer(time.Second).C:
			fmt.Println("b timeout")
		}
		wg.Done()
	}()
	go func() {
		select {
		case a := <-cchan:
			fmt.Println("c:", a)
		case <-time.NewTimer(time.Second).C:
			fmt.Println("c timeout")
		}
		wg.Done()
	}()
	select {
	case achan <- 1: // if one action selected, the others don't try
	case bchan <- 2:
	case cchan <- 3:
	}
	wg.Wait()
	// output:
	// c: 3
	// a timeout
	// b timeout
}

// hchan defined in runtime/chan.go
type hchan struct {
	qcount   uint           // total data in the queue, the buffered chan is stored in a circular queue
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

// though not common, but len() and cap() can apply to channel, which is a non-blocking action.
// they are retrieved from the 'qcount' and 'dataqsiz' fields of 'hchan'.
// but the result should be deemed as ephemeral and not reliable.
// it can be used as a hint to determine whether we should send or receive to avoid blocking.
func testLenCapChan() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second / 2)
		fmt.Println(len(ch)) // return 0 for unbuffered chan
		fmt.Println(cap(ch)) // 0
		fmt.Println("got", <-ch)
	}()
	ch <- 1
	time.Sleep(time.Second)
	fmt.Println()

	ch = make(chan int, 5)
	go func() {
		time.Sleep(time.Second / 2)
		fmt.Println(len(ch)) // 3
		fmt.Println(cap(ch)) // 5
		fmt.Println("got", <-ch)
		fmt.Println("got", <-ch)
		fmt.Println(len(ch)) // 1
		fmt.Println(cap(ch)) // 5
		fmt.Println("got", <-ch)
		fmt.Println("got", <-ch)
		fmt.Println("got", <-ch)
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	time.Sleep(time.Second)
}

func testChanClosedWhenSending() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer func() {
			if a := recover(); a != nil {
				fmt.Println("recover1:", a)
			}
		}()
		// when current goroutine is blocked on sending, while another goroutine closes the channel,
		// current goroutine will panic "send on closed channel"
		ch <- 1
	}()
	go func() {
		defer wg.Done()
		defer func() {
			if a := recover(); a != nil {
				fmt.Println("recover2:", a)
			}
		}()
		ch <- 2
	}()
	fmt.Println(<-ch)
	time.Sleep(time.Second)
	close(ch)
	wg.Wait()
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

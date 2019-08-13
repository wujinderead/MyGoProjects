package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//testCancel()
	//testTimeout()
	//testCancelLevel()
	//testTimeoutLevel()
	//testDeferCancel()
	//testAlreadyCancel()
	testWithValue()
}

func testCancel() {
	num := 10
	timeout := 5
	var wg sync.WaitGroup
	wg.Add(num)
	// when timeout, all <-ctx.Done() can receive from channel
	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= 10; i++ {
		go func(ctx context.Context, i int) {
			select {
			case <-time.After(time.Duration(i * int(time.Second))):
				fmt.Println(i, "done  :", time.Now(), ctx.Err()) // when non-canceled, ctx.Err() is nil
			case <-ctx.Done():
				fmt.Println(i, "cancel:", time.Now(), ctx.Err()) // after cancel, ctx.Err() is not nil (context canceled)
			}
			wg.Done()
		}(ctx, i)
	}
	<-time.After(time.Duration(timeout * int(time.Second)))
	// cancel() to close ctx.Down(), that makes all goroutines blocked on <-ctx.Down() receive immediately
	cancel()
	wg.Wait()
	fmt.Println("exit", time.Now())
}

func testTimeout() {
	done := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ticker := time.NewTicker(time.Second)
	go func() {
	outer:
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("get event:", t)
			case <-ctx.Done():
				fmt.Println("timeout:", time.Now(), ctx.Err()) // after cancel, ctx.Err() is non-nil (context deadline exceeded)
				ticker.Stop()
				break outer
			}
		}
		done <- struct{}{}
	}()
	// we can cancel() manually, or we just wait to timeout, it will cancel automatically
	if false {
		cancel()
	}
	<-done
	fmt.Println("exit", time.Now())
}

func testCancelLevel() {
	gp, gpcancel := context.WithCancel(context.Background())
	p1, _ := context.WithCancel(gp)
	p2, p2cancel := context.WithCancel(gp)
	c11, c11cancel := context.WithCancel(p1)
	c12, _ := context.WithCancel(p1)
	c21, _ := context.WithCancel(p2)
	c22, _ := context.WithCancel(p2)

	fn := func(ctx context.Context, wg *sync.WaitGroup, name string) {
		<-ctx.Done()
		t := time.Now()
		fmt.Println(name, "done at", t.UnixNano())
		wg.Done()
	}

	wg := new(sync.WaitGroup)
	wg.Add(7)
	go fn(gp, wg, "gp ")
	go fn(p1, wg, "p1 ")
	go fn(p2, wg, "p2 ")
	go fn(c11, wg, "c11")
	go fn(c12, wg, "c12")
	go fn(c21, wg, "c21")
	go fn(c22, wg, "c22")
	<-time.After(time.Second)
	c11cancel() // c11 cancel itself, and remove c11 from parent (p1)'s map
	<-time.After(time.Second)
	p2cancel() // p2 cancel itself and remove p2 from parent (gp)'s map; also cancel its children c21, c22
	<-time.After(time.Second)
	gpcancel() // gp cancel itself, and cancel its child p1; p1 cancel itself, and cancel its child c12
	wg.Wait()
}

func testTimeoutLevel() {
	gp, _ := context.WithTimeout(context.Background(), 3*time.Second) // gp timeout at 3
	p1, _ := context.WithTimeout(gp, 10*time.Second)
	p2, _ := context.WithTimeout(gp, 2*time.Second)  // p2 timeout at 2
	c11, _ := context.WithTimeout(p1, 1*time.Second) // c11 timeout at 1
	c12, _ := context.WithTimeout(p1, 10*time.Second)
	c21, _ := context.WithTimeout(p2, 10*time.Second)
	c22, _ := context.WithTimeout(p2, 10*time.Second)

	fn := func(ctx context.Context, wg *sync.WaitGroup, name string) {
		<-ctx.Done()
		t := time.Now()
		fmt.Println(name, "done at", t.UnixNano(), ctx.Err())
		wg.Done()
	}

	// c11 timeout at 1, and cancel c11
	// p2 timeout at 2, and cancel p2, c21, c22
	// gp timeout at 3, and cancel gp, p1, c12
	wg := new(sync.WaitGroup)
	wg.Add(7)
	go fn(gp, wg, "gp ")
	go fn(p1, wg, "p1 ")
	go fn(p2, wg, "p2 ")
	go fn(c11, wg, "c11")
	go fn(c12, wg, "c12")
	go fn(c21, wg, "c21")
	go fn(c22, wg, "c22")
	wg.Wait()
}

func testDeferCancel() {
	// defer cancel() to release resource if the Context is not canceled.
	// if the Context has been canceled, cancel() is no op.
	c1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	c2, cancel2 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel2()

	var wg sync.WaitGroup
	wg.Add(2)
	fn := func(ctx context.Context, name string) {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println(name, "done", time.Now())
		case <-ctx.Done():
			fmt.Println(name, "cancel", time.Now())
		}
		wg.Done()
	}
	// c1 timeout at 1, c1 is canceled, defer cancel1() is no op.
	// c2 job done at 2 (before timeout 3), call `defer cancel2()` to close c2.Done() to release resources.
	go fn(c1, "c1")
	go fn(c2, "c2")
	wg.Wait()
}

func testAlreadyCancel() {
	fmt.Println("start at", time.Now())
	parent, cancel := context.WithCancel(context.Background())
	cancel()
	child, _ := context.WithTimeout(parent, time.Second)
	<-child.Done() // unblock immediately. since parent is already canceled, child is also canceled immediately
	fmt.Println("end at", time.Now())
}

func testWithValue() {
	// valueContext is used to transfer kv pairs between processes and APIs.
	type mystring string
	cv1 := context.WithValue(context.Background(), mystring("hahaha"), "lalala")
	cv2 := context.WithValue(context.Background(), "hahaha", "lalala")
	fn := func(ctx context.Context) {
		if value := ctx.Value(mystring("hahaha")); value != nil {
			fmt.Println(value.(string))
		} else {
			fmt.Println("no such key")
		}
	}
	fn(cv1)
	fn(cv2)
}

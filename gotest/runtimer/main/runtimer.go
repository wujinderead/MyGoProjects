package main

import (
	"fmt"
	"runtime"
)

func main() {
	testRuntimeGet()
	testRuntimeStackMem()
}

func testRuntimeGet() {
	fmt.Println("go os:", runtime.GOOS)
	fmt.Println("go arch:", runtime.GOARCH)
	fmt.Println("go root:", runtime.GOROOT())
	fmt.Println("num cpu:", runtime.NumCPU())
	fmt.Println("num goroutine:", runtime.NumGoroutine())
	fmt.Println("version:", runtime.Version())
}

func testRuntimeStackMem() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	// get current goroutine trace
	stack := make([]byte, 1024)
	n := runtime.Stack(stack, false)
	fmt.Println("=== trace current:", string(stack[:n]))

	// get all goroutine trace, will 'stop the world'
	n = runtime.Stack(stack, true)
	fmt.Println("=== trace all:", string(stack[:n]))

	<-ch
	// get current memory stats
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Println("mem stats:", stats)
}

func testRuntime() {
	// 出让当前运行权，类似Java的Thread.yiled()
	runtime.Gosched()

	// 运行垃圾回收，其他goroutine会暂停
	runtime.GC()

	// 将当前gouroutine绑定到它当前运行在的OS线程上，当前gouroutine会始终运行在此线程上，
	// 其他的goroutine不会调度到此线程上，直到调用UnlockOSThread解绑。
	// 如果绑定解除之前goroutine退出了，线程会随着goroutine一起结束。
	// 一般用途：调用OS服务；或调用其他语言的服务，需要用到per-thread state的
	runtime.LockOSThread()
	runtime.UnlockOSThread()

	// 在KeepAlive之前的程序，这个变量都不会被GC回收。
	// 例如 a=1;xxxx;xxxx;xxxx;KeepAlive(a); 保证在xxxx;xxxx;xxxx;这些语句中a不会被回收
	runtime.KeepAlive(1)

	// 提供用户实现的finalizer
	runtime.SetFinalizer(1, func() {})
}

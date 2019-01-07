package runtimer

import (
	"testing"
	"fmt"
	"runtime"
)

func TestRuntimeGet(t *testing.T) {
	fmt.Println("go os:", runtime.GOOS)
	fmt.Println("go arch:", runtime.GOARCH)
	fmt.Println("go root:", runtime.GOROOT())
	fmt.Println("num cpu:", runtime.NumCPU())
	fmt.Println("num goroutine:", runtime.NumGoroutine())
	fmt.Println("version:", runtime.Version())

	// get current goroutine trace
	stack := make([]byte, 1024)
	n := runtime.Stack(stack, false)
	fmt.Println("trace:", string(stack[:n]))

	// get all goroutine trace, will 'stop the world'
	n = runtime.Stack(stack, true)
	fmt.Println("trace:", string(stack[:n]))

	// get current memory stats
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Println("mem stats:", stats)
}

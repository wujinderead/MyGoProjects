package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	//testRuntimeGet()
	//testRuntimeStackMem()
	//testGC()
	testPprofRuntime()
	//testTrace()
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

/*
go build -o /tmp/aaa /home/xzy/golang/gotest/runtimer/main/runtimer.go

# default GOGC=100
GODEBUG=gctrace=1 /tmp/aaa

# gc trace information
gc 49 @0.079s 13%: 0.045+0.17+0.005 ms clock, 0.36+0.050/0.085/0+0.041 ms cpu, 8->9->1 MB, 9 MB goal, 8 P

# fields explained
gc 49                                 the 49th gc
@0.079s                               0.079s since program start
13%                                   13% cpu time spent on gc
0.045+0.17+0.005 ms clock             wall time for 1st stw, current mark and scan, 2nd stw
0.36+0.050/0.085/0+0.041 ms cpu       cpu time for 1st stw, current mark and scan, 2nd stw
8->9->1 MB,                           heap size when gc start -> when gc end -> alive objects size
9 MB goal                             target heap size for next gc
8 P                                   number of P

# scavenge trace information
scvg: 0 MB released
scvg: inuse: 4, idle: 58, sys: 63, released: 52, consumed: 10 (MB)

# GOGC=200, gc is less frequent, but the heap size to trigger gc is increased
GOGC=200 GODEBUG=gctrace=1 /tmp/aaa

# trace schedule status
GODEBUG=schedtrace=500,scheddetail=1 /tmp/aaa

# current threads information
  SCHED 283ms: gomaxprocs=8 idleprocs=6 threads=16 spinningthreads=1 idlethreads=8 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
# details of P, M, G
  P0: status=0 schedtick=1793 syscalltick=96 m=-1 runqsize=0 gfreecnt=6
  P1: status=0 schedtick=5640 syscalltick=93 m=-1 runqsize=0 gfreecnt=10
  P2: status=0 schedtick=5164 syscalltick=153 m=-1 runqsize=0 gfreecnt=25
  P3: status=0 schedtick=2341 syscalltick=104 m=-1 runqsize=0 gfreecnt=31
  P4: status=1 schedtick=1013 syscalltick=93 m=8 runqsize=0 gfreecnt=7
  P5: status=0 schedtick=379 syscalltick=38 m=-1 runqsize=0 gfreecnt=0
  P6: status=0 schedtick=181 syscalltick=20 m=-1 runqsize=0 gfreecnt=0
  P7: status=0 schedtick=119 syscalltick=6 m=-1 runqsize=0 gfreecnt=0
  M15: p=-1 curg=97 mallocing=0 throwing=0 preemptoff= locks=2 dying=0 spinning=false blocked=false lockedg=-1
  M14: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M13: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M12: p=6 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 spinning=true blocked=false lockedg=-1
  M11: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M10: p=-1 curg=35 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M9: p=-1 curg=71 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M8: p=4 curg=9 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 spinning=false blocked=true lockedg=-1
  M7: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M6: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M5: p=-1 curg=81 mallocing=0 throwing=0 preemptoff= locks=2 dying=0 spinning=false blocked=false lockedg=-1
  M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M2: p=-1 curg=113 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 spinning=false blocked=false lockedg=-1
  M0: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 spinning=false blocked=true lockedg=-1
  G1: status=4(wait for GC cycle) m=-1 lockedm=-1
  G2: status=4(force gc (idle)) m=-1 lockedm=-1
  G3: status=4(GC sweep wait) m=-1 lockedm=-1
  G4: status=4(sleep) m=-1 lockedm=-1
  G5: status=4(finalizer wait) m=-1 lockedm=-1
  G51: status=4(timer goroutine (idle)) m=-1 lockedm=-1
  G7: status=6() m=-1 lockedm=-1
  G8: status=6() m=-1 lockedm=-1
  G9: status=4(GC worker (idle)) m=8 lockedm=-1
  G10: status=4(GC worker (idle)) m=-1 lockedm=-1
  G17: status=4(GC worker (idle)) m=-1 lockedm=-1
  G551: status=6() m=-1 lockedm=-1
  G113: status=3(timer goroutine (idle)) m=2 lockedm=-1
  ......
*/

func testGC() {
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	displayMemStat(stats)
	fmt.Println()

	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		ss := make([]byte, 1024*1024)
		ss[0] = 'a'
		ss[1024*1024-1] = 'd'
		go func() {
			s := make([]byte, 1024*1024)
			s[0] = 'a'
			s[1024*1024-1] = 'd'
			s = nil
			wg.Done()
		}()
	}
	wg.Wait()
	runtime.ReadMemStats(stats)
	displayMemStat(stats)
	fmt.Println()

	for i := 0; i < 500; i++ {
		s := make([]byte, 1024*1024)
		s[0] = 'a'
		s[1024*1024-1] = 'd'
		if i%2 == 0 {
			s = nil
		}
	}
	runtime.GC() // force GC
	runtime.ReadMemStats(stats)
	displayMemStat(stats)
	fmt.Println()
}

/*
# display trace on browser
go build -o /tmp/aaa runtimer/main/runtimer.go
/tmp/aaa
go tool trace /tmp/trace1
*/
func testTrace() {
	f, err := os.OpenFile("/tmp/trace1", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file err: err")
		return
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		fmt.Println("trace start err:", err)
		return
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		ss := make([]byte, 1024*1024)
		ss[0] = 'a'
		ss[1024*1024-1] = 'd'
		go func() {
			s := make([]byte, 1024*1024)
			s[0] = 'a'
			s[1024*1024-1] = 'd'
			s = nil
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
# output different memory profiles:
    goroutine     goroutineProfile
    threadcreate  threadcreateProfile
    heap          heapProfile
    allocs        allocsProfile
    block         blockProfile
    mutex         mutexProfile

go build -o /tmp/aaa runtimer/main/runtimer.go
/tmp/aaa         # binary executable file

# we can just use 'go tool pprof /tmp/cpu', but referring the binary is better
go tool pprof -http :8080 /tmp/aaa /tmp/cpu
go tool pprof -http :8080 /tmp/aaa /tmp/goroutine
go tool pprof -http :8080 /tmp/aaa /tmp/threadcreate
go tool pprof -http :8080 /tmp/aaa /tmp/heap
go tool pprof -http :8080 /tmp/aaa /tmp/allocs
go tool pprof -http :8080 /tmp/aaa /tmp/block
go tool pprof -http :8080 /tmp/aaa /tmp/mutex
*/
func testPprofRuntime() {
	// write cpu profile
	f, err := os.OpenFile("/tmp/cpu", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		fmt.Println("write cpu profile err:", err)
		return
	}
	defer pprof.StopCPUProfile()

	// concurrent quick sort for a long array
	r := rand.New(rand.NewSource(time.Now().Unix()))
	arr := r.Perm(1e7)
	//quickSortDoneChan(arr, 0, len(arr), nil)
	quickSortWaitGroup(arr, 0, len(arr), nil)

	// write pprof profile
	for _, s := range []string{"goroutine", "threadcreate", "heap", "allocs", "block", "mutex"} {
		f, err := os.OpenFile("/tmp/"+s, os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
		if err != nil {
			fmt.Println("open file err:", err)
			return
		}
		defer f.Close()
		p := pprof.Lookup(s)
		err = p.WriteTo(f, 0)
		if err != nil {
			fmt.Printf("write %s profile err: %s\n", s, err)
			continue
		}
	}
}

func displayMemStat(s *runtime.MemStats) {
	fmt.Println("Alloc        :", str(s.Alloc))        // Alloc is bytes of allocated heap objects.
	fmt.Println("TotalAlloc   :", str(s.TotalAlloc))   // TotalAlloc is cumulative bytes allocated for heap objects.
	fmt.Println("Sys          :", str(s.Sys))          // Sys is the total bytes of memory obtained from the OS.
	fmt.Println("Lookups      :", s.Lookups)           // Lookups is the number of pointer lookups performed by the runtime.
	fmt.Println("Mallocs      :", s.Mallocs)           // Mallocs is the cumulative count of heap objects allocated.
	fmt.Println("Frees        :", s.Frees)             // Frees is the cumulative count of heap objects freed.
	fmt.Println("HeapAlloc    :", str(s.HeapAlloc))    // HeapAlloc is bytes of allocated heap objects.
	fmt.Println("HeapSys      :", str(s.HeapSys))      // HeapSys is bytes of heap memory obtained from the OS.
	fmt.Println("HeapIdle     :", str(s.HeapIdle))     // HeapIdle is bytes in idle (unused) spans.
	fmt.Println("HeapInuse    :", str(s.HeapInuse))    // HeapInuse is bytes in in-use spans.
	fmt.Println("HeapReleased :", str(s.HeapReleased)) // HeapReleased is bytes of physical memory returned to the OS.
	fmt.Println("HeapObjects  :", s.HeapObjects)       // HeapObjects is the number of allocated heap objects.
	fmt.Println("StackInuse   :", str(s.StackInuse))   // StackInuse is bytes in stack spans.
	fmt.Println("StackSys     :", str(s.StackSys))     // StackSys is bytes of stack memory obtained from the OS.
	fmt.Println("MSpanInuse   :", str(s.MSpanInuse))   // MSpanInuse is bytes of allocated mspan structures.
	fmt.Println("MSpanSys     :", str(s.MSpanSys))     // MSpanSys is bytes of memory obtained from the OS for mspan structures.
	fmt.Println("MCacheInuse  :", str(s.MCacheInuse))  // MCacheInuse is bytes of allocated mcache structures.
	fmt.Println("MCacheSys    :", str(s.MCacheSys))    // MCacheSys is bytes of memory obtained from the OS for mcache structures.
	fmt.Println("BuckHashSys  :", str(s.BuckHashSys))  // BuckHashSys is bytes of memory in profiling bucket hash tables.
	fmt.Println("GCSys        :", str(s.GCSys))        // GCSys is bytes of memory in garbage collection metadata.
	fmt.Println("OtherSys     :", str(s.OtherSys))     // OtherSys is bytes of memory in miscellaneous off-heap runtime allocations.
	fmt.Println("NextGC       :", str(s.NextGC))       // NextGC is the target heap size of the next GC cycle.
	fmt.Println("PauseTotalNs :", s.PauseTotalNs)      // PauseTotalNs is the cumulative nanoseconds in GC stop-the-world pauses since the program started.
	fmt.Println("PauseNs      :", s.PauseNs)           // PauseNs is a circular buffer of recent GC stop-the-world pause times in nanoseconds.
	fmt.Println("NumGC        :", s.NumGC)             // NumGC is the number of completed GC cycles.
	fmt.Println("NumForcedGC  :", s.NumForcedGC)       // NumForcedGC is the number of GC cycles that were forced by the application calling the GC function.
	fmt.Println("GCCPUFraction:", s.GCCPUFraction)     // GCCPUFraction is the fraction of this program's available CPU time used by the GC since the program started.
}

func str(s uint64) string {
	if s == 0 {
		return "0"
	} else if s < 1024 {
		return fmt.Sprintf("%4d", s)
	} else if s < 1024*1024 {
		return fmt.Sprintf("%4.1f KB", float64(s)/1024)
	} else if s < 1024*1024*1024 {
		return fmt.Sprintf("%4.1f MB", float64(s)/(1024*1024))
	} else {
		return fmt.Sprintf("%4.1f GB", float64(s)/(1024*1024*1024))
	}
}

func execTop() {
	pid := syscall.Getpid()
	cmd := exec.Command("top", "-b", "-p", strconv.FormatInt(int64(pid), 10))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("top err:", err)
		return
	}
	time.Sleep(1 * time.Second)              // wait to print
	err = cmd.Process.Signal(syscall.SIGINT) // interrupt top process
	if err != nil {
		fmt.Println("interrupt err:", err)
		return
	}
}

func quickSortWaitGroup(arr []int, start, end int, pwg *sync.WaitGroup) {
	if end-start < 12 {
		insertionSort(arr, start, end)
		if pwg != nil {
			pwg.Done()
		}
		return
	}
	var wg sync.WaitGroup
	wg.Add(2)
	pivot := partition(arr, start, end)
	go quickSortWaitGroup(arr, start, pivot, &wg)
	go quickSortWaitGroup(arr, pivot+1, end, &wg)
	wg.Wait()
	if pwg != nil {
		pwg.Done()
	}
}

// sort array from arr[start: end], start is inclusive, end is exclusive
func quickSortDoneChan(arr []int, start, end int, pdone chan struct{}) {
	if end-start < 12 {
		insertionSort(arr, start, end)
		if pdone != nil {
			pdone <- struct{}{}
		}
		return
	}
	done := make(chan struct{})
	pivot := partition(arr, start, end)
	go quickSortDoneChan(arr, start, pivot, done)
	go quickSortDoneChan(arr, pivot+1, end, done)
	<-done
	<-done
	if pdone != nil {
		pdone <- struct{}{}
	}
}

// merge two sorted array arr[start: mid], arr[mid: end]
func partition(arr []int, start, end int) int {
	mid := start + (end-start)/2
	// find the median value of (arr[start], arr[mid], arr[end-1]),
	// use it as pivot and place it at arr[end-1]
	if arr[start] > arr[end-1] {
		arr[start], arr[end-1] = arr[end-1], arr[start]
	}
	if arr[mid] < arr[start] {
		arr[start], arr[end-1] = arr[end-1], arr[start]
	} else if arr[mid] < arr[end-1] {
		arr[mid], arr[end-1] = arr[end-1], arr[mid]
	}
	// partition
	i := start - 1
	p := arr[end-1]
	for j := start; j < end-1; j++ {
		if arr[j] < p {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[end-1] = arr[end-1], arr[i+1]
	return i + 1
}

func insertionSort(arr []int, start, end int) {
	for i := start + 1; i < end; i++ {
		for j := i - 1; j >= start && arr[j+1] < arr[j]; j-- {
			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}
}

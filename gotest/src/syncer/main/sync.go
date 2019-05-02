package main

import (
	"sync/atomic"
	"fmt"
	"sync"
	"time"
	"unsafe"
	"runtime"
	"math/rand"
)

func main() {
	//testOnce()
	//testMutex()
	//testCond()
	//testAtomic()
	//testAtomicValue()
	testAtomicDifferentType()
}

func testOnce() {
	var once sync.Once
	a := 0
	for i:=0; i<3; i++ {
		once.Do(func() {
			a += 1
		})
	}
	fmt.Print(a)
}

func testMutex() {
	var mutex sync.Mutex
	a := 0
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		a += 1
	}()
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		a += 1
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func testCond() {
	lock := new(sync.Mutex)
	cond := sync.NewCond(lock)
	a := 1
	start := time.Now().UnixNano()
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		lock.Lock()
		fmt.Println("a1 lock at", time.Now().UnixNano()-start)
		for a!=3 {
			cond.Wait()
			fmt.Println("a1 wake at", time.Now().UnixNano()-start)
		}
		lock.Unlock()
		fmt.Println("a1 unlock at", time.Now().UnixNano()-start)
		wg.Done()
	}()
	go func() {
		lock.Lock()
		fmt.Println("a2 lock at", time.Now().UnixNano()-start)
		for a!=3 {
			cond.Wait()
			fmt.Println("a1 wake at", time.Now().UnixNano()-start)
		}
		lock.Unlock()
		fmt.Println("a2 unlock at", time.Now().UnixNano()-start)
		wg.Done()
	}()
	time.Sleep(50*time.Millisecond)
	go func() {
		lock.Lock()
		fmt.Println("t1 lock at", time.Now().UnixNano()-start)
		a += 1
		cond.Broadcast()
		time.Sleep(100*time.Millisecond)
		fmt.Println("t1 unlock at", time.Now().UnixNano()-start)
		lock.Unlock()
		wg.Done()
	}()
	go func() {
		lock.Lock()
		fmt.Println("t2 lock at", time.Now().UnixNano()-start)
		a += 1
		cond.Broadcast()
		time.Sleep(100*time.Millisecond)
		fmt.Println("t2 unlock at", time.Now().UnixNano()-start)
		lock.Unlock()
		wg.Done()
	}()
	j:=0
	for i:=0; i<5; i++ {
		go func() {
			wg.Wait()
			j+=1
			fmt.Println(j, "wait end")
		}()
	}
	time.Sleep(3000*time.Millisecond)
	wg.Wait()
	fmt.Println("a=", a)
}

func testAtomic() {
	type pojo struct {
		a int
		b string
	}
	var inter32 int32
	var inter64 int64
	var uinter32 uint32
	var uinter64 uint64
	var uintptrer uintptr
	var structer unsafe.Pointer
	var apojo = pojo{1,"aaa"}
	var bpojo = pojo{2,"bbb"}

	fmt.Printf("apojo: %p = %d, bpojo: %p = %d\n",
		&apojo, uintptr(unsafe.Pointer(&apojo)), &bpojo, uintptr(unsafe.Pointer(&bpojo)))
	atomic.StoreInt32(&inter32, -123)
	atomic.StoreInt64(&inter64, -123)
	atomic.StoreUint32(&uinter32, 0x12345678)
	atomic.StoreUint64(&uinter64, 0x12345678)
	atomic.StoreUintptr(&uintptrer, uintptr(unsafe.Pointer(&apojo)))
	atomic.StorePointer(&structer, unsafe.Pointer(&apojo))

	fmt.Println(atomic.LoadInt32(&inter32))
	fmt.Println(atomic.LoadInt64(&inter64))
	fmt.Println(atomic.LoadUint32(&uinter32))
	fmt.Println(atomic.LoadUint64(&uinter64))
	fmt.Println(atomic.LoadUintptr(&uintptrer))
	fmt.Println(atomic.LoadPointer(&structer), (*pojo)(structer))
	fmt.Println()

	fmt.Println(atomic.CompareAndSwapInt32(&inter32, -123, -456))
	fmt.Println(atomic.CompareAndSwapInt64(&inter64, -123, -456))
	fmt.Println(atomic.CompareAndSwapUint32(&uinter32, 0x12345678, 0x98765432))
	fmt.Println(atomic.CompareAndSwapUint64(&uinter64, 0x12345678, 0x98765432))
	fmt.Println(atomic.CompareAndSwapUintptr(&uintptrer, uintptr(unsafe.Pointer(&apojo)), uintptr(unsafe.Pointer(&bpojo))))
	fmt.Println(atomic.CompareAndSwapPointer(&structer, unsafe.Pointer(&apojo), unsafe.Pointer(&bpojo)))
	fmt.Println()

	fmt.Println(atomic.SwapInt32(&inter32, -789))
	fmt.Println(atomic.SwapInt64(&inter64, -789))
	fmt.Println(atomic.SwapUint32(&uinter32, 123456789))
	fmt.Println(atomic.SwapUint64(&uinter64, 123456789))
	fmt.Println(atomic.SwapUintptr(&uintptrer, uintptr(unsafe.Pointer(&apojo))))
	fmt.Println((*pojo)(atomic.SwapPointer(&structer, unsafe.Pointer(&apojo))))
	fmt.Println()

	fmt.Println(inter32)
	fmt.Println(inter64)
	fmt.Println(uinter32)
	fmt.Println(uinter64)
	fmt.Println(uintptrer)
	fmt.Println((*pojo)(structer))
	fmt.Println()

	fmt.Println(atomic.AddInt32(&inter32, 100))
	fmt.Println(atomic.AddInt64(&inter64, 100))
	fmt.Println(atomic.AddUint32(&uinter32, 100))
	fmt.Println(atomic.AddUint64(&uinter64, 100))
	fmt.Println(atomic.AddUintptr(&uintptrer, 100))
	fmt.Println()
}

func testAtomicValue() {
	type pojo struct {
		a int
		b string
	}
	var holder atomic.Value
	fmt.Println(runtime.GOMAXPROCS(6))
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<100; i++ {
		// atomic value can store and load simultaneously
		go func() {
			holder.Store(pojo{rand.Int(), "aaa"})
		}()
		go func() {
			holder.Load()
		}()
	}
	time.Sleep(100*time.Millisecond)
	fmt.Println(holder.Load().(pojo))
}

func testAtomicDifferentType() {
	type a struct {
		aa int
		bb string
	}
	var v atomic.Value
	v.Store(a{1, "bbb"})
	v.Store(a{2, "ccc"})
	v.Store(&a{3, "ddd"})  // panic here, value should be consistent with previous type
}

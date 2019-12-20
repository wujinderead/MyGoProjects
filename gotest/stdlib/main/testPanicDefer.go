package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	//testPanic()
	//testGoroutinePanic()
	//testTwice()
	//testCantRecover()
	//testBothDefer()
	testDeferScope()
	//testOncePanic()
}

func testPanic() {
	panic1()
	panic2()
	panic3()
}

func panic1() {
	defer func() {
		a := recover()
		if str, ok := a.(string); ok {
			log.Println("got str err:", str)
		} else if inter, ok := a.(int); ok {
			log.Println("got int err:", inter)
		} else {
			panic(a)
		}
	}()
	panic(123)
}

func panic2() {
	defer func() {
		a := recover()
		if str, ok := a.(string); ok {
			log.Println("got str err:", str)
		} else if inter, ok := a.(int); ok {
			log.Println("got int err:", inter)
		} else {
			panic(a)
		}
	}()
	panic("morelia")
}

func panic3() {
	defer func() {
		a := recover()
		if str, ok := a.(string); ok {
			log.Println("got str err:", str)
		} else if inter, ok := a.(int); ok {
			log.Println("got int err:", inter)
		} else {
			panic(a)
		}
	}()
	panic(123.45)
}

func testGoroutinePanic() {
	ch := make(chan struct{})
	go func() {
		defer func() {
			a := recover()
			if str, ok := a.(string); ok {
				log.Println("got str err:", str)
			}
			ch <- struct{}{}
		}()
		panic("panic1")
	}()
	go func() {
		defer func() {
			// _ = recover()  // runs ok when uncomment this, otherwise, panic in goroutine makes the whole process exit
			ch <- struct{}{}
		}()
		panic("panic2")
	}()
	for i := 0; i < 2; i++ {
		<-ch
	}
	log.Println("main done")
}

func testRecoverTwice() {
	// output:
	// can't recover: a
	// i can recover: a
	defer func() {
		if i := recover(); i != nil { // recover in outer defer
			fmt.Println("i can recover:", i.(string))
		}
	}()
	defer func() {
		if i := recover(); i != nil {
			fmt.Println("can't recover:", i.(string))
			panic(i) // panic in inner defer
		}
	}()
	panic("a")
}

func testCantRecover() {
	defer func() {
		if i := recover(); i != nil {
			fmt.Println("can't recover:", i.(string))
			panic(i) // the panic trace stack will print both
		}
	}()
	panic("a")
}

func testBothDefer() {
	defer func() {
		fmt.Println("main defer")
		// can't recover other goroutine's panic,
		// because this defer won't be executed when other goroutine panicked
		r := recover()
		fmt.Println(r)
	}()
	go func() {
		defer func() {
			fmt.Println("inner defer")
			// current goroutine's panic will cause the whole process exit,
			// the only chance to fix is to recover() here.

			// err := recover()
		}()
		time.Sleep(10 * time.Millisecond)
		// panic in goroutine makes the whole process exit.
		// only the panic goroutine's defer func will be executed,
		// other goroutines' defer (including main routine) won't be executed
		panic("go1 panic")
	}()
	// block endless. not like 'for{}' which cost system resource, 'select{}' yield and block
	select {}
}

func testDeferScope() {
	// output: 2 4 5 3 7 9 8 6 1
	defer fmt.Println(1)
	fmt.Println(2)
	func() { // defer adjust to func's scope
		defer fmt.Println(3)
		fmt.Println(4)
		defer fmt.Println(5)
	}()
	{ // this scope just adjust to variables
		defer fmt.Println(6)
		fmt.Println(7)
		defer fmt.Println(8)
	}
	defer fmt.Println(9)
}

func testOncePanic() {
	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(2)
	f := func() {
		time.Sleep(1000)
		fmt.Println("once done")
		panic("aaa")
	}
	go func() {
		defer wg.Done()
		defer func() { // Once consider f() done even f() panics, however, it's caller's responsibility to recover
			a := recover()
			fmt.Println("r1:", a)
		}()
		once.Do(f)
		fmt.Println("1 done")
	}()
	go func() {
		defer wg.Done()
		defer func() {
			a := recover()
			fmt.Println("r2:", a)
		}()
		once.Do(f)
		fmt.Println("2 done")
	}()
	wg.Wait()
	once.Do(f) // no op
}

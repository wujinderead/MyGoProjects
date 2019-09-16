package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//testPanic()
	//testGoroutinePanic()
	//testTwice()
	//testCantRecover()
	testBothDefer()
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
	}()
	go func() {
		defer func() {
			fmt.Println("inner defer")
		}()
		time.Sleep(10 * time.Millisecond)
		// panic in goroutine makes the whole process exit.
		// only the panic goroutine's defer func will be executed,
		// other goroutines' defer (including main routine) won't be executed
		panic("go1 panic")
	}()
	var a int
	for {
		a++
	}
}

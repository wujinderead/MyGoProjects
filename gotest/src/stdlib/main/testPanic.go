package main

import (
	"log"
)

func main() {
	//testPanic()
	testGoroutinePanic()
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

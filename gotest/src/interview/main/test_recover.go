package main

import (
	"fmt"
	"reflect"
)

func main1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("++++")
			fmt.Println(err, reflect.TypeOf(err).Kind().String())
		} else {
			fmt.Println("fatal")
		}
	}()

	// only the last panic is catched by recover()
	defer panic(55)
	defer func() {
		panic(func() string {
			return "defer panic"
		})
	}()
	panic("string panic")
}

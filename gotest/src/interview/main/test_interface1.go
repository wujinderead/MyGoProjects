package main

import (
	"fmt"
)

type People1 interface {
	Show()
}

type Student1 struct{}

func (stu *Student1) Show() {

}

func live() People1 {
	var stu *Student1
	fmt.Println("stu: ", stu)
	return stu
}

func main() {
	a := live()
	fmt.Println("a:", a)
	if a == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
	/*
		result: BBBBBB
		the 'data' ptr of interface is null, but the interface is not null
		golang interface:
	type eface struct {      //空接口
	    _type *_type         //类型信息
	    data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
	}
	type iface struct {      //带有方法的接口
	    tab  *itab           //存储type信息还有结构实现方法的集合
	    data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
	}
	*/
}

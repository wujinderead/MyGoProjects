package stdlib

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

type comparator interface {
	Compare(interface{}) int
}

type Int int

func (a Int) Compare(other interface{}) int {
	return int(a) - int(other.(Int))
}

func TestInterfaceCompare(t *testing.T) {
	a := Int(8)
	b := Int(16)
	if a.Compare(b) > 0 {
		fmt.Println("bigger")
	} else {
		fmt.Println("smaller")
	}
}

func TestNilInterface(t *testing.T) {
	var a interface{} = nil
	var b interface{} = 8
	var c interface{} = new(int16)
	fmt.Println(a == nil)
	fmt.Println(b.(int))
	fmt.Println(c.(*int16))
	l := list.New()
	l.PushBack(nil)
	l.PushBack(2)
	l.PushBack(2.3)
	niler := l.Front()
	fmt.Println(niler.Value == nil)
	fmt.Println(niler.Value != nil)
	l.Remove(niler)
	inter := l.Front()
	fmt.Println(inter.Value.(int))
	l.Remove(inter)
	floater := l.Front()
	fmt.Println(floater.Value.(float64))
	l.Remove(floater)
	fmt.Println(l.Front())
}

func TestNilInterface1(t *testing.T) {
	// b is nil
	var b interface{} = nil
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b == nil)
	fmt.Println()

	// a is not nil, a.type is *int, a.value is nil
	var p *int = nil
	var a interface{} = p
	if v, ok := a.(*int); ok {
		fmt.Println("a=", a, ", a == nil ?, ", a==nil)
		fmt.Println(reflect.TypeOf(a))
		fmt.Println(v)
	}
	fmt.Println()

	var c interface{} = new(int16)
	if v, ok := c.(*int16); ok {
		fmt.Println(reflect.TypeOf(c))
		fmt.Println(v)
		fmt.Println(*v)
	}
	fmt.Println()
}

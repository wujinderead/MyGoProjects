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
	var a *string = nil        // a, b, c, d == nil
	var b []string = nil
	var c *int = nil
	var d map[int]int = nil
	var ai interface{} = a     // ai, bi, ci, di != nil (interface struct is non-nil, only data is nil)
	var bi interface{} = b
	var ci interface{} = c
	var di interface{} = d
	var ei interface{} = nil   // ei ==nil
	fmt.Println(a==nil, b==nil, c==nil, d==nil)
	fmt.Println(ai==nil, bi==nil, ci==nil, di==nil, ei==nil)
	// *string, []string, *int, map[int]int; type is retained
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b),
		reflect.TypeOf(c), reflect.TypeOf(d))
	// *string, []string, *int, map[int]int; type is retained
	fmt.Println(reflect.TypeOf(ai), reflect.TypeOf(bi),
		reflect.TypeOf(ci), reflect.TypeOf(di), reflect.TypeOf(ei))
	// type is <nil>
	fmt.Println(reflect.TypeOf(ei))
}

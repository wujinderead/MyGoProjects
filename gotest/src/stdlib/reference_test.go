package stdlib

import (
	"fmt"
	"testing"
)

func TestReference(t *testing.T) {
	type aa struct {
		a int
		b string
	}
	a := aa{1, "aa"}
	pa := &aa{2, "bb"}
	func (x aa) {
		x.a = 3
		x.b = "tt"
	}(a)
	fmt.Println(a)
	func (x *aa) {
		x.a = 4
		x.b = "dd"
	}(pa)
	fmt.Println(pa)

	var inter = 3
	var pint = &inter
	func (x int) {
		x = 4
	}(inter)
	fmt.Println(inter)
	func (x *int) {
		*x = 4
	}(pint)
	fmt.Println(inter)

	arr := []int{1, 2, 3}
	func (x []int) {
		x[0] = 100
	}(arr)
	fmt.Println(arr)
	func (x *int) {
		*x = 4
	}(&arr[1])
	fmt.Println(arr)
	func (x *[]int) {
		(*x)[0] = 99
	}(&arr)
	fmt.Println(arr)
	func (x int) {
		x = 7
	}(arr[2])
	fmt.Println(arr)
}

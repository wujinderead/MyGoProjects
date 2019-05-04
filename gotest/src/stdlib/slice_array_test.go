package stdlib

import (
	"sync"
	"testing"
	"fmt"
	"reflect"
)

func TestAppend(t *testing.T) {
	x := []string{"start"}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		y := append(x, "hello", "world")
		t.Log(cap(y), len(y))
	}()
	go func() {
		defer wg.Done()
		z := append(x, "goodbye", "bob")
		t.Log(cap(z), len(z))
	}()
	wg.Wait()

}

func TestAppend1(t *testing.T) {
	x := make([]string, 0, 6)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		y := append(x, "hello", "world")
		t.Log(len(y))
	}()
	go func() {
		defer wg.Done()
		z := append(x, "goodbye", "bob")
		t.Log(len(z))
	}()
	wg.Wait()
}

func TestFixedSizeArray(t *testing.T)  {
	a := []int{1,2,3}
	b := [3]int{1,2,3}
	bp := &b
	c := [...]int{1,2,3}    // fixed-size array not specify length
	_ = bp[1]               // pointer of fixed size array can be indexed directly
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(b), reflect.TypeOf(c))
}
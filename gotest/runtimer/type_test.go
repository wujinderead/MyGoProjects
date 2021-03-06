package runtimer

import (
	"crypto/elliptic"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"testing"
	"unsafe"
)

type airport struct {
	name string
	city string
	iata [3]byte
	iaco [4]byte
}

func TestTypes(t *testing.T) {
	// for most types, not only the common _type is stored, but also other information of the type
	{
		var slicer interface{} = make([]airport, 3, 6)
		efacer := (*eface)(unsafe.Pointer(&slicer))
		slicet := (*slicetype)(unsafe.Pointer(efacer._type)) // extended to from *_type to *slicetype
		fmt.Println("reflect type    :", reflect.TypeOf(slicer).String())
		fmt.Println("slice size      :", slicet.typ.size) // 24 for slice
		fmt.Println("slice hash      :", slicet.typ.hash)
		fmt.Println("slice kind      :", slicet.typ.kind, reflect.Kind(slicet.typ.kind))
		fmt.Println("slice str       :", slicet.typ.str)
		fmt.Println("slice ptrToThis :", slicet.typ.ptrToThis)

		// slice element (airport) type information
		fmt.Println("slice elem size      :", slicet.elem.size) // airport size: 40
		fmt.Println("slice elem hash      :", slicet.elem.hash)
		fmt.Println("slice elem kind      :", slicet.elem.kind, reflect.Kind(slicet.elem.kind))
		fmt.Println("slice elem str       :", slicet.elem.str)
		fmt.Println("slice elem ptrToThis :", slicet.elem.ptrToThis)
		fmt.Println()
	}

	{
		var arrayer interface{} = [...]airport{{"a", "b", [...]byte{1, 2, 3}, [...]byte{1, 2, 3, 4}}, {}, {}}
		efacer := (*eface)(unsafe.Pointer(&arrayer))
		arrayt := (*arraytype)(unsafe.Pointer(efacer._type)) // extended to from *_type to *arraytype
		fmt.Println("reflect type    :", reflect.TypeOf(arrayer).String())
		fmt.Println("array size      :", arrayt.typ.size) // 120 = 3×40
		fmt.Println("array hash      :", arrayt.typ.hash)
		fmt.Println("array kind      :", arrayt.typ.kind, reflect.Kind(arrayt.typ.kind))
		fmt.Println("array str       :", arrayt.typ.str)
		fmt.Println("array ptrToThis :", arrayt.typ.ptrToThis)

		// array element (airport) type information
		fmt.Println("array elem size      :", arrayt.elem.size) // [3]airport size: 120
		fmt.Println("array elem hash      :", arrayt.elem.hash)
		fmt.Println("array elem kind      :", arrayt.elem.kind, reflect.Kind(arrayt.elem.kind))
		fmt.Println("array elem str       :", arrayt.elem.str)
		fmt.Println("array elem ptrToThis :", arrayt.elem.ptrToThis)
		fmt.Println("array len            :", arrayt.len)

		// array still need a underlying slice
		fmt.Println("array slice size      :", arrayt.slice.size)
		fmt.Println("array slice hash      :", arrayt.slice.hash)
		fmt.Println("array slice kind      :", arrayt.slice.kind, reflect.Kind(arrayt.slice.kind))
		fmt.Println("array slice str       :", arrayt.slice.str)
		fmt.Println("array slice ptrToThis :", arrayt.slice.ptrToThis)
		fmt.Println()
	}

	{
		var pointer interface{} = &airport{"a", "b", [...]byte{1, 2, 3}, [...]byte{1, 2, 3, 4}}
		efacer := (*eface)(unsafe.Pointer(&pointer))
		pointert := (*ptrtype)(unsafe.Pointer(efacer._type))                 // extended to from *_type to *ptrtype
		fmt.Println("reflect type      :", reflect.TypeOf(pointer).String()) // *airport
		fmt.Println("pointer size      :", pointert.typ.size)                // 8 for pointer
		fmt.Println("pointer hash      :", pointert.typ.hash)
		fmt.Println("pointer kind      :", pointert.typ.kind, reflect.Kind(pointert.typ.kind))
		fmt.Println("pointer str       :", pointert.typ.str)
		fmt.Println("pointer ptrToThis :", pointert.typ.ptrToThis)

		// pointer element (airport) type information
		fmt.Println("pointer elem size      :", pointert.elem.size) // airport size: 40
		fmt.Println("pointer elem hash      :", pointert.elem.hash)
		fmt.Println("pointer elem kind      :", pointert.elem.kind, reflect.Kind(pointert.elem.kind))
		fmt.Println("pointer elem str       :", pointert.elem.str)
		fmt.Println("pointer elem ptrToThis :", pointert.elem.ptrToThis) // *airport addr
		fmt.Println()
	}
}

func TestFuncs(t *testing.T) {
	{
		// funcer1 and funcer2 are the same function type with different body, so they are of the same *functype
		var funcer1 interface{} = func(elliptic.Curve, io.Reader) ([]uint8, *big.Int, *big.Int, error) {
			return []uint8{}, nil, nil, nil
		}
		efacer := (*eface)(unsafe.Pointer(&funcer1))
		funcert := (*functype)(unsafe.Pointer(efacer._type))                   // extended to from *_type to *functype
		fmt.Println("reflect type     :", reflect.TypeOf(funcer1).String())    // *airport
		fmt.Println("funcer size      :", funcert.typ.size, efacer._type.size) // 8 for func
		fmt.Println("funcer hash      :", funcert.typ.hash)
		fmt.Println("funcer kind      :", funcert.typ.kind, reflect.Kind(funcert.typ.kind))
		fmt.Println("funcer str       :", funcert.typ.str)
		fmt.Println("funcer ptrToThis :", funcert.typ.ptrToThis)

		fmt.Println("funcer incount   :", funcert.inCount)
		fmt.Println("funcer outcount  :", funcert.outCount)
		fmt.Println("funcer dotdotdot :", funcert.dotdotdot())
		for i := range funcert.in() {
			fmt.Println("in", i, ":", reflect.Kind(funcert.in()[i].kind))
		}
		for i := range funcert.out() {
			fmt.Println("out", i, ":", reflect.Kind(funcert.out()[i].kind))
		}
		fmt.Println()
	}

	{
		// for a func defined in file, first assign is to a func variable
		var funcer2 interface{} = elliptic.GenerateKey
		efacer := (*eface)(unsafe.Pointer(&funcer2))
		funcert := (*functype)(unsafe.Pointer(efacer._type)) // extended to from *_type to *functype
		// type: func(elliptic.Curve, io.Reader) ([]uint8, *big.Int, *big.Int, error)
		fmt.Println("reflect type     :", reflect.TypeOf(funcer2).String())
		fmt.Println("funcer size      :", funcert.typ.size, efacer._type.size) // 8 for func
		fmt.Println("funcer hash      :", funcert.typ.hash)
		fmt.Println("funcer kind      :", funcert.typ.kind, reflect.Kind(funcert.typ.kind))
		fmt.Println("funcer str       :", funcert.typ.str)
		fmt.Println("funcer ptrToThis :", funcert.typ.ptrToThis)

		fmt.Println("funcer incount   :", funcert.inCount)
		fmt.Println("funcer outcount  :", funcert.outCount)
		fmt.Println("funcer dotdotdot :", funcert.dotdotdot())
		for i := range funcert.in() {
			fmt.Println("in", i, ":", reflect.Kind(funcert.in()[i].kind))
		}
		for i := range funcert.out() {
			fmt.Println("out", i, ":", reflect.Kind(funcert.out()[i].kind))
		}
		fmt.Println()
	}
}

func TestChans(t *testing.T) {
	var chaners = []interface{}{make(chan<- int), make(<-chan string), make(chan []int, 2)}
	for i := range chaners {
		efacer := (*eface)(unsafe.Pointer(&chaners[i]))
		chanert := (*chantype)(unsafe.Pointer(efacer._type))
		fmt.Println("reflect type      :", reflect.TypeOf(chaners[i]).String())
		fmt.Println("chanert size      :", chanert.typ.size, efacer._type.size) // 8 for chan, actually *runtime.hchan
		fmt.Println("chanert hash      :", chanert.typ.hash)
		fmt.Println("chanert kind      :", chanert.typ.kind, reflect.Kind(chanert.typ.kind))
		fmt.Println("chanert str       :", chanert.typ.str)
		fmt.Println("chanert ptrToThis :", chanert.typ.ptrToThis)

		// 1 for <-chan (receive only), 2 for chan<- (send only), 3 for chan (double direction)
		fmt.Println("chanert direction :", chanert.dir)
		fmt.Println("chanert elem size      :", chanert.elem.size)
		fmt.Println("chanert elem hash      :", chanert.elem.hash)
		fmt.Println("chanert elem kind      :", chanert.elem.kind, reflect.Kind(chanert.elem.kind))
		fmt.Println("chanert elem str       :", chanert.elem.str)
		fmt.Println("chanert elem ptrToThis :", chanert.elem.ptrToThis)
		fmt.Println()
	}
}

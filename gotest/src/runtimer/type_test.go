package runtimer

import (
	"fmt"
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

	{
		var funcer interface{} = func(a int, b string) (rune, uint64) {
			return 'r', 123456
		}
		efacer := (*eface)(unsafe.Pointer(&funcer))
		funcert := (*functype)(unsafe.Pointer(efacer._type))                   // extended to from *_type to *ptrtype
		fmt.Println("reflect type      :", reflect.TypeOf(funcer).String())    // *airport
		fmt.Println("funcer size      :", funcert.typ.size, efacer._type.size) // 8 for pointer
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
	}
}

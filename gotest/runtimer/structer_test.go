package runtimer

import (
	"fmt"
	"math/big"
	"reflect"
	"sync"
	"testing"
	"unsafe"
)

func TestEmptyStruct(t *testing.T) {
	// struct{} is a special type with size 0, all struct{}{} is the same reference of a constant
	fmt.Println(reflect.TypeOf(struct{}{})) // struct {}
	array := [3]struct{}{{}, {}, {}}
	ints := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(reflect.TypeOf(array), reflect.TypeOf(array).Size()) // [3] struct {}, 0
	fmt.Println(reflect.TypeOf(ints), reflect.TypeOf(ints).Size())   // [7]int, 56
	fmt.Println(unsafe.Pointer(&array[0]))                           // 0x66a320, all the same, indicate that all struct{}{} are the same reference
	fmt.Println(unsafe.Pointer(&array[1]))
	fmt.Println(unsafe.Pointer(&array[2]))
	a, b := struct{}{}, struct{}{}
	pa, pb := &a, &b
	fmt.Println(unsafe.Pointer(&a), reflect.TypeOf(pa))
	fmt.Println(unsafe.Pointer(&b), reflect.TypeOf(pb))
	fmt.Println()

	var ia interface{} = a
	typ := (*eface)(unsafe.Pointer(&ia))._type
	fmt.Println("typ size      :", typ.size) // 0
	fmt.Println("typ ptrdata   :", typ.ptrdata)
	fmt.Println("typ hash      :", typ.hash)
	fmt.Println("typ align     :", typ.align)
	fmt.Println("typ fieldalign:", typ.fieldalign)
	fmt.Println("typ kind      :", typ.kind, reflect.Kind(typ.kind))
	fmt.Println("typ str       :", typ.str)
	fmt.Println("typ ptrToThis :", typ.ptrToThis)
	fmt.Printf("eface data: %x\n", uintptr((*eface)(unsafe.Pointer(&ia)).data)) // 0x66a320, stll the same address
	fmt.Println()

	var ipa interface{} = pa
	typ = (*eface)(unsafe.Pointer(&ipa))._type
	fmt.Println("typ size      :", typ.size)
	fmt.Println("typ ptrdata   :", typ.ptrdata)
	fmt.Println("typ hash      :", typ.hash)
	fmt.Println("typ align     :", typ.align)
	fmt.Println("typ fieldalign:", typ.fieldalign)
	fmt.Println("typ kind      :", typ.kind, reflect.Kind(typ.kind))
	fmt.Println("typ str       :", typ.str)
	fmt.Println("typ ptrToThis :", typ.ptrToThis)
}

func TestEmptyStructValueMap(t *testing.T) {
	mapper := make(map[int64]struct{})
	var mapEface interface{} = mapper
	efacer := (*eface)(unsafe.Pointer(&mapEface))
	typ := efacer._type

	mt := (*maptype)(unsafe.Pointer(typ)) // *_type to *maptype
	fmt.Println("typ size      :", mt.typ.size)
	fmt.Println("typ kind      :", mt.typ.kind, reflect.Kind(mt.typ.kind))
	fmt.Println("typ ptrToThis :", mt.typ.ptrToThis)
	fmt.Println("indirectkey:", mt.indirectkey())
	fmt.Println("indirectvalue:", mt.indirectvalue())
	fmt.Println()

	// key size 8, value size 0
	// bucket size 80 bytes, tophash (8 bytes), 8 keys (64 bytes), 8 values (0 bytes), *overflow (8 bytes)
	fmt.Println("key size:", mt.keysize)
	fmt.Println("value size:", mt.valuesize)
	fmt.Println("bucket size:", mt.bucketsize)
	fmt.Println("flag:", mt.flags)
	fmt.Println()
}

func TestStructFields(t *testing.T) {
	type stct struct {
		a int
		b string
		c *big.Int
		d byte // 1 byte but will be aligned to 8 bytes
		f interface{}
	}
	var stcter interface{} = stct{}
	efacer := (*eface)(unsafe.Pointer(&stcter))
	stctt := (*structtype)(unsafe.Pointer(efacer._type)) // extend from *_type to *stcttype
	fmt.Println("reflect type    :", reflect.TypeOf(stcter).String())
	fmt.Println("stctt size      :", stctt.typ.size) // 8+16+8+8+16
	fmt.Println("stctt hash      :", stctt.typ.hash)
	fmt.Println("stctt kind      :", stctt.typ.kind, reflect.Kind(stctt.typ.kind))
	fmt.Println("stctt str       :", stctt.typ.str)
	fmt.Println("stctt ptrToThis :", stctt.typ.ptrToThis)

	fmt.Println("pkg path        :", stctt.pkgPath.name())
	for i := range stctt.fields {
		field := stctt.fields[i] // a structfield instance
		fmt.Println(i, "name:", field.name.name())
		fmt.Println(i, "offset:", field.offset())
		fmt.Println(i, "kind:", reflect.Kind(field.typ.kind))
	}
}

func TestUnnamedStruct(t *testing.T) {
	a := struct {
		int
		string
	}{10, "aaa"}
	fmt.Println(reflect.TypeOf(a)) // struct { int; string }
	fmt.Println(a.int, a.string)   // accessed by a.int, a.string

	b := struct {
		int
		string
	}{10, "aaa"}
	fmt.Println(reflect.TypeOf(b)) // struct { int; string }
	fmt.Println(b.int, b.string)

	b1 := struct {
		a int
		string
	}{10, "aaa"}
	fmt.Println(reflect.TypeOf(b1)) // struct { a int; string }

	b2 := struct {
		a int
		b string
	}{10, "aaa"}
	fmt.Println(reflect.TypeOf(b2)) // struct { a int; b string }

	// 'struct{int;string}' seems very like 'struct{int int;string string}', but they are not same type
	b3 := struct {
		int    int
		string string
	}{10, "aaa"}
	fmt.Println(b3.int, b3.string)
	fmt.Println(reflect.TypeOf(b3))

	var efa interface{} = a
	var efb interface{} = b
	var efb1 interface{} = b1
	var efb2 interface{} = b2
	var efb3 interface{} = b3
	ta := (*eface)(unsafe.Pointer(&efa))._type
	tb := (*eface)(unsafe.Pointer(&efb))._type
	tb1 := (*eface)(unsafe.Pointer(&efb1))._type
	tb2 := (*eface)(unsafe.Pointer(&efb2))._type
	tb3 := (*eface)(unsafe.Pointer(&efb3))._type
	fmt.Println(unsafe.Pointer(ta), ta.hash) // a.type == b.type != b1.type != b2.type != b3.type
	fmt.Println(unsafe.Pointer(tb), tb.hash) // struct{int;string} is not struct{a int;string}
	fmt.Println(unsafe.Pointer(tb1), tb1.hash)
	fmt.Println(unsafe.Pointer(tb2), tb2.hash)
	fmt.Println(unsafe.Pointer(tb3), tb3.hash)

	// c := struct {int;string;int}{}     // duplicated field not allowed
	c := struct { // it's like struct{ int int; string string; a int; Mutex *sync.Mutex }
		int
		string
		a int
		*sync.Mutex
	}{10, "aaa", 38, &sync.Mutex{}}
	fmt.Println(reflect.TypeOf(c)) // struct { int; string; a int; *sync.Mutex }
	fmt.Println(c.int, c.string, c.a, c.Mutex)
}

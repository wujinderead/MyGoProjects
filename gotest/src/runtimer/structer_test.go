package runtimer

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestEmptyStruct(t *testing.T) {
	// struct{} is a special type with size 0, all struct{}{} is the same object in constant pool
	fmt.Println(reflect.TypeOf(struct{}{})) // struct {}
	array := [3]struct{}{{}, {}, {}}
	ints := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(reflect.TypeOf(array), reflect.TypeOf(array).Size()) // [3] struct {}, 0
	fmt.Println(reflect.TypeOf(ints), reflect.TypeOf(ints).Size())   // [7]int, 56
	fmt.Println(unsafe.Pointer(&array[0]))                           // 0x66a320, all the same, indicate that struct{}{} is in the constant pool
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

package runtimer

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

type flyable interface {
	fly()
	speed() int
}
type bird struct {
	inter  int64
	uinter uint32
	name   string
}

func (bird bird) fly()       {}
func (bird bird) speed() int { return 88 }

type plane int

func (plane plane) fly()       {}
func (plane plane) speed() int { return 999 }

func TestInterfaceCopy(t *testing.T) {
	v1 := bird{8, 123, "aaa"}
	var face1, face3 flyable
	face1 = v1     // struct assigned to interface
	face2 := face1 // copy interface
	face3 = &v1    // pointer assigned to interface
	inspect := func(n *flyable, u *bird) {
		word := uintptr(unsafe.Pointer(n)) + uintptr(unsafe.Sizeof(&u))
		value := (**bird)(unsafe.Pointer(word))
		fmt.Printf("Addr face: %p Word Value: %p  Ptr Value: %v\n",
			n, *value, **value)
	}
	fmt.Printf("bird addr: %p\n", &v1)
	inspect(&face1, &v1)
	inspect(&face2, &v1)
	inspect(&face3, &v1)
	/*  example result:
	bird addr: 0xc00004e1c0
	Addr face: 0xc00004e1d0 Word Value: 0xc00004e1f0  Ptr Value: {aaa}
	Addr face: 0xc00004e200 Word Value: 0xc00004e1f0  Ptr Value: {aaa}
	Addr face: 0xc00004e1e0 Word Value: 0xc00004e1c0  Ptr Value: {aaa}  */
	/*  interface has 2 part, the type and the value.
	when bird struct assigned to face1, the type is bird, the value is the addr of copy of 'v1'
	when face2 copy face1, the value is the same as face1
	when bird pointer assigned to face3, the type is *bird, the value is the addr of 'v1' */
}

func TestIfaceEface(t *testing.T) {
	v1 := bird{8, 123, "aaa"}
	v2 := plane(777)
	var bird1 flyable = v1     // iface: value assigned to particular face
	var bird2 interface{} = v1 // eface: value assigned to empty face
	var bird3 interface{} = &v1
	var plane1 flyable = &v2     // iface: pointer assigned to particular face
	var plane2 interface{} = &v2 // eface: pointer assigned to empty face
	var plane3 interface{} = v2
	displayIface(unsafe.Pointer(&bird1))
	displayEface(unsafe.Pointer(&bird2))
	displayEface(unsafe.Pointer(&bird3))
	displayIface(unsafe.Pointer(&plane1))
	displayEface(unsafe.Pointer(&plane2))
	displayEface(unsafe.Pointer(&plane3))
}

func displayIface(p unsafe.Pointer) {
	ifacer := (*iface)(p)
	tab := ifacer.tab
	fmt.Println("itab inter pkgname: ", tab.inter.pkgpath.name(), tab.inter.pkgpath.nameLen(), tab.inter.pkgpath)
	fmt.Println("itab inter imethod[]:", tab.inter.mhdr) // interface methods
	fmt.Println("itab fun:", tab.fun[0], "0x"+strconv.FormatUint(uint64(tab.fun[0]), 16))
	fmt.Println("itab hash:", tab.hash, "0x"+strconv.FormatUint(uint64(tab.hash), 16))
	fmt.Println("itab inter typ size      :", tab.inter.typ.size)
	fmt.Println("itab inter typ ptrdata   :", tab.inter.typ.ptrdata)
	fmt.Println("itab inter typ hash      :", tab.inter.typ.hash)
	fmt.Println("itab inter typ align     :", tab.inter.typ.align)
	fmt.Println("itab inter typ fieldalign:", tab.inter.typ.fieldalign)
	fmt.Println("itab inter typ kind      :", tab.inter.typ.kind, reflect.Kind(tab.inter.typ.kind))
	fmt.Println("itab inter typ str       :", tab.inter.typ.str)
	fmt.Println("itab inter typ ptrToThis :", tab.inter.typ.ptrToThis)
	fmt.Println("itab _type size      :", tab._type.size)
	fmt.Println("itab _type ptrdata   :", tab._type.ptrdata)
	fmt.Println("itab _type hash      :", tab._type.hash)
	fmt.Println("itab _type align     :", tab._type.align)
	fmt.Println("itab _type fieldalign:", tab._type.fieldalign)
	fmt.Println("itab _type kind      :", tab._type.kind, reflect.Kind(tab._type.kind))
	fmt.Println("itab _type str       :", tab._type.str)
	fmt.Println("itab _type ptrToThis :", tab._type.ptrToThis)
	fmt.Println()
}

func displayEface(p unsafe.Pointer) {
	efacer := (*eface)(p)
	typ := efacer._type
	fmt.Println("typ size      :", typ.size)
	fmt.Println("typ ptrdata   :", typ.ptrdata)
	fmt.Println("typ hash      :", typ.hash)
	fmt.Println("typ align     :", typ.align)
	fmt.Println("typ fieldalign:", typ.fieldalign)
	fmt.Println("typ kind      :", typ.kind, reflect.Kind(typ.kind))
	fmt.Println("typ str       :", typ.str)
	fmt.Println("typ ptrToThis :", typ.ptrToThis)
	fmt.Println()
}

func TestEfaceData(t *testing.T) {
	v1 := bird{8, 123, "aaa"}
	v2 := bird{5, 666, "bbb"}
	var bird1 flyable = v1     // iface: value assigned to particular face
	var bird2 interface{} = v1 // eface: value assigned to empty face
	var bird3 interface{} = &v2
	displayEfaceData(unsafe.Pointer(&bird1))
	displayEfaceData(unsafe.Pointer(&bird2))
	displayEfaceData(unsafe.Pointer(&bird3))
	fmt.Printf("v1: %p, v2: %p\n", &v1, &v2)
	// value for interface bird1, bird2 is v1's copy (value assigned to interface)
	// value for interface bird3 is exactly v2 (pointer assigned to interface)
}

func displayEfaceData(p unsafe.Pointer) {
	// suppose on 64bit machine pointer length 8
	upp := (*unsafe.Pointer)(unsafe.Pointer(uintptr(p) + uintptr(8)))
	if (*eface)(p).data != *upp {
		panic("get interface data error!")
	}
	fmt.Println("eface.data (unsafe pointer): ", *upp)
	up := *upp // eface.data
	bp := (*bird)(up)
	fmt.Printf("bp: %p\n", bp)
	fmt.Println(bp.inter)
	fmt.Println(bp.uinter)
	fmt.Println(bp.name)
}

func TestTypeEqual(t *testing.T) {
	v1 := bird{8, 123, "aaa"}
	v2 := bird{5, 666, "bbb"}
	var bird1 flyable = v1
	var bird2 flyable = &v2
	var bird3 interface{} = v1
	var bird4 interface{} = &v2
	if1 := (*iface)(unsafe.Pointer(&bird1))
	fmt.Printf("if1 *type: %p, hash: %d\n", if1.tab._type, if1.tab.hash)
	fmt.Printf("if1 face &type: %p, hash: %d\n", &if1.tab.inter.typ, if1.tab.inter.typ.hash)
	if2 := (*iface)(unsafe.Pointer(&bird2))
	fmt.Printf("if2 *type: %p, hash: %d\n", if2.tab._type, if2.tab.hash)
	fmt.Printf("if2 face &type: %p, hash: %d\n", &if2.tab.inter.typ, if2.tab.inter.typ.hash)
	ef3 := (*eface)(unsafe.Pointer(&bird3))
	fmt.Printf("ef3 *type: %p, hash: %d\n", ef3._type, ef3._type.hash)
	ef4 := (*eface)(unsafe.Pointer(&bird4))
	fmt.Printf("ef4 *type: %p, hash: %d\n", ef4._type, ef4._type.hash)
}

func TestConvertInterface(t *testing.T) {
	var bird1 flyable = &bird{-2, 123, "ccc"}
	type iface1 struct {
		_type *_type
		data  unsafe.Pointer
	}
	type iface2 struct {
		typ  unsafe.Pointer
		data unsafe.Pointer
	}
	up1 := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&bird1)) + uintptr(8)))
	up2 := (*iface1)(unsafe.Pointer(&bird1)).data
	up3 := (*iface2)(unsafe.Pointer(&bird1)).data
	// pointer of fixed-sized array can index directly
	// e.g., a = &[3]int{1,2,3}, b = a[2]
	up4 := (*[2]unsafe.Pointer)(unsafe.Pointer(&bird1))[1]
	fmt.Println(up1, up2, up3, up4)
	fmt.Println(up1 == up2, up2 == up3, up3 == up4)
}

func TestInterfaceData(t *testing.T) {
	p1 := &bird{-2, 123, "ccc"}
	var bird1 interface{} = p1
	var bird2 interface{} = bird{12, 456, "dddd"}
	var bird3 interface{} = &p1
	fmt.Println(reflect.TypeOf(bird1), reflect.TypeOf(bird2), reflect.TypeOf(bird3))
	typ1 := (*eface)(unsafe.Pointer(&bird1))._type
	data1 := (*eface)(unsafe.Pointer(&bird1)).data
	typ2 := (*eface)(unsafe.Pointer(&bird2))._type
	data2 := (*eface)(unsafe.Pointer(&bird2)).data
	data3 := (*eface)(unsafe.Pointer(&bird3)).data
	fmt.Println("typ1 size:", typ1.size, "typ2 size:", typ2.size)
	fmt.Println((*bird)(data1).uinter)     // bird1 is *bird, eface.data is also *bird
	fmt.Println((*bird)(data2).uinter)     // bird2 is bird, eface.data is *bird
	fmt.Println((*(**bird)(data3)).uinter) // bird3 is **bird, eface.data is **bird
}

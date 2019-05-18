package runtimer

import (
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

// stringStruct in runtime/string.go
type stringStruct struct {
	str unsafe.Pointer
	len int
}

func TestStringStruct(t *testing.T) {
	str := "abcde"
	fmt.Println("ascii len:", len(str), ", bytes:", []byte(str))
	ps := (*stringStruct)(unsafe.Pointer(&str))
	fmt.Println("header len:", ps.len)
	fmt.Println("str addr:", strconv.FormatUint(uint64(uintptr(ps.str)), 16))
	for i := 0; i < ps.len; i++ {
		pb := (*byte)(unsafe.Pointer(uintptr(ps.str) + uintptr(i)))
		fmt.Print(*pb, ", ")
		// string created by a := "aaa" is created in the constant pool.
		// it is immutable, thus the 'str' pointer is also marked as immutable
		// it can be read but cannot be wrote. if write to it, panic rises.
		// *pb += 1
	}
	fmt.Println()
	byter := []byte(str)
	// a copy of the string bytes
	fmt.Printf("bytes: %p, %v\n", &byter[0], byter)
	fmt.Println()

	// same string reference the same underlying stringStruct
	str1 := "abcde"
	fmt.Println("ascii len:", len(str1), ", bytes:", []byte(str1))
	ps = (*stringStruct)(unsafe.Pointer(&str1))
	fmt.Println("header len:", ps.len)
	fmt.Println("str addr:", strconv.FormatUint(uint64(uintptr(ps.str)), 16))

	// []byte(string) makes a copy of the string bytes, in heap
	byter = []byte(str1)
	fmt.Printf("bytes: %p, %v\n", &byter[0], byter)
	fmt.Println()
}

func TestStringCreateFromBytes(t *testing.T) {
	str2 := "赵客缦胡缨吴钩霜雪明"
	ps := (*stringStruct)(unsafe.Pointer(&str2))
	fmt.Println("str2 header len:", ps.len)
	fmt.Println("str2 addr:", uintptr(ps.str))
	for i := 0; i < ps.len; i++ {
		pb := (*byte)(unsafe.Pointer(uintptr(ps.str) + uintptr(i)))
		fmt.Print(uintptr(unsafe.Pointer(pb)), ", ")
		// immutable
		// *pb += 1
	}
	fmt.Println()
	fmt.Println()

	byter := []byte{232, 181, 181, 229, 174, 162, 231, 188, 166,
		232, 131, 161, 231, 188, 168, 229, 144, 180,
		233, 146, 169, 233, 156, 156, 233, 155, 170, 230, 152, 142} // same content as str2
	for i := range byter {
		// print byter address
		fmt.Print(uintptr(unsafe.Pointer(&byter[i])), ", ")
	}
	fmt.Println()

	// string created during runtime, is allocated in heap
	str3 := string(byter)
	// this string is equal to the string created in constant pool
	fmt.Println("str2:", str2, ", str3:", str3, ", ==?", str3 == str2)
	ps = (*stringStruct)(unsafe.Pointer(&str3))
	fmt.Println("str3 header len:", ps.len)
	fmt.Println("str3 addr:", uintptr(ps.str)) // a copy of byter
	for i := 0; i < ps.len; i++ {
		pb := (*byte)(unsafe.Pointer(uintptr(ps.str) + uintptr(i)))
		fmt.Print(uintptr(unsafe.Pointer(pb)), ", ") // copy of byter, different addresses with byter
		// string created during heap can be modified
		if i%3 == 2 {
			*pb += 1
		}
	}
	fmt.Println()
	fmt.Println("new string:", str3) // string has been changed

	// the string length can also be changed
	ps.len -= 9
	fmt.Println("new len string:", str3)
	fmt.Println()
}

func TestStringFromStringHeader(t *testing.T) {
	byter := []byte("赵客缦胡缨吴钩霜雪明")
	fmt.Println("byter addr:", uintptr(unsafe.Pointer(&byter[0])))

	// string created by stringStruct
	strStrcut := &stringStruct{unsafe.Pointer(&byter[0]), len(byter)}
	strp := (*string)(unsafe.Pointer(strStrcut))

	// by this operation, a new copy of 'strStrcut' is created,
	// however, we just create a new header, underlying []byte is still 'byter'
	str := *strp
	fmt.Println("string:", str) // 赵客缦胡缨吴钩霜雪明
	ps := (*stringStruct)(unsafe.Pointer(&str))
	fmt.Println("str addr:", uintptr(ps.str)) // same as byter addr

	// modify length and change underlying bytes
	strStrcut.len = 12
	byter[2] += 1
	// *strp is also a new string header
	fmt.Println("new string:", *strp) // 赶客缦胡
}

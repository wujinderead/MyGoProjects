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

func TestStringAdd(t *testing.T) {
	str := "abcde"
	fmt.Println("len:", len(str), ", bytes:", []byte(str))
	ps := (*stringStruct)(unsafe.Pointer(&str))
	fmt.Println("header len:", ps.len)
	fmt.Println("str addr:", strconv.FormatUint(uint64(uintptr(ps.str)), 16))

	// same as str, as it can be determined during compiling
	str1 := "abc" + "de"
	fmt.Println("len:", len(str1), ", bytes:", []byte(str1))
	ps = (*stringStruct)(unsafe.Pointer(&str1))
	fmt.Println("header len:", ps.len)
	fmt.Println("str addr:", strconv.FormatUint(uint64(uintptr(ps.str)), 16)) // same as str

	// can not determined during compiling, it is allocated on heap during runtime
	str2 := "abc" + string([]byte{100, 101})
	fmt.Println("len:", len(str2), ", bytes:", []byte(str2))
	ps = (*stringStruct)(unsafe.Pointer(&str2))
	fmt.Println("header len:", ps.len)
	fmt.Println("str addr:", strconv.FormatUint(uint64(uintptr(ps.str)), 16)) // different from str
	fmt.Println("str==str2:", str == str2)                                    // equal
}

func TestStringTruncate(t *testing.T) {
	str := "赵客缦胡缨吴钩霜雪明"
	fmt.Println("len:", len(str)) // 30, length of utf8 bytes
	fmt.Println("bytes:", []byte(str))
	fmt.Println("runes:", []rune(str))
	ps := (*stringStruct)(unsafe.Pointer(&str))
	fmt.Println("header len:", ps.len)        // 30
	fmt.Println("str addr:", uintptr(ps.str)) // 5623127
	fmt.Println()

	// truncation are based on byte, not rune
	// fmt.Println(&str[0])   // cannot take address of str[0], to prevent write
	fmt.Println("str[0]:", str[0])     // 232, first bytes
	fmt.Println("str[0:4]:", str[0:6]) // 赵客, first 6 bytes

	// truncated string use the same underlying bytes as the original string, only length changed
	truc := str[0:6]
	ps = (*stringStruct)(unsafe.Pointer(&truc))
	fmt.Println("header len:", ps.len)        // 6
	fmt.Println("str addr:", uintptr(ps.str)) // 5623127
	fmt.Println()

	truc = str[6:12]
	ps = (*stringStruct)(unsafe.Pointer(&truc))
	fmt.Println("header len:", ps.len)        // 6
	fmt.Println("str addr:", uintptr(ps.str)) // 5623133
	fmt.Println()

	// string allocated on heap during runtime
	str2 := string([]byte("赵客缦胡缨吴钩霜雪明"))
	fmt.Println("len:", len(str2), ", bytes:", []byte(str2))
	ps = (*stringStruct)(unsafe.Pointer(&str2))
	fmt.Println("header len:", ps.len)         // 30
	fmt.Println("str2 addr:", uintptr(ps.str)) // 824633763136
	fmt.Println()

	// truncated string mechanism same as constant pool string
	trunc1 := str2[0:6]
	ps = (*stringStruct)(unsafe.Pointer(&trunc1))
	fmt.Println("header len:", ps.len)        // 30
	fmt.Println("str addr:", uintptr(ps.str)) // 824633763136
	fmt.Println()

	trunc2 := str2[6:12]
	ps = (*stringStruct)(unsafe.Pointer(&trunc2))
	fmt.Println("header len:", ps.len)        // 30
	fmt.Println("str addr:", uintptr(ps.str)) // 824633763142
	fmt.Println()

	// only that string allocated on heap can be modified
	ps = (*stringStruct)(unsafe.Pointer(&str2))
	b := (*byte)(unsafe.Pointer(uintptr(ps.str) + uintptr(2))) // modify str2[2]
	*b += 1
	ps = (*stringStruct)(unsafe.Pointer(&str2))
	b = (*byte)(unsafe.Pointer(uintptr(ps.str) + uintptr(8))) // modify str2[8]
	*b += 1
	fmt.Println("str2:", str2)     // 赶客缧胡缨吴钩霜雪明
	fmt.Println("trunc1:", trunc1) // 赶客
	fmt.Println("trunc2:", trunc2) // 缧胡
}

func TestEmptyString(t *testing.T) {
	str := ""
	fmt.Println("len:", len(str), ", bytes:", []byte(str))
	ps := (*stringStruct)(unsafe.Pointer(&str))
	fmt.Println("header len:", ps.len)        // 0
	fmt.Println("str addr:", uintptr(ps.str)) // 0

	// same as str, as it can be determined during compiling
	str1 := string([]byte{})
	fmt.Println("len:", len(str1), ", bytes:", []byte(str1))
	ps = (*stringStruct)(unsafe.Pointer(&str1))
	fmt.Println("header len:", ps.len)        // 0
	fmt.Println("str addr:", uintptr(ps.str)) // 0

	// for v=str, addr==0
	// for v="", addr!=0
	for _, v := range []string{str, ""} {
		fmt.Println("len:", len(v), ", bytes:", []byte(v))
		ps = (*stringStruct)(unsafe.Pointer(&v))
		fmt.Println("header len:", ps.len)
		fmt.Println("str addr:", uintptr(ps.str))
	}
}

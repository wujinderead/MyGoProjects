package stdlib

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

var (
	xia_ke_xing = "赵客缦胡缨，吴钩霜雪明，银鞍照白马，飒沓如流星。" +
		"十步杀一人，千里不留行，事了拂衣去，深藏身与名。" +
		"闲过信陵饮，脱剑膝前横，将炙啖朱亥，持觞劝侯嬴。" +
		"三杯吐然诺，五岳倒为轻，眼花耳热后，意气素霓生。" +
		"救赵挥金槌，邯郸先震惊，千秋二壮士，烜赫大梁城。" +
		"纵死侠骨香，不惭世上英，谁能书阁下，白首太玄经。"

	ascii = `1234567890-=qwertyuiop[]\asdfghjkl;'zxcvbnm,./ZXCVBNM<>?ASDFGHJKL:"QWERTYUIOP{}|!@#$%^&*()_+`

	single = "a"
)

func Test_align1(t *testing.T) {
	func() {
		type ds struct {
			a [5]uint8
			b [2]uint8
		}
		s := ds{[5]uint8{1, 2, 3, 4, 5}, [2]uint8{1, 2}}
		s1 := &ds{[5]uint8{45, 6, 255, 254, 5}, [2]uint8{123, 2}}
		p := new(ds)
		p.a = [5]uint8{45, 6, 0, 254, 53}
		p.b = [2]uint8{4, 63}
		fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(s1), unsafe.Sizeof(p))
	}()
	func() {
		type ds struct {
			a [5]uint8
			b [2]uint
		}
		s := ds{[5]uint8{1, 2, 3, 4, 5}, [2]uint{1, 2}}
		s1 := &ds{[5]uint8{45, 6, 255, 254, 5}, [2]uint{123, 2}}
		p := new(ds)
		p.a = [5]uint8{45, 6, 0, 254, 53}
		p.b = [2]uint{4, 63}
		fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(s1), unsafe.Sizeof(p))
	}()
	func() {
		type ds struct {
			a [7]uint8
			b [2]uint
		}
		s := ds{[7]uint8{1, 2, 3, 4, 5}, [2]uint{1, 2}}
		s1 := &ds{[7]uint8{45, 6, 255, 254, 5}, [2]uint{123, 2}}
		p := new(ds)
		p.a = [7]uint8{45, 6, 0, 254, 53}
		p.b = [2]uint{4, 63}
		fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(s1), unsafe.Sizeof(p))
	}()
	func() {
		type ds struct {
			a [9]uint8
			b [2]uint
		}
		s := ds{[9]uint8{1, 2, 3, 4, 5}, [2]uint{1, 2}}
		s1 := &ds{[9]uint8{45, 6, 255, 254, 5}, [2]uint{123, 2}}
		p := new(ds)
		p.a = [9]uint8{45, 6, 0, 254, 53}
		p.b = [2]uint{4, 63}
		fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(s1), unsafe.Sizeof(p))
	}()
}

func Test_align2(t *testing.T) {
	func() {
		type ds struct {
			c string
			b [2]uint8
		}
		s := ds{"test", [2]uint8{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
	func() {
		type ds struct {
			c string
			b [2]uint
		}
		s := ds{"test123", [2]uint{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
	func() {
		type ds struct {
			c string
			b [2]uint
		}
		s := ds{"test", [2]uint{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
	func() {
		type ds struct {
			c string
			b [2]uint8
		}
		s := ds{"test123", [2]uint8{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
	func() {
		type ds struct {
			c *string
			b [2]uint
		}
		a := "test"
		s := ds{&a, [2]uint{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
	func() {
		type ds struct {
			c *string
			b [2]uint8
		}
		a := "test123"
		s := ds{&a, [2]uint8{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
}

// on this darwin 64-bit machine
// the unsafe.Size() of any pointer is 8 (64-bit)
// the unsafe.Size() of any string is 16 (128-bit)
// the unsafe.Size() of any array reference is 24 (192-bit)
func testString(t *testing.T, str string) {
	runes := []rune(str)
	point := &str
	byter := []byte(str)
	fmt.Println(unsafe.Sizeof(str))
	fmt.Println(unsafe.Sizeof(runes))
	fmt.Println(unsafe.Sizeof(point))
	fmt.Println(unsafe.Sizeof(byter))

	fmt.Printf("%d %d %d\n", len(str), len(byter), len(runes))
}

func TestInt8(t *testing.T) {
	var a int8 = 53
	var p = &a
	var arr1 = []int8{12, -7, 87}
	var arr2 = []int8{12, -7, 87, 12, -7, 87, 12, -7, 87, 12, -7, 87}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(arr1))
	fmt.Println(unsafe.Sizeof(arr2))
}

func TestInt16(t *testing.T) {
	var a int16 = 53
	var p = &a
	var arr1 = []int16{12, -7, 87}
	var arr2 = []int16{12, -7, 87, 12, -7, 87, 12, -7, 87, 12, -7, 87}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(arr1))
	fmt.Println(unsafe.Sizeof(arr2))
}

func TestFloat32(t *testing.T) {
	var a float32 = 53
	var p = &a
	var arr1 = []float32{12.7, -7.9, 8.78}
	var arr2 = []float32{12, -7, 87, 12, -7, 87, 12, -7, 87, 12, -7, 87}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(arr1))
	fmt.Println(unsafe.Sizeof(arr2))
}

func TestFloat64(t *testing.T) {
	var a float64 = 53
	var p = &a
	var arr1 = []float64{12.7, -7.9, 8.78}
	var arr2 = []float64{12, -7, 87, 12, -7, 87, 12, -7, 87, 12, -7, 87}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(arr1))
	fmt.Println(unsafe.Sizeof(arr2))
}

func TestComplex64(t *testing.T) {
	var a complex64 = complex(4.5, 7.8)
	var p = &a
	var arr1 = []complex64{a, a, a}
	var arr2 = []complex64{a, a, a, a, a, a, a, a, a}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(arr1))
	fmt.Println(unsafe.Sizeof(arr2))
}

func TestComplex128(t *testing.T) {
	var a complex128 = complex(4.5, 7.8)
	var p = &a
	var arr1 = []complex128{a, a, a}
	var arr2 = []complex128{a, a, a, a, a, a, a, a, a}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(arr1))
	fmt.Println(unsafe.Sizeof(arr2))
}

func TestString(t *testing.T) {
	testString(t, xia_ke_xing)
	testString(t, ascii)
	testString(t, single)
}

func TestStringModify(t *testing.T) {
	runes := []rune(xia_ke_xing)
	//point := &xia_ke_xing
	byter := []byte(xia_ke_xing)

	// can we modify string through modifying the underline array?
	fmt.Println(hex.EncodeToString(byter))
	for _, v := range runes {
		fmt.Print(v, " ", string(v), " ")
		fmt.Printf("%x\n", v)
	}
}

func TestBytesUnsafe(t *testing.T) {
	byter := []byte(xia_ke_xing)
	base_p := &byter[0]
	for i := 0; i < len(byter); i++ {
		byte_p := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(base_p)) + uintptr(i)))
		if i%3 == 2 {
			*byte_p = *byte_p + 1
		}
	}
	fmt.Printf("%x\n", byter)
	fmt.Println(string(byter))
}

type arrayStr struct {
	a string
	b []int8
	c []float64
}

// in this machine, the size of string is always 16, the size of any array is always 24
func TestArraySize(t *testing.T) {
	a := arrayStr{
		"hahahahahahhahahah",
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]float64{0.1, 0.2, 0.432, 56.123, 456.4, 7.8, 8.4, 13.8},
	}
	b := arrayStr{
		"a",
		[]int8{1},
		[]float64{0.1},
	}
	c := arrayStr{
		"",
		[]int8{},
		[]float64{},
	}
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(a.a), unsafe.Sizeof(a.b), unsafe.Sizeof(a.c))
	fmt.Println(unsafe.Sizeof(b), unsafe.Sizeof(b.a), unsafe.Sizeof(b.b), unsafe.Sizeof(b.c))
	fmt.Println(unsafe.Sizeof(c), unsafe.Sizeof(c.a), unsafe.Sizeof(c.b), unsafe.Sizeof(c.c))
}

func TestArrayConvert(t *testing.T) {
	type ULL uint64
	a := []uint64{0x123456, 0x123465, 0x123489, 0x1234560, 0x123456d, 0x123456a, 0x123456f,}
	fmt.Println(unsafe.Sizeof(a))
	pa := &a
	fmt.Println(reflect.TypeOf(pa))
	fmt.Println(reflect.TypeOf(a))

	pull := (*[]ULL)(unsafe.Pointer(pa))
	ull := *pull
	fmt.Println(reflect.TypeOf(pull))
	fmt.Println(reflect.TypeOf(ull))
	fmt.Println(len(ull))
	for i, _ := range ull {
		fmt.Println(ull[i])
	}
	fmt.Println(ull)

	pul := (*[]uint32)(unsafe.Pointer(pa))
	ul := *pul
	fmt.Println(reflect.TypeOf(pul))
	fmt.Println(reflect.TypeOf(ul))
	fmt.Println(len(ul))
	for i, _ := range ul {
		fmt.Println(ul[i])
	}
	fmt.Println(ul)
}

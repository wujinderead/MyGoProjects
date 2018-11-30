package util

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
	"golang.org/x/text/encoding/unicode"
	"strconv"
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

// rune is int32 actually
// the string in golang is underlyingly stored as UTF-8
// when a utf-8 character is converted to rune, the rune is the unicode of the character
func TestRuneUtf8(t *testing.T) {
	var h string = "赵跷"                       // actual chinese
	var a string = "\xe8\xb5\xb5\xe8\xb7\xb7" // bytes representation
	var aa string = "\u8d75\u8df7"            // unicode representation
	var b0 rune = []rune(a)[0]                // 8d75
	var b1 rune = []rune(a)[1]                // 8df7
	fmt.Printf("%s %x\n", h, []byte(h))
	fmt.Printf("%s %x\n", a, []byte(a))
	fmt.Printf("%s %x\n", aa, []byte(aa))
	fmt.Printf("%d %x %s\n", b0, b0, string(b0))
	fmt.Println(unsafe.Sizeof(b0))
	fmt.Printf("%d %x %s\n", b1, b1, string(b1))

	// can not modify a string by converting it to []rune, string is const
	runes := []rune(h)
	runes[0], runes[1] = runes[1], runes[0]
	fmt.Println(h, " ", string(runes))
	runes[1] = []rune("\u8df4")[0]
	fmt.Println(h, " ", string(runes))

	// can not modify a string by converting it to []byte, string is const
	byter := []byte(h)
	byter[2] = 0xa4
	fmt.Println(h, " ", string(byter))
}

// why the unsafe.Sizeof(string) is 16, what exactly is the meaning of the 16-bytes
func TestStringUnsafe(t *testing.T) {
	var h string = "赵跷" // actual chinese
	//var a string = "\xe8\xb5\xb5\xe8\xb7\xb7"  // bytes representation
	var a string = "\xe8\xb5\xb5\xe8\xb7\xb7" // bytes representation
	//var aa string = "\u8d75\u8df7"  // unicode representation

	p := &h
	fmt.Println(p, " ", reflect.TypeOf(p))

	for i := 0; i < int(unsafe.Sizeof(h)); i++ {
		byte1 := *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)))
		fmt.Printf("%x\n", byte1)
	}

	fmt.Println(a, " ", &a, " ", reflect.TypeOf(&a))
	for i := 0; i < int(unsafe.Sizeof(xia_ke_xing)); i++ {
		byte1 := *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&xia_ke_xing)) + uintptr(i)))
		fmt.Printf("%x\n", byte1)
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

// in this mahine, the size of string is always 16, the size of any array is always 24
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

func TestUniCode(t *testing.T) {
	a := []rune{'\u2070', '\u00b9', '\u00b2', '\u00b3', '\u2074', '\u2075', '\u2076', '\u2077', '\u2078', '\u2079'}
	fmt.Println(a)
	for _, r := range a {
		fmt.Println("a" + string(r))
	}
	/*
	[]byte("😀☯商ñA") return utf8 byte[]     '0x f09f9880e298afe59586c3b141'
	[]rune("😀☯商ñA") return unicode int32[] {128512, 9775, 21830, 241, 65}
	   +------+---------------------+------------+------------+
	   | char |   unicode           |  utf8      |  utf16     |
	   +------+---------------------+------------+------------+
	   | 😀   |   128512 (U+1f660)  |  f09f9880  |  d83dde00  |
	   | ☯    |   9775   (U+262F)   |  e298af    |  262f      |
	   | 商   |   21830  (U+5546)   |  e59586    |  5546      |
	   | ñ    |   241    (U+00F1)   |  c3b1      |  00f1      |
	   | A    |   65     (U+0041)   |  41        |  0041      |
	   +------+---------------------+------------+------------+ */
	for _, a := range []rune{128512, 9775, 21830, 241, 65} {
		s := string(a)
		fmt.Println(s, hex.EncodeToString([]byte(s)))
	}

	// len(string) is the underlying utf8 bytes len
	b := "\xf0\x9f\x98\x80\xe2\x98\xaf\xe5\x95\x86\xc3\xb1\x41"  // "𐌡໔加ñA"
	// when str is transferred to []rune, the number is the number of characters
	c := []rune(b)
	fmt.Println(b, len(b), len(c)) // 13, 5
	for _, ch := range b {         // range string is to range []rune
		fmt.Println(ch, strconv.FormatInt(int64(ch), 16))
	}

	utf16 := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	byter, _ := utf16.NewEncoder().Bytes([]byte(b))
	fmt.Println(hex.EncodeToString(byter))
}

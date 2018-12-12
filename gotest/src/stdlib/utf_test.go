package stdlib

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

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

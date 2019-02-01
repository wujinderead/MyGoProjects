package stdlib

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"reflect"
	"strconv"
	"testing"
	"unicode/utf16"
	"unicode/utf8"
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

// test superscript
func TestUpperMark(t *testing.T) {
	a := []rune{'\u2070', '\u00b9', '\u00b2', '\u00b3', '\u2074', '\u2075', '\u2076', '\u2077', '\u2078', '\u2079'}
	fmt.Println(a)
	for _, r := range a {
		fmt.Println("a" + string(r))
	}
}

func TestUniCode(t *testing.T) {
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

	// "😀☯商ñA" in utf16: d83dde00 262f 5546 00f1 0041
	utf16E := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	byter, _ := utf16E.NewEncoder().Bytes([]byte(b))
	fmt.Println(hex.EncodeToString(byter))
}

func TestUnicode_Utf8(t *testing.T) {
	// test whether []byte, string, or rune is valid utf8-encoded
	byter := []byte("😀☯商ñA😀")
	fmt.Println(byter, utf8.Valid(byter))
	fmt.Println(string(byter), utf8.ValidString(string(byter)))
	byter[1] = 143
	byter[5] = 66
	fmt.Println(byter, utf8.Valid(byter))
	fmt.Println(string(byter), utf8.ValidString(string(byter)))
	fmt.Println()

	valid := 'a'
	invalid := rune(0xfffffee)
	fmt.Println(valid, utf8.ValidRune(valid))
	fmt.Println(invalid, utf8.ValidRune(invalid))
	fmt.Println()

	// decode a utf8 rune at the at the beginning or ending of string or []byte
	rune, size := utf8.DecodeLastRune([]byte{41, 226, 152, 175})
	fmt.Println(rune, size, string(rune))
	rune, size = utf8.DecodeLastRune([]byte{41, 226, 66, 175})
	fmt.Println(rune, rune==utf8.RuneError, size, string(rune))

	rune, size = utf8.DecodeLastRuneInString(string([]byte{41, 226, 152, 175}))
	fmt.Println(rune, size, string(rune))
	rune, size = utf8.DecodeLastRuneInString(string([]byte{41, 226, 66, 175}))
	fmt.Println(rune, rune==utf8.RuneError, size, string(rune))

	rune, size = utf8.DecodeRune([]byte{226, 152, 175, 68})
	fmt.Println(rune, size, string(rune))
	rune, size = utf8.DecodeRune([]byte{226, 66, 175, 68})
	fmt.Println(rune, rune==utf8.RuneError, size, string(rune))

	rune, size = utf8.DecodeRuneInString(string([]byte{226, 152, 175, 68}))
	fmt.Println(rune, size, string(rune))
	rune, size = utf8.DecodeRuneInString(string([]byte{226, 66, 175, 68}))
	fmt.Println(rune, rune==utf8.RuneError, size, string(rune))
	fmt.Println()

	// write a rune in utf8 encoding in bytes
	// f09f9880 00 e59586 0000
	byter = make([]byte, 10)
	fmt.Println(utf8.EncodeRune(byter, '😀'))
	fmt.Println(utf8.EncodeRune(byter[5:], '商'))
	fmt.Println(hex.EncodeToString(byter))
	fmt.Println()

	// count rune number in string or bytes
	fmt.Println(hex.EncodeToString([]byte("😀☯商ñA")), utf8.RuneCount([]byte("😀☯商ñA")))
	byter = []byte{240, 159, 152, 128, 226, 152, 175, 229, 149, 134, 195, 177, 65, 240, 159, 152, 128}
	fmt.Println(byter, utf8.RuneCountInString(string(byter)))
	fmt.Println()

	// rune len in utf8 encoding
	fmt.Println('😀', utf8.RuneLen('😀'))
	fmt.Println('商', utf8.RuneLen('商'))
	fmt.Println('ñ', utf8.RuneLen('ñ'))
	fmt.Println()
}


func TestUnicode_Utf16(t *testing.T) {
	/*
	   +------+---------------------+--------------+--------------+
	   | char |   unicode           |  utf8        |  utf16       |
	   +------+---------------------+--------------+--------------+
	   |      | rune (int32) in go  | string in go | char in java |
	   +------+---------------------+------------+----------------+
	   | 😀   |   128512 (U+1f660)  |  f09f9880    |  d83dde00    |
	   | ☯    |   9775   (U+262F)   |  e298af      |  262f        |
	   | 商   |   21830  (U+5546)   |  e59586      |  5546        |
	   | ñ    |   241    (U+00F1)   |  c3b1        |  00f1        |
	   | A    |   65     (U+0041)   |  41          |  0041        |
	   +------+---------------------+--------------+--------------+       */
	encoded := utf16.Encode([]rune{'A', 'ñ', '商', '☯', '😀'})
	fmt.Println("A: ", strconv.FormatUint(uint64(encoded[0]), 16))
	fmt.Println("ñ: ", strconv.FormatUint(uint64(encoded[1]), 16))
	fmt.Println("商: ", strconv.FormatUint(uint64(encoded[2]), 16))
	fmt.Println("☯: ", strconv.FormatUint(uint64(encoded[3]), 16))
	// 😀 exceeds 16-bit and need two uint16 to represent
	fmt.Println("😀: ", strconv.FormatUint(uint64(encoded[4]), 16), strconv.FormatUint(uint64(encoded[5]), 16))

	decoded := utf16.Decode(encoded)
	fmt.Println("decoded: ", decoded)
}

func TestExtend(t *testing.T) {
	// simplified chinese
	gb18030 := simplifiedchinese.GB18030.NewEncoder()
	gbk := simplifiedchinese.GBK.NewDecoder()
	gb2312 := simplifiedchinese.HZGB2312.NewEncoder()
	utf8str := "赵客缦胡缨，吴钩霜雪明。"
	utf8bytes := []byte(utf8str)
	fmt.Println("utf8 :", hex.EncodeToString(utf8bytes))

	var byter []byte
	var str string
	byter, _ = gb18030.Bytes(utf8bytes)
	fmt.Println("18030:", hex.EncodeToString(byter))
	str, _ = gb18030.String(utf8str)
	fmt.Println("18030:", str)

	byter, _ = gbk.Bytes(utf8bytes)
	fmt.Println("gbk  :", hex.EncodeToString(byter))
	str, _ = gbk.String(utf8str)
	fmt.Println("gbk  :", str)

	byter, _ = gb2312.Bytes(utf8bytes)
	fmt.Println("2312 :", hex.EncodeToString(byter))
	str, _ = gb2312.String(utf8str)
	fmt.Println("2312 :", str)
	fmt.Println()

	// traditional chinese
	utf8str = "愛爾蘭政府網站繁體中文"
	utf8bytes = []byte(utf8str)
	fmt.Println("utf8 :", hex.EncodeToString(utf8bytes))
	big5 := traditionalchinese.Big5.NewEncoder()
	byter, _ = big5.Bytes(utf8bytes)
	fmt.Println("big5 :", hex.EncodeToString(byter))
	str, _ = big5.String(utf8str)
	fmt.Println("big5 :", str)
}
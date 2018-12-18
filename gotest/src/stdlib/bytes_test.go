package stdlib

import (
	"bytes"
	"fmt"
	"testing"
	"unicode"
)

func TestCompare(t *testing.T) {
	// Interpret Compare's result by comparing it to zero.
	var a []byte = []byte("seafood")
	var b []byte = []byte("test")
	var c []byte = []byte{0x74, 0x65, 0x73, 0x74}
	var d []byte = []byte{116, 101, 115, 116}
	var e []byte = []byte{'t', 'e', 's', 't'}
	var uni1 []byte = []byte("a我的天哪")
	var uni2 []byte = []byte("a届かない恋")
	var uni3 []byte = []byte("aáéñ")
	t.Log("seafood test: ", bytes.Compare(a, b))
	t.Log("test {0x74, 0x65, 0x73, 0x74}", bytes.Compare(b, c))
	t.Log("test: {116, 101, 115, 116}", bytes.Compare(b, d))
	t.Log("test: {'t', 'e', 's', 't'}", bytes.Compare(b, e))
	t.Log("我的天哪 to bytes", uni1)
	t.Log("届かない恋 to bytes", uni2)
	t.Log("áéñ to bytes", uni3)
}

func TestContain(t *testing.T) {
	t.Log(bytes.Contains([]byte("seafood"), []byte("foo")))
	t.Log(bytes.Contains([]byte("seafood"), []byte("bar")))
	t.Log(bytes.Contains([]byte("seafood"), []byte("")))
	t.Log(bytes.Contains([]byte(""), []byte("")))
}

func TestCount(t *testing.T) {
	t.Log(bytes.Count([]byte("cheese"), []byte("e")))
	t.Log(bytes.Count([]byte("cheeese"), []byte("ee"))) // counts the number of non-overlapping instances
	t.Log(bytes.Count([]byte("cheeeese"), []byte("ee")))
	t.Log(bytes.Count([]byte("five"), []byte(""))) // before & after each rune
}

func TestFold(t *testing.T) {
	t.Logf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	t.Logf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))
	f1 := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	t.Logf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f1))
}

func TestHasPrefix(t *testing.T) {
	t.Log(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
	t.Log(bytes.HasPrefix([]byte("Gopher"), []byte("g")))
	t.Log(bytes.HasPrefix([]byte("Gopher"), []byte("")))
}

func TestHasSuffix(t *testing.T) {
	t.Log(bytes.HasSuffix([]byte("Amigo"), []byte("go")))
	t.Log(bytes.HasSuffix([]byte("Amigo"), []byte("O")))
	t.Log(bytes.HasSuffix([]byte("Amigo"), []byte("Ami")))
	t.Log(bytes.HasSuffix([]byte("Amigo"), []byte("")))
}

func TestAll(t *testing.T) {
	t.Log(bytes.Index([]byte("chicken"), []byte("ken")))
	t.Log(bytes.Index([]byte("chicken"), []byte("dmr")))
	t.Log(bytes.IndexAny([]byte("chicken"), "aeiouy"))
	t.Log(bytes.IndexAny([]byte("crwth"), "aeiouy"))
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	t.Log(bytes.IndexFunc([]byte("Hello, 世界"), f))
	t.Log(bytes.IndexFunc([]byte("Hello, world"), f))
	t.Log(bytes.IndexRune([]byte("chicken"), 'k'))
	t.Log(bytes.IndexRune([]byte("chicken"), 'd'))
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	t.Logf("%s", bytes.Join(s, []byte(", ")))
	t.Log(bytes.Index([]byte("go gopher"), []byte("go")))
	t.Log(bytes.LastIndex([]byte("go gopher"), []byte("go")))
	t.Log(bytes.LastIndex([]byte("go gopher"), []byte("rodent")))
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	t.Logf("%s", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))
	t.Logf("ba%s", bytes.Repeat([]byte("na"), 2))
	t.Logf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	t.Logf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))
	t.Logf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	t.Logf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	t.Logf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	t.Logf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
}

func TestRune(t *testing.T) {
	fmt.Println(bytes.IndexRune([]byte("aa我的世界"), '我'))
	aa := bytes.Runes([]byte("aa我的世界"))
	for _, c := range aa {
		fmt.Printf("%c ", c)
	}
	for _, d:= range "aa对我的我" {
		fmt.Printf("%c ", d)
	}
	for _, e := range "aaa" {
		fmt.Printf("%c ", e)
	}
}

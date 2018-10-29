package test

import (
	"golang.org/x/crypto/ripemd160"
	"hash"
	"testing"
	"encoding/hex"
	"bytes"
)

type mdTest struct {
	out string
	in  string
}

var vectors = [...]mdTest{
	{"9c1185a5c5e9fc54612808977ee8f548b2258d31", ""},
	{"0bdc9d2d256b3ee9daae347be6f4dc835a467ffe", "a"},
	{"8eb208f7e05d987a9b044a8e98c6b087f15a0bfc", "abc"},
	{"5d0689ef49d2fae572b881b123a85ffa21595f36", "message digest"},
	{"f71c27109c692c1b56bbdceb5b9d2865b3708dbc", "abcdefghijklmnopqrstuvwxyz"},
	{"12a053384a9c0c88e405a06c27dcf49ada62eb2b", "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq"},
	{"b0e20b6e3116640286ed3a87a5713079b21f5189", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"},
	{"9b752e45573d4b39f4dbd3323cab82bf63326bfb", "12345678901234567890123456789012345678901234567890123456789012345678901234567890"},
}

func TestRipemd160(t *testing.T) {
	var hasher hash.Hash = ripemd160.New()
	for _, vec := range vectors {
		hasher.Write([]byte(vec.in))
		hashed := hasher.Sum(nil)
		t.Logf("hashed: %x, in: \"%s\"", hashed, vec.in)
		expected, _ := hex.DecodeString(vec.out)
		if bytes.Compare(hashed, expected) != 0 {
			t.Errorf("check err for \"%s\".", vec.in)
		}
		hasher.Reset()
	}

}

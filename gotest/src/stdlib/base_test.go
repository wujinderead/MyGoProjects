package stdlib

import (
	"encoding/base64"
	"testing"
)

type testpair struct {
	decoded, encoded string
}

var pairs = []testpair{
	// RFC 3548 examples
	{"\x14\xfb\x9c\x03\xd9\x7e", "FPucA9l+"},
	{"\x14\xfb\x9c\x03\xd9", "FPucA9k="},
	{"\x14\xfb\x9c\x03", "FPucAw=="},

	// RFC 4648 examples
	{"", ""},
	{"f", "Zg=="},
	{"fo", "Zm8="},
	{"foo", "Zm9v"},
	{"foob", "Zm9vYg=="},
	{"fooba", "Zm9vYmE="},
	{"foobar", "Zm9vYmFy"},
	{"foobar\n", "Zm9vYmFyCg=="},

	// Wikipedia examples
	{"sure.", "c3VyZS4="},
	{"sure", "c3VyZQ=="},
	{"sur", "c3Vy"},
	{"su", "c3U="},
	{"leasure.", "bGVhc3VyZS4="},
	{"easure.", "ZWFzdXJlLg=="},
	{"asure.", "YXN1cmUu"},
	{"sure.", "c3VyZS4="},
}

func testBase64Encode(tested, expected string , t *testing.T) {
	var stdEncoding = base64.StdEncoding
	src := []byte(tested)
	enLen := stdEncoding.EncodedLen(len(src))
	t.Log("encode len: ", enLen)
	dst := make([]byte, enLen)
	stdEncoding.Encode(dst, src)
	t.Logf("src: %s, dst: %s", src, dst)
	if string(dst) != expected {
		t.Error("encode err.")
	}
}

func TestBase64Encode(t *testing.T) {
	for _, pair := range pairs {
		testBase64Encode(pair.decoded, pair.encoded, t)
	}
}

func TestBase64EncodeToString(t *testing.T) {
	var stdEncoding = base64.StdEncoding
	for _, pair := range pairs {
		result := stdEncoding.EncodeToString([]byte(pair.decoded))
		t.Logf("src: %s, dst: %s", []byte(pair.decoded), result)
		if (result != pair.encoded) {
			t.Error("encode err.")
		}
	}
}

func testBase64Decode(tested, expected string , t *testing.T) {
	var stdEncoding = base64.StdEncoding
	src := []byte(tested)
	deLen := stdEncoding.DecodedLen(len(src))
	t.Log("decode len: ", deLen)
	dst := make([]byte, deLen)
	n, _ := stdEncoding.Decode(dst, src)
	t.Logf("src: %s, dst: %s", src, dst[:n])
	if string(dst[:n]) != expected {
		t.Error("decode err.")
	}
}

func TestBase64Decode(t *testing.T) {
	for _, pair := range pairs {
		testBase64Decode(pair.encoded, pair.decoded, t)
	}
}

func TestBase64DecodeToString(t *testing.T) {
	var stdEncoding = base64.StdEncoding
	for _, pair := range pairs {
		result, err := stdEncoding.DecodeString(pair.encoded)
		if err != nil {
			t.Error("decode err", err)
		}
		t.Logf("src: %s, dst: %s", pair.encoded, result)
		if string(result) != pair.decoded {
			t.Error("encode err.")
		}
	}
}
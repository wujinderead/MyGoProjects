package v4

import "testing"

func TestBase58Encode(t *testing.T) {
	t.Log(string(Base58Encode([]byte{})))
	t.Log(string(Base58Encode([]byte("test"))))
	t.Log(string(Base58Encode([]byte("the quick fox jump over the lazy dog."))))
}

func TestBase58Decode(t *testing.T) {
	t.Log(string(Base58Decode([]byte{})))
	t.Log(string(Base58Decode([]byte("3yZe7d"))))
	t.Log(string(Base58Decode([]byte("4uHL6VBwGzMF8hxrcbvingBhRfh8xSnDwbdKTjnBo334AnnKwvM"))))
}

func TestReverseBytes(t *testing.T) {
	a := []byte{}
	ReverseBytes(a)
	t.Log(string(a))
	a = []byte("a")
	ReverseBytes(a)
	t.Log(string(a))
	a = []byte("ab")
	ReverseBytes(a)
	t.Log(string(a))
	a = []byte("abc")
	ReverseBytes(a)
	t.Log(string(a))
}

package factorization

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	"testing"
)

func TestNewPrivateKey(t *testing.T) {
	key, err := NewPrivateKey(rand.Reader, 192)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Printf("p len: %d, v: %x\n", key.P.BitLen(), key.P.Bytes())
	fmt.Printf("q len: %d, v: %x\n", key.Q.BitLen(), key.Q.Bytes())
	fmt.Printf("n len: %d, v: %x\n", key.N.BitLen(), key.N.Bytes())
	fmt.Printf("d len: %d, v: %x\n", key.D.BitLen(), key.D.Bytes())
	fmt.Printf("lamdba len: %d, v: %x\n", key.Lcm.BitLen(), key.Lcm.Bytes())

	msg := make([]byte, 10)
	_, err = rand.Read(msg)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	msg[0] &= 0x7f

	m := new(big.Int).SetBytes(msg)
	encrypted := new(big.Int).Exp(m, key.N, key.N)
	decrypted := new(big.Int).Exp(encrypted, key.D, key.PQ)
	fmt.Printf("m len: %d, v: %x\n", m.BitLen(), m.Bytes())
	fmt.Printf("en len: %d, v: %x\n", encrypted.BitLen(), encrypted.Bytes())
	fmt.Printf("de len: %d, v: %x\n", decrypted.BitLen(), decrypted.Bytes())

	m1 := new(big.Int).Exp(encrypted, key.Dp, key.P)
	m2 := new(big.Int).Exp(encrypted, key.Dq, key.Q)
	h := new(big.Int).Sub(m1, m2)
	if h.Sign() < 0 {
		h.Add(h, key.P)
	}
	h.Mul(h, key.Qinv)
	h.Mod(h, key.P)
	de := new(big.Int).Mul(h, key.Q)
	de.Add(de, m2)
	fmt.Printf("dp len: %d, v: %x\n", key.Dp.BitLen(), key.Dp.Bytes())
	fmt.Printf("dq len: %d, v: %x\n", key.Dq.BitLen(), key.Dq.Bytes())
	fmt.Printf("qinv len: %d, v: %x\n", key.Qinv.BitLen(), key.Qinv.Bytes())
	fmt.Printf("m1 len: %d, v: %x\n", m1.BitLen(), m1.Bytes())
	fmt.Printf("m2 len: %d, v: %x\n", m2.BitLen(), m2.Bytes())
	fmt.Printf("h len: %d, v: %x\n", h.BitLen(), h.Bytes())
	fmt.Printf("de len: %d, v: %x\n", de.BitLen(), de.Bytes())
}

func Test_bigIntEndian(t *testing.T) {
	b := int64(123456789123456789)
	a := new(big.Int).SetInt64(b)
	le := make([]byte, 8)
	be := make([]byte, 8)
	binary.LittleEndian.PutUint64(le, uint64(b))
	binary.BigEndian.PutUint64(be, uint64(b))
	fmt.Printf("%x\n", a.Bytes())
	fmt.Printf("%x\n", le)
	fmt.Printf("%x\n", be)
}

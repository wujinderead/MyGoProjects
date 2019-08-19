package cipher

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"math/big"
	"testing"
)

func Test_rsa(t *testing.T) {
	priv, err := rsa.GenerateMultiPrimeKey(rand.Reader, 4, 768)
	if err != nil {
		t.Errorf("error generate rsa priv key")
		t.Fail()
	}
	primes := priv.Primes
	for _, prime := range primes {
		fmt.Printf("prime: %x\n", prime.Bytes())
	}
	fmt.Printf("D: %x\n", priv.D.Bytes())
	fmt.Printf("E: %x\n", priv.E)
	fmt.Printf("N: %x\n", priv.N.Bytes())
	fmt.Printf("Dp: %x\n", priv.Precomputed.Dp.Bytes())
	fmt.Printf("Dq: %x\n", priv.Precomputed.Dq.Bytes())
	fmt.Printf("qInv: %x\n", priv.Precomputed.Qinv.Bytes())
	for _, crt := range priv.Precomputed.CRTValues {
		fmt.Printf("R: %x\n", crt.R.Bytes())
		fmt.Printf("Exp: %x\n", crt.Exp.Bytes())
		fmt.Printf("inv: %x\n", crt.Coeff.Bytes())
	}
	message := make([]byte, 80)
	rand.Reader.Read(message)
	msg := new(big.Int).SetBytes(message)
	fmt.Printf("msg: %x\n", msg.Bytes())

	c := new(big.Int).Exp(msg, new(big.Int).SetInt64(int64(priv.E)), priv.N)
	fmt.Printf("encrypted: %x\n", c.Bytes())

	m := new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
	fmt.Printf("m1: %x\n", m.Bytes())
	m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
	fmt.Printf("m2: %x\n", m2.Bytes())
	m.Sub(m, m2)
	if m.Sign() < 0 {
		m.Add(m, priv.Primes[0])
	}
	m.Mul(m, priv.Precomputed.Qinv)
	m.Mod(m, priv.Primes[0])
	fmt.Printf("h: %x\n", m.Bytes())
	m.Mul(m, priv.Primes[1])
	m.Add(m, m2)
	fmt.Printf("m = msg ? %v\n", bytes.Equal(m.Bytes(), msg.Bytes()))

	for i, value := range priv.Precomputed.CRTValues {
		prime := priv.Primes[2+i]
		m2.Exp(c, value.Exp, prime)
		fmt.Printf("mi: %x\n", m2.Bytes())
		m2.Sub(m2, m)
		m2.Mul(m2, value.Coeff)
		m2.Mod(m2, prime)
		fmt.Printf("h: %x\n", m2.Bytes())
		if m2.Sign() < 0 {
			m2.Add(m2, prime)
		}
		m2.Mul(m2, value.R)
		m.Add(m, m2)
		fmt.Printf("m = msg ? %v\n", bytes.Equal(m.Bytes(), msg.Bytes()))
	}
}

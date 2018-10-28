package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"math/big"
	"testing"
)

func Test_rsa(t *testing.T) {
	priv, err := rsa.GenerateMultiPrimeKey(rand.Reader, 3, 512)
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
	message := make([]byte, 32)
	rand.Reader.Read(message)
	fmt.Printf("msg: %x\n", message)
	c := new(big.Int).SetBytes(message)
	fmt.Printf("c: %x\n", c.Bytes())

	m := new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
	m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
	m.Sub(m, m2)
	if m.Sign() < 0 {
		m.Add(m, priv.Primes[0])
	}
	m.Mul(m, priv.Precomputed.Qinv)
	m.Mod(m, priv.Primes[0])
	m.Mul(m, priv.Primes[1])
	m.Add(m, m2)

	for i, values := range priv.Precomputed.CRTValues {
		prime := priv.Primes[2+i]
		m2.Exp(c, values.Exp, prime)
		m2.Sub(m2, m)
		m2.Mul(m2, values.Coeff)
		m2.Mod(m2, prime)
		if m2.Sign() < 0 {
			m2.Add(m2, prime)
		}
		m2.Mul(m2, values.R)
		m.Add(m, m2)
	}
}

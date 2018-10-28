package factorization

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
)

type PublicKey struct {
	N *big.Int
}

type PrivateKey struct {
	PublicKey
	P, Q, D               *big.Int
	Lcm, PQ, Dp, Dq, Qinv *big.Int
}

func NewPrivateKey(random io.Reader, bits int) (*PrivateKey, error) {
	if bits%192 != 0 || bits <= 0 {
		return nil, errors.New("bits length are not multiple of 192")
	}
	one := new(big.Int).SetInt64(1)
	bitLen := bits / 3
	p, q := new(big.Int), new(big.Int)
	for {
		var err error
		p, err = rand.Prime(random, bitLen)
		if err != nil {
			return nil, err
		}
		q, err = rand.Prime(random, bitLen)
		if err != nil {
			return nil, err
		}
		if p.Cmp(q) != 0 {
			break
		}
	}
	pq := new(big.Int).Mul(p, q)
	n := new(big.Int).Mul(pq, p)
	p_1 := new(big.Int).Sub(p, one)
	q_1 := new(big.Int).Sub(q, one)
	lcm := new(big.Int).Mul(p_1, q_1)
	gcd := new(big.Int).GCD(nil, nil, p_1, q_1)
	fmt.Println("gcd: ", gcd.String())
	lcm.Div(lcm, gcd)
	d := new(big.Int).ModInverse(n, lcm)
	dp := new(big.Int).Mod(d, p_1)
	dq := new(big.Int).Mod(d, q_1)
	qinv := new(big.Int).ModInverse(q, p)
	return &PrivateKey{PublicKey{n}, p, q, d, lcm, pq, dp, dq, qinv}, nil
}

package ecc

import "math/big"

var (
	Zero     = new(big.Int).SetInt64(0)
	ONE      = new(big.Int).SetInt64(1)
	TWO      = new(big.Int).SetInt64(2)
	NEG_ONE  = new(big.Int).SetInt64(-1)
	Infinity = &EcPoint{Zero, Zero}
)

// a/b mod p
func ModFraction(a, b, p *big.Int) *big.Int {
	b_inv := new(big.Int).ModInverse(b, p)
	b_inv.Mul(b_inv, a)
	b_inv.Mod(b_inv, p)
	return b_inv
}

func ModFractionInt64(a, b int64, p *big.Int) *big.Int {
	aa := new(big.Int).SetInt64(a)
	bb := new(big.Int).SetInt64(b)
	return ModFraction(aa, bb, p)
}

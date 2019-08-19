package util

import (
	"math/big"
)

var (
	P25519, SQRTNEG486664, N486662, ONE, D *big.Int
)

func init() {
	N19 := new(big.Int).SetInt64(19)
	P25519 = new(big.Int).SetInt64(1)
	P25519.Lsh(P25519, 255)
	P25519.Sub(P25519, N19)
	NEG486664 := new(big.Int).SetInt64(-486664)
	N486662 = new(big.Int).SetInt64(486662)
	SQRTNEG486664 = new(big.Int).ModSqrt(NEG486664, P25519)
	ONE = new(big.Int).SetInt64(1)
	D, _ = new(big.Int).SetString("37095705934669439343138083508754565189542113879843219016388785533085940283555", 10)
}

func Curve25519ToEd25519(u, v *big.Int) (x, y *big.Int) {
	// (u, v) = ((1+y)/(1-y), sqrt(-486664)*u/x)
	x = new(big.Int).ModInverse(v, P25519) // 1/v
	x.Mul(x, u)                            // u/v
	x.Mul(x, SQRTNEG486664)                // sqrt(-486664)*u/v
	x.Mod(x, P25519)                       // x ends
	u_1 := new(big.Int).Add(u, ONE)        // u+1
	u_1.ModInverse(u_1, P25519)            // 1/(u+1)
	y = new(big.Int).Sub(u, ONE)           // u-1
	y.Mul(y, u_1)                          // (u-1)/(u+1)
	y.Mod(y, P25519)                       // y ends
	return
}

func Ed25519ToCurve25519(x, y *big.Int) (u, v *big.Int) {
	// (x, y) = (sqrt(-486664)*u/v, (u-1)/(u+1))
	_1_y := new(big.Int).Sub(ONE, y)       // 1-y
	_1_y.ModInverse(_1_y, P25519)          // 1/(1-y)
	u = new(big.Int).Add(ONE, y)           // 1+y
	u.Mul(u, _1_y)                         // (1+y)/(1-y)
	u.Mod(u, P25519)                       // u ends
	v = new(big.Int).ModInverse(x, P25519) // 1/x
	v.Mul(u, v)                            // u/x
	v.Mul(SQRTNEG486664, v)                // sqrt(-486664)*u/x
	v.Mod(v, P25519)
	return
}

func Curve25519U2V(u *big.Int) (v, v1 *big.Int) {
	// v^2 = u^3 + A*u^2 + u
	u2 := new(big.Int).Mul(u, u)      // u^2
	v = new(big.Int).Mul(N486662, u2) // 486662u^2
	v.Add(v, u)                       // 486662u^2 + u
	u2.Mul(u2, u)                     // u^3
	v.Add(u2, v)                      // u^3 + 486662u^2 + u
	v.ModSqrt(v, P25519)              // sqrt(u^3 + A*u^2 + u)
	v.Mod(v, P25519)                  // first v
	v1 = new(big.Int).Sub(P25519, v)  // second v
	return
}

func Ed25519Y2X(y *big.Int) (x *big.Int) {
	// -x² + y² = 1 + dx²y²
	// x² = (y²-1)/(dy²+1)
	x = new(big.Int).Mul(y, y)    // y²
	tmp := new(big.Int).Mul(D, x) // dy²
	tmp.Add(tmp, ONE)             // dy²+1
	tmp.ModInverse(tmp, P25519)   // 1/(dy²+1)
	tmp.Mod(tmp, P25519)
	x.Sub(x, ONE)        //  y²-1
	x.Mul(x, tmp)        //  (y²-1)/(dy²+1)
	x.ModSqrt(x, P25519) // sqrt( (y²-1)/(dy²+1) )
	x.Mod(x, P25519)
	return
}

func Curve25519XToEd25519Y(u *big.Int) (y *big.Int) {
	// (x, y) = (sqrt(-486664)*u/v, (u-1)/(u+1))
	u_1 := new(big.Int).Add(u, ONE) // u+1
	u_1.ModInverse(u_1, P25519)     // 1/(u+1)
	y = new(big.Int).Sub(u, ONE)    // u-1
	y.Mul(y, u_1)                   // (u-1)/(u+1)
	y.Mod(y, P25519)                // y ends
	return
}

func Ed25519YToCurve25519X(y *big.Int) (u *big.Int) {
	// (u, v) = ((1+y)/(1-y), sqrt(-486664)*u/x)
	y.Mod(y, P25519)                 // y ends
	_1_y := new(big.Int).Sub(ONE, y) // 1-y
	_1_y.ModInverse(_1_y, P25519)    // 1/(1-y)
	u = new(big.Int).Add(ONE, y)     // 1+y
	u.Mul(u, _1_y)                   // (1+y)/(1-y)
	u.Mod(u, P25519)                 // u ends
	return
}

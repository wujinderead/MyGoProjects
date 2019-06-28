package ecc

import (
	"crypto/elliptic"
	"math/big"
)

type FpCurve EcCurve

func (curve *FpCurve) IsOnCurve(p *EcPoint) bool {
	// y² ≡ x³ + Ax + B (mod P)
	y2 := new(big.Int).Mul(p.Y, p.Y) // y²
	y2.Mod(y2, curve.P)              // y² mod P

	x3 := new(big.Int).Mul(p.X, p.X)
	x3.Mul(x3, p.X) // x³

	ax := new(big.Int).Mul(curve.A, p.X)
	x3.Add(x3, ax)
	x3.Add(x3, curve.B)
	x3.Mod(x3, curve.P)

	return y2.Cmp(x3) == 0
}

func (curve *FpCurve) Add(p, q *EcPoint) *EcPoint {
	if p.Equals(Infinity) {
		return q.Copy()
	}
	if q.Equals(Infinity) {
		return p.Copy()
	}
	m, tmp := new(big.Int), new(big.Int)
	if p.Equals(q) { // double
		if (p.Y).Cmp(ZERO) == 0 { // Yp=Yq=0, i.e. (a, 0) + (a, 0) = Infinity
			return NewPoint()
		} else {
			// m = (3Xp² + A) (2Yp)^-1 mod P
			tmp.Mul(p.X, p.X)                      // Xp²
			tmp.Mul(tmp, new(big.Int).SetInt64(3)) // 3Xp²
			tmp.Add(tmp, curve.A)                  // (3Xp² + A)
			tmp.Mod(tmp, curve.P)
			m = m.Lsh(p.Y, 1)        // 2Yp
			m.ModInverse(m, curve.P) // (2Yp)^-1 mod P
			m.Mul(m, tmp)
			m.Mod(m, curve.P)
		}
	} else { // add
		if (p.X).Cmp(q.X) == 0 { // (a, b) + (a, P-b) = Infinity
			return NewPoint()
		} else {
			// m = (Yp - Yq) (Xp - Xq)^-1 mod P
			m.Sub(p.X, q.X)
			m.ModInverse(m, curve.P)
			tmp = tmp.Sub(p.Y, q.Y)
			m.Mul(m, tmp)
			m.Mod(m, curve.P)
		}
	}
	// Xr = (m² - Xp - Xq) mod P
	Xr := new(big.Int).Mul(m, m)
	Xr.Sub(Xr, p.X)
	Xr.Sub(Xr, q.X)
	Xr.Mod(Xr, curve.P)
	// Yr = (Yp + m(Xr - Xp)) mod P
	Yr := new(big.Int).Sub(Xr, p.X)
	Yr.Mul(Yr, m)
	Yr.Add(Yr, p.Y)
	Yr.Mod(Yr, curve.P)
	// return (Xr, -Yr). if Yr = 0, -Yr = 0; else -Yr = P - Yr
	if Yr.Cmp(ZERO) != 0 {
		Yr.Sub(curve.P, Yr)
	}
	return &EcPoint{Xr, Yr}
}

func (curve *FpCurve) ScalaMult(p *EcPoint, k []byte) *EcPoint {
	sum := NewPoint()
	var mask byte = 0x80
	for _, b := range k {
		for i := 0; i < 8; i++ {
			sum = curve.Add(sum, sum)
			if (b & mask) == mask { // check bit is 1 or not
				sum = curve.Add(sum, p)
			}
			b <<= 1
		}
	}
	return sum
}

func (curve *FpCurve) ScalaMultBase(k []byte) *EcPoint {
	return curve.ScalaMult(&EcPoint{curve.X, curve.Y}, k)
}

func (curve *FpCurve) ToGoNative() *elliptic.CurveParams {
	return &elliptic.CurveParams{
		P:  curve.P,              // the order of the underlying field
		N:  curve.Order,          // the order of the base point
		B:  curve.B,              // the constant of the curve equation
		Gx: curve.X, Gy: curve.Y, // (x,y) of the base point
		BitSize: curve.P.BitLen(), // the size of the underlying field
		Name:    "",               // the canonical name of the curve
	}
}

package ecc

import (
	"fmt"
	"math/big"
)

type EcCurve struct {
	head                       *EcCurveHead
	Seed, P, A, B, X, Y, Order *big.Int
}

type EcPoint struct {
	X, Y *big.Int
}

// for instantiate an EcPoint
func NewPoint() *EcPoint {
	return &EcPoint{new(big.Int).SetInt64(0), new(big.Int).SetInt64(0)}
}

func (point *EcPoint) Equals(p *EcPoint) bool {
	return point.X.Cmp(p.X) == 0 && point.Y.Cmp(p.Y) == 0
}

func (point *EcPoint) Copy() *EcPoint {
	return &EcPoint{new(big.Int).Set(point.X), new(big.Int).Set(point.Y)}
}

func (point *EcPoint) ToStr() string {
	return fmt.Sprintf("[%x, %x]", point.X.Bytes(), point.Y.Bytes())
}

func parseEcCurve(head *EcCurveHead, data []byte) *EcCurve {
	ret := &EcCurve{}
	ret.head = head
	if head.seedLen > 0 {
		ret.Seed = new(big.Int).SetBytes(data[0:head.seedLen])
	}
	ret.P = new(big.Int).SetBytes(data[head.seedLen : head.seedLen+head.paramLen])
	ret.A = new(big.Int).SetBytes(data[head.seedLen+head.paramLen : head.seedLen+head.paramLen*2])
	ret.B = new(big.Int).SetBytes(data[head.seedLen+head.paramLen*2 : head.seedLen+head.paramLen*3])
	ret.X = new(big.Int).SetBytes(data[head.seedLen+head.paramLen*3 : head.seedLen+head.paramLen*4])
	ret.Y = new(big.Int).SetBytes(data[head.seedLen+head.paramLen*4 : head.seedLen+head.paramLen*5])
	ret.Order = new(big.Int).SetBytes(data[head.seedLen+head.paramLen*5 : head.seedLen+head.paramLen*6])
	return ret
}

func (curve *EcCurve) IsOnCurve(p *EcPoint) bool {
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

func (curve *EcCurve) Add(p, q *EcPoint) *EcPoint {
	if p.Equals(Infinity) {
		return q.Copy()
	}
	if q.Equals(Infinity) {
		return p.Copy()
	}
	m, tmp := new(big.Int), new(big.Int)
	if p.Equals(q) { // double
		if (p.Y).Cmp(Zero) == 0 { // Yp=Yq=0, i.e. (a, 0) + (a, 0) = Infinity
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
	if Yr.Cmp(Zero) != 0 {
		Yr.Sub(curve.P, Yr)
	}
	return &EcPoint{Xr, Yr}
}

func (curve *EcCurve) ScalaMult(p *EcPoint, k []byte) *EcPoint {
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

func (curve *EcCurve) ScalaMultBase(k []byte) *EcPoint {
	return curve.ScalaMult(&EcPoint{curve.X, curve.Y}, k)
}

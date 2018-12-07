package ecc

import (
	"math/big"
)

type F2mCurve struct {
	*EcCurve
	p []int
}

func (curve *F2mCurve) gadd(a, b *big.Int) *big.Int {
	return new(big.Int).Xor(a, b)  // since we're in GF(2^m), addition is an XOR
}

func (curve *F2mCurve) gaddSelf(a, b *big.Int) {
	a.Xor(a, b)
}

func (curve *F2mCurve) gmul(x, y *big.Int) *big.Int {
	if x == y {
		pro := bn_gf2m_sqr(x.Bits())
		bn_gf2m_mod_arr_self(pro, curve.p)
		return new(big.Int).SetBits(pro)
	}
	pro := bn_gf2m_mul(x.Bits(), y.Bits())
	bn_gf2m_mod_arr_self(pro, curve.p)
	return new(big.Int).SetBits(pro)
}

func (curve *F2mCurve) gmulinv(a *big.Int) *big.Int {
	// in GF(2^m), a^(2^m-1)=1, so a*a^(2^m-2)=1, so a^-1=a(2^m-2)
	inv := bn_gf2m_mod_inv_vartime(a.Bits(), curve.P.Bits())
	return new(big.Int).SetBits(inv)
}

// elliptic curve y²+xy=x³+Ax²+B over finite binary field
func (curve *F2mCurve) IsOnCurve(x, y *big.Int) bool {
	left := curve.gadd(x, y)       // x+y
	left = curve.gmul(left, y)     // (x+y)*y
	x2 := curve.gmul(x, x)         // x²
	right := curve.gmul(x2, x)     // x³
	Ax2 := curve.gmul(curve.A, x2) // Ax²
	curve.gaddSelf(right, Ax2)
	curve.gaddSelf(right, curve.B) // x³+Ax²+B
	return left.Cmp(right) == 0
}

func (curve *F2mCurve) Add(p, q *EcPoint) *EcPoint {
	if p.Equals(Infinity) {
		return q.Copy()
	}
	if q.Equals(Infinity) {
		return p.Copy()
	}
	if p.Equals(q) { // double
		if (p.X).Cmp(Zero) == 0 { // (0, b) + (0, b) = Infinity
			return NewPoint()
		} else {
			// x3 = (x1+y1/x1)²+(x1+y1/x1)+A
			// y3 = (x1+y1/x1)³+(x1+A+1)*(x1+y1/x1)+A+y1
			m := curve.gmulinv(p.X)
			m = curve.gmul(m, p.Y)
			curve.gaddSelf(m, p.X)           // x1+y1/x1
			x3 := curve.gmul(m, m)           // (x1+y1/x1)²
			y3 := curve.gmul(m, x3)          // (x1+y1/x1)³
			curve.gaddSelf(x3, m)
			curve.gaddSelf(x3, curve.A)      // (x1+y1/x1)²+(x1+y1/x1)+A
			n := curve.gadd(p.X, curve.A)
			curve.gaddSelf(n, ONE)           // (x1+A+1)
			n = curve.gmul(n, m)             // (x1+A+1)*(x1+y1/x1)
			curve.gaddSelf(y3, n)
			curve.gaddSelf(y3, curve.A)
			curve.gaddSelf(y3, p.Y)          // (x1+y1/x1)³+(x1+A+1)*(x1+y1/x1)+A+y1
			return &EcPoint{x3, y3}
		}
	} else { // add
		if (p.X).Cmp(q.X) == 0 { // (a, b) + (a, b+a) = Infinity
			return NewPoint()
		} else {
			// x3 = ((y1+y2)/(x1+x2))²+((y1+y2)/(x1+x2))+x1+x2+A
			// y3 = ((y1+y2)/(x1+x2))³+(x2+A+1)*((y1+y2)/(x1+x2))+x1+x2+A+y1
			m := curve.gadd(p.X, q.X)
			n := curve.gadd(p.Y, q.Y)
			m = curve.gmulinv(m)
			m = curve.gmul(m, n)             // (y1+y2)/(x1+x2)
			x3 := curve.gmul(m, m)           // ((y1+y2)/(x1+x2))²
			y3 := curve.gmul(m, x3)          // ((y1+y2)/(x1+x2))³
			curve.gaddSelf(x3, m)
			curve.gaddSelf(x3, p.X)
			curve.gaddSelf(x3, q.X)
			curve.gaddSelf(x3, curve.A)      // ((y1+y2)/(x1+x2))²+((y1+y2)/(x1+x2))+x1+x2+A
			n = curve.gadd(q.X, curve.A)
			curve.gaddSelf(n, ONE)           // (x2+A+1)
			n = curve.gmul(n, m)             // (x2+A+1)*((y1+y2)/(x1+x2))
			curve.gaddSelf(y3, n)
			curve.gaddSelf(y3, p.X)
			curve.gaddSelf(y3, q.X)
			curve.gaddSelf(y3, curve.A)
			curve.gaddSelf(y3, p.Y)          // ((y1+y2)/(x1+x2))³+(x2+A+1)*((y1+y2)/(x1+x2))+x1+x2+A+y1
			return &EcPoint{x3, y3}
		}
	}

}

func (curve *F2mCurve) ScalaMult(p *EcPoint, k []byte) *EcPoint {
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

func (curve *F2mCurve) ScalaMultBase(k []byte) *EcPoint {
	return curve.ScalaMult(&EcPoint{curve.X, curve.Y}, k)
}

// reference implements for gf2m multiply
func (curve *F2mCurve) gmulReference(x, y *big.Int) *big.Int {
	product := new(big.Int).SetInt64(0)
	a := new(big.Int).Set(x)
	size := curve.P.BitLen()-1
	for i:=0; i<size; i++ {
		// if b least bit is 1, then add the corresponding a to p
		// final product is sum of all a's corresponding to odd b's
		if y.Bit(i)==1 {
			product.Xor(product, a)
		}
		a.Lsh(a, 1)
		if a.Bit(size) == 1 {
			a.Xor(a, curve.P)
		}
	}
	return product
}
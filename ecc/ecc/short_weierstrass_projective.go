package ecc

import (
	"math/big"
)

// in this file, we implements short weierstrass elliptic curve arithmetic in projective coordinates
func (curve *FpCurve) AddProjective(p1, p2 *EcPoint) *EcPoint {
	x1, y1, x2, y2 := p1.X, p1.Y, p2.X, p2.Y
	z1 := zForAffine(x1, y1)
	z2 := zForAffine(x2, y2)
	return curve.affineFromProjective(curve.addProjective(x1, y1, z1, x2, y2, z2))
}

// see https://hyperelliptic.org/EFD/g1p/auto-shortw-projective.html#addition-add-1998-cmo-2
func (curve *FpCurve) addProjective(x1, y1, z1, x2, y2, z2 *big.Int) (*big.Int, *big.Int, *big.Int) {
	x3, y3, z3 := new(big.Int), new(big.Int), new(big.Int)
	if z1.Sign() == 0 {
		x3.Set(x2)
		y3.Set(y2)
		z3.Set(z2)
		return x3, y3, z3
	}
	if z2.Sign() == 0 {
		x3.Set(x1)
		y3.Set(y1)
		z3.Set(z1)
		return x3, y3, z3
	}

	y1z2 := new(big.Int).Mul(y1, z2) // Y1Z2 = Y1*Z2
	y1z2.Mod(y1z2, curve.P)
	x1z2 := new(big.Int).Mul(x1, z2) // X1Z2 = X1*Z2
	x1z2.Mod(x1z2, curve.P)

	u := new(big.Int).Mul(y2, z1)
	u.Mod(u, curve.P)
	u.Sub(u, y1z2) // u = Y2*Z1-Y1Z2

	v := new(big.Int).Mul(x2, z1)
	v.Mod(v, curve.P)
	v.Sub(v, x1z2) // v = X2*Z1-X1Z2

	if u.Sign() == 0 && v.Sign() == 0 { // u=v=0 means adding two same point
		return curve.doubleProjective(x1, y1, z1)
	}

	uu := new(big.Int).Mul(u, u)     // uu = u²
	vv := new(big.Int).Mul(v, v)     // vv = v²
	z1z2 := new(big.Int).Mul(z1, z2) // Z1Z2 = Z1*Z2
	vvv := new(big.Int).Mul(v, vv)   // vvv =v*vv
	R := new(big.Int).Mul(vv, x1z2)  // R = vv*X1Z2
	A := new(big.Int).Mul(uu, z1z2)
	A.Sub(A, vvv)
	A.Sub(A, R)
	A.Sub(A, R) // A = uu*Z1Z2-vvv-2*R

	x3.Mul(v, A) // X3 = v*A
	x3.Mod(x3, curve.P)

	y3.Sub(R, A)
	y3.Mul(y3, u)
	z3.Mul(vvv, y1z2)
	y3.Sub(y3, z3) // Y3 = u*(R-A)-vvv*Y1Z2
	y3.Mod(y3, curve.P)

	z3.Mul(vvv, z1z2) // Z3 = vvv*Z1Z2
	z3.Mod(z3, curve.P)
	return x3, y3, z3
}

func (curve *FpCurve) DoubleProjective(point *EcPoint) *EcPoint {
	x1, y1 := point.X, point.Y
	z1 := zForAffine(x1, y1)
	return curve.affineFromProjective(curve.doubleProjective(x1, y1, z1))
}

// see https://hyperelliptic.org/EFD/g1p/auto-shortw-projective.html#doubling-dbl-2007-bl
func (curve *FpCurve) doubleProjective(x, y, z *big.Int) (*big.Int, *big.Int, *big.Int) {
	XX := new(big.Int).Mul(x, x) // XX = X1²
	ZZ := new(big.Int).Mul(z, z) // ZZ = Z1²
	w := new(big.Int).Mul(curve.A, ZZ)
	w.Add(w, XX)
	w.Add(w, XX)
	w.Add(w, XX) // w = a*ZZ+3*XX
	s := new(big.Int).Mul(y, z)
	s.Lsh(s, 1)                    // s = 2*Y1*Z1
	ss := new(big.Int).Mul(s, s)   // ss = s*s
	sss := new(big.Int).Mul(s, ss) // sss = s*ss
	R := new(big.Int).Mul(y, s)    // R = Y1*s
	RR := new(big.Int).Mul(R, R)   // RR = R*R
	B := new(big.Int).Mul(x, R)
	B.Lsh(B, 1) // B = 2*X1*R
	h := new(big.Int).Mul(w, w)
	h.Sub(h, B)
	h.Sub(h, B)                  // h = w²-2*B
	x3 := new(big.Int).Mul(h, s) // X3 = h*s
	x3.Mod(x3, curve.P)
	y3 := new(big.Int).Sub(B, h)
	y3.Mul(y3, w)
	y3.Sub(y3, RR)
	y3.Sub(y3, RR) // Y3 = w*(B-h)-2*RR
	y3.Mod(y3, curve.P)
	sss.Mod(sss, curve.P) // Z3 = sss
	return x3, y3, sss
}

func (curve *FpCurve) ScalaMultProjective(point *EcPoint, k []byte) *EcPoint {
	Bx, By := point.X, point.Y
	Bz := new(big.Int).SetInt64(1)
	x, y, z := new(big.Int), new(big.Int), new(big.Int)

	for _, byte := range k {
		for bitNum := 0; bitNum < 8; bitNum++ {
			x, y, z = curve.doubleProjective(x, y, z)
			if byte&0x80 == 0x80 {
				x, y, z = curve.addProjective(Bx, By, Bz, x, y, z)
			}
			byte <<= 1
		}
	}

	return curve.affineFromProjective(x, y, z)
}

func (curve *FpCurve) ScalaMultBaseProjective(k []byte) *EcPoint {
	return curve.ScalaMultProjective(&EcPoint{curve.X, curve.Y}, k)
}

func (curve *FpCurve) affineFromProjective(x, y, z *big.Int) *EcPoint {
	if z.Sign() == 0 {
		return NewPoint()
	}
	zinv := new(big.Int).ModInverse(z, curve.P)
	xOut := new(big.Int).Mul(x, zinv)
	xOut.Mod(xOut, curve.P)
	yOut := new(big.Int).Mul(y, zinv)
	yOut.Mod(yOut, curve.P)
	return &EcPoint{xOut, yOut}
}

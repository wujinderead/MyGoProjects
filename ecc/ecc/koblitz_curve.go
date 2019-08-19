package ecc

import "math/big"

// in this file, we implements koblitz curve arithmetic in Jacobian coordinates
// koblitz curve is elliptic curve y²=x³+ax+b with a=0
type KoblitzCurve FpCurve

func (curve *KoblitzCurve) Add(p1, p2 *EcPoint) *EcPoint {
	x1, y1, x2, y2 := p1.X, p1.Y, p2.X, p2.Y
	z1 := zForAffine(x1, y1)
	z2 := zForAffine(x2, y2)
	return (*FpCurve)(curve).affineFromJacobian(curve.addJacobian(x1, y1, z1, x2, y2, z2))
}

// see https://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-0.html#addition-add-2007-bl
func (curve *KoblitzCurve) addJacobian(x1, y1, z1, x2, y2, z2 *big.Int) (*big.Int, *big.Int, *big.Int) {
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

	z1z1 := new(big.Int).Mul(z1, z1) // Z1Z1 = Z1²
	z1z1.Mod(z1z1, curve.P)
	z2z2 := new(big.Int).Mul(z2, z2) // Z2Z2 = Z2²
	z2z2.Mod(z2z2, curve.P)

	u1 := new(big.Int).Mul(x1, z2z2) // U1 = X1*Z2Z2
	u1.Mod(u1, curve.P)
	u2 := new(big.Int).Mul(x2, z1z1) // U2 = X2*Z1Z1
	u2.Mod(u2, curve.P)
	h := new(big.Int).Sub(u2, u1) // H = U2-U1
	xEqual := h.Sign() == 0
	if h.Sign() == -1 {
		h.Add(h, curve.P)
	}
	i := new(big.Int).Lsh(h, 1)
	i.Mul(i, i)                 // I = (2*H)²
	j := new(big.Int).Mul(h, i) // J = H*I

	s1 := new(big.Int).Mul(y1, z2)
	s1.Mul(s1, z2z2) // S1 = Y1*Z2*Z2Z2
	s1.Mod(s1, curve.P)
	s2 := new(big.Int).Mul(y2, z1)
	s2.Mul(s2, z1z1) // S2 = Y2*Z1*Z1Z1
	s2.Mod(s2, curve.P)
	r := new(big.Int).Sub(s2, s1)
	if r.Sign() == -1 {
		r.Add(r, curve.P)
	}
	yEqual := r.Sign() == 0
	if xEqual && yEqual { // H=r=0 means adding two same point
		return curve.doubleJacobian(x1, y1, z1)
	}
	r.Lsh(r, 1)                  // r = 2*(S2-S1)
	v := new(big.Int).Mul(u1, i) // V = U1*I

	x3.Set(r)
	x3.Mul(x3, x3)
	x3.Sub(x3, j)
	x3.Sub(x3, v)
	x3.Sub(x3, v) // X3 = r²-J-2*V
	x3.Mod(x3, curve.P)

	y3.Set(r)
	v.Sub(v, x3)
	y3.Mul(y3, v)
	s1.Mul(s1, j)
	s1.Lsh(s1, 1)
	y3.Sub(y3, s1) // Y3 = r*(V-X3)-2*S1*J
	y3.Mod(y3, curve.P)

	z3.Add(z1, z2)
	z3.Mul(z3, z3)
	z3.Sub(z3, z1z1)
	z3.Sub(z3, z2z2)
	z3.Mul(z3, h) // Z3 = ((Z1+Z2)²-Z1Z1-Z2Z2)*H
	z3.Mod(z3, curve.P)

	return x3, y3, z3
}

func (curve *KoblitzCurve) Double(point *EcPoint) *EcPoint {
	x1, y1 := point.X, point.Y
	z1 := zForAffine(x1, y1)
	return (*FpCurve)(curve).affineFromJacobian(curve.doubleJacobian(x1, y1, z1))
}

// see https://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-0.html#doubling-dbl-2009-l
func (curve *KoblitzCurve) doubleJacobian(x, y, z *big.Int) (*big.Int, *big.Int, *big.Int) {
	A := new(big.Int).Mul(x, x) // A = X1²
	A.Mod(A, curve.P)
	B := new(big.Int).Mul(y, y) // B = Y1²
	B.Mod(B, curve.P)
	C := new(big.Int).Mul(B, B) // C = B²
	C.Mod(C, curve.P)
	D := new(big.Int).Mul(x, B)
	D.Lsh(D, 2) // D = 2*((X1+B)²-A-C) = 4*X1*B
	E := new(big.Int).Lsh(A, 1)
	E.Add(E, A) // E = 3*A

	x3 := new(big.Int).Mul(E, E) // F = E²
	x3.Sub(x3, D)
	x3.Sub(x3, D) // X3 = F-2*D
	x3.Mod(x3, curve.P)

	y3 := new(big.Int).Sub(D, x3)
	y3.Mul(y3, E)
	c8 := new(big.Int).Lsh(C, 3)
	y3.Sub(y3, c8) // Y3 = E*(D-X3)-8*C
	y3.Mod(y3, curve.P)

	z3 := new(big.Int).Lsh(y, 1)
	z3.Mul(z3, z) // Z3 = 2*Y1*Z1
	z3.Mod(z3, curve.P)
	return x3, y3, z3
}

func (curve *KoblitzCurve) ScalaMult(point *EcPoint, k []byte) *EcPoint {
	Bx, By := point.X, point.Y
	Bz := new(big.Int).SetInt64(1)
	x, y, z := new(big.Int), new(big.Int), new(big.Int)

	for _, byte := range k {
		for bitNum := 0; bitNum < 8; bitNum++ {
			x, y, z = curve.doubleJacobian(x, y, z)
			if byte&0x80 == 0x80 {
				x, y, z = curve.addJacobian(Bx, By, Bz, x, y, z)
			}
			byte <<= 1
		}
	}

	return (*FpCurve)(curve).affineFromJacobian(x, y, z)
}

func (curve *KoblitzCurve) ScalaMultBase(k []byte) *EcPoint {
	return curve.ScalaMult(&EcPoint{curve.X, curve.Y}, k)
}

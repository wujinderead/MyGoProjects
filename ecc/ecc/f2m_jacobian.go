package ecc

import "math/big"

// in this file, we implements short weierstrass elliptic curve arithmetic
// over binary finite field in Jacobian coordinates: x=X/Z², y=Y/Z³
func (curve *F2mCurve) AddJacobian(p1, p2 *EcPoint) *EcPoint {
	x1, y1, x2, y2 := p1.X, p1.Y, p2.X, p2.Y
	z1 := zForAffine(x1, y1)
	z2 := zForAffine(x2, y2)
	return curve.affineFromJacobian(curve.addJacobian(x1, y1, z1, x2, y2, z2))
}

// see http://hyperelliptic.org/EFD/g12o/auto-shortw-jacobian.html#addition-add-2005-dl
func (curve *F2mCurve) addJacobian(x1, y1, z1, x2, y2, z2 *big.Int) (*big.Int, *big.Int, *big.Int) {
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

	O1 := curve.gmul(z1, z1) // O1 = Z1²
	O2 := curve.gmul(z2, z2) // O2 = Z2²
	A := curve.gmul(x1, O2)  // A = X1*O2
	B := curve.gmul(x2, O1)  // B = X2*O1
	C := curve.gmul(y1, O2)
	C = curve.gmul(C, z2) // C = Y1*O2*Z2
	D := curve.gmul(y2, O1)
	D = curve.gmul(D, z1)               // D = Y2*O1*Z1
	E := curve.gadd(A, B)               // E = A+B
	F := curve.gadd(C, D)               // F = C+D
	if E.Sign() == 0 && F.Sign() == 0 { // E=F=0 means adding two same point
		return curve.doubleJacobian(x1, y1, z1)
	}
	G := curve.gmul(E, z1) // G = E*Z1
	tmp := curve.gmul(F, x2)
	H := curve.gmul(G, y2)
	curve.gaddSelf(H, tmp) // H = F*X2+G*Y2
	z3 = curve.gmul(G, z2) // Z3 = G*Z2
	I := curve.gadd(F, z3) // I = F+Z3
	x3 = curve.gmul(z3, z3)
	x3 = curve.gmul(x3, curve.A)
	tmp = curve.gmul(F, I)
	curve.gaddSelf(x3, tmp)
	tmp = curve.gmul(E, E)
	tmp = curve.gmul(tmp, E)
	curve.gaddSelf(x3, tmp) // X3 = a2*Z3²+F*I+E*E²
	tmp = curve.gmul(I, x3)
	y3 = curve.gmul(G, G)
	y3 = curve.gmul(y3, H)
	curve.gaddSelf(y3, tmp) // Y3 = I*X3+G²*H
	return x3, y3, z3
}

func (curve *F2mCurve) DoubleJacobian(point *EcPoint) *EcPoint {
	x1, y1 := point.X, point.Y
	z1 := zForAffine(x1, y1)
	return curve.affineFromJacobian(curve.doubleJacobian(x1, y1, z1))
}

// see http://hyperelliptic.org/EFD/g12o/auto-shortw-jacobian.html#doubling-dbl-2005-dl
func (curve *F2mCurve) doubleJacobian(x, y, z *big.Int) (*big.Int, *big.Int, *big.Int) {
	A := curve.gmul(x, x) // A = X1²
	B := curve.gmul(A, A) // B = A²
	C := curve.gmul(z, z) // C = Z1²
	D := curve.gmul(C, C) // D = C²
	x3 := curve.gmul(D, D)
	x3 = curve.gmul(x3, curve.B)
	curve.gaddSelf(x3, B)  // X3 = B+a6*D²
	z3 := curve.gmul(x, C) // Z3 = X1*C
	tmp := curve.gmul(B, z3)
	y3 := curve.gmul(y, z)
	curve.gaddSelf(y3, A)
	curve.gaddSelf(y3, z3)
	y3 = curve.gmul(y3, x3)
	curve.gaddSelf(y3, tmp) // Y3 = B*Z3+(A+Y1*Z1+Z3)*X3
	return x3, y3, z3
}

func (curve *F2mCurve) ScalaMultJacobian(point *EcPoint, k []byte) *EcPoint {
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

	return curve.affineFromJacobian(x, y, z)
}

func (curve *F2mCurve) ScalaMultBaseJacobian(k []byte) *EcPoint {
	return curve.ScalaMultJacobian(&EcPoint{curve.X, curve.Y}, k)
}

func (curve *F2mCurve) affineFromJacobian(x, y, z *big.Int) *EcPoint {
	if z.Sign() == 0 {
		return NewPoint()
	}
	zinv := curve.gmulinv(z)
	zinvsq := curve.gmul(zinv, zinv)
	xOut := curve.gmul(x, zinvsq)
	zinvsq = curve.gmul(zinvsq, zinv)
	yOut := curve.gmul(y, zinvsq)
	return &EcPoint{xOut, yOut}
}

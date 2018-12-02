package ecc

import "math/big"

// in this file, we implements short weierstrass elliptic curve arithmetic
// over binary finite field in Lopez-Dahab coordinates: x=X/Z, y=Y/Z²
func (curve *F2mCurve) AddLopezDahab(p1, p2 *EcPoint) *EcPoint {
	x1, y1, x2, y2 := p1.X, p1.Y, p2.X, p2.Y
	z1 := zForAffine(x1, y1)
	z2 := zForAffine(x2, y2)
	return curve.affineFromLopezDahab(curve.addLopezDahab(x1, y1, z1, x2, y2, z2))
}

// see http://hyperelliptic.org/EFD/g12o/auto-shortw-lopezdahab.html#addition-add-2005-dl
func (curve *F2mCurve) addLopezDahab(x1, y1, z1, x2, y2, z2 *big.Int) (*big.Int, *big.Int, *big.Int) {
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

	A := curve.gmul(x1, z2)           // A = X1*Z2
	B := curve.gmul(x2, z1)           // B = X2*Z1
	E := curve.gadd(A, B)             // E = A+B
	G := curve.gmul(z2, z2)
	G = curve.gmul(G, y1)             // G = Y1*Z2²
	H := curve.gmul(z1, z1)
	H = curve.gmul(H, y2)             // H = Y2*Z1²
	I := curve.gadd(G, H)             // I = G+H
	if E.Sign()==0 && I.Sign()==0 {   // E=I=0 means adding two same point
		return curve.doubleLopezDahab(x1, y1, z1)
	}
	C := curve.gmul(A, A)             // C = A²
	D := curve.gmul(B, B)             // D = B²
	F := curve.gadd(C, D)             // F = C+D
	J := curve.gmul(I, E)             // J = I*E

	z3 = curve.gmul(F, z1)
	z3 = curve.gmul(z3, z2)           // Z3 = F*Z1*Z2

	tmp := curve.gadd(H, D)
	tmp = curve.gmul(tmp, A)
	x3 = curve.gadd(C, G)
	x3 = curve.gmul(x3, B)
	curve.gaddSelf(x3, tmp)           // X3 = A*(H+D)+B*(C+G)

	tmp = curve.gmul(A, J)
	y3 = curve.gmul(F, G)
	curve.gaddSelf(y3, tmp)
	y3 = curve.gmul(y3, F)
	tmp = curve.gadd(J, z3)
	tmp = curve.gmul(tmp, x3)
	curve.gaddSelf(y3, tmp)           // Y3 = (A*J+F*G)*F+(J+Z3)*X3
	return x3, y3, z3
}

func (curve *F2mCurve) DoubleLopezDahab(point *EcPoint) *EcPoint {
	x1, y1 := point.X, point.Y
	z1 := zForAffine(x1, y1)
	return curve.affineFromLopezDahab(curve.doubleLopezDahab(x1, y1, z1))
}

// see http://hyperelliptic.org/EFD/g12o/auto-shortw-lopezdahab.html#doubling-dbl-2005-dl
func (curve *F2mCurve) doubleLopezDahab(x, y, z *big.Int) (*big.Int, *big.Int, *big.Int) {
	A := curve.gmul(z, z)          // A = Z1²
	B := curve.gmul(A, A)
	B = curve.gmul(B, curve.B)     // B = a6*A²
	C := curve.gmul(x, x)          // C = X1²
	z3 := curve.gmul(A, C)         // Z3 = A*C
	x3 := curve.gmul(C, C)
	curve.gaddSelf(x3, B)          // X3 = C²+B
	tmp := curve.gmul(curve.A, z3)
	y3 := curve.gmul(y, y)
	curve.gaddSelf(y3, tmp)
	curve.gaddSelf(y3, B)
	y3 = curve.gmul(y3, x3)
	tmp = curve.gmul(z3, B)
	curve.gaddSelf(y3, tmp)        // Y3 = (Y1²+a2*Z3+B)*X3+Z3*B
	return x3, y3, z3
}

func (curve *F2mCurve) ScalaMultLopezDahab(point *EcPoint, k []byte) *EcPoint {
	Bx, By := point.X, point.Y
	Bz := new(big.Int).SetInt64(1)
	x, y, z := new(big.Int), new(big.Int), new(big.Int)

	for _, byte := range k {
		for bitNum := 0; bitNum < 8; bitNum++ {
			x, y, z = curve.doubleLopezDahab(x, y, z)
			if byte&0x80 == 0x80 {
				x, y, z = curve.addLopezDahab(Bx, By, Bz, x, y, z)
			}
			byte <<= 1
		}
	}

	return curve.affineFromLopezDahab(x, y, z)
}

func (curve *F2mCurve) ScalaMultBaseLopezDahab(k []byte) *EcPoint {
	return curve.ScalaMultLopezDahab(&EcPoint{curve.X, curve.Y}, k)
}


func (curve *F2mCurve) affineFromLopezDahab(x, y, z *big.Int) *EcPoint {
	if z.Sign() == 0 {
		return NewPoint()
	}
	zinv := curve.gmulinv(z)
	xOut := curve.gmul(x, zinv)
	zinv = curve.gmul(zinv, zinv)
	yOut := curve.gmul(y, zinv)
	return &EcPoint{xOut, yOut}
}
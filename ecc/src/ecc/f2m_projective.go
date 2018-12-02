package ecc

import "math/big"

// in this file, we implements short weierstrass elliptic curve arithmetic
// over binary finite field in projective coordinates: x=X/Z, y=Y/Z
func (curve *F2mCurve) AddProjective(p1, p2 *EcPoint) *EcPoint {
	x1, y1, x2, y2 := p1.X, p1.Y, p2.X, p2.Y
	z1 := zForAffine(x1, y1)
	z2 := zForAffine(x2, y2)
	return curve.affineFromProjective(curve.addProjective(x1, y1, z1, x2, y2, z2))
}

// see http://hyperelliptic.org/EFD/g12o/auto-shortw-projective.html#addition-add-2008-bl
func (curve *F2mCurve) addProjective(x1, y1, z1, x2, y2, z2 *big.Int) (*big.Int, *big.Int, *big.Int) {
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

	y1z2 := curve.gmul(y1, z2)        // Y1Z2 = Y1*Z2
	x1z2 := curve.gmul(x1, z2)        // X1Z2 = X1*Z2

	A := curve.gmul(z1, y2)
	curve.gaddSelf(A, y1z2)           // A = Y1Z2+Z1*Y2

	B := curve.gmul(z1, x2)
	curve.gaddSelf(B, x1z2)           // B = X1Z2+Z1*X2

	if A.Sign()==0 && B.Sign()==0 {   // A=B=0 means adding two same point
		return curve.doubleProjective(x1, y1, z1)
	}

	AB := curve.gadd(A, B)            // AB = A+B
	C := curve.gmul(B, B)             // C = B²
	D := curve.gmul(z1, z2)           // D = Z1*Z2
	E := curve.gmul(B, C)             // E = B*C

	tmp := curve.gmul(A, AB)
	F := curve.gmul(curve.A, C)
	curve.gaddSelf(F, tmp)
	F = curve.gmul(F, D)
	curve.gaddSelf(F, E)              // F = (A*AB+a2*C)*D+E

	x3 = curve.gmul(B, F)             // X3 = B*F

	tmp = curve.gmul(A, x1z2)
	y3 = curve.gmul(B, y1z2)
	curve.gaddSelf(y3, tmp)
	y3 = curve.gmul(y3, C)
	tmp = curve.gmul(AB, F)
	curve.gaddSelf(y3, tmp)           // Y3 = C*(A*X1Z2+B*Y1Z2)+AB*F

	z3 = curve.gmul(E, D)             // Z3 = E*D

	return x3, y3, z3
}

func (curve *F2mCurve) DoubleProjective(point *EcPoint) *EcPoint {
	x1, y1 := point.X, point.Y
	z1 := zForAffine(x1, y1)
	return curve.affineFromProjective(curve.doubleProjective(x1, y1, z1))
}

// see http://hyperelliptic.org/EFD/g12o/auto-shortw-projective.html#doubling-dbl-2008-bl
func (curve *F2mCurve) doubleProjective(x, y, z *big.Int) (*big.Int, *big.Int, *big.Int) {
	A := curve.gmul(x, x)         // A = X1²
	B := curve.gmul(y, z)
	curve.gaddSelf(B, A)          // B = A+Y1*Z1
	C := curve.gmul(x, z)         // C = X1*Z1
	BC := curve.gadd(B, C)        // BC = B+C
	D := curve.gmul(C, C)         // D = C²
	tmp := curve.gmul(B, BC)
	E := curve.gmul(curve.A, D)
	curve.gaddSelf(E, tmp)        // E = B*BC+a2*D
	x3 := curve.gmul(C, E)        // X3 = C*E
	tmp = curve.gmul(BC, E)       
	y3 := curve.gmul(A, A)        
	y3 = curve.gmul(y3, C)        
	curve.gaddSelf(y3, tmp)       // Y3 = BC*E+A²*C
	z3 := curve.gmul(C, D)        // Z3 = C*D
	return x3, y3, z3
}

func (curve *F2mCurve) ScalaMultProjective(point *EcPoint, k []byte) *EcPoint {
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

func (curve *F2mCurve) ScalaMultBaseProjective(k []byte) *EcPoint {
	return curve.ScalaMultProjective(&EcPoint{curve.X, curve.Y}, k)
}

func (curve *F2mCurve) affineFromProjective(x, y, z *big.Int) *EcPoint {
	if z.Sign() == 0 {
		return NewPoint()
	}
	zinv := curve.gmulinv(z)
	xOut := curve.gmul(x, zinv)
	yOut := curve.gmul(y, zinv)
	return &EcPoint{xOut, yOut}
}
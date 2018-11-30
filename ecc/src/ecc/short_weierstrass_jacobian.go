package ecc

import (
	"math/big"
)

// in this file, we implements short weierstrass elliptic curve arithmetic in Jacobian coordinates
func (curve *FpCurve) AddJacobian(p1, p2 *EcPoint) *EcPoint {
	x1, y1, x2, y2 := p1.X, p1.Y, p2.X, p2.Y
	z1 := zForAffine(x1, y1)
	z2 := zForAffine(x2, y2)
	return curve.affineFromJacobian(curve.addJacobian(x1, y1, z1, x2, y2, z2))
}

// see https://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian.html#addition-add-2007-bl
func (curve *FpCurve) addJacobian(x1, y1, z1, x2, y2, z2 *big.Int) (*big.Int, *big.Int, *big.Int) {
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
	h := new(big.Int).Sub(u2, u1)    // H = U2-U1
	xEqual := h.Sign() == 0
	if h.Sign() == -1 {
		h.Add(h, curve.P)
	}
	i := new(big.Int).Lsh(h, 1)
	i.Mul(i, i)                    // I = (2*H)²
	j := new(big.Int).Mul(h, i)    // J = H*I

	s1 := new(big.Int).Mul(y1, z2)
	s1.Mul(s1, z2z2)               // S1 = Y1*Z2*Z2Z2
	s1.Mod(s1, curve.P)
	s2 := new(big.Int).Mul(y2, z1)
	s2.Mul(s2, z1z1)               // S2 = Y2*Z1*Z1Z1
	s2.Mod(s2, curve.P)
	r := new(big.Int).Sub(s2, s1)
	if r.Sign() == -1 {
		r.Add(r, curve.P)
	}
	yEqual := r.Sign() == 0
	if xEqual && yEqual {          // H=r=0 means adding two same point
		return curve.doubleJacobian(x1, y1, z1)
	}
	r.Lsh(r, 1)                    // r = 2*(S2-S1)
	v := new(big.Int).Mul(u1, i)   // V = U1*I

	x3.Set(r)
	x3.Mul(x3, x3)
	x3.Sub(x3, j)
	x3.Sub(x3, v)
	x3.Sub(x3, v)         // X3 = r²-J-2*V
	x3.Mod(x3, curve.P)

	y3.Set(r)
	v.Sub(v, x3)
	y3.Mul(y3, v)
	s1.Mul(s1, j)
	s1.Lsh(s1, 1)
	y3.Sub(y3, s1)        // Y3 = r*(V-X3)-2*S1*J
	y3.Mod(y3, curve.P)

	z3.Add(z1, z2)
	z3.Mul(z3, z3)
	z3.Sub(z3, z1z1)
	z3.Sub(z3, z2z2)
	z3.Mul(z3, h)         // Z3 = ((Z1+Z2)²-Z1Z1-Z2Z2)*H
	z3.Mod(z3, curve.P)

	return x3, y3, z3
}

func (curve *FpCurve) DoubleJacobian(point *EcPoint) *EcPoint {
	x1, y1 := point.X, point.Y
	z1 := zForAffine(x1, y1)
	return curve.affineFromJacobian(curve.doubleJacobian(x1, y1, z1))
}

// see https://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian.html#doubling-dbl-2007-bl
func (curve *FpCurve) doubleJacobian(x, y, z *big.Int) (*big.Int, *big.Int, *big.Int) {
	XX := new(big.Int).Mul(x, x)      // XX = X1²
	YY := new(big.Int).Mul(y, y)      // YY = Y1²
	ZZ := new(big.Int).Mul(z, z)      // ZZ = Z1²
	YYYY := new(big.Int).Mul(YY, YY)  // YYYY = YY²
	M := new(big.Int).Mul(ZZ, ZZ)
	M.Mul(M, curve.A)
	M.Add(M, XX)
	M.Add(M, XX)
	M.Add(M, XX)                      // M = 3*XX+a*ZZ²
	S := new(big.Int).Mul(x, YY)
	S.Lsh(S, 2)                       // S = 4*X1*YY

	x3 := new(big.Int).Mul(M, M)
	x3.Sub(x3, S)
	x3.Sub(x3, S)                     // X3 = T = M²-2*S
	x3.Mod(x3, curve.P)

	y3 := new(big.Int).Lsh(YYYY, 3)
	S.Sub(S, x3)
	S.Mul(S, M)
	y3.Sub(S, y3)                     // Y3 = M*(S-T)-8*YYYY
	y3.Mod(y3, curve.P)

	z3 := new(big.Int).Mul(y, z)
	z3.Lsh(z3, 1)                     // Z3 = 2*Y1*Z1
	z3.Mod(z3, curve.P)
	return x3, y3, z3
}

func (curve *FpCurve) ScalaMultJacobian(point *EcPoint, k []byte) *EcPoint {
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

func (curve *FpCurve) ScalaMultBaseJacobian(k []byte) *EcPoint {
	return curve.ScalaMultJacobian(&EcPoint{curve.X, curve.Y}, k)
}


func (curve *FpCurve) affineFromJacobian(x, y, z *big.Int) *EcPoint {
	if z.Sign() == 0 {
		return NewPoint()
	}

	zinv := new(big.Int).ModInverse(z, curve.P)
	zinvsq := new(big.Int).Mul(zinv, zinv)

	xOut := new(big.Int).Mul(x, zinvsq)
	xOut.Mod(xOut, curve.P)
	zinvsq.Mul(zinvsq, zinv)
	yOut := new(big.Int).Mul(y, zinvsq)
	yOut.Mod(yOut, curve.P)
	return &EcPoint{xOut, yOut}
}
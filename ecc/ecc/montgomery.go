package ecc

import (
	"fmt"
	"math/big"
	"sync"
)

var (
	initMt                                 sync.Once
	curve25519, m221, m383, curve448, m511 *MtCurve
)

type MtCurve struct {
	P, A, Bx, By, Order, B *big.Int
	Name, Pstr             string
}

// always assume that B=1
// By² ≡ x³ + Ax² + x (mod P)
func (curve *MtCurve) IsOnCurve(x, y *big.Int) bool {
	y2 := new(big.Int).Mul(y, y)
	y2.Mod(y2, curve.P)
	x3 := new(big.Int).Mul(x, x)
	Ax2 := new(big.Int).Mul(x3, curve.A)
	x3.Mul(x3, x)
	x3.Add(x3, Ax2)
	x3.Add(x3, x)
	x3.Mod(x3, curve.P)
	fmt.Println("left : ", y2.String())
	fmt.Println("right: ", x3.String())
	return y2.Cmp(x3) == 0
}

func initCurve25519() {
	p, _ := new(big.Int).SetString("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed", 16)
	bx := new(big.Int).SetInt64(9)
	by, _ := new(big.Int).SetString("20ae19a1b8a086b4e01edd2c7748d14c923d4d7e6d7c61b229e9c5a27eced3d9", 16)
	order, _ := new(big.Int).SetString("1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed", 16)
	a := new(big.Int).SetInt64(486662)
	curve25519 = &MtCurve{}
	curve25519.P = p
	curve25519.Bx = bx
	curve25519.By = by
	curve25519.A = a
	curve25519.B = ONE
	curve25519.Order = order
	curve25519.Name = "Curve25519"
	curve25519.Pstr = "p = 2²⁵⁵ - 19 = 1 mod 4"
}

func initM221() {
	p, _ := new(big.Int).SetString("1ffffffffffffffffffffffffffffffffffffffffffffffffffffffd", 16)
	bx := new(big.Int).SetInt64(4)
	by, _ := new(big.Int).SetString("f7acdd2a4939571d1cef14eca37c228e61dbff10707dc6c08c5056d", 16)
	a := new(big.Int).SetInt64(117050)
	order, _ := new(big.Int).SetString("40000000000000000000000000015a08ed730e8a2f77f005042605b", 16)
	m221 = &MtCurve{}
	m221.P = p
	m221.Bx = bx
	m221.By = by
	m221.A = a
	m221.B = ONE
	m221.Order = order
	m221.Name = "M-221"
	m221.Pstr = "p = 2²²¹ - 3 = 1 mod 4"
}

func initM383() {
	p, _ := new(big.Int).SetString("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff45", 16)
	bx := new(big.Int).SetInt64(12)
	by, _ := new(big.Int).SetString("1ec7ed04aaf834af310e304b2da0f328e7c165f0e8988abd3992861290f617aa1f1b2e7d0b6e332e969991b62555e77e", 16)
	a := new(big.Int).SetInt64(2065150)
	order, _ := new(big.Int).SetString("10000000000000000000000000000000000000000000000006c79673ac36ba6e7a32576f7b1b249e46bbc225be9071d7", 16)
	m383 = &MtCurve{}
	m383.P = p
	m383.Bx = bx
	m383.By = by
	m383.A = a
	m383.B = ONE
	m383.Order = order
	m383.Name = "M-383"
	m383.Pstr = "p = 2³⁸³ - 187 = 1 mod 4"
}

func initM511() {
	p, _ := new(big.Int).SetString("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff45", 16)
	bx := new(big.Int).SetInt64(5)
	by, _ := new(big.Int).SetString("2fbdc0ad8530803d28fdbad354bb488d32399ac1cf8f6e01ee3f96389b90c809422b9429e8a43dbf49308ac4455940abe9f1dbca542093a895e30a64af056fa5", 16)
	a := new(big.Int).SetInt64(530438)
	order, _ := new(big.Int).SetString("100000000000000000000000000000000000000000000000000000000000000017b5feff30c7f5677ab2aeebd13779a2ac125042a6aa10bfa54c15bab76baf1b", 16)
	m511 = &MtCurve{}
	m511.P = p
	m511.Bx = bx
	m511.By = by
	m511.A = a
	m511.B = ONE
	m511.Order = order
	m511.Name = "M-511"
	m511.Pstr = "p = 2⁵¹¹  - 187 = 1 mod 4"
}

func initCurve448() {
	p, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffeffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	bx := new(big.Int).SetInt64(5)
	by, _ := new(big.Int).SetString("7d235d1295f5b1f66c98ab6e58326fcecbae5d34f55545d060f75dc28df3f6edb8027e2346430d211312c4b150677af76fd7223d457b5b1a", 16)
	a := new(big.Int).SetInt64(156326)
	order, _ := new(big.Int).SetString("3fffffffffffffffffffffffffffffffffffffffffffffffffffffff7cca23e9c44edb49aed63690216cc2728dc58f552378c292ab5844f3", 16)
	curve448 = &MtCurve{}
	curve448.P = p
	curve448.Bx = bx
	curve448.By = by
	curve448.A = a
	curve448.Order = order
	curve448.B = ONE
	curve448.Name = "Curve448"
	curve448.Pstr = "p = 2⁴⁴⁸ - 2²²⁴ - 1 = 3 mod 4"
}

func initAllMontgomery() {
	initCurve25519()
	initM221()
	initM383()
	initM511()
	initCurve448()
}

func Curve25519() *MtCurve {
	initMt.Do(initAllMontgomery)
	return curve25519
}

func M221() *MtCurve {
	initMt.Do(initAllMontgomery)
	return m221
}

func M383() *MtCurve {
	initMt.Do(initAllMontgomery)
	return m383
}

func M511() *MtCurve {
	initMt.Do(initAllMontgomery)
	return m511
}

func Curve448() *MtCurve {
	initMt.Do(initAllMontgomery)
	return curve448
}

func (curve *MtCurve) Add(p, q *EcPoint) *EcPoint {
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
			// m = (3Xp² + 2AXp + 1) (2Yp)^-1 mod P
			tmp.Mul(p.X, p.X)                      // Xp²
			tmp.Mul(tmp, new(big.Int).SetInt64(3)) // 3Xp²
			Ax := new(big.Int).Mul(curve.A, p.X)   // AXp
			Ax.Lsh(Ax, 1)                          // 2AXp
			tmp.Add(tmp, Ax)                       // 3Xp² + 2AXp
			tmp.Add(tmp, new(big.Int).SetInt64(1)) // 3Xp² + 2AXp + 1
			m = m.Lsh(p.Y, 1)                      // 2Yp
			m.ModInverse(m, curve.P)               // (2Yp)^-1 mod P
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
	// Xr = (m² - A -Xp - Xq) mod P
	Xr := new(big.Int).Mul(m, m)
	Xr.Sub(Xr, p.X)
	Xr.Sub(Xr, q.X)
	Xr.Sub(Xr, curve.A)
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

func (curve *MtCurve) ScalaMult(p *EcPoint, k []byte) *EcPoint {
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

func (curve *MtCurve) ScalaMultBase(k []byte) *EcPoint {
	return curve.ScalaMult(&EcPoint{curve.Bx, curve.By}, k)
}

//	use montgomery ladder to perform scalar multiply:
//  input point P and scalar k (m bits), return the scalar multiplication product kP
//
//  R0 ← 0
//  R1 ← P
//  for i from m-1 downto 0 do
//   	if ki = 0 then
//       	R1 ← point_add(R0, R1)
//       	R0 ← point_double(R0)
//		else
//       	R0 ← point_add(R0, R1)
//       	R1 ← point_double(R1)
//  return R0
func (curve *MtCurve) nonuniformScalaMultProjective(p *EcPoint, k []byte) *EcPoint {
	x, y, z := p.X, p.Y, zForAffine(p.X, p.Y)
	x0, z0 := new(big.Int), new(big.Int)
	x1, z1 := new(big.Int).Set(p.X), zForAffine(p.X, p.Y)
	for _, b := range k {
		for i := 0; i < 8; i++ {
			if b&0x80 == 0x80 {
				x0, z0 = curve.diffAddProjective(x, z, x0, z0, x1, z1) // R0 = R0 ⊕ R1
				x1, z1 = curve.doubleProjective(x1, z1)                // R1 = 2R1
			} else {
				x1, z1 = curve.diffAddProjective(x, z, x0, z0, x1, z1) // R1 = R0 ⊕ R1
				x0, z0 = curve.doubleProjective(x0, z0)                // R0 = 2R0
			}
			b <<= 1
		}
	}
	x, y, z = curve.recoverY(x, y, x0, z0, x1, z1)
	return curve.affineFromProjective(x, y, z)
}

//  uniformed montgomery ladder, use constant-time conditional swap other than using 'if else'.
//  furthermore, the add and double functions contain some common intermediate results,
//  so they shouldn't called separately but merged.
//
//  input scalar k (l bits, l-1 to 0) and point P, return the scalar multiplication product kP
//  x2, z2 = 1, 0      // initial as O (infinite point)
//  x3, z3 = P.X, 1    // initial as P
//  for i = l−1 down to 0 do
//  	cswap(k(i+1)^k(i), x2, x3)             // k(i) is the i-th bit of scalar k
//  	x2, x3 = double(x2), add(x2, x3)
//  cswap(k(0), x2, x3)
//  return x2
func (curve *MtCurve) ScalaMultProjective(p *EcPoint, k []byte) *EcPoint {
	x1 := new(big.Int).Set(p.X)                                  // x1=p.X, z1=1; x1 is the difference, is always p
	x2, z2 := new(big.Int).SetInt64(1), new(big.Int).SetInt64(0) // x2 initial as O
	x3, z3 := new(big.Int).Set(p.X), zForAffine(p.X, p.Y)        // x3 initial as P
	a24 := new(big.Int).SetInt64((curve.A.Int64() - 2) / 4)
	var m byte = 0 // initial former bit as 0
	for _, b := range k {
		for i := 0; i < 8; i++ {
			ki := (b & 0x80) >> 7 // current bit
			m ^= ki               // the swap argument is current_bit xor former_bit
			cswap(m, x2, x3)
			cswap(m, z2, z3) // conditional swap x2, x3
			m = ki           // set to current bit; in next loop, it become former bit

			A := new(big.Int).Add(x2, z2)
			AA := new(big.Int).Mul(A, A)
			B := new(big.Int).Sub(x2, z2)
			BB := new(big.Int).Mul(B, B)
			E := new(big.Int).Sub(AA, BB)
			C := new(big.Int).Add(x3, z3)
			D := new(big.Int).Sub(x3, z3)
			DA := new(big.Int).Mul(D, A)
			CB := new(big.Int).Mul(C, B)

			x3.Add(DA, CB)
			x3.Mul(x3, x3)
			x3.Mod(x3, curve.P) // x3 = z1*(DA+CB)², z1 is always 1

			z3.Sub(DA, CB)
			z3.Mul(z3, z3)
			z3.Mul(z3, x1)
			z3.Mod(z3, curve.P) // z3 = x1*(DA-CB)²

			x2.Mul(AA, BB)
			x2.Mod(x2, curve.P) // x2 = AA*BB

			z2.Mul(a24, E)
			z2.Add(z2, AA)
			z2.Mul(z2, E)
			z2.Mod(z2, curve.P) // z2 = E*(AA+a24*E)
			b <<= 1
		}
	}
	cswap(m, x2, x3)
	cswap(m, z2, z3)
	x, y, z := curve.recoverY(p.X, p.Y, x2, z2, x3, z3)
	return curve.affineFromProjective(x, y, z)
}

// input and output only x coordinate
func (curve *MtCurve) ScalaMultProjectiveU(x1 *big.Int, k []byte) (x *big.Int) {
	x2, z2 := new(big.Int).SetInt64(1), new(big.Int).SetInt64(0) // x2 initial as O
	x3, z3 := new(big.Int).Set(x1), zForAffine(x1, x1)           // x3 initial as P
	a24 := new(big.Int).SetInt64((curve.A.Int64() - 2) / 4)
	var m byte = 0 // initial former bit as 0
	for _, b := range k {
		for i := 0; i < 8; i++ {
			ki := (b & 0x80) >> 7 // current bit
			m ^= ki               // the swap argument is current_bit xor former_bit
			cswap(m, x2, x3)
			cswap(m, z2, z3) // conditional swap x2, x3
			m = ki           // set to current bit; in next loop, it become former bit
			A := new(big.Int).Add(x2, z2)
			AA := new(big.Int).Mul(A, A)
			B := new(big.Int).Sub(x2, z2)
			BB := new(big.Int).Mul(B, B)
			E := new(big.Int).Sub(AA, BB)
			C := new(big.Int).Add(x3, z3)
			D := new(big.Int).Sub(x3, z3)
			DA := new(big.Int).Mul(D, A)
			CB := new(big.Int).Mul(C, B)
			x3.Add(DA, CB)
			x3.Mul(x3, x3)
			x3.Mod(x3, curve.P) // x3 = z1*(DA+CB)², z1 is always 1
			z3.Sub(DA, CB)
			z3.Mul(z3, z3)
			z3.Mul(z3, x1)
			z3.Mod(z3, curve.P) // z3 = x1*(DA-CB)²
			x2.Mul(AA, BB)
			x2.Mod(x2, curve.P) // x2 = AA*BB
			z2.Mul(a24, E)
			z2.Add(z2, AA)
			z2.Mul(z2, E)
			z2.Mod(z2, curve.P) // z2 = E*(AA+a24*E)
			b <<= 1
		}
	}
	cswap(m, x2, x3)
	cswap(m, z2, z3)

	// convert projective coordinate x2/z2 to affine coordinate
	if z2.Sign() == 0 {
		return new(big.Int)
	}
	z2.ModInverse(z2, curve.P)
	x2.Mul(x2, z2)
	x2.Mod(x2, curve.P)
	return x2
}

// constant-time conditional swap, i.e., if b==1, swap two numbers, if b==0, do not swap.
//    m = mask(b)                   // mask() generate all-0 or all-1 bits according to b
//    v = m and (x0 xor x1)
//    x0 = x0 xor v
//    x1 = x1 xor v
func cswap(b byte, x0, x1 *big.Int) {
	m := ^big.Word(b) + 1 // if b=0, ^b=11111111, ^b+1=0; if b=1, ^b=11111110, ^b+1=11111111
	v := new(big.Int).Xor(x1, x0)
	bits := v.Bits() // changing Bits() will modify the big.Int directly
	for i := range bits {
		bits[i] = bits[i] & m
	}
	x0.Xor(x0, v)
	x1.Xor(x1, v)
}

//  Give point P(xP, yP), through montgomery ladder, we can get scalar multiplication product of
//  kP = (Xₖ, Zₖ) and (k+1)P = (Xₖ₊₁, Zₖ₊₁) in projective coordinates. One can compute the y-coordinate of Q,
//  if (xP : yP : 1) = P , (XQ : ZQ) = x(Q), and (X⊕ : Z⊕) = x(P⊕Q) are known. Note that (k+1)P = kP + P,
//  so we can recover the y-coordinate of the scalar multiplication result.
//  Algorithm 5 in "Montgomery curves and their arithmetic: The case of large characteristic fields".
//  The parameters of this function:
//     xP, yP: the affine coordinate of point P
//     Xq, Zq: the projective coordinate of kP
//     Xa, Za: the projective coordinate of (k+1)P
//  return (X', Y', Z'), the new projective coordinates of kP
func (curve *MtCurve) recoverY(xP, yP, Xq, Zq, Xa, Za *big.Int) (X, Y, Z *big.Int) {
	v1 := new(big.Int).Mul(xP, Zq)
	v2 := new(big.Int).Add(Xq, v1)
	v3 := new(big.Int).Sub(Xq, v1)
	v3.Mul(v3, v3)
	v3.Mul(v3, Xa)
	v1.Mul(curve.A, Zq)
	v1.Lsh(v1, 1)
	v2.Add(v2, v1)
	v4 := new(big.Int).Mul(xP, Xq)
	v4.Add(v4, Zq)
	v2.Mul(v2, v4)
	v1.Mul(v1, Zq)
	v2.Sub(v2, v1)
	v2.Mul(v2, Za)
	Y = new(big.Int).Sub(v2, v3)
	Y.Mod(Y, curve.P)
	v1.Mul(curve.B, yP)
	v1.Lsh(v1, 1)
	v1.Mul(v1, Zq)
	v1.Mul(v1, Za)
	X = new(big.Int).Mul(v1, Xq)
	X.Mod(X, curve.P)
	Z = new(big.Int).Mul(v1, Zq)
	Z.Mod(Z, curve.P)
	return
}

func (curve *MtCurve) ScalaMultBaseProjective(k []byte) *EcPoint {
	return curve.ScalaMultProjective(&EcPoint{curve.Bx, curve.By}, k)
}

func (curve *MtCurve) nonuniformScalaMultBaseProjective(k []byte) *EcPoint {
	return curve.nonuniformScalaMultProjective(&EcPoint{curve.Bx, curve.By}, k)
}

func (curve *MtCurve) affineFromProjectiveX(x, z *big.Int) (xOut *big.Int) {
	if z.Sign() == 0 {
		return new(big.Int)
	}
	zinv := new(big.Int).ModInverse(z, curve.P)
	xOut = new(big.Int).Mul(x, zinv)
	xOut.Mod(xOut, curve.P)
	return
}

func (curve *MtCurve) affineFromProjective(x, y, z *big.Int) *EcPoint {
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

func (curve *MtCurve) doubleProjective(x1, z1 *big.Int) (x3, z3 *big.Int) {
	// X3 = (X1²-Z1²)²
	// Z3 = 4*X1*Z1*(X1²+a*X1*Z1+Z1²)
	x3 = new(big.Int).Mul(x1, x1)
	tmp := new(big.Int).Mul(z1, z1)
	z3 = new(big.Int).Add(x3, tmp)
	x3.Sub(x3, tmp)
	x3.Mul(x3, x3)
	x3.Mod(x3, curve.P)
	tmp.Mul(x1, z1)
	z3.Mul(z3, tmp)
	tmp.Mul(tmp, tmp)
	tmp.Mul(tmp, curve.A)
	z3.Add(z3, tmp)
	z3.Lsh(z3, 2)
	z3.Mod(z3, curve.P)
	return
}

// differential adding is to compute P+Q, given not only P nad Q, but also P-Q.
// (x1, z1) is the differential, (x2, z2) and (x3, z3) are the points to be added
func (curve *MtCurve) diffAddProjective(x1, z1, x2, z2, x3, z3 *big.Int) (x5, z5 *big.Int) {
	x5, z5 = new(big.Int), new(big.Int)
	if z2.Sign() == 0 {
		x5.Set(x3)
		z5.Set(z3)
		return
	}
	if z3.Sign() == 0 {
		x5.Set(x2)
		z5.Set(z2)
		return
	}
	// A = X2+Z2
	// B = X2-Z2
	// C = X3+Z3
	// D = X3-Z3
	// DA = D*A
	// CB = C*B
	// X5 = Z1*(DA+CB)²
	// Z5 = X1*(DA-CB)²
	A := new(big.Int).Add(x2, z2)
	B := new(big.Int).Sub(x2, z2)
	C := new(big.Int).Add(x3, z3)
	D := new(big.Int).Sub(x3, z3)
	DA := new(big.Int).Mul(D, A)
	CB := new(big.Int).Mul(C, B)
	x5 = new(big.Int).Add(DA, CB)
	x5.Mul(x5, x5)
	x5.Mul(x5, z1)
	x5.Mod(x5, curve.P)
	z5 = new(big.Int).Sub(DA, CB)
	z5.Mul(z5, z5)
	z5.Mul(z5, x1)
	z5.Mod(z5, curve.P)
	return
}

// form 1:
// the birational equivalence between montgomery curve and twisted edwards curve:
// to birational map montgomery curve (Bv² = u³ + Au² + u) to twisted edwards curve: (ax² + y² = 1 + Dx²y²)
// coefficient:            point convert:
//    a = (A+2)/B             x = u/v
//    D = (A-2)/B             y = ±(u-1)/(u+1)
//
// if we want to convert montgomery curve to edwards curve (i.e., a=1), we need B=A+2,
// thus, to convert (v² = u³ + Au² + u) to (x² + y² = 1 + Dx²y²),
// we first convert (v² = u³ + Au² + u) to (Bv'² = u³ + Au² + u), by setting v'=v/sqrt(B),
// then to edwards curve, the coefficients become: a=1, D=(A-2)/(A+2),
// and point conversion become: x = u/v' = sqrt(B)*u/v
//
// similarly, if we want a twisted edwards curve, i.e., a = (A+2)/B ≠ 1,
// we need to choose the proper value of B.
// however, we must make sure that sqrt(B) is existed in the finite field Fp.
//
// example:
// to convert curve25519 (v² = u³ + 486662u² + u) to twisted edwards curve with a=-1,
// we need B = -(A+2) = -486664, then D = (A-2)/B = -486660/486664 = -121665/121666,
// i.e.  -x² + y² = 1 - 121665/121666 x²y²,
// with x = sqrt(-486664)*u/v, y = ±(u-1)/(u+1)
func (curve *MtCurve) ToEdwardsCurveForm1(a *big.Int) (*EdCurve, *big.Int) {
	edwards := new(EdCurve)
	edwards.Name = "Edwards form1 ed.a=(mt.A+2)/mt.B of " + curve.Name
	edwards.A = a
	Aadd2 := new(big.Int).Add(curve.A, TWO)
	Asub2 := new(big.Int).Sub(curve.A, TWO)
	B := new(big.Int).Div(Aadd2, a)
	edwards.D = ModFraction(Asub2, B, curve.P)
	edwards.P = curve.P
	edwards.Order = curve.Order
	sqrtB := new(big.Int).ModSqrt(B, curve.P)
	if sqrtB != nil {
		p1, _ := curve.ToEdwardsPointForm1(sqrtB, &EcPoint{curve.Bx, curve.By})
		edwards.Bx = p1.X
		edwards.By = p1.Y
	}
	return edwards, sqrtB
}

// form 1
// x = sqrt(B)*u/v, y = ±(u-1)/(u+1)
func (curve *MtCurve) ToEdwardsPointForm1(sqrtB *big.Int, p *EcPoint) (p1, p2 *EcPoint) {
	uAdd1 := new(big.Int).Add(p.X, ONE)
	uSub1 := new(big.Int).Sub(p.X, ONE)
	p1, p2 = NewPoint(), NewPoint()
	p1.Y = ModFraction(uSub1, uAdd1, curve.P) // y = (u-1)/(u+1)
	uSub1.Neg(uSub1)
	p2.Y = ModFraction(uSub1, uAdd1, curve.P) // y = (1-u)/(u+1)
	p1.X = ModFraction(p.X, p.Y, curve.P)
	p1.X.Mul(p1.X, sqrtB) // x = sqrt(B) u/v
	p1.X.Mod(p1.X, curve.P)
	p2.X.Set(p1.X)
	return
}

// form 2:
// to birational map montgomery curve (Bv² = u³ + Au² + u) to twisted edwards curve: (ax² + y² = 1 + Dx²y²)
// coefficient:            point convert:
//    a = (A-2)/B             x = u/v
//    D = (A+2)/B             y = ±(u+1)/(u-1)
//
// if we want to convert montgomery curve to edwards curve (i.e., a=1), we need B=A-2,
// thus, to convert (v² = u³ + Au² + u) to (x² + y² = 1 + Dx²y²),
// we first convert (v² = u³ + Au² + u) to (Bv'² = u³ + Au² + u), by setting v'=v/sqrt(B),
// then to edwards curve, the coefficients become: a=1, D=(A+2)/(A-2),
// and point conversion become: x = u/v' = sqrt(B)*u/v
//
// similarly, if we want a twisted edwards curve, i.e., a = (A-2)/B ≠ 1,
// we need to choose the proper value of B.
// however, we must make sure that sqrt(B) is existed in the finite field Fp.
//
// example:
// to convert curve448 (v² = u³ + 156326² + u) to edwards curve with a=1,
// we need B = A-2 = 156324, then D = (A+2)/B = 156328/156324 = 39082/39081,
// i.e.  x² + y² = 1 + 39082/39081 x²y²,
// with x = sqrt(156324)*u/v, y = ±(u+1)/(u-1)
func (curve *MtCurve) ToEdwardsCurveForm2(a *big.Int) (*EdCurve, *big.Int) {
	edwards := new(EdCurve)
	edwards.Name = "Edwards form1 ed.a=(mt.A-2)/mt.B of " + curve.Name
	edwards.A = a
	Aadd2 := new(big.Int).Add(curve.A, TWO)
	Asub2 := new(big.Int).Sub(curve.A, TWO)
	B := new(big.Int).Div(Asub2, a)
	edwards.D = ModFraction(Aadd2, B, curve.P)
	edwards.P = curve.P
	edwards.Order = curve.Order
	sqrtB := new(big.Int).ModSqrt(B, curve.P)
	if sqrtB != nil {
		p1, _ := curve.ToEdwardsPointForm1(sqrtB, &EcPoint{curve.Bx, curve.By})
		edwards.Bx = p1.X
		edwards.By = p1.Y
	}
	return edwards, sqrtB
}

// form 2
// x = sqrt(B)*u/v, y = ±(u+1)/(u-1)
func (curve *MtCurve) ToEdwardsPointForm2(sqrtB *big.Int, p *EcPoint) (p1, p2 *EcPoint) {
	uAdd1 := new(big.Int).Add(p.X, ONE)
	uSub1 := new(big.Int).Sub(p.X, ONE)
	p1, p2 = NewPoint(), NewPoint()
	p1.Y = ModFraction(uAdd1, uSub1, curve.P) // y = (u+1)/(u-1)
	uSub1.Neg(uSub1)
	p2.Y = ModFraction(uAdd1, uSub1, curve.P) // y = (u+1)/(1-u)
	p1.X = ModFraction(p.X, p.Y, curve.P)
	p1.X.Mul(p1.X, sqrtB) // x = sqrt(B) u/v
	p1.X.Mod(p1.X, curve.P)
	p2.X.Set(p1.X)
	return
}

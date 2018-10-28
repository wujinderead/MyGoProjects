package ecc

import (
	"fmt"
	"math/big"
	"sync"
)

var (
	initEd                                                       sync.Once
	e222, curve1174, ed25519, e382, curve41417, edwards448, e521 *EdCurve
)

type EdCurve struct {
	// order is the order of base point
	A, P, D, Bx, By, Order *big.Int
	Name, Pstr             string
}

// ax² + y² ≡ 1 + Dx²y² (mod P)
func (curve *EdCurve) IsOnCurve(x, y *big.Int) bool {
	y2 := new(big.Int).Mul(y, y)
	x2 := new(big.Int).Mul(x, x)
	right := new(big.Int).Mul(x2, y2)
	right.Mul(right, curve.D)
	right.Add(right, ONE)
	right.Mod(right, curve.P)
	x2.Mul(curve.A, x2)
	left := new(big.Int).Add(x2, y2)
	left.Mod(left, curve.P)
	fmt.Println("left : ", left.String())
	fmt.Println("right: ", right.String())
	return left.Cmp(right) == 0
}

func initCurve1174() {
	p, _ := new(big.Int).SetString("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7", 16)
	bx, _ := new(big.Int).SetString("37fbb0cea308c479343aee7c029a190c021d96a492ecd6516123f27bce29eda", 16)
	by, _ := new(big.Int).SetString("6b72f82d47fb7cc6656841169840e0c4fe2dee2af3f976ba4ccb1bf9b46360e", 16)
	d := new(big.Int).SetInt64(-1174)
	order, _ := new(big.Int).SetString("1fffffffffffffffffffffffffffffff77965c4dfd307348944d45fd166c971", 16)
	curve1174 = &EdCurve{}
	curve1174.P = p
	curve1174.Bx = bx
	curve1174.By = by
	curve1174.D = d
	curve1174.A = ONE
	curve1174.Order = order
	curve1174.Name = "Curve1174"
	curve1174.Pstr = "p = 2²⁵¹ - 9 = 3 mod 4"
}

func initE222() {
	p, _ := new(big.Int).SetString("3fffffffffffffffffffffffffffffffffffffffffffffffffffff8b", 16)
	bx, _ := new(big.Int).SetString("19b12bb156a389e55c9768c303316d07c23adab3736eb2bc3eb54e51", 16)
	by := new(big.Int).SetInt64(28)
	d := new(big.Int).SetInt64(160102)
	order, _ := new(big.Int).SetString("ffffffffffffffffffffffffffff70cbc95e932f802f31423598cbf", 16)
	e222 = &EdCurve{}
	e222.P = p
	e222.Bx = bx
	e222.By = by
	e222.D = d
	e222.A = ONE
	e222.Order = order
	e222.Name = "E-222"
	e222.Pstr = "p = 2²²² - 117 = 3 mod 4"
}

func initE382() {
	p, _ := new(big.Int).SetString("3fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff97", 16)
	bx, _ := new(big.Int).SetString("196f8dd0eab20391e5f05be96e8d20ae68f840032b0b64352923bab85364841193517dbce8105398ebc0cc9470f79603", 16)
	by := new(big.Int).SetInt64(17)
	d := new(big.Int).SetInt64(-67254)
	order, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffd5fb21f21e95eee17c5e69281b102d2773e27e13fd3c9719", 16)
	e382 = &EdCurve{}
	e382.P = p
	e382.Bx = bx
	e382.By = by
	e382.D = d
	e382.A = ONE
	e382.Order = order
	e382.Name = "E-382"
	e382.Pstr = "p = 2³⁸² - 105 = 3 mod 4"
}

func initCurve41417() {
	p, _ := new(big.Int).SetString("3fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffef", 16)
	bx, _ := new(big.Int).SetString("1a334905141443300218c0631c326e5fcd46369f44c03ec7f57ff35498a4ab4d6d6ba111301a73faa8537c64c4fd3812f3cbc595", 16)
	by := new(big.Int).SetInt64(34)
	d := new(big.Int).SetInt64(3617)
	order, _ := new(big.Int).SetString("7ffffffffffffffffffffffffffffffffffffffffffffffffffeb3cc92414cf706022b36f1c0338ad63cf181b0e71a5e106af79", 16)
	curve41417 = &EdCurve{}
	curve41417.P = p
	curve41417.Bx = bx
	curve41417.By = by
	curve41417.D = d
	curve41417.A = ONE
	curve41417.Order = order
	curve41417.Name = "Curve41417"
	curve41417.Pstr = "p = 2⁴¹⁴ - 17 = 3 mod 4"
}

func initEdwards448() {
	p, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffeffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	bx, _ := new(big.Int).SetString("297ea0ea2692ff1b4faff46098453a6a26adf733245f065c3c59d0709cecfa96147eaaf3932d94c63d96c170033f4ba0c7f0de840aed939f", 16)
	by := new(big.Int).SetInt64(19)
	d := new(big.Int).SetInt64(-39081)
	order, _ := new(big.Int).SetString("3fffffffffffffffffffffffffffffffffffffffffffffffffffffff7cca23e9c44edb49aed63690216cc2728dc58f552378c292ab5844f3", 16)
	edwards448 = &EdCurve{}
	edwards448.P = p
	edwards448.Bx = bx
	edwards448.By = by
	edwards448.D = d
	edwards448.A = ONE
	edwards448.Order = order
	edwards448.Name = "Ed448-Goldilocks"
	edwards448.Pstr = "p = 2⁴⁴⁸ - 2²²⁴ - 1 = 3 mod 4"
}

func initEd25519() {
	p, _ := new(big.Int).SetString("7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed", 16)
	bx, _ := new(big.Int).SetString("15112221349535400772501151409588531511454012693041857206046113283949847762202", 10)
	by, _ := new(big.Int).SetString("46316835694926478169428394003475163141307993866256225615783033603165251855960", 10)
	d, _ := new(big.Int).SetString("37095705934669439343138083508754565189542113879843219016388785533085940283555", 10)
	order, _ := new(big.Int).SetString("1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed", 16)
	ed25519 = &EdCurve{}
	ed25519.P = p
	ed25519.Bx = bx
	ed25519.By = by
	ed25519.D = d
	ed25519.A = NEG_ONE
	ed25519.Order = order
	ed25519.Name = "Ed25519"
	ed25519.Pstr = "p = 2²⁵⁵-19 = 1 mod 4"
}

func initE521() {
	p, _ := new(big.Int).SetString("1ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)
	bx, _ := new(big.Int).SetString("752cb45c48648b189df90cb2296b2878a3bfd9f42fc6c818ec8bf3c9c0c6203913f6ecc5ccc72434b1ae949d568fc99c6059d0fb13364838aa302a940a2f19ba6c", 16)
	by := new(big.Int).SetInt64(12)
	d := new(big.Int).SetInt64(-376014)
	order, _ := new(big.Int).SetString("7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd15b6c64746fc85f736b8af5e7ec53f04fbd8c4569a8f1f4540ea2435f5180d6b", 16)
	e521 = &EdCurve{}
	e521.P = p
	e521.Bx = bx
	e521.By = by
	e521.D = d
	e521.A = ONE
	e521.Order = order
	e521.Name = "E-521"
	e521.Pstr = "p = 2⁵²¹ - 1 = 3 mod 4"
}

func initAllEdwards() {
	initCurve1174()
	initCurve41417()
	initE222()
	initE382()
	initE521()
	initEdwards448()
	initEd25519()
}

func Curve1174() *EdCurve {
	initEd.Do(initAllEdwards)
	return curve1174
}

func Curve41417() *EdCurve {
	initEd.Do(initAllEdwards)
	return curve41417
}

func E222() *EdCurve {
	initEd.Do(initAllEdwards)
	return e222
}

func E382() *EdCurve {
	initEd.Do(initAllEdwards)
	return e382
}

func E521() *EdCurve {
	initEd.Do(initAllEdwards)
	return e521
}

func Edwards448() *EdCurve {
	initEd.Do(initAllEdwards)
	return edwards448
}

func Ed25519() *EdCurve {
	initEd.Do(initAllEdwards)
	return ed25519
}

func (curve *EdCurve) Add(p, q *EcPoint) *EcPoint {
	if p.Equals(Infinity) {
		return q.Copy()
	}
	if q.Equals(Infinity) {
		return p.Copy()
	}
	x1, x2, y1, y2 := p.X, q.X, p.Y, q.Y
	if y1.Cmp(y2) == 0 && new(big.Int).Add(x1, x2).Cmp(curve.P) == 0 {
		return NewPoint()
	}
	x1y2 := new(big.Int).Mul(x1, y2) // x1y2
	x2y1 := new(big.Int).Mul(x2, y1) // x2y1
	y1y2 := new(big.Int).Mul(y1, y2) // y1y2
	x1x2 := new(big.Int).Mul(x1, x2) // x1x2
	one := new(big.Int).SetInt64(1)
	dx1x2y2y2 := new(big.Int).Mul(x1x2, y1y2) // x1x2y1y2
	dx1x2y2y2.Mul(dx1x2y2y2, curve.D)         // dx1x2y1y2
	x3 := new(big.Int).Add(one, dx1x2y2y2)    // 1 + dx1x2y1y2
	x3.ModInverse(x3, curve.P)                // (1 + dx1x2y1y2)^-1
	x1y2.Add(x1y2, x2y1)                      // x1y2 + x2y1
	x3.Mul(x3, x1y2)                          // (x1y2 + x2y1)(1 + dx1x2y1y2)^-1
	x3.Mod(x3, curve.P)
	y3 := new(big.Int).Sub(one, dx1x2y2y2) // 1 - dx1x2y1y2
	y3.ModInverse(y3, curve.P)             // (1 - dx1x2y1y2)^-1
	x1x2.Mul(x1x2, curve.A)                // ax1x2
	y1y2.Sub(y1y2, x1x2)                   // y1y2 - ax1x2
	y3.Mul(y3, y1y2)                       // (y1y2 - ax1x2)(1 - dx1x2y1y2)^-1
	y3.Mod(y3, curve.P)
	return &EcPoint{x3, y3}
}

func (curve *EdCurve) ScalaMult(p *EcPoint, k []byte) *EcPoint {
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

func (curve *EdCurve) ScalaMultBase(k []byte) *EcPoint {
	return curve.ScalaMult(&EcPoint{curve.Bx, curve.By}, k)
}

func (curve *EdCurve) affineFromProjective(x, y, z *big.Int) *EcPoint {
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

func (curve *EdCurve) Double(p *EcPoint) *EcPoint {
	z1 := zForAffine(p.X, p.Y)
	return curve.affineFromProjective(curve.doubleProjective(p.X, p.Y, z1))
}

func (curve *EdCurve) doubleProjective(x1, y1, z1 *big.Int) (*big.Int, *big.Int, *big.Int) {
	B := new(big.Int).Add(x1, y1)
	B.Mul(B, B)
	C := new(big.Int).Mul(x1, x1)
	D := new(big.Int).Mul(y1, y1)
	E := new(big.Int).Mul(curve.A, C)
	F := new(big.Int).Add(E, D)
	H := new(big.Int).Mul(z1, z1)
	H.Lsh(H, 1)
	J := new(big.Int).Sub(F, H)
	x3 := new(big.Int).Sub(B, C)
	x3.Sub(x3, D)
	x3.Mul(x3, J)
	y3 := new(big.Int).Sub(E, D)
	y3.Mul(F, y3)
	z3 := new(big.Int).Mul(F, J)

	x3.Mod(x3, curve.P)
	y3.Mod(y3, curve.P)
	z3.Mod(z3, curve.P)
	return x3, y3, z3
}

func (curve *EdCurve) addProjective(x1, y1, z1, x2, y2, z2 *big.Int) (*big.Int, *big.Int, *big.Int) {
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

	A := new(big.Int).Mul(z1, z2)
	B := new(big.Int).Mul(A, A)
	C := new(big.Int).Mul(x1, x2)
	D := new(big.Int).Mul(y1, y2)
	E := new(big.Int).Mul(C, D)
	E.Mul(E, curve.D)
	F := new(big.Int).Sub(B, E)
	G := new(big.Int).Add(B, E)
	x1y2 := new(big.Int).Mul(x1, y2)
	x2y1 := new(big.Int).Mul(x2, y1)
	x1y2.Add(x1y2, x2y1)
	x3.Mul(A, F)
	x3.Mul(x3, x1y2)

	C.Mul(curve.A, C)
	D.Sub(D, C)
	y3.Mul(A, G)
	y3.Mul(y3, D)
	z3.Mul(F, G)

	x3.Mod(x3, curve.P)
	y3.Mod(y3, curve.P)
	z3.Mod(z3, curve.P)
	return x3, y3, z3
}

func (curve *EdCurve) AddProjective(p1, p2 *EcPoint) *EcPoint {
	z1 := zForAffine(p1.X, p1.Y)
	z2 := zForAffine(p2.X, p2.Y)
	return curve.affineFromProjective(curve.addProjective(p1.X, p1.Y, z1, p2.X, p2.Y, z2))
}

func (curve *EdCurve) ScalaMultProjective(p *EcPoint, k []byte) *EcPoint {
	Bz := new(big.Int).SetInt64(1)
	x, y, z := new(big.Int), new(big.Int), new(big.Int)

	for _, abyte := range k {
		for bitNum := 0; bitNum < 8; bitNum++ {
			x, y, z = curve.doubleProjective(x, y, z)
			if abyte&0x80 == 0x80 {
				x, y, z = curve.addProjective(p.X, p.Y, Bz, x, y, z)
			}
			abyte <<= 1
		}
	}

	return curve.affineFromProjective(x, y, z)
}

func (curve *EdCurve) ScalaMultBaseProjective(k []byte) *EcPoint {
	return curve.ScalaMultProjective(&EcPoint{curve.Bx, curve.By}, k)
}

// from edwards "ax² + y² ≡ 1 + Dx²y² mod p" to montgomery: "Bv² ≡ u³ + Au² + u mod p"
//
// form1: A=2(a+d)/(a-d)   B=4/(a-d)
// form2: A=2(a+d)/(d-a)   B=4/(d-a)
func (curve *EdCurve) ToMontgomeryCurveForm1() (*MtCurve, *big.Int) {
	a_d := new(big.Int).Sub(curve.A, curve.D)   // a-d
	a_d.ModInverse(a_d, curve.P)                // 1/(a-d)
	montA := new(big.Int).Add(curve.A, curve.D) //a+d
	montA.Lsh(montA, 1)                         // 2(a+d)
	montA.Mul(montA, a_d)                       // 2(a+d)/(a-d)
	montA.Mod(montA, curve.P)
	montB := new(big.Int).Lsh(a_d, 2) // 4/(a-d)
	montB.Mod(montB, curve.P)

	mont := &MtCurve{}
	mont.P = curve.P
	mont.A = montA
	mont.B = montB
	mont.Name = "Montgomery form of " + curve.Name
	mont.Order = curve.Order

	sqrtB := new(big.Int).ModSqrt(montB, curve.P)
	if sqrtB != nil {
		p1, _ := curve.ToMontgomeryPointForm1(sqrtB, &EcPoint{curve.Bx, curve.By})
		mont.Bx = p1.X
		mont.By = p1.Y
	}
	return mont, sqrtB
}

// form1
// mont.(u, v) = ( (1+y)/(1-y) , sqrt(mont.B) * u/x )
// mont.(u, v) = ( (1-y)/(1+y) , sqrt(mont.B) * u/x )
func (curve *EdCurve) ToMontgomeryPointForm1(sqrtB *big.Int, p *EcPoint) (p1, p2 *EcPoint) {
	oneSubY := new(big.Int).Sub(ONE, p.Y) // 1-y
	oneAddY := new(big.Int).Add(ONE, p.Y) // 1+y
	p1, p2 = NewPoint(), NewPoint()
	p1.X = ModFraction(oneAddY, oneSubY, curve.P) // (1+y)/(1-y)
	p1.Y = ModFraction(p1.X, p.X, curve.P)        // u/x
	p1.Y.Mul(p1.Y, sqrtB)                         // sqrtB * u/x
	p1.Y.Mod(p1.Y, curve.P)

	p2.X = ModFraction(oneSubY, oneAddY, curve.P) // (1-y)/(1+y)
	p2.Y = ModFraction(p2.X, p.X, curve.P)        // u/x
	p2.Y.Mul(p2.Y, sqrtB)                         // sqrtB * u/x
	p2.Y.Mod(p2.Y, curve.P)
	return
}

// from edwards "ax² + y² ≡ 1 + Dx²y² mod p" to montgomery: "Bv² ≡ u³ + Au² + u mod p"
//
// form1: A=2(a+d)/(a-d)   B=4/(a-d)
// form2: A=2(a+d)/(d-a)   B=4/(d-a)
func (curve *EdCurve) ToMontgomeryCurveForm2() (*MtCurve, *big.Int) {
	a_d := new(big.Int).Sub(curve.D, curve.A)   // d-a
	a_d.ModInverse(a_d, curve.P)                // 1/(d-a)
	montA := new(big.Int).Add(curve.A, curve.D) //a+d
	montA.Lsh(montA, 1)                         // 2(a+d)
	montA.Mul(montA, a_d)                       // 2(a+d)/(d-a)
	montA.Mod(montA, curve.P)
	montB := new(big.Int).Lsh(a_d, 2) // 4/(d-a)
	montB.Mod(montB, curve.P)

	mont := &MtCurve{}
	mont.P = curve.P
	mont.A = montA
	mont.B = montB
	mont.Name = "Montgomery form of " + curve.Name
	mont.Order = curve.Order

	sqrtB := new(big.Int).ModSqrt(montB, curve.P)
	if sqrtB != nil {
		p1, _ := curve.ToMontgomeryPointForm1(sqrtB, &EcPoint{curve.Bx, curve.By})
		mont.Bx = p1.X
		mont.By = p1.Y
	}
	return mont, sqrtB
}

// form2
// mont.(u, v) = ( (y+1)/(y-1) , sqrt(mont.B) * u/x )
// mont.(u, v) = ( (y-1)/(y+1) , sqrt(mont.B) * u/x )
func (curve *EdCurve) ToMontgomeryPointForm2(sqrtB *big.Int, p *EcPoint) (p1, p2 *EcPoint) {
	yAddOne := new(big.Int).Add(p.Y, ONE) // y+1
	ySubOne := new(big.Int).Sub(p.Y, ONE) // y-1
	p1, p2 = NewPoint(), NewPoint()
	p1.X = ModFraction(yAddOne, ySubOne, curve.P) // (y+1)/(y-1)
	p1.Y = ModFraction(p1.X, p.X, curve.P)        // u/x
	p1.Y.Mul(p1.Y, sqrtB)                         // sqrtB * u/x
	p1.Y.Mod(p1.Y, curve.P)

	p2.X = ModFraction(ySubOne, yAddOne, curve.P) // (y-1)/(y+1)
	p2.Y = ModFraction(p2.X, p.X, curve.P)        // u/x
	p2.Y.Mul(p2.Y, sqrtB)                         // sqrtB * u/x
	p2.Y.Mod(p2.Y, curve.P)
	return
}

package ecc

import (
	"math/big"
	"unsafe"
)

var (
	uintlen = unsafe.Sizeof(uint(1))

	x25519base = [32]byte{9} // x=9 in little-endian as {9, 0, ..., 0}
	a24_25519  = new(big.Int).SetInt64((Curve25519().A.Int64() - 2) / 4)
	p25519     = Curve25519().P
	zero25519  = [32]byte{}

	x448base = [56]byte{5} // x=5 in little-endian as {5, 0, ..., 0}
	a24_448  = new(big.Int).SetInt64((Curve448().A.Int64() - 2) / 4)
	p448     = Curve448().P
	zero448  = [56]byte{}
)

// scala multiplication in rfc7748, the input and output are both little endian
func Curve25519ScalaMultBase(dst, in *[32]byte) {
	Curve25519ScalaMult(dst, in, &x25519base)
}

func Curve25519ScalaMult(dst, in, base *[32]byte) {
	// process in
	var e [32]byte
	copy(e[:], in[:])
	e[0] &= 248
	e[31] &= 127
	e[31] |= 64

	// set base point to *big.Int
	c := make([]uintptr, 3)
	c[0] = uintptr(unsafe.Pointer(base))
	c[1] = uintptr(len(base)) / uintlen
	c[2] = c[1]
	bw := *(*[]big.Word)(unsafe.Pointer(&c[0]))
	x1 := new(big.Int).SetBits(bw)

	// the same as (*MtCurve).ScalaMultProjectiveU1(x1 *big.Int, k []byte), except that k is little endian
	x2, z2 := new(big.Int).SetInt64(1), new(big.Int).SetInt64(0) // x2 initial as O
	x3, z3 := new(big.Int).Set(x1), zForAffine(x1, x1)           // x3 initial as P
	var m byte = 0                                               // initial former bit as 0
	for i := len(e) - 1; i >= 0; i-- {                           // the input is little-endian
		b := e[i]
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
			x3.Mod(x3, p25519) // x3 = z1*(DA+CB)², z1 is always 1
			z3.Sub(DA, CB)
			z3.Mul(z3, z3)
			z3.Mul(z3, x1)
			z3.Mod(z3, p25519) // z3 = x1*(DA-CB)²
			x2.Mul(AA, BB)
			x2.Mod(x2, p25519) // x2 = AA*BB
			z2.Mul(a24_25519, E)
			z2.Add(z2, AA)
			z2.Mul(z2, E)
			z2.Mod(z2, p25519) // z2 = E*(AA+a24*E)
			b <<= 1
		}
	}
	cswap(m, x2, x3)
	cswap(m, z2, z3)

	// convert projective coordinate x2/z2 to affine coordinate
	copy(dst[:], zero25519[:]) // clear dst first
	if z2.Sign() == 0 {
		return // the result is zero, just return 0
	}
	z2.ModInverse(z2, p25519)
	x2.Mul(x2, z2)
	x2.Mod(x2, p25519)

	// from big.Int to little endian bytes
	words := x2.Bits()
	byters := *(*[]byte)(unsafe.Pointer(&words))
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&byters)) + uintptr(8))) = len(words) * int(uintlen)
	copy(dst[:], byters)
}

func Curve448ScalaMultBase(dst, in *[56]byte) {
	Curve448ScalaMult(dst, in, &x448base)
}

// the only differences between Curve448ScalaMult and Curve448ScalaMult are:
// (1). prime field; (2). a24 value; (3). process input scala; (4). input and output length.
func Curve448ScalaMult(dst, in, base *[56]byte) {
	// process in
	var e [56]byte
	copy(e[:], in[:])
	e[0] &= 252
	e[55] |= 128

	// set base point to *big.Int
	c := make([]uintptr, 3)
	c[0] = uintptr(unsafe.Pointer(base))
	c[1] = uintptr(len(base)) / uintlen
	c[2] = c[1]
	bw := *(*[]big.Word)(unsafe.Pointer(&c[0]))
	x1 := new(big.Int).SetBits(bw)

	// the same as (*MtCurve).ScalaMultProjectiveU1(x1 *big.Int, k []byte), except that k is little endian
	x2, z2 := new(big.Int).SetInt64(1), new(big.Int).SetInt64(0) // x2 initial as O
	x3, z3 := new(big.Int).Set(x1), zForAffine(x1, x1)           // x3 initial as P
	var m byte = 0                                               // initial former bit as 0
	for i := len(e) - 1; i >= 0; i-- {                           // the input is little-endian
		b := e[i]
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
			x3.Mod(x3, p448) // x3 = z1*(DA+CB)², z1 is always 1
			z3.Sub(DA, CB)
			z3.Mul(z3, z3)
			z3.Mul(z3, x1)
			z3.Mod(z3, p448) // z3 = x1*(DA-CB)²
			x2.Mul(AA, BB)
			x2.Mod(x2, p448) // x2 = AA*BB
			z2.Mul(a24_448, E)
			z2.Add(z2, AA)
			z2.Mul(z2, E)
			z2.Mod(z2, p448) // z2 = E*(AA+a24*E)
			b <<= 1
		}
	}
	cswap(m, x2, x3)
	cswap(m, z2, z3)

	// convert projective coordinate x2/z2 to affine coordinate
	copy(dst[:], zero448[:]) // clear dst first
	if z2.Sign() == 0 {
		return // the result is zero, just return 0
	}
	z2.ModInverse(z2, p448)
	x2.Mul(x2, z2)
	x2.Mod(x2, p448)

	// from big.Int to little endian bytes
	words := x2.Bits()
	byters := *(*[]byte)(unsafe.Pointer(&words))
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&byters)) + uintptr(8))) = len(words) * int(uintlen)
	copy(dst[:], byters)
}

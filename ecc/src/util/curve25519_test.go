package util

import (
	"crypto"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/ed25519"
	"math/big"
	"reflect"
	"testing"
	"bytes"
	"ecc"
	"crypto/sha512"
)

// curve25519
func TestCurve25519(t *testing.T) {
	var priv1, priv2, pub1, pub2, share1, share2 [32]byte
	rand.Read(priv1[:])
	rand.Read(priv2[:])
	fmt.Println(reflect.TypeOf(priv1).Kind().String())
	fmt.Println(reflect.TypeOf(priv2).Kind().String())
	fmt.Printf("priv1: %x\n", priv1)
	fmt.Printf("priv2: %x\n", priv2)
	curve25519.ScalarBaseMult(&pub1, &priv1)
	curve25519.ScalarBaseMult(&pub2, &priv2)
	fmt.Printf("pub1: %x\n", pub1)
	fmt.Printf("pub2: %x\n", pub2)
	curve25519.ScalarMult(&share2, &priv2, &pub1)
	curve25519.ScalarMult(&share1, &priv1, &pub2)
	fmt.Printf("share1: %x\n", share1)
	fmt.Printf("share2: %x\n", share2)

	// the order of in and base is important, the following is wrong
	curve25519.ScalarMult(&share2, &pub1, &priv2)
	curve25519.ScalarMult(&share1, &pub2, &priv1)
	fmt.Printf("share1: %x\n", share1)
	fmt.Printf("share2: %x\n", share2)
}

func TestEd25519(t *testing.T) {
	reader := rand.Reader
	/*pub, priv, err := ed25519.GenerateKey(reader)
	if err != nil {
		fmt.Println("genkey err: ", err.Error())
	}*/
	pub, err := hex.DecodeString("206d3b36a46545d00fd417375df871f546de840cf1e1fb6ea055358cf92949e2")
	priv, err := hex.DecodeString("86ce3bc7c2bd0be6b71bf9eb9bc6a9acdb97ad7f5aaf7ca2fcfadbbf7be2f216206d3b36a46545d00fd417375df871f546de840cf1e1fb6ea055358cf92949e2")
	fmt.Println(reflect.TypeOf(pub).Kind().String())
	fmt.Println(reflect.TypeOf(priv).Kind().String())
	fmt.Printf("pub: %x\n", []byte(pub))
	fmt.Printf("priv: %x\n", []byte(priv))
	msg := []byte("wang sheng tao.")
	sign, err := ed25519.PrivateKey(priv).Sign(reader, msg, crypto.Hash(0))
	if err != nil {
		fmt.Println("sign err: ", err.Error())
	}
	fmt.Printf("sign: %x\n", sign)
	fmt.Println(ed25519.Verify(ed25519.PublicKey(pub), msg, sign))
}

func TestD(t *testing.T) {
	// -121665/121666 = 37095705934669439343138083508754565189542113879843219016388785533085940283555 mod 2²⁵⁵-19
	POS121666 := new(big.Int).SetInt64(121666)
	NEG121665 := new(big.Int).SetInt64(-121665)
	d := new(big.Int).ModInverse(POS121666, P25519)
	d.Mul(d, NEG121665)
	d.Mod(d, P25519)
	fmt.Println(d.String())
}

func TestEd25519Base(t *testing.T) {
	// 4/5 = 46316835694926478169428394003475163141307993866256225615783033603165251855960 mod 2²⁵⁵-19
	POS4 := new(big.Int).SetInt64(4)
	POS5 := new(big.Int).SetInt64(5)
	d := new(big.Int).ModInverse(POS5, P25519)
	d.Mul(d, POS4)
	d.Mod(d, P25519)
	fmt.Println(d.String())
	fmt.Println(hex.EncodeToString(d.Bytes()))
}

func TestConvertPoint(t *testing.T) {
	// set numbers
	N19 := new(big.Int).SetInt64(19)
	P25519 := new(big.Int).SetInt64(1)
	P25519.Lsh(P25519, 255)
	P25519.Sub(P25519, N19)
	fmt.Printf("25519: %x\n", P25519.Bytes())
	NEG486664 := new(big.Int).SetInt64(-486664)
	N486662 := new(big.Int).SetInt64(486662)
	SQRTNEG486664 := new(big.Int).ModSqrt(NEG486664, P25519)
	fmt.Printf("sqrt: %x\n", SQRTNEG486664.Bytes())
	ONE := new(big.Int).SetInt64(1)
	D, _ := new(big.Int).SetString("37095705934669439343138083508754565189542113879843219016388785533085940283555", 10)
	{
		fmt.Println("\nfrom curve25519 (u,v) ed25519 to (x, y)")
		// u = 9
		// v = 14781619447589544791020593568409986887264606134616475288964881837755586237401
		// (x, y) = (sqrt(-486664)*u/v, (u-1)/(u+1))
		u := new(big.Int).SetInt64(9)
		v, _ := new(big.Int).SetString("14781619447589544791020593568409986887264606134616475288964881837755586237401", 10)
		x := new(big.Int).ModInverse(v, P25519) // 1/v
		x.Mul(x, u)                             // u/v
		x.Mul(x, SQRTNEG486664)                 // sqrt(-486664)*u/v
		x.Mod(x, P25519)                        // x ends
		u_1 := new(big.Int).Add(u, ONE)         // u+1
		u_1.ModInverse(u_1, P25519)             // 1/(u+1)
		y := new(big.Int).Sub(u, ONE)           // u-1
		y.Mul(y, u_1)                           // (u-1)/(u+1)
		y.Mod(y, P25519)                        // y ends
		fmt.Println("x: ", x.String())
		fmt.Println("y: ", y.String())

		fmt.Println("use function: ")
		x, y = Curve25519ToEd25519(u, v)
		fmt.Println("x: ", x.String())
		fmt.Println("y: ", y.String())
	}

	{
		fmt.Println("\nfrom ed25519 (x,y) to curve25519 (u, v)")
		// x = 15112221349535400772501151409588531511454012693041857206046113283949847762202
		// y = 46316835694926478169428394003475163141307993866256225615783033603165251855960
		// (u, v) = ((1+y)/(1-y), sqrt(-486664)*u/x)
		x, _ := new(big.Int).SetString("15112221349535400772501151409588531511454012693041857206046113283949847762202", 10)
		y, _ := new(big.Int).SetString("46316835694926478169428394003475163141307993866256225615783033603165251855960", 10)
		_1_y := new(big.Int).Sub(ONE, y)        // 1-y
		_1_y.ModInverse(_1_y, P25519)           // 1/(1-y)
		u := new(big.Int).Add(ONE, y)           // 1+y
		u.Mul(u, _1_y)                          // (1+y)/(1-y)
		u.Mod(u, P25519)                        // u ends
		v := new(big.Int).ModInverse(x, P25519) // 1/x
		v.Mul(u, v)                             // u/x
		v.Mul(SQRTNEG486664, v)                 // sqrt(-486664)*u/x
		v.Mod(v, P25519)
		fmt.Println("u: ", u.String())
		fmt.Println("v: ", v.String())

		fmt.Println("use function: ")
		u, v = Ed25519ToCurve25519(x, y)
		fmt.Println("u: ", u.String())
		fmt.Println("v: ", v.String())
	}

	{
		fmt.Println("\nfrom curve25519 (u,?) to (u,v)")
		u := new(big.Int).SetInt64(9)
		u2 := new(big.Int).Mul(u, u)       // u^2
		v := new(big.Int).Mul(N486662, u2) // 486662u^2
		v.Add(v, u)                        // 486662u^2 + u
		u2.Mul(u2, u)                      // u^3
		v.Add(u2, v)                       // u^3 + 486662u^2 + u
		v.ModSqrt(v, P25519)               // sqrt(u^3 + A*u^2 + u)
		v.Mod(v, P25519)                   // first v
		v1 := new(big.Int).Sub(P25519, v)  //second v
		fmt.Println("v : ", v.String())
		fmt.Println("v1: ", v1.String())
		v, v1 = Curve25519U2V(u)
		fmt.Println("v : ", v.String())
		fmt.Println("v1: ", v1.String())
	}

	{
		fmt.Println("\nfrom ed25519 (?,y) to (x,y)")
		y, _ := new(big.Int).SetString("46316835694926478169428394003475163141307993866256225615783033603165251855960", 10)
		// -x² + y² = 1 + dx²y²
		// x² = (y²-1)/(dy²+1)
		x := new(big.Int).Mul(y, y)   // y²
		tmp := new(big.Int).Mul(D, x) // dy²
		tmp.Add(tmp, ONE)             // dy²+1
		tmp.ModInverse(tmp, P25519)   // 1/(dy²+1)
		tmp.Mod(tmp, P25519)
		x.Sub(x, ONE)        //  y²-1
		x.Mul(x, tmp)        //  (y²-1)/(dy²+1)
		x.ModSqrt(x, P25519) // sqrt( (y²-1)/(dy²+1) )
		x.Mod(x, P25519)
		fmt.Println("x : ", x.String())
		x = Ed25519Y2X(y)
		fmt.Println("x : ", x.String())
	}
}

// curve25519 base_point (bu=9, bv)   *   secret_K  =  (U, ?)        U: curve25519 public key
//                           ||                          ||
// ed25519 base_point  (bx, by=4/5)   *   secret_K  =  (?, Y)        Y: ed25519 public key
// y = (u-1)/(u+1)
// u = (1+y)/(1-y)
func TestCurve25519AndEd25519Arithmetic(t *testing.T) {
	reader := rand.Reader
	seed := make([]byte, ed25519.SeedSize)
	n, err := reader.Read(seed)
	if n < 32 || err != nil {
		fmt.Println("error read seed: ", err)
		t.FailNow()
	}
	fmt.Printf("seed: %x\n", seed)
	dgst := sha512.Sum512(seed)
	var mltr [32]byte
	copy(mltr[:], dgst[:])
	fmt.Printf("dgst: %x\n", dgst[:])
	fmt.Printf("mltr: %x\n", mltr[:])
	privKey := ed25519.NewKeyFromSeed(seed)
	privs := []byte(privKey)[0:32]
	pubs := []byte(privKey)[32:]
	fmt.Printf("priv: %x\n", privs)
	fmt.Printf("pubs: %x\n", pubs)

	var rslt [32]byte
	curve25519.ScalarBaseMult(&rslt, &mltr)
	fmt.Printf("rslt: %x\n", rslt)

	// curve25519 and ed25519 are little-endian, while big.Int is big-endian.
	// so every time we see a big.Int, reverse it.
	reverse := func(arr []byte) []byte {
		for i, j:= 0, len(arr)-1; i<j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		return arr
	}

	y := Curve25519XToEd25519Y(new(big.Int).SetBytes(reverse(rslt[:])));
	fmt.Printf("y   : %x\n", reverse(y.Bytes()))

	// when reverse(pubs) equals y, u=u1; reverse(pubs) and y may have minor difference
	u := Ed25519YToCurve25519X(new(big.Int).SetBytes(reverse(pubs)))
	u1 := Ed25519YToCurve25519X(y)
	fmt.Printf("u   : %x\n", reverse(u.Bytes()))
	fmt.Printf("u1  : %x\n", reverse(u1.Bytes()))
	fmt.Println(bytes.Equal(pubs, reverse(y.Bytes())), bytes.Equal(rslt[:], reverse(u.Bytes())))
}

func TestX25519Rfc7748(t *testing.T) {
	scala1, _ := hex.DecodeString("77076d0a7318a57d3c16c17251b26645df4c2f87ebc0992ab177fba51db92c2a")
	scala2, _ := hex.DecodeString("5dab087e624a8a4b79e17f8b83800ee66f3bb1292618b6fd1c2f8b27ff88e0eb")
	pub1Ex, _ := hex.DecodeString("8520f0098930a754748b7ddcb43ef75a0dbf3a0d26381af4eba4a98eaa9b4e6a")
	pub2Ex, _ := hex.DecodeString("de9edb7d7b7dc1b4d35b61c2ece435373f8343c85b78674dadfc7e146f882b4f")
	shareE, _ := hex.DecodeString("4a5d9d5ba4ce2de1728e3bf480350f25e07e21c947d19e3376f09b3c1e161742")
	fmt.Printf("scala1: %x\n", scala1)
	fmt.Printf("scala2: %x\n", scala2)
	var priv1, priv2, pub1, pub2, share1, share2 [32]byte
	copy(priv1[:], scala1)
	copy(priv2[:], scala2)
	curve25519.ScalarBaseMult(&pub1, &priv1)
	curve25519.ScalarBaseMult(&pub2, &priv2)
	fmt.Printf("pub1: %x\n", pub1)
	fmt.Printf("pub2: %x\n", pub2)
	curve25519.ScalarMult(&share2, &priv2, &pub1)
	curve25519.ScalarMult(&share1, &priv1, &pub2)
	fmt.Printf("share1: %x\n", share1)
	fmt.Printf("share2: %x\n", share2)
	fmt.Println(bytes.Equal(pub1[:], pub1Ex))
	fmt.Println(bytes.Equal(pub2[:], pub2Ex))
	fmt.Println(bytes.Equal(share1[:], shareE))
	fmt.Println(bytes.Equal(share1[:], share2[:]))
}

func TestEquivalence25519(t *testing.T) {
	// 4/5 = 46316835694926478169428394003475163141307993866256225615783033603165251855960 mod 2²⁵⁵-19
	fourFive := ecc.ModFractionInt64(4, 5, P25519)
	fmt.Println("4/5: ", fourFive)
	{
		d := ecc.ModFractionInt64(121665, 121666, P25519)
		fmt.Println("d: ", d.String())
		a := new(big.Int).SetInt64(1)
		aSubD := new(big.Int).Sub(a, d) // a-d
		fmt.Println("a-d: ", aSubD.String())
		aSubD.ModInverse(aSubD, P25519) // 1/(a-d)
		fmt.Println("1/(a-d): ", aSubD.String())
		A := new(big.Int).Add(a, d) // a+d
		fmt.Println("a+d: ", A.String())
		A.Lsh(A, 1) // 2(a+d)
		fmt.Println("2(a+d): ", A.String())
		A.Mul(A, aSubD) // 2(a+d)/(a-d)
		fmt.Println("2(a+d)(a-d): ", A.String())
		A.Mod(A, P25519)
		fmt.Println("A: ", A.String())
		B := new(big.Int).Lsh(aSubD, 2)
		B.Mod(B, P25519)
		fmt.Println("B: ", B.String())
		B.ModInverse(B, P25519)
		fmt.Println("B^-1: ", B.String())
		POS2 := new(big.Int).SetInt64(2)
		aa := new(big.Int).Add(A, POS2)
		aa.Mul(aa, B)
		aa.Mod(aa, P25519)
		fmt.Println("a: ", aa.String())
		dd := new(big.Int).Sub(A, POS2)
		dd.Mul(dd, B)
		dd.Mod(dd, P25519)
		fmt.Println("d: ", dd.String())
		bSqrt := new(big.Int).ModSqrt(B, P25519)
		fmt.Println("sqrtB: ", bSqrt)
		fmt.Println()
	}
	{
		// -121665/121666 = 37095705934669439343138083508754565189542113879843219016388785533085940283555 mod 2²⁵⁵-19
		d := ecc.ModFractionInt64(-121665, 121666, P25519)
		fmt.Println("d: ", d.String())
		a := new(big.Int).SetInt64(-1)
		aSubD := new(big.Int).Sub(a, d) // a-d
		fmt.Println("a-d: ", aSubD.String())
		aSubD.ModInverse(aSubD, P25519) // 1/(a-d)
		fmt.Println("1/(a-d): ", aSubD.String())
		A := new(big.Int).Add(a, d) // a+d
		fmt.Println("a+d: ", A.String())
		A.Lsh(A, 1) // 2(a+d)
		fmt.Println("2(a+d): ", A.String())
		A.Mul(A, aSubD) // 2(a+d)/(a-d)
		fmt.Println("2(a+d)(a-d): ", A.String())
		A.Mod(A, P25519)
		fmt.Println("A: ", A.String())
		B := new(big.Int).Lsh(aSubD, 2)
		B.Mod(B, P25519)
		fmt.Println("B: ", new(big.Int).Sub(B, P25519).String())
		B.ModInverse(B, P25519)
		fmt.Println("B^-1: ", B.String())
		POS2 := new(big.Int).SetInt64(2)
		aa := new(big.Int).Add(A, POS2)
		aa.Mul(aa, B)
		aa.Mod(aa, P25519)
		aa.Sub(aa, P25519)
		fmt.Println("a: ", aa.String())
		dd := new(big.Int).Sub(A, POS2)
		dd.Mul(dd, B)
		dd.Mod(dd, P25519)
		fmt.Println("d: ", dd.String())
		bSqrt := new(big.Int).ModSqrt(B, P25519)
		fmt.Println("sqrtB: ", bSqrt)
	}
}

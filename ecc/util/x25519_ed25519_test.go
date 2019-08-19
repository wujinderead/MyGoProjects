package util

import (
	"crypto"
	"crypto/rand"
	"crypto/sha512"
	"ecc/ecc"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/ed25519"
	"reflect"
	"testing"
)

// https://tools.ietf.org/html/rfc7748
func TestX25519Rfc7748PrivPub(t *testing.T) {
	// test vector from rfc7748
	it1, _ := hex.DecodeString("422c8e7a6227d7bca1350b3e2bb7279f7897b87bb6854b783c60e80311ae3079")
	it1000, _ := hex.DecodeString("684cf59ba83309552800ef566f2f4d3c1c3887c49360e3875f2eb94d99532c51")
	it1000000, _ := hex.DecodeString("7c3911e0ab2586fd864497297e575e6f3bc601c0883c30df5f4dd2d24f665424")

	// initialize base and scala both as [9 0 0 0...]
	var base, scala, prod [32]byte
	base[0] = 9
	scala[0] = 9

	// do 1 iteration
	curve25519.ScalarMult(&prod, &scala, &base)
	fmt.Printf("1 iteration: %x\n", prod[:])
	assert.Equal(t, it1, prod[:])

	// do 1000 iteration
	for i := 2; i <= 1000; i++ {
		// use product as new scala, use old scala as new base
		base, prod, scala = scala, base, prod
		curve25519.ScalarMult(&prod, &scala, &base)
	}
	fmt.Printf("1000 iteration: %x\n", prod[:])
	assert.Equal(t, it1000, prod[:])

	// do 1000000 iteration, this takes dozens of seconds time
	for i := 1001; i <= 1000000; i++ {
		// use product as new scala, use old scala as new base
		base, prod, scala = scala, base, prod
		curve25519.ScalarMult(&prod, &scala, &base)
	}
	fmt.Printf("1000000 iteration: %x\n", prod[:])
	assert.Equal(t, it1000000, prod[:])
}

func TestX25519Rfc7748DiffieHellman(t *testing.T) {
	// test golang.org/x/crypto/curve25519 package
	// test vectors from in rfc7748, all in little-endian
	scala1, _ := hex.DecodeString("77076d0a7318a57d3c16c17251b26645df4c2f87ebc0992ab177fba51db92c2a")
	scala2, _ := hex.DecodeString("5dab087e624a8a4b79e17f8b83800ee66f3bb1292618b6fd1c2f8b27ff88e0eb")
	pub1Ex, _ := hex.DecodeString("8520f0098930a754748b7ddcb43ef75a0dbf3a0d26381af4eba4a98eaa9b4e6a")
	pub2Ex, _ := hex.DecodeString("de9edb7d7b7dc1b4d35b61c2ece435373f8343c85b78674dadfc7e146f882b4f")
	shareE, _ := hex.DecodeString("4a5d9d5ba4ce2de1728e3bf480350f25e07e21c947d19e3376f09b3c1e161742")
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
	assert.Equal(t, pub1[:], pub1Ex, "pub1 equal")
	assert.Equal(t, pub2[:], pub2Ex, "pub2 equal")
	assert.Equal(t, share1[:], shareE, "share equal")
	assert.Equal(t, share1[:], share2[:], "share 1,2 equal")

	// test ecc package
	// 'golang.org/x/crypto/curve25519' input and output are little-endian
	// 'ecc' scala multiplication input is big-endian
	pub1Ecc := ecc.Curve25519().ScalaMultBase(reverse(process(scala1)))
	pub2Ecc := ecc.Curve25519().ScalaMultBase(reverse(process(scala2)))
	share1Ecc := ecc.Curve25519().ScalaMult(pub1Ecc, reverse(process(scala2)))
	share2Ecc := ecc.Curve25519().ScalaMult(pub2Ecc, reverse(process(scala1)))
	fmt.Printf("ecc pub1: %x\n", reverse(pub1Ecc.X.Bytes()))
	fmt.Printf("ecc pub2: %x\n", reverse(pub2Ecc.X.Bytes()))
	fmt.Printf("ecc share1: %x\n", reverse(share1Ecc.X.Bytes()))
	fmt.Printf("ecc share2: %x\n", reverse(share2Ecc.X.Bytes()))
	assert.Equal(t, reverse(pub1Ecc.X.Bytes()), pub1Ex, "pub1 equal")
	assert.Equal(t, reverse(pub2Ecc.X.Bytes()), pub2Ex, "pub2 equal")
	assert.Equal(t, reverse(share1Ecc.X.Bytes()), shareE, "share equal")
	assert.Equal(t, share1Ecc.X.Bytes(), share2Ecc.X.Bytes(), "share 1,2 equal")

	// test scala multiplication in projective coordinates
	pub1Pro := ecc.Curve25519().ScalaMultBaseProjective(reverse(process(scala1)))
	pub2Pro := ecc.Curve25519().ScalaMultBaseProjective(reverse(process(scala2)))
	share1Pro := ecc.Curve25519().ScalaMultProjective(pub1Ecc, reverse(process(scala2)))
	share2Pro := ecc.Curve25519().ScalaMultProjective(pub2Ecc, reverse(process(scala1)))
	assert.Equal(t, reverse(pub1Pro.X.Bytes()), pub1Ex, "pub1 equal")
	assert.Equal(t, reverse(pub2Pro.X.Bytes()), pub2Ex, "pub2 equal")
	assert.Equal(t, reverse(share1Pro.X.Bytes()), shareE, "share equal")
	assert.Equal(t, share1Pro.X.Bytes(), share2Pro.X.Bytes(), "share 1,2 equal")
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

func TestEd25519GenerateKey(t *testing.T) {
	var seed [32]byte
	_, _ = rand.Read(seed[:])
	privKey := ed25519.NewKeyFromSeed(seed[:])

	// ed25519 key is seed + public key
	fmt.Println("seed:", hex.EncodeToString(privKey[:32]))
	fmt.Println("pubK:", hex.EncodeToString(privKey[32:]))

	// re-perform a key generate process
	h := sha512.New()
	h.Write(seed[:])
	digest := h.Sum(nil)

	var priK = digest[:32]
	priK[0] &= 248
	priK[31] &= 127
	priK[31] |= 64
	pubPoint := ecc.Ed25519().ScalaMultBaseProjective(reverse(priK))
	y := reverse(pubPoint.Y.Bytes())
	x := reverse(pubPoint.X.Bytes())
	fmt.Println("y   :", hex.EncodeToString(y))
	fmt.Println("x   :", x[0])
	// copy the least significant bit of pubPoint.X to the most significant bit of pubPoint.Y
	if x[0]&1 == 1 { // the least significant bit of pubPoint.X is 1
		y[31] |= 0x80
	} else {
		y[31] &= 0x7f
	}
	fmt.Println("newy:", hex.EncodeToString(y))
	assert.Equal(t, y, []byte(privKey[32:]), "public key equal")
}

func process(scala []byte) []byte {
	e := make([]byte, len(scala))
	copy(e, scala)
	e[0] &= 248
	e[31] &= 127
	e[31] |= 64
	return e
}
func reverse(b []byte) []byte {
	a := make([]byte, len(b))
	copy(a, b)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

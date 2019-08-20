package util

import (
	"crypto"
	"crypto/rand"
	"crypto/sha512"
	"ecc/ecc"
	"ecc/util/edwards25519"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/ed25519"
	"math/big"
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
	seed, _ := hex.DecodeString("5c747028926820d3669df3c3ee7c04818ea44c3f9b79430a97bf6b3efeba86c6")
	pubex, _ := hex.DecodeString("0479a40adbb1703c07a9e301ce8ca1074350634949c67e4df651e21d3c9bc5cd")
	priv := ed25519.NewKeyFromSeed(seed)
	pub := priv.Public().(ed25519.PublicKey)
	assert.Equal(t, []byte(pub), pubex, "pub key equal")
	assert.Equal(t, []byte(priv[:ed25519.SeedSize]), seed, "priv is actually seed")

	msg := []byte("The quick brown fox jumps over the lazy dog.")

	// ed25519 sign and verify
	// rand.Reader is unused, so ed25519 generates the same signature for same message
	sign, _ := ed25519.PrivateKey(priv).Sign(reader, msg, crypto.Hash(0))
	fmt.Printf("sign: %x\n", sign)
	fmt.Println(ed25519.Verify(ed25519.PublicKey(pub), msg, sign))

	// generate ed25519 signature manually by following the procedure in ed25519 package
	// sha512 to digest seed
	h := sha512.New()
	h.Write(seed)
	digest := h.Sum(nil)

	// part seed digest to 2 parts
	frontDigest := digest[:32]
	backDigest := digest[32:]

	// use frontDigest as secret
	frontDigest[0] &= 248
	frontDigest[31] &= 63 // why 63?
	frontDigest[31] |= 64

	// digest (backDigest + message) as md
	h.Reset()
	h.Write(backDigest)
	h.Write(msg)
	md := h.Sum(nil)
	// convert md to a number, then reduce it within ed25519 prime order (l = 2^252 + 27742317777372353535851937790883648493)
	mdn := new(big.Int).SetBytes(reverse(md))
	mdn.Mod(mdn, ecc.Ed25519().Order)
	// scala multiply md with base point, generate R
	R := ecc.Ed25519().ScalaMultBaseProjective(mdn.Bytes())
	// encode R
	encodedR := ed25519Encode(R)
	fmt.Printf("R: %x\n", encodedR)

	// digest (encodeR + publicKey + message) as hram
	h.Reset()
	h.Write(encodedR[:])
	h.Write(pub)
	h.Write(msg)
	hram := h.Sum(nil)
	// convert hram to number and reduce
	hramn := new(big.Int).SetBytes(reverse(hram))
	hramn.Mod(hramn, ecc.Ed25519().Order)

	// compute s as (hramn * secret + mdn) mod l
	secret := new(big.Int).SetBytes(reverse(frontDigest))
	s := new(big.Int).Mul(secret, hramn)
	s.Add(s, mdn)
	s.Mod(s, ecc.Ed25519().Order)
	fmt.Printf("s: %x\n", reverse(s.Bytes()))
}

// reference procedure of Sign in 'golang.org/x/crypto/curve25519'
func TestEd25519SignReference(t *testing.T) {
	seed, _ := hex.DecodeString("5c747028926820d3669df3c3ee7c04818ea44c3f9b79430a97bf6b3efeba86c6")
	privateKey := ed25519.NewKeyFromSeed(seed)

	message := []byte("The quick brown fox jumps over the lazy dog.")

	h := sha512.New()
	h.Write(privateKey[:32])

	var digest1, messageDigest, hramDigest [64]byte
	var expandedSecretKey [32]byte
	h.Sum(digest1[:0])
	copy(expandedSecretKey[:], digest1[:])
	expandedSecretKey[0] &= 248
	expandedSecretKey[31] &= 63
	expandedSecretKey[31] |= 64

	h.Reset()
	h.Write(digest1[32:])
	h.Write(message)
	h.Sum(messageDigest[:0])

	var messageDigestReduced [32]byte
	edwards25519.ScReduce(&messageDigestReduced, &messageDigest)
	var R edwards25519.ExtendedGroupElement
	edwards25519.GeScalarMultBase(&R, &messageDigestReduced)

	var encodedR [32]byte
	R.ToBytes(&encodedR)
	fmt.Println("R:", hex.EncodeToString(encodedR[:]))

	h.Reset()
	h.Write(encodedR[:])
	h.Write(privateKey[32:])
	h.Write(message)
	h.Sum(hramDigest[:0])
	var hramDigestReduced [32]byte
	edwards25519.ScReduce(&hramDigestReduced, &hramDigest)

	var s [32]byte
	edwards25519.ScMulAdd(&s, &hramDigestReduced, &expandedSecretKey, &messageDigestReduced)
	fmt.Println("s:", hex.EncodeToString(s[:]))

	signature := make([]byte, 64)
	copy(signature[:], encodedR[:])
	copy(signature[32:], s[:])
}

func TestEd25519GenerateKey(t *testing.T) {
	var seed [32]byte
	_, _ = rand.Read(seed[:])
	privKey := ed25519.NewKeyFromSeed(seed[:])

	// ed25519 key is seed + public key
	fmt.Println("seed:", hex.EncodeToString(privKey[:32]))
	fmt.Println("pubK:", hex.EncodeToString(privKey[32:]))

	// re-perform a key generate process
	// sha512 to digest the seed
	h := sha512.New()
	h.Write(seed[:])
	digest := h.Sum(nil)

	// use front 32 bytes of the digest as the real secret key
	var priK = digest[:32]

	// scala multiply secret key and base point, and generate public key
	pubPoint := ecc.Ed25519().ScalaMultBaseProjective(reverse(process(priK)))

	// convert the product ec point to ed25519 format,
	// i.e., copy the least significant bit of point.X to the most significant bit of point.Y
	eccPub := ed25519Encode(pubPoint)
	fmt.Println("eccP:", hex.EncodeToString(eccPub))
	assert.Equal(t, eccPub, []byte(privKey[32:]), "public key equal")
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

func ed25519Encode(point *ecc.EcPoint) []byte {
	y := reverse(point.Y.Bytes())
	x := reverse(point.X.Bytes())
	// copy the least significant bit of point.X to the most significant bit of point.Y
	if x[0]&1 == 1 { // if the least significant bit of point.X is 1
		y[31] |= 0x80 // set the most significant bit of point.Y to 1
	} else { // if 0
		y[31] &= 0x7f // set 0
	}
	return y
}

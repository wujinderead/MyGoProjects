package util

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/sha512"
	"ecc"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/ed25519"
	"math/big"
	"reflect"
	"testing"
)

// curve25519
func TestCurve25519(t *testing.T) {
	var priv1, priv2, pub1, pub2, share1, share2 [32]byte
	_, _ = rand.Read(priv1[:])
	_, _ = rand.Read(priv2[:])
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
	assert.Equal(t, share1, share2, "share equal")

	// the order of in and base is important, the following is wrong
	curve25519.ScalarMult(&share2, &pub1, &priv2)
	curve25519.ScalarMult(&share1, &pub2, &priv1)
	fmt.Printf("share1: %x\n", share1)
	fmt.Printf("share2: %x\n", share2)
	assert.NotEqual(t, share1, share2, "share not equal")
}

// https://tools.ietf.org/html/rfc7748
func TestX25519Rfc7748(t *testing.T) {
	// test golang.org/x/crypto/curve25519 package
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
	process := func(scala []byte) []byte {
		e := make([]byte, len(scala))
		copy(e, scala)
		e[0] &= 248
		e[31] &= 127
		e[31] |= 64
		return e
	}
	reverse := func(b []byte) []byte {
		i, j := 0, len(b)-1
		for i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
		return b
	}
	// ecc scala multiplication input is big-endian, rfc7748 input and output are little-endian
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
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		return arr
	}

	y := Curve25519XToEd25519Y(new(big.Int).SetBytes(reverse(rslt[:])))
	fmt.Printf("y   : %x\n", reverse(y.Bytes()))

	// when reverse(pubs) equals y, u=u1; reverse(pubs) and y may have minor difference
	u := Ed25519YToCurve25519X(new(big.Int).SetBytes(reverse(pubs)))
	u1 := Ed25519YToCurve25519X(y)
	fmt.Printf("u   : %x\n", reverse(u.Bytes()))
	fmt.Printf("u1  : %x\n", reverse(u1.Bytes()))
	fmt.Println(bytes.Equal(pubs, reverse(y.Bytes())), bytes.Equal(rslt[:], reverse(u.Bytes())))
}

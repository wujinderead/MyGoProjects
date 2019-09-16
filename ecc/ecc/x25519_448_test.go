package ecc

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestX25519Rfc7748(t *testing.T) {
	// diffie-hellman test
	scala1, _ := hex.DecodeString("77076d0a7318a57d3c16c17251b26645df4c2f87ebc0992ab177fba51db92c2a")
	scala2, _ := hex.DecodeString("5dab087e624a8a4b79e17f8b83800ee66f3bb1292618b6fd1c2f8b27ff88e0eb")
	pub1Ex, _ := hex.DecodeString("8520f0098930a754748b7ddcb43ef75a0dbf3a0d26381af4eba4a98eaa9b4e6a")
	pub2Ex, _ := hex.DecodeString("de9edb7d7b7dc1b4d35b61c2ece435373f8343c85b78674dadfc7e146f882b4f")
	shareE, _ := hex.DecodeString("4a5d9d5ba4ce2de1728e3bf480350f25e07e21c947d19e3376f09b3c1e161742")
	var priv1, priv2, pub1, pub2, share1, share2 [32]byte
	copy(priv1[:], scala1)
	copy(priv2[:], scala2)
	Curve25519ScalaMultBase(&pub1, &priv1)
	Curve25519ScalaMultBase(&pub2, &priv2)
	fmt.Printf("pub1: %x\n", pub1)
	fmt.Printf("pub2: %x\n", pub2)
	Curve25519ScalaMult(&share2, &priv2, &pub1)
	Curve25519ScalaMult(&share1, &priv1, &pub2)
	fmt.Printf("share1: %x\n", share1)
	fmt.Printf("share2: %x\n", share2)
	assert.Equal(t, pub1[:], pub1Ex, "pub1 equal")
	assert.Equal(t, pub2[:], pub2Ex, "pub2 equal")
	assert.Equal(t, share1[:], shareE, "share equal")
	assert.Equal(t, share1[:], share2[:], "share 1,2 equal")

	// scala multiply test
	it1, _ := hex.DecodeString("422c8e7a6227d7bca1350b3e2bb7279f7897b87bb6854b783c60e80311ae3079")
	it1000, _ := hex.DecodeString("684cf59ba83309552800ef566f2f4d3c1c3887c49360e3875f2eb94d99532c51")

	// initialize base and scala both as base point [9 0 0 0...]
	var base, scala, prod [32]byte
	copy(base[:], x25519base[:])
	copy(scala[:], x25519base[:])

	// do 1 iteration
	Curve25519ScalaMult(&prod, &scala, &base)
	fmt.Printf("1 iteration: %x\n", prod[:])
	assert.Equal(t, it1, prod[:])

	// do 1000 iteration
	for i := 2; i <= 1000; i++ {
		// use product as new scala, use old scala as new base
		base, prod, scala = scala, base, prod
		Curve25519ScalaMult(&prod, &scala, &base)
	}
	fmt.Printf("1000 iteration: %x\n", prod[:])
	assert.Equal(t, it1000, prod[:])
}

func TestX448Rfc7748(t *testing.T) {
	// diffie-hellman test
	scala1, _ := hex.DecodeString("9a8f4925d1519f5775cf46b04b5800d4ee9ee8bae8bc5565d498c28dd9c9baf574a9419744897391006382a6f127ab1d9ac2d8c0a598726b")
	scala2, _ := hex.DecodeString("1c306a7ac2a0e2e0990b294470cba339e6453772b075811d8fad0d1d6927c120bb5ee8972b0d3e21374c9c921b09d1b0366f10b65173992d")
	pub1Ex, _ := hex.DecodeString("9b08f7cc31b7e3e67d22d5aea121074a273bd2b83de09c63faa73d2c22c5d9bbc836647241d953d40c5b12da88120d53177f80e532c41fa0")
	pub2Ex, _ := hex.DecodeString("3eb7a829b0cd20f5bcfc0b599b6feccf6da4627107bdb0d4f345b43027d8b972fc3e34fb4232a13ca706dcb57aec3dae07bdc1c67bf33609")
	shareE, _ := hex.DecodeString("07fff4181ac6cc95ec1c16a94a0f74d12da232ce40a77552281d282bb60c0b56fd2464c335543936521c24403085d59a449a5037514a879d")
	var priv1, priv2, pub1, pub2, share1, share2 [56]byte
	copy(priv1[:], scala1)
	copy(priv2[:], scala2)
	Curve448ScalaMultBase(&pub1, &priv1)
	Curve448ScalaMultBase(&pub2, &priv2)
	fmt.Printf("pub1: %x\n", pub1)
	fmt.Printf("pub2: %x\n", pub2)
	Curve448ScalaMult(&share2, &priv2, &pub1)
	Curve448ScalaMult(&share1, &priv1, &pub2)
	fmt.Printf("share1: %x\n", share1)
	fmt.Printf("share2: %x\n", share2)
	assert.Equal(t, pub1[:], pub1Ex, "pub1 equal")
	assert.Equal(t, pub2[:], pub2Ex, "pub2 equal")
	assert.Equal(t, share1[:], shareE, "share equal")
	assert.Equal(t, share1[:], share2[:], "share 1,2 equal")

	// scala multiply test
	it1, _ := hex.DecodeString("3f482c8a9f19b01e6c46ee9711d9dc14fd4bf67af30765c2ae2b846a4d23a8cd0db897086239492caf350b51f833868b9bc2b3bca9cf4113")
	it1000, _ := hex.DecodeString("aa3b4749d55b9daf1e5b00288826c467274ce3ebbdd5c17b975e09d4af6c67cf10d087202db88286e2b79fceea3ec353ef54faa26e219f38")

	// initialize base and scala both as base point [5 0 0 0...]
	var base, scala, prod [56]byte
	copy(base[:], x448base[:])
	copy(scala[:], x448base[:])

	// do 1 iteration
	Curve448ScalaMult(&prod, &scala, &base)
	fmt.Printf("1 iteration: %x\n", prod[:])
	assert.Equal(t, it1, prod[:])

	// do 1000 iteration
	for i := 2; i <= 1000; i++ {
		// use product as new scala, use old scala as new base
		base, prod, scala = scala, base, prod
		Curve448ScalaMult(&prod, &scala, &base)
	}
	fmt.Printf("1000 iteration: %x\n", prod[:])
	assert.Equal(t, it1000, prod[:])
}

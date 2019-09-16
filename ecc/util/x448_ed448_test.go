package util

import (
	"ecc/ecc"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestX448Rfc7748PrivPub(t *testing.T) {
	curve448 := ecc.Curve448()
	// test vector from rfc7748
	it1, _ := hex.DecodeString("3f482c8a9f19b01e6c46ee9711d9dc14fd4bf67af30765c2ae2b846a4d23a8cd0db897086239492caf350b51f833868b9bc2b3bca9cf4113")
	it1000, _ := hex.DecodeString("aa3b4749d55b9daf1e5b00288826c467274ce3ebbdd5c17b975e09d4af6c67cf10d087202db88286e2b79fceea3ec353ef54faa26e219f38")

	// initialize base and scala both as [5 0 0 0...]
	scala := make([]byte, 56)
	scala[0] = 5

	// after 1 iteration
	base := new(big.Int).SetBytes(reverse(scala))
	prod := curve448.ScalaMultProjectiveU(base, reverse(x448process(scala)))
	fmt.Printf("1 iteration: %x\n", reverse(prod.Bytes()))
	assert.Equal(t, it1, reverse(prod.Bytes()))

	// after 1000 iteration
	for i := 2; i <= 1000; i++ {
		// use old scala as new base
		base.SetBytes(reverse(scala)) // scala is little endian, reverse to big-endian to big.Int
		// use product as new scala
		scala = make([]byte, 56)
		copy(scala, reverse(prod.Bytes())) // set scala as little-endian
		prod = curve448.ScalaMultProjectiveU(base, reverse(x448process(scala)))
	}
	fmt.Printf("1000 iteration: %x\n", reverse(prod.Bytes()))
	assert.Equal(t, it1000, reverse(prod.Bytes()))
}

func TestX448Rfc7748DiffieHellman(t *testing.T) {
	// test vectors from in rfc7748, all in little-endian
	scala1, _ := hex.DecodeString("9a8f4925d1519f5775cf46b04b5800d4ee9ee8bae8bc5565d498c28dd9c9baf574a9419744897391006382a6f127ab1d9ac2d8c0a598726b")
	scala2, _ := hex.DecodeString("1c306a7ac2a0e2e0990b294470cba339e6453772b075811d8fad0d1d6927c120bb5ee8972b0d3e21374c9c921b09d1b0366f10b65173992d")
	pub1Ex, _ := hex.DecodeString("9b08f7cc31b7e3e67d22d5aea121074a273bd2b83de09c63faa73d2c22c5d9bbc836647241d953d40c5b12da88120d53177f80e532c41fa0")
	pub2Ex, _ := hex.DecodeString("3eb7a829b0cd20f5bcfc0b599b6feccf6da4627107bdb0d4f345b43027d8b972fc3e34fb4232a13ca706dcb57aec3dae07bdc1c67bf33609")
	shareE, _ := hex.DecodeString("07fff4181ac6cc95ec1c16a94a0f74d12da232ce40a77552281d282bb60c0b56fd2464c335543936521c24403085d59a449a5037514a879d")

	// test ecc package
	// 'ecc' scala multiplication input is big-endian
	baseU := new(big.Int).SetInt64(5)
	curve448 := ecc.Curve448()
	pub1 := curve448.ScalaMultProjectiveU(baseU, reverse(x448process(scala1)))
	pub2 := curve448.ScalaMultProjectiveU(baseU, reverse(x448process(scala2)))
	share1 := curve448.ScalaMultProjectiveU(pub1, reverse(x448process(scala2)))
	share2 := curve448.ScalaMultProjectiveU(pub2, reverse(x448process(scala1)))
	fmt.Printf("ecc pub1: %x\n", reverse(pub1.Bytes()))
	fmt.Printf("ecc pub2: %x\n", reverse(pub2.Bytes()))
	fmt.Printf("ecc share1: %x\n", reverse(share1.Bytes()))
	fmt.Printf("ecc share2: %x\n", reverse(share2.Bytes()))
	assert.Equal(t, reverse(pub1.Bytes()), pub1Ex, "pub1 equal")
	assert.Equal(t, reverse(pub2.Bytes()), pub2Ex, "pub2 equal")
	assert.Equal(t, reverse(share1.Bytes()), shareE, "share equal")
	assert.Equal(t, share1.Bytes(), share2.Bytes(), "share 1,2 equal")
}

func x448process(scala []byte) []byte {
	e := make([]byte, len(scala))
	copy(e, scala)
	e[0] &= 252
	e[55] |= 128
	return e
}

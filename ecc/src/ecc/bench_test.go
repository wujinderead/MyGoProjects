package ecc

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

func BenchmarkEcCurve_ScalaMult(b *testing.B) {
	byter, _ := hex.DecodeString("b76ab22849759e9d31f1b656923c8dfef45e7d82fd82b383ba501f29b528c3d1")
	curve := _EC_X9_62_PRIME_256V1
	p256 := elliptic.P256()
	point := &EcPoint{p256.Params().Gx, p256.Params().Gx}
	for i := 0; i < b.N; i++ {
		curve.ScalaMult(point, byter)
	}
}

func BenchmarkEcCurve_GoNativeP256ScalaMult(b *testing.B) {
	byter, _ := hex.DecodeString("b76ab22849759e9d31f1b656923c8dfef45e7d82fd82b383ba501f29b528c3d1")
	p256 := elliptic.P256()
	point := &EcPoint{p256.Params().Gx, p256.Params().Gx}
	for i := 0; i < b.N; i++ {
		p256.ScalarMult(point.X, point.Y, byter)
	}
}

func BenchmarkEdCurve_ScalaMultBase(b *testing.B) {
	byter, _ := hex.DecodeString("c0f9ac4c2dd74f37937dd1a06872705c3c299cf57b8efba8985f00a6f1f392f6")
	ed := Edwards448()
	for i := 0; i < b.N; i++ {
		ed.ScalaMultBase(byter)
	}
}

func BenchmarkEdCurve_ScalaMultBaseProjective(b *testing.B) {
	byter, _ := hex.DecodeString("c0f9ac4c2dd74f37937dd1a06872705c3c299cf57b8efba8985f00a6f1f392f6")
	ed := Edwards448()
	for i := 0; i < b.N; i++ {
		ed.ScalaMultBaseProjective(byter)
	}
}

func TestEquationEc(t *testing.T) {
	curve := _EC_X9_62_PRIME_256V1
	p256 := elliptic.P256()
	byter := make([]byte, 32)
	_, err := rand.Reader.Read(byter)
	if err != nil {
		fmt.Println("gen rand err: ", err.Error())
		return
	}
	fmt.Println("rand: ", hex.EncodeToString(byter))
	x, y := p256.ScalarBaseMult(byter)
	m := curve.ScalaMultBase(byter)
	fmt.Println("x: ", x.String())
	fmt.Println("y: ", y.String())
	fmt.Println(m.X.Cmp(x), m.Y.Cmp(y))
}

func TestEquationEd(t *testing.T) {
	byter := make([]byte, 32)
	_, err := rand.Reader.Read(byter)
	if err != nil {
		fmt.Println("gen rand err: ", err.Error())
		return
	}
	fmt.Println("rand: ", hex.EncodeToString(byter))
	ed := Ed25519()
	m1 := ed.ScalaMultBase(byter)
	m2 := ed.ScalaMultBaseProjective(byter)
	fmt.Println("x: ", m1.X.String())
	fmt.Println("y: ", m1.Y.String())
	fmt.Println(m1.X.Cmp(m2.X), m1.Y.Cmp(m2.Y))
}

func TestEquationJava(t *testing.T) {
	curve := _EC_SECG_PRIME_256K1
	s, _ := new(big.Int).SetString("27519130846651076635714601172252979491810019324250252554655444934306184823446", 10)
	point := curve.ScalaMultBase(s.Bytes())
	fmt.Println("X: ", point.X.String())
	fmt.Println("Y: ", point.Y.String())
}

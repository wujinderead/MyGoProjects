package ecc

import (
	"crypto/elliptic"
	"encoding/hex"
	"testing"
)

func BenchmarkEcCurve_ScalaMult(b *testing.B) {
	byter, _ := hex.DecodeString("b76ab22849759e9d31f1b656923c8dfef45e7d82fd82b383ba501f29b528c3d1")
	curve, _ := GetFpCurve("prime256v1")
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

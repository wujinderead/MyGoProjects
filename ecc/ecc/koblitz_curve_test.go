package ecc

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestKoblitzCurve_Arithmetic(t *testing.T) {
	k1, k2 := make([]byte, 32), make([]byte, 32)
	rand.Read(k1)
	rand.Read(k2)
	fmt.Printf("k1: %s, k2: %s\n", hex.EncodeToString(k1), hex.EncodeToString(k2))
	initEd.Do(initEcCurves)
	for i, name := range KoblitzCurveNames {
		curve, err := GetFpCurve(name)
		if err != nil {
			t.Errorf("koblitz curve '%s' not esist.\n", name)
			continue
		}
		fmt.Printf("\n%d %s: \n", i, name)
		kz := (*KoblitzCurve)(curve)
		base := &EcPoint{curve.X, curve.Y}

		// test multiply
		mul1 := curve.ScalaMultBase(k1)
		mul2 := curve.ScalaMultBase(k2)
		kz_mul1 := kz.ScalaMult(base, k1)
		kz_mul2 := kz.ScalaMult(base, k2)
		fmt.Println("mul1: ", mul1)
		fmt.Println("mul2: ", mul2)
		fmt.Println(mul1.Equals(kz_mul1), mul2.Equals(kz_mul2))

		// test double
		dbl := curve.Add(mul1, mul1)
		kz_dbl := kz.Double(mul1)
		fmt.Println("dbl: ", dbl)
		fmt.Println(dbl.Equals(kz_dbl))

		// test add
		sum := curve.Add(mul2, dbl)
		kz_sum := kz.Add(mul2, dbl)
		fmt.Println("sum: ", sum)
		fmt.Println(sum.Equals(kz_sum))
	}
}

func TestKoblitzCurve_Add_Double(t *testing.T) {
	initEd.Do(initEcCurves)
	for i, name := range KoblitzCurveNames {
		curve, err := GetFpCurve(name)
		if err != nil {
			t.Errorf("koblitz curve '%s' not esist.\n", name)
			continue
		}
		fmt.Printf("\n%d %s: \n", i, name)
		kz := (*KoblitzCurve)(curve)
		p1 := &EcPoint{curve.X, curve.Y}

		// test add and double
		p2 := curve.Add(p1, p1)
		p3 := curve.Add(p2, p1)
		p4 := curve.Add(p2, p2)
		p5 := curve.Add(p2, p3)
		kz_p2 := kz.Double(p1)
		kz_p3 := kz.Add(p1, p2)
		kz_p4 := kz.Double(p2)
		kz_p5 := kz.Add(p2, p3)
		fmt.Println(p2.Equals(kz_p2), p3.Equals(kz_p3), p4.Equals(kz_p4), p5.Equals(kz_p5))
	}
}

func TestEquation_KoblitzCurve(t *testing.T) {
	initEd.Do(initEcCurves)
	for i, name := range KoblitzCurveNames {
		_, k1, px1, py1 := getOpensslEcPrivateKey(name) // the scalar multiply from of openssl
		_, k2, px2, py2 := getOpensslEcPrivateKey(name) // the scalar multiply from of openssl
		curve, err := GetFpCurve(name)
		if err != nil {
			t.Errorf("koblitz curve '%s' not esist.\n", name)
			continue
		}
		fmt.Printf("\n%d %s: \n", i, name)
		kz := (*KoblitzCurve)(curve)
		mul1 := kz.ScalaMultBase(k1)
		mul2 := kz.ScalaMultBase(k2)
		fmt.Printf("k1: %s\nk2: %s\n", hex.EncodeToString(k1), hex.EncodeToString(k2))
		fmt.Println("mul1: ", mul1)
		fmt.Println("mul2: ", mul2)
		fmt.Println(mul1.X.Cmp(px1), mul1.Y.Cmp(py1))
		fmt.Println(mul2.X.Cmp(px2), mul2.Y.Cmp(py2))
	}
}

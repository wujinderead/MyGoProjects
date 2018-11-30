package ecc

import (
	"encoding/hex"
	"testing"
	"crypto/rand"
	"fmt"
	"crypto/elliptic"
)

func TestAneg3Curve_Arithmetic(t *testing.T) {
	k1, k2 := make([]byte, 32), make([]byte, 32)
	rand.Read(k1)
	rand.Read(k2)
	fmt.Printf("k1: %s, k2: %s\n", hex.EncodeToString(k1), hex.EncodeToString(k2))
	initEd.Do(initEcCurves)
	for i, name := range Aneg3CurveNames {
		curve, err := GetFpCurve(name)
		if err != nil {
			t.Errorf("normal curve '%s' not esist.\n", name)
			continue
		}
		fmt.Printf("\n%d %s: \n", i, name)
		a3 := (*Aneg3Curve)(curve)
		base := &EcPoint{curve.X, curve.Y}

		// test multiply
		mul1 := curve.ScalaMultBase(k1)
		mul2 := curve.ScalaMultBase(k2)
		a3_mul1 := a3.ScalaMult(base, k1)
		a3_mul2 := a3.ScalaMult(base, k2)
		fmt.Println("mul1: ", mul1)
		fmt.Println("mul2: ", mul2)
		fmt.Println(mul1.Equals(a3_mul1), mul2.Equals(a3_mul2))

		// test double
		dbl := curve.Add(mul1, mul1)
		a3_dbl := a3.Double(mul1)
		fmt.Println("dbl: ", dbl)
		fmt.Println(dbl.Equals(a3_dbl))

		// test add
		sum := curve.Add(mul2, dbl)
		a3_sum := a3.Add(mul2, dbl)
		fmt.Println("sum: ", sum)
		fmt.Println(sum.Equals(a3_sum))
	}
}

func TestAneg3Curve_Add_Double(t *testing.T) {
	initEd.Do(initEcCurves)
	for i, name := range Aneg3CurveNames {
		curve, err := GetFpCurve(name)
		if err != nil {
			t.Errorf("normal curve '%s' not esist.\n", name)
			continue
		}
		fmt.Printf("\n%d %s: \n", i, name)
		a3 := (*Aneg3Curve)(curve)
		p1 := &EcPoint{curve.X, curve.Y}

		// test add and double
		p2 := curve.Add(p1, p1)
		p3 := curve.Add(p2, p1)
		p4 := curve.Add(p2, p2)
		p5 := curve.Add(p2, p3)
		a3_p2 := a3.Double(p1)
		a3_p3 := a3.Add(p1, p2)
		a3_p4 := a3.Double(p2)
		a3_p5 := a3.Add(p2, p3)
		fmt.Println(p2.Equals(a3_p2), p3.Equals(a3_p3), p4.Equals(a3_p4), p5.Equals(a3_p5))
	}
}

func TestEquation_Aneg3Curve(t *testing.T) {
	{
		curve, _ := GetFpCurve("prime256v1")
		p256 := elliptic.P256()
		byter := make([]byte, 32)
		_, err := rand.Reader.Read(byter)
		if err != nil {
			fmt.Println("gen rand err: ", err.Error())
			return
		}
		fmt.Println("rand: ", hex.EncodeToString(byter))
		x, y := p256.ScalarBaseMult(byter)  // golang native scala multiply (in Jacobian coordinates)
		m := curve.ScalaMultBase(byter)     // direct arithmetic
		n := (*Aneg3Curve)(curve).ScalaMult(&EcPoint{curve.X, curve.Y}, byter) // in projective coordinates
		fmt.Println("x: ", x.String())
		fmt.Println("y: ", y.String())
		fmt.Println(m.X.Cmp(x), m.Y.Cmp(y))
		fmt.Println(n.X.Cmp(x), n.Y.Cmp(y))
	}
	{
		curve, _ := GetFpCurve("secp521r1")
		p521 := elliptic.P521()
		byter := make([]byte, 32)
		_, err := rand.Reader.Read(byter)
		if err != nil {
			fmt.Println("gen rand err: ", err.Error())
			return
		}
		fmt.Println("rand: ", hex.EncodeToString(byter))
		x, y := p521.ScalarBaseMult(byter)  // golang native scala multiply
		m := curve.ScalaMultBase(byter)     // our scala multiply
		n := (*Aneg3Curve)(curve).ScalaMult(&EcPoint{curve.X, curve.Y}, byter) // in projective coordinates
		fmt.Println("x: ", x.String())
		fmt.Println("y: ", y.String())
		fmt.Println(m.X.Cmp(x), m.Y.Cmp(y))
		fmt.Println(n.X.Cmp(x), n.Y.Cmp(y))
	}
}

func TestEquation_Aneg3Curve_Openssl(t *testing.T) {
	initEd.Do(initEcCurves)
	for i, name := range Aneg3CurveNames {
		_, k1, px1, py1 := getOpensslEcPrivateKey(name)  // the scalar multiply from of openssl
		_, k2, px2, py2 := getOpensslEcPrivateKey(name)  // the scalar multiply from of openssl
		curve, err := GetFpCurve(name)
		if err != nil {
			t.Errorf("koblitz curve '%s' not esist.\n", name)
			continue
		}
		fmt.Printf("\n%d %s: \n", i, name)
		kz := (*Aneg3Curve)(curve)
		mul1 := kz.ScalaMultBase(k1)
		mul2 := kz.ScalaMultBase(k2)
		fmt.Printf("k1: %s\nk2: %s\n", hex.EncodeToString(k1), hex.EncodeToString(k2))
		fmt.Println("mul1: ", mul1)
		fmt.Println("mul2: ", mul2)
		fmt.Println(mul1.X.Cmp(px1), mul1.Y.Cmp(py1))
		fmt.Println(mul2.X.Cmp(px2), mul2.Y.Cmp(py2))
	}
}

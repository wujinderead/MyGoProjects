package ecc

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

func TestF2mCurve_IsOnCurve(t *testing.T) {
	initEd.Do(initEcCurves)
	for i, name := range F2mCurveNames {
		fmt.Printf("\n%d %s: \n", i, name)
		curve, err := GetF2mCurve(name)
		if err != nil {
			fmt.Printf("error getting binary field curve '%s'\n", name)
			continue
		}
		if !curve.IsOnCurve(curve.X, curve.Y) {
			fmt.Printf("binary field curve '%s' base is not on curve\n", name)
		}
	}
}

func TestEquation_F2m_Openssl(t *testing.T) {
	initEd.Do(initEcCurves)
	for i, name := range F2mCurveNames {
		_, k1, px1, py1 := getOpensslEcPrivateKey(name)  // the scalar multiply from of openssl
		_, k2, px2, py2 := getOpensslEcPrivateKey(name)  // the scalar multiply from of openssl
		curve, err := GetF2mCurve(name)
		if err != nil {
			t.Errorf("koblitz curve '%s' not esist.\n", name)
			continue
		}
		fmt.Printf("\n%d %s: \n", i, name)
		fmt.Printf("k1: %s\nk2: %s\n", hex.EncodeToString(k1), hex.EncodeToString(k2))
		mul1 := &EcPoint{px1, py1}
		mul2 := &EcPoint{px2, py2}
		f2m_mul1 := curve.ScalaMultBase(k1)
		f2m_mul2 := curve.ScalaMultBase(k2)
		pro_mul1 := curve.ScalaMultBaseProjective(k1)
		pro_mul2 := curve.ScalaMultBaseProjective(k2)
		jcb_mul1 := curve.ScalaMultBaseJacobian(k1)
		jcb_mul2 := curve.ScalaMultBaseJacobian(k2)
		lpz_mul1 := curve.ScalaMultBaseLopezDahab(k1)
		lpz_mul2 := curve.ScalaMultBaseLopezDahab(k2)
		fmt.Println("mul1: ", mul1)
		fmt.Println("mul2: ", mul2)
		fmt.Println(mul1.Equals(f2m_mul1), mul2.Equals(f2m_mul2))
		fmt.Println(mul1.Equals(pro_mul1), mul2.Equals(pro_mul2))
		fmt.Println(mul1.Equals(jcb_mul1), mul2.Equals(jcb_mul2))
		fmt.Println(mul1.Equals(lpz_mul1), mul2.Equals(lpz_mul2))
	}
}

func TestF2mCurve_arithmetic_F2m(t *testing.T) {
	initEd.Do(initEcCurves)
	for i, name := range F2mCurveNames {
		curve, _ := GetF2mCurve("c2tnb239v1")
		fmt.Printf("\n%d %s: \n", i, name)
		curve.gmul(curve.X, curve.Y)
		fmt.Println(curve.IsOnCurve(curve.X, curve.Y))

		xinv := curve.gmulinv(curve.X)
		xinv_1 := curve.gmulinv(new(big.Int).Sub(curve.X, ONE))
		xinv_2 := curve.gmulinv(new(big.Int).Add(curve.X, TWO))
		fmt.Println(curve.gmul(xinv, curve.X))
		fmt.Println(curve.gmul(xinv_1, new(big.Int).Sub(curve.X, ONE)))
		fmt.Println(curve.gmul(xinv_2, new(big.Int).Add(curve.X, TWO)))

		p1 := &EcPoint{curve.X, curve.Y}
		p2 := curve.Add(p1, p1)
		p3 := curve.Add(p1, p2)
		fmt.Println(curve.IsOnCurve(p1.X, p1.Y))
		fmt.Println(curve.IsOnCurve(p2.X, p2.Y))
		fmt.Println(curve.IsOnCurve(p3.X, p3.Y))

		pro_p2 := curve.DoubleProjective(p1)
		pro_p2_add := curve.AddProjective(p1, p1)
		pro_p3 := curve.AddProjective(p1, p2)
		fmt.Println(p2.Equals(pro_p2), p2.Equals(pro_p2_add), p3.Equals(pro_p3))

		jcb_p2 := curve.DoubleJacobian(p1)
		jcb_p2_add := curve.AddJacobian(p1, p1)
		jcb_p3 := curve.AddJacobian(p1, p2)
		fmt.Println(p2.Equals(jcb_p2), p2.Equals(jcb_p2_add), p3.Equals(jcb_p3))

		lpz_p2 := curve.DoubleLopezDahab(p1)
		lpz_p2_add := curve.AddLopezDahab(p1, p1)
		lpz_p3 := curve.AddLopezDahab(p1, p2)
		fmt.Println(p2.Equals(lpz_p2), p2.Equals(lpz_p2_add), p3.Equals(lpz_p3))

		p100 := curve.ScalaMultBase([]byte{100})
		pro_p100 := curve.ScalaMultBaseProjective([]byte{100})
		jcb_p100 := curve.ScalaMultBaseJacobian([]byte{100})
		lpz_p100 := curve.ScalaMultBaseLopezDahab([]byte{100})
		fmt.Println(curve.IsOnCurve(p100.X, p100.Y))
		fmt.Println(p100.Equals(pro_p100), p100.Equals(jcb_p100), p100.Equals(lpz_p100))
	}
}

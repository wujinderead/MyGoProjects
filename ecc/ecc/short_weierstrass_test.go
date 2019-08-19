package ecc

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestEquation_ShortWProjective_Openssl(t *testing.T) {
	initEd.Do(initEcCurves)
	for i, name := range FpCurveNames {
		_, k1, px1, py1 := getOpensslEcPrivateKey(name) // the scalar multiply from of openssl
		_, k2, px2, py2 := getOpensslEcPrivateKey(name) // the scalar multiply from of openssl
		curve, err := GetFpCurve(name)
		if err != nil {
			t.Errorf("koblitz curve '%s' not esist.\n", name)
			continue
		}
		fmt.Printf("\n%d %s: \n", i, name)
		mul1 := curve.ScalaMultBaseProjective(k1)
		mul2 := curve.ScalaMultBaseProjective(k2)
		fmt.Printf("k1: %s\nk2: %s\n", hex.EncodeToString(k1), hex.EncodeToString(k2))
		fmt.Println("mul1: ", mul1)
		fmt.Println("mul2: ", mul2)
		fmt.Println(mul1.X.Cmp(px1), mul1.Y.Cmp(py1))
		fmt.Println(mul2.X.Cmp(px2), mul2.Y.Cmp(py2))
	}
}

func TestEquation_ShortWJacobian_Openssl(t *testing.T) {
	initEd.Do(initEcCurves)
	for i, name := range FpCurveNames {
		_, k1, px1, py1 := getOpensslEcPrivateKey(name) // the scalar multiply from of openssl
		_, k2, px2, py2 := getOpensslEcPrivateKey(name) // the scalar multiply from of openssl
		curve, err := GetFpCurve(name)
		if err != nil {
			t.Errorf("koblitz curve '%s' not esist.\n", name)
			continue
		}
		fmt.Printf("\n%d %s: \n", i, name)
		mul1 := curve.ScalaMultBaseJacobian(k1)
		mul2 := curve.ScalaMultBaseJacobian(k2)
		fmt.Printf("k1: %s\nk2: %s\n", hex.EncodeToString(k1), hex.EncodeToString(k2))
		fmt.Println("mul1: ", mul1)
		fmt.Println("mul2: ", mul2)
		fmt.Println(mul1.X.Cmp(px1), mul1.Y.Cmp(py1))
		fmt.Println(mul2.X.Cmp(px2), mul2.Y.Cmp(py2))
	}
}

// 5 types of arithmetic are implemented:
// 'a=0', 'a=-3', 'trivial' in Jacobian coordinates; 'a=-3', 'trivial' in projective coordinates
// in the 3 types of Jacobian implementation, the 'add' is implemented the same, while 'double' is different
// in the 2 types of projective implementation, the 'add' is implemented the same, while 'double' is different
func TestEquation_Trivial_Koblitz_Aneg3(t *testing.T) {
	initEd.Do(initEcCurves)
	types := []string{"trivial", "AnegThree", "koblitz"}
	for type_i, names := range [][]string{TrivialCurveNames, Aneg3CurveNames, KoblitzCurveNames} {
		for i, name := range names {
			curve, err := GetFpCurve(name)
			if err != nil {
				t.Errorf("koblitz curve '%s' not esist.\n", name)
				continue
			}
			fmt.Printf("\n%s[%d] %s: \n", types[type_i], i, name)

			p1 := &EcPoint{curve.X, curve.Y}
			p2 := curve.Add(p1, p1) // arithmetic based on algebraic
			p3 := curve.Add(p2, p1)
			p4 := curve.Add(p2, p2)
			p5 := curve.Add(p2, p3)

			kz := (*KoblitzCurve)(curve)
			kz_p2 := kz.Double(p1) // arithmetic based on a=0 Jacobian
			kz_p3 := kz.Add(p1, p2)
			kz_p4 := kz.Double(p2)
			kz_p5 := kz.Add(p2, p3)
			fmt.Println("kz : ", p2.Equals(kz_p2), p3.Equals(kz_p3), p4.Equals(kz_p4), p5.Equals(kz_p5))

			a3j := curve.ToGoNative() // arithmetic based on a=-3 Jacobian
			a3j_p2x, a3j_p2y := a3j.Double(p1.X, p1.Y)
			a3j_p3x, a3j_p3y := a3j.Add(p1.X, p1.Y, p2.X, p2.Y)
			a3j_p4x, a3j_p4y := a3j.Double(p2.X, p2.Y)
			a3j_p5x, a3j_p5y := a3j.Add(p3.X, p3.Y, p2.X, p2.Y)
			fmt.Println("a3j: ", p2.Equals(&EcPoint{a3j_p2x, a3j_p2y}), p3.Equals(&EcPoint{a3j_p3x, a3j_p3y}),
				p4.Equals(&EcPoint{a3j_p4x, a3j_p4y}), p5.Equals(&EcPoint{a3j_p5x, a3j_p5y}))

			trj_p2 := curve.DoubleJacobian(p1) // arithmetic based on trivial Jacobian
			trj_p3 := curve.AddJacobian(p1, p2)
			trj_p4 := curve.DoubleJacobian(p2)
			trj_p5 := curve.AddJacobian(p2, p3)
			fmt.Println("trj: ", p2.Equals(trj_p2), p3.Equals(trj_p3), p4.Equals(trj_p4), p5.Equals(trj_p5))

			a3p := (*Aneg3Curve)(curve)
			a3p_p2 := a3p.Double(p1) // arithmetic based on a=-3 projective
			a3p_p3 := a3p.Add(p1, p2)
			a3p_p4 := a3p.Double(p2)
			a3p_p5 := a3p.Add(p2, p3)
			fmt.Println("a3p: ", p2.Equals(a3p_p2), p3.Equals(a3p_p3), p4.Equals(a3p_p4), p5.Equals(a3p_p5))

			trp_p2 := curve.DoubleJacobian(p1) // arithmetic based on trivial projective
			trp_p3 := curve.AddJacobian(p1, p2)
			trp_p4 := curve.DoubleJacobian(p2)
			trp_p5 := curve.AddJacobian(p2, p3)
			fmt.Println("trp: ", p2.Equals(trp_p2), p3.Equals(trp_p3), p4.Equals(trp_p4), p5.Equals(trp_p5))
		}
	}
}

package ecc

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	curve := Curve25519()
	method, _ := reflect.TypeOf(curve).MethodByName("ToEdwardCurveForm1")
	results := method.Func.Call([]reflect.Value{reflect.ValueOf(curve), reflect.ValueOf(ONE)})
	ed := results[0].Interface().(*EdCurve)
	sq := results[1].Interface().(*big.Int)
	fmt.Println(ed.Name)
	fmt.Println(sq)
}

func TestMtCurve_ToEdwardCurve(t *testing.T) {
	initAllMontgomery()
	for _, curve := range []*MtCurve{curve25519, m221, m383, curve383187, m511, curve448} {
		fmt.Println("MtCurve name: ", curve.Name)
		toCurve1, _ := reflect.TypeOf(curve).MethodByName("ToEdwardsCurveForm1")
		toCurve2, _ := reflect.TypeOf(curve).MethodByName("ToEdwardsCurveForm2")
		toPoint1, _ := reflect.TypeOf(curve).MethodByName("ToEdwardsPointForm1")
		toPoint2, _ := reflect.TypeOf(curve).MethodByName("ToEdwardsPointForm2")
		for _, args := range [][]interface{}{
			{toCurve1, ONE, toPoint1},
			{toCurve1, NEG_ONE, toPoint1},
			{toCurve2, ONE, toPoint2},
			{toCurve2, NEG_ONE, toPoint2}} {
			curveMethod := args[0].(reflect.Method)
			a := args[1].(*big.Int)
			pointMethod := args[2].(reflect.Method)
			results := curveMethod.Func.Call([]reflect.Value{reflect.ValueOf(curve), reflect.ValueOf(a)})
			ed := results[0].Interface().(*EdCurve)
			sqrtB := results[1].Interface().(*big.Int)
			if sqrtB != nil {
				fmt.Printf("y² ≡ x³ + %sx² + x (mod P)\n", curve.A.String())
				fmt.Printf("%sx² + y² ≡ 1 + %s x²y² (mod P)\n", ed.A.String(), ed.D.String())
				fmt.Printf("sqrtB = %s\n", sqrtB.String())
				points := pointMethod.Func.Call(
					[]reflect.Value{reflect.ValueOf(curve),
						reflect.ValueOf(sqrtB),
						reflect.ValueOf(&EcPoint{curve.Bx, curve.By})})
				p1, p2 := points[0].Interface().(*EcPoint), points[1].Interface().(*EcPoint)
				fmt.Println(ed.IsOnCurve(p1.X, p1.Y), ed.IsOnCurve(p2.X, p2.Y))
				mt579 := curve.ScalaMultBase([]byte{5, 7, 9})
				points = pointMethod.Func.Call(
					[]reflect.Value{reflect.ValueOf(curve),
						reflect.ValueOf(sqrtB),
						reflect.ValueOf(&EcPoint{mt579.X, mt579.Y})})
				ed579_1, ed579_2 := points[0].Interface().(*EcPoint), points[1].Interface().(*EcPoint)
				ed.Bx, ed.By = p1.X, p1.Y
				mul1 := ed.ScalaMultBase([]byte{5, 7, 9})
				ed.Bx, ed.By = p2.X, p2.Y
				mul2 := ed.ScalaMultBase([]byte{5, 7, 9})
				fmt.Println("ba1: ", p1)
				fmt.Println("ba2: ", p2)
				fmt.Println("ed1: ", ed579_1)
				fmt.Println("ed2: ", ed579_2)
				fmt.Println("ml1: ", mul1)
				fmt.Println("ml2: ", mul2)
				fmt.Println(mul1.Equals(ed579_1), mul2.Equals(ed579_2))
			}
			fmt.Println()
		}
	}
}

func TestEdCurve_ToMontgomeryCurve(t *testing.T) {
	initAllEdwards()
	for _, curve := range []*EdCurve{e222, curve1174, ed25519, e382, curve41417, edwards448, e521} {
		fmt.Println("EdCurve name: ", curve.Name)
		toCurve1, _ := reflect.TypeOf(curve).MethodByName("ToMontgomeryCurveForm1")
		toCurve2, _ := reflect.TypeOf(curve).MethodByName("ToMontgomeryCurveForm2")
		toPoint1, _ := reflect.TypeOf(curve).MethodByName("ToMontgomeryPointForm1")
		toPoint2, _ := reflect.TypeOf(curve).MethodByName("ToMontgomeryPointForm2")
		for _, args := range [][]interface{}{
			{toCurve1, toPoint1},
			{toCurve2, toPoint2}} {
			curveMethod := args[0].(reflect.Method)
			pointMethod := args[1].(reflect.Method)
			results := curveMethod.Func.Call([]reflect.Value{reflect.ValueOf(curve)})
			mt := results[0].Interface().(*MtCurve)
			sqrtB := results[1].Interface().(*big.Int)
			if sqrtB != nil {
				fmt.Printf("%sx² + y² ≡ 1 + %s x²y² (mod P)\n", curve.A.String(), curve.D.String())
				fmt.Printf("y² ≡ x³ + %sx² + x (mod P)\n", mt.A.String())
				fmt.Printf("sqrtB = %s\n", sqrtB.String())
				points := pointMethod.Func.Call(
					[]reflect.Value{reflect.ValueOf(curve),
						reflect.ValueOf(sqrtB),
						reflect.ValueOf(&EcPoint{curve.Bx, curve.By})})
				p1, p2 := points[0].Interface().(*EcPoint), points[1].Interface().(*EcPoint)
				fmt.Println(mt.IsOnCurve(p1.X, p1.Y), mt.IsOnCurve(p2.X, p2.Y))
				ed579 := curve.ScalaMultBase([]byte{5, 7, 9})
				points = pointMethod.Func.Call(
					[]reflect.Value{reflect.ValueOf(curve),
						reflect.ValueOf(sqrtB),
						reflect.ValueOf(&EcPoint{ed579.X, ed579.Y})})
				mt579_1, mt579_2 := points[0].Interface().(*EcPoint), points[1].Interface().(*EcPoint)
				mt.Bx, mt.By = p1.X, p1.Y
				mul1 := mt.ScalaMultBase([]byte{5, 7, 9})
				mt.Bx, mt.By = p2.X, p2.Y
				mul2 := mt.ScalaMultBase([]byte{5, 7, 9})
				fmt.Println("ba1: ", p1)
				fmt.Println("ba2: ", p2)
				fmt.Println("ed1: ", mt579_1)
				fmt.Println("ed2: ", mt579_2)
				fmt.Println("ml1: ", mul1)
				fmt.Println("ml2: ", mul2)
				fmt.Println(mul1.Equals(mt579_1), mul2.Equals(mt579_2))
			}
			fmt.Println()
		}
	}
}

func TestMtSqrtB(t *testing.T) {
	initAllMontgomery()
	for _, curve := range []*MtCurve{curve25519, m221, m383, curve383187, m511, curve448} {
		fmt.Println("name: ", curve.Name)
		a := curve.A.Int64()
		fmt.Println("curve a: ", a)
		fmt.Println("sqrt  a+2: ", new(big.Int).ModSqrt(new(big.Int).SetInt64(a+2), curve.P))
		fmt.Println("sqrt  a-2: ", new(big.Int).ModSqrt(new(big.Int).SetInt64(a-2), curve.P))
		fmt.Println("sqrt -a+2: ", new(big.Int).ModSqrt(new(big.Int).SetInt64(-a+2), curve.P))
		fmt.Println("sqrt -a-2: ", new(big.Int).ModSqrt(new(big.Int).SetInt64(-a-2), curve.P))
	}
}

func TestEdSqrtB(t *testing.T) {
	initAllEdwards()
	for _, curve := range []*EdCurve{e222, curve1174, ed25519, e382, curve41417, edwards448, e521} {
		fmt.Println("name: ", curve.Name)
		_, sqb1 := curve.ToMontgomeryCurveForm1()
		_, sqb2 := curve.ToMontgomeryCurveForm2()
		fmt.Println("form1 sqB: ", sqb1)
		fmt.Println("form2 sqB: ", sqb2)
	}
}

func TestEdCurve_ToMontgomeryPointForm1(t *testing.T) {
	curve := Ed25519()
	p := &EcPoint{curve.Bx, curve.By}
	fmt.Println("ed base: ", p)
	B := new(big.Int).SetInt64(-486664)
	sqrtB := new(big.Int).ModSqrt(B, curve.P)
	fmt.Println("sqrtB: ", sqrtB)
	oneSubY := new(big.Int).Sub(ONE, p.Y) // 1-y
	oneAddY := new(big.Int).Add(ONE, p.Y) // 1+y
	p1, p2 := NewPoint(), NewPoint()
	p1.X = ModFraction(oneAddY, oneSubY, curve.P) // (1+y)/(1-y)
	fmt.Println("p1x: ", p1.X)
	p1.Y = ModFraction(p1.X, p.X, curve.P)        // u/x
	fmt.Println("p1y: ", p1.Y)
	p1.Y.Mul(p1.Y, sqrtB)                         // sqrtB * u/x
	fmt.Println("p1y: ", p1.Y)
	p1.Y.Mod(p1.Y, curve.P)
	fmt.Println("p1y: ", p1.Y)

	p2.X = ModFraction(oneSubY, oneAddY, curve.P) // (1+y)/(1-y)
	fmt.Println("p2x: ", p2.X)
	p2.Y = ModFraction(p2.X, p.X, curve.P)        // u/x
	fmt.Println("p2y: ", p2.Y)
	p2.Y.Mul(p2.Y, sqrtB)                         // sqrtB * u/x
	fmt.Println("p2y: ", p2.Y)
	p2.Y.Mod(p2.Y, curve.P)
	fmt.Println("p2y: ", p2.Y)

	mt := Curve25519()
	fmt.Println(mt.IsOnCurve(p1.X, p1.Y), mt.IsOnCurve(p1.X, p1.Y))
}

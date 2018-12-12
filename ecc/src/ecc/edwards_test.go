package ecc

import (
	"fmt"
	"math/big"
	"testing"
)

func TestEdCurve_IsOnCurve(t *testing.T) {
	initAllEdwards()
	for _, curve := range []*EdCurve{e222, curve1174, ed25519, e382, curve41417, edwards448, e521} {
		four := new(big.Int).SetInt64(4)
		res := new(big.Int)
		fmt.Println(curve.Name, curve.Pstr, ", p mod 4 =", res.Mod(curve.P, four).String())
		fmt.Println(curve.IsOnCurve(curve.Bx, curve.By))
	}
}

func TestEdCurve_ScalaMult(t *testing.T) {
	initAllEdwards()
	part := new(big.Int)
	for _, curve := range []*EdCurve{e222, curve1174, ed25519, e382, curve41417, edwards448, e521} {
		fmt.Println("\n", curve.Name, ":")
		fmt.Println("ba", (&EcPoint{curve.Bx, curve.By}).ToStr())
		fmt.Printf("P: %x\n", curve.P.Bytes())
		for i := int64(1); i < 6; i++ {
			fmt.Println("\n", i, ":")
			part.SetInt64(i)
			p := curve.ScalaMultBase(curve.Order.Bytes())
			fmt.Println("base * order = ", p.ToStr())
			re := curve.ScalaMultBase(new(big.Int).Sub(curve.Order, part).Bytes())
			pa := curve.ScalaMultBase(part.Bytes())
			fmt.Println("re", re.ToStr())
			fmt.Println("pa", pa.ToStr())
			fmt.Println(pa.Y.Cmp(re.Y), new(big.Int).Add(pa.X, re.X).Cmp(curve.P))
		}
	}
}

func TestEdCurve_AddProjective(t *testing.T) {
	initAllEdwards()
	for _, curve := range []*EdCurve{e222, curve1174, ed25519, e382, curve41417, edwards448, e521} {
		fmt.Println("\n", curve.Name, ":")
		fmt.Println("ba", (&EcPoint{curve.Bx, curve.By}).ToStr())
		fmt.Printf("P: %x\n", curve.P.Bytes())
		aff_2p := curve.Add(&EcPoint{curve.Bx, curve.By}, &EcPoint{curve.Bx, curve.By})
		aff_3p := curve.Add(aff_2p, &EcPoint{curve.Bx, curve.By})
		aff_4p := curve.Add(aff_2p, aff_2p)
		aff_5p := curve.Add(aff_2p, aff_3p)
		aff_6p := curve.Add(aff_3p, aff_3p)
		aff_7p := curve.Add(aff_2p, aff_5p)
		pro_2p := curve.Double(&EcPoint{curve.Bx, curve.By})
		pro_2p_add := curve.AddProjective(&EcPoint{curve.Bx, curve.By}, &EcPoint{curve.Bx, curve.By})
		pro_3p := curve.AddProjective(pro_2p, &EcPoint{curve.Bx, curve.By})
		pro_4p := curve.Double(pro_2p)
		pro_4p_add := curve.AddProjective(pro_2p, pro_2p)
		pro_5p := curve.AddProjective(pro_2p, pro_3p)
		pro_6p := curve.Double(pro_3p)
		pro_6p_add := curve.AddProjective(pro_3p, pro_3p)
		pro_7p := curve.AddProjective(pro_2p, pro_5p)
		fmt.Println("aff_2p: ", aff_2p.ToStr())
		fmt.Println("pro_2p: ", pro_2p.ToStr())
		fmt.Println("aff_3p: ", aff_3p.ToStr())
		fmt.Println("pro_3p: ", pro_3p.ToStr())
		fmt.Println("aff_4p: ", aff_4p.ToStr())
		fmt.Println("pro_4p: ", pro_4p.ToStr())
		fmt.Println("aff_5p: ", aff_5p.ToStr())
		fmt.Println("pro_5p: ", pro_5p.ToStr())
		fmt.Println("aff_6p: ", aff_6p.ToStr())
		fmt.Println("pro_6p: ", pro_6p.ToStr())
		fmt.Println("aff_7p: ", aff_7p.ToStr())
		fmt.Println("pro_7p: ", pro_7p.ToStr())
		fmt.Println(aff_2p.Equals(pro_2p), aff_3p.Equals(pro_3p), aff_4p.Equals(pro_4p),
			aff_5p.Equals(pro_5p), aff_6p.Equals(pro_6p), aff_7p.Equals(pro_7p),
			pro_2p.Equals(pro_2p_add), pro_4p.Equals(pro_4p_add), pro_6p.Equals(pro_6p_add))
	}
}
func TestEdCurve_ScalaMultProjective(t *testing.T) {
	initAllEdwards()
	part := new(big.Int)
	for _, curve := range []*EdCurve{e222, curve1174, ed25519, e382, curve41417, edwards448, e521} {
		fmt.Println("\n", curve.Name, ":")
		fmt.Println("ba", (&EcPoint{curve.Bx, curve.By}).ToStr())
		fmt.Printf("P: %x\n", curve.P.Bytes())
		p := curve.ScalaMultBase(curve.Order.Bytes())
		fmt.Println("affine base * order = ", p.ToStr())
		p1 := curve.ScalaMultBaseProjective(curve.Order.Bytes())
		fmt.Println("projct base * order = ", p1.ToStr())
		for i := int64(1); i < 6; i++ {
			fmt.Println("\n", i, ":")
			part.SetInt64(i)
			re := curve.ScalaMultBase(new(big.Int).Sub(curve.Order, part).Bytes())
			pa := curve.ScalaMultBase(part.Bytes())
			re1 := curve.ScalaMultBaseProjective(new(big.Int).Sub(curve.Order, part).Bytes())
			pa1 := curve.ScalaMultBaseProjective(part.Bytes())
			fmt.Println("re aff", re.ToStr())
			fmt.Println("pa aff", pa.ToStr())
			fmt.Println("re pro", re1.ToStr())
			fmt.Println("pa pro", pa1.ToStr())
			fmt.Println(re1.Equals(re), pa.Equals(pa1))
		}
	}
}

package ecc

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMtCurve_IsOnCurve(t *testing.T) {
	initAllMontgomery()
	for _, curve := range []*MtCurve{curve25519, m221, m383, curve383187, m511, curve448} {
		four := new(big.Int).SetInt64(4)
		res := new(big.Int)
		fmt.Println(curve.Name, curve.Pstr, ", p mod 4 =", res.Mod(curve.P, four).String())
		fmt.Println(curve.IsOnCurve(curve.Bx, curve.By))
	}
}

func TestMtCurve_Add(t *testing.T) {
	initAllMontgomery()
	x1, _ := new(big.Int).SetString("9", 10)
	y1, _ := new(big.Int).SetString("43114425171068552920764898935933967039370386198203806730763910166200978582548", 10)
	x2, _ := new(big.Int).SetString("14847277145635483483963372537557091634710985132825781088887140890597596352251", 10)
	y2, _ := new(big.Int).SetString("48981431527428949880507557032295310859754924433568441600873610210018059225738", 10)
	ax, _ := new(big.Int).SetString("12697861248284385512127539163427099897745340918349830473877503196793995869202", 10)
	ay, _ := new(big.Int).SetString("39113539887452079713994524130201898724087778094240617142109147539155741236674", 10)
	sum := curve25519.Add(&EcPoint{x1, y1}, &EcPoint{x2, y2})
	fmt.Printf("%x, %x\n", ax.Bytes(), ay.Bytes())
	fmt.Printf("%x, %x\n", sum.X.Bytes(), sum.Y.Bytes())
}

func TestMtCurve_ScalaMult(t *testing.T) {
	initAllMontgomery()
	part := new(big.Int)
	for _, curve := range []*MtCurve{curve25519, m221, m383, curve383187, m511, curve448} {
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
			fmt.Println(pa.X.Cmp(re.X), new(big.Int).Add(pa.Y, re.Y).Cmp(curve.P))
		}
	}
}

func TestXX(t *testing.T) {
	b, _ := new(big.Int).SetString("355293926785568175264127502063783334808976399387714271831880898435169088786967410002932673765864550910142774147268105838985595290606362", 10)
	fmt.Printf("%x", b.Bytes())
}

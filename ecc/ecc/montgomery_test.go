package ecc

import (
	"crypto/rand"
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

func TestMtCurve_doubleProjective_diffAddprojective(t *testing.T) {
	initAllMontgomery()
	for _, curve := range []*MtCurve{curve25519, m221, m383, curve383187, m511, curve448} {
		fmt.Println("\n", curve.Name, ":")
		p := &EcPoint{curve.Bx, curve.By}
		fmt.Println("ba:", p.ToStr())
		fmt.Printf("P: %x\n", curve.P.Bytes())
		p2 := curve.Add(p, p)
		p3 := curve.Add(p2, p)
		p4 := curve.Add(p2, p2)
		p5 := curve.Add(p2, p3)
		p6 := curve.Add(p3, p3)
		p7 := curve.Add(p3, p4)
		z1 := zForAffine(curve.Bx, curve.By)
		x1 := p.X

		x2, z2 := curve.doubleProjective(x1, z1)
		X2 := curve.affineFromProjectiveX(x2, z2)
		fmt.Println(p2.X.Cmp(X2))

		x3, z3 := curve.diffAddProjective(x1, z1, x2, z2, x1, z1)
		X3 := curve.affineFromProjectiveX(x3, z3)
		fmt.Println(p3.X.Cmp(X3))

		x4, z4 := curve.diffAddProjective(x2, z2, x1, z1, x3, z3)
		X4 := curve.affineFromProjectiveX(x4, z4)
		fmt.Println(p4.X.Cmp(X4))

		for _, x5args := range [][]*big.Int{
			{x1, z1, x2, z2, x3, z3},
			{x1, z1, x3, z3, x2, z2}} {
			x5, z5 := curve.diffAddProjective(x5args[0], x5args[1], x5args[2], x5args[3], x5args[4], x5args[5])
			X5 := curve.affineFromProjectiveX(x5, z5)
			fmt.Println("x5z5: ", x5.Text(16), z5.Text(16))
			fmt.Println(p5.X.Cmp(X5))
			for _, x6args := range [][]*big.Int{
				{x2, z2, x4, z4, x2, z2},
				{x4, z4, x5, z5, x1, z1}} {
				x6, z6 := curve.diffAddProjective(x6args[0], x6args[1], x6args[2], x6args[3], x6args[4], x6args[5])
				X6 := curve.affineFromProjectiveX(x6, z6)
				fmt.Println("x6z6: ", x6.Text(16), z6.Text(16))
				fmt.Println(p6.X.Cmp(X6))
				for _, x7args := range [][]*big.Int{
					{x1, z1, x4, z4, x3, z3},
					{x3, z3, x5, z5, x2, z2},
					{x5, z5, x1, z1, x6, z6}} {
					x7, z7 := curve.diffAddProjective(x7args[0], x7args[1], x7args[2], x7args[3], x7args[4], x7args[5])
					X7 := curve.affineFromProjectiveX(x7, z7)
					fmt.Println("x7z7: ", x7.Text(16), z7.Text(16))
					fmt.Println(p7.X.Cmp(X7))
				}
			}
		}
	}
}

func TestMtCurve_ScalaMultProjective(t *testing.T) {
	initAllMontgomery()
	reader := rand.Reader
	for _, curve := range []*MtCurve{curve25519, m221, m383, curve383187, m511, curve448} {
		fmt.Println("\n", curve.Name, ":")
		size := (curve.P.BitLen() - 8) / 8
		byter := make([]byte, size)
		for i := 0; i < 10; i++ {
			reader.Read(byter)
			p := curve.ScalaMultBase(byter)
			pp := curve.ScalaMultBaseProjective(byter)
			fmt.Println(p)
			fmt.Println(pp)
			fmt.Println(p.X.Cmp(pp.X), p.Y.Cmp(pp.Y))
		}
	}
}

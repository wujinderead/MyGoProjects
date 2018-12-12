package ecc

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestGf2m_mul(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var a = big.Word(rand.Uint64())
	var b = big.Word(rand.Uint64())
	var r1, r0 = new(big.Word), new(big.Word)
	gf2m_mul_1x1(r1, r0, a, b)
	bigmul := mulBigInt(new(big.Int).SetBits([]big.Word{a}), new(big.Int).SetBits([]big.Word{b}))
	fmt.Println("r1: ", strconv.FormatUint(uint64(*r1), 2))
	fmt.Println("r0: ", strconv.FormatUint(uint64(*r0), 2))
	funcmul := new(big.Int).SetBits([]big.Word{*r0, *r1})
	fmt.Println("equals: ", bigmul.Cmp(funcmul))

	r := make([]big.Word, 4)
	a1 := big.Word(rand.Uint64())
	a0 := big.Word(rand.Uint64())
	b1 := big.Word(rand.Uint64())
	b0 := big.Word(rand.Uint64())
	gf2m_mul_2x2(r, a1, a0, b1, b0)
	bigmul = mulBigInt(new(big.Int).SetBits([]big.Word{a0, a1}), new(big.Int).SetBits([]big.Word{b0, b1}))
	fmt.Println("r3: ", strconv.FormatUint(uint64(r[3]), 2))
	fmt.Println("r2: ", strconv.FormatUint(uint64(r[2]), 2))
	fmt.Println("r1: ", strconv.FormatUint(uint64(r[1]), 2))
	fmt.Println("r0: ", strconv.FormatUint(uint64(r[0]), 2))
	funcmul = new(big.Int).SetBits(r)
	fmt.Println("equals: ", bigmul.Cmp(funcmul))
}

func TestBn_gf2m_mul_sqr(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i:=1; i<=13; i++ {
		for j:=1; j<=16; j++ {
			a := make([]big.Word, i)
			b := make([]big.Word, j)
			fillRandom(a)
			fillRandom(b)
			r1 := new(big.Int).SetBits(bn_gf2m_mul(a, b))
			r2 := mulBigInt(new(big.Int).SetBits(a), new(big.Int).SetBits(b))
			fmt.Printf("i=%d, j=%d, eq=%d, bitlen=%d\n", i, j, r1.Cmp(r2), r1.BitLen())
			sq1 := new(big.Int).SetBits(bn_gf2m_sqr(b))
			sq2 := mulBigInt(new(big.Int).SetBits(b), new(big.Int).SetBits(b))
			fmt.Printf("j=%d, eq=%d, bitlen=%d\n", j, sq1.Cmp(sq2), sq1.BitLen())
		}
	}
}

func TestBn_gf2m_poly2arr(t *testing.T) {
	initEcCurves()
	for i, name := range F2mCurveNames {
		curve, _ := GetF2mCurve(name)
		fmt.Printf("\n%d %s: \n", i, name)
		p := bn_gf2m_poly2arr(curve.P.Bits())
		fmt.Println(curve.P.Text(2))
		fmt.Println(p)
	}
}

func TestBn_gf2m_mod_arr(t *testing.T) {
	initEcCurves()
	for i, name := range F2mCurveNames {
		curve, _ := GetF2mCurve(name)
		fmt.Printf("\n%d %s: \n", i, name)

		p := bn_gf2m_poly2arr(curve.P.Bits())
		aa := curve.X
		bb := curve.Y

		cc := curve.gmulReference(aa, bb)
		pd := bn_gf2m_mul(aa.Bits(), bb.Bits())
		re := new(big.Int).SetBits(bn_gf2m_mod_arr(pd, p))
		fmt.Println(cc)
		fmt.Println(new(big.Int).SetBits(pd))
		fmt.Println(re)
		fmt.Println("equals: ", cc.Cmp(re))
		bn_gf2m_mod_arr_self(pd, p)
		sf := new(big.Int).SetBits(pd)
		fmt.Println(sf)
		fmt.Println("equals: ", cc.Cmp(sf))
	}
}

func TestBn_gf2m_mod_inv_vartime(t *testing.T) {
	initEcCurves()
	rander := rand.New(rand.NewSource(time.Now().Unix()))
	for i, name := range F2mCurveNames {
		curve, _ := GetF2mCurve(name)
		fmt.Printf("\n%d %s: \n", i, name)
		p := bn_gf2m_poly2arr(curve.P.Bits())
		for i:=0; i<10; i++ {
			rdi := new(big.Int).Rand(rander, curve.P)
			inv := bn_gf2m_mod_inv_vartime(rdi.Bits(), curve.P.Bits())
			pro := bn_gf2m_mul(rdi.Bits(), inv)
			bn_gf2m_mod_arr_self(pro, p)
			fmt.Println(new(big.Int).SetBits(pro))
		}
	}
}

func TestBn_numbits(t *testing.T) {
	a, _ := new(big.Int).SetString("1a2b3d4e5f668899734adf", 16)
	fmt.Println(a.BitLen())
	fmt.Println(bn_num_bits(a.Bits()))
	fmt.Println(bn_num_bits_word(big.Word(0xffffccccf)))
}

func mulBigInt(a, b *big.Int) *big.Int {
	re := new(big.Int)
	for i:=0; i<b.BitLen(); i++ {
		if b.Bit(i) == 1 {
			re.Xor(re, a)
		}
		a.Lsh(a, 1)
	}
	return re
}

func fillRandom(arr []big.Word) {
	for i, _ := range arr {
		arr[i] = big.Word(rand.Uint64())
	}
}
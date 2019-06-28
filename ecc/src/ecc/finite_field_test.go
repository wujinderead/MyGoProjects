package ecc

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"
)

var upper = []rune{'⁰', '¹', '²', '³', '⁴', '⁵', '⁶', '⁷', '⁸', '⁹'}

func TestBls12Parameters(t *testing.T) {
	// y² = x³ + 4 over Fq
	q, _ := new(big.Int).SetString("1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab", 16)
	r, _ := new(big.Int).SetString("73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001", 16)
	// order of the elliptic curve is retrieved from SageMath: E=EllipticCurve(GF(q),[0,4]); hex(E.order())
	n, _ := new(big.Int).SetString("1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb15400008c0000000000aaab", 16)
	g1x, _ := new(big.Int).SetString("17f1d3a73197d7942695638c4fa9ac0fc3688c4f9774b905a14e3a3f171bac586c55e83ff97a1aeffb3af00adb22c6bb", 16)
	g1y, _ := new(big.Int).SetString("8b3f481e3aaa0f1a09e30ed741d8ae4fcf5e095d5d00af600db18cb2c04b3edd03cc744a2888ae40caa232946c5e7e1", 16)
	fmt.Println("q =", q.BitLen(), new(big.Int).Mod(q, FOUR), q.String())
	fmt.Println("r =", r.BitLen(), new(big.Int).Mod(r, FOUR), r.String())
	fmt.Println("n =", n.BitLen(), n.String())
	fq := FpCurve{head: nil, P: q, A: ZERO, B: FOUR, X: nil, Y: nil, Order: n}
	ndivr := new(big.Int).Div(n, r)
	nmodr := new(big.Int).Mod(n, r)
	g1 := &EcPoint{g1x, g1y}
	fmt.Println("g1 on curve:", fq.IsOnCurve(g1))
	fmt.Println("n/r =", ndivr, "mod:", nmodr) // n mod r = 0, i.e., r | n
	qq := new(big.Int).Set(q)
	for i := 2; i <= 12; i++ {
		qq.Mul(qq, q)
		qq.Mod(qq, r)
		fmt.Println("q"+toUpper(i), "mod r =", qq) // q^12 mod r = 1, i.e., r | (q¹² - 1)
	}

	fmt.Println("g1 * r =", fq.ScalaMult(g1, r.Bytes())) // G1*r=0, i.e., G1 is in subgroup of order r

	x, _ := new(big.Int).SetString("-d201000000010000", 16) // generator
	co1 := new(big.Int).Sub(x, ONE)
	co1.Mul(co1, co1)
	co1.Div(co1, THREE)
	fmt.Println("cofator1:", co1, co1.Cmp(ndivr)) //  (x-1)²/3 = n/r, the cofactor of order r is (x-1)²/3
}

func testFiniteField(p, k int, irreducible []int) {
	fmt.Println(string(upper))
	// check if it is an irreducible polynomial
	// a candidate polynomial is x^k + [0,p)x^(k-1) + [0,p)x^(k-2) + [0,p)x + [0,p)
	fmt.Println("irreducible polynomial:", polyStr(k, irreducible))

	n, bases := getOrder(p, k)
	fmt.Println(n, bases)
	if !checkIrreducible(p, k, irreducible) {
		fmt.Println("polynomial reducible!")
		return
	}

	set := set(make(map[int]struct{}))
	cur := make([]int, k+1)
	cur[0] = 1
	fmt.Println("λ⁰ =", cur, "=", decimal(cur, bases))
	set.add(0)
	set.add(1)
	for i := 1; i < n-1; i++ {
		multiplyAndMod(p, k, cur, irreducible)
		dec := decimal(cur, bases)
		fmt.Println("λ"+toUpper(i), "=", cur, "=", dec)
		if set.contains(dec) {
			break
		}
		set.add(dec)
	}
	fmt.Println("set size:", set.size())
}

func TestFiniteField(t *testing.T) {
	testFiniteField(5, 3, []int{1, 2, 0, 1})    // GF(5³), irreducible polynomial x³ + 3x¹ + 3
	testFiniteField(7, 4, []int{3, 4, 5, 0, 1}) // GF(7⁴), irreducible polynomial x⁴ + 5x² + 4x¹ + 3
}

func multiplyAndMod(p, k int, cur, irreducible []int) {
	for i := k; i > 0; i-- {
		cur[i] = cur[i-1]
	}
	cur[0] = 0
	if cur[k] == 0 {
		return
	}
	if cur[k] > 1 {
		newirre := make([]int, k)
		for i := 0; i < k; i++ {
			newirre[i] = cur[k] * irreducible[i]
		}
		for i := 0; i < k; i++ {
			cur[i] = (cur[i] - newirre[i]) % p
			if cur[i] < 0 {
				cur[i] += p
			}
		}
	} else { // subtract directly
		for i := 0; i < k; i++ {
			cur[i] = (cur[i] - irreducible[i]) % p
			if cur[i] < 0 {
				cur[i] += p
			}
		}
	}
}

func toUpper(a int) string {
	astr := strconv.Itoa(a)
	runes := make([]rune, len(astr))
	for i := range astr {
		runes[i] = upper[astr[i]-'0']
	}
	return string(runes)
}

func polyStr(k int, effis []int) string {
	str := ""
	for i := range effis {
		effi := effis[k-i]
		exp := k - i
		if effi == 0 {
			continue
		}
		if effi != 1 {
			str += fmt.Sprint(effi)
		}
		if exp == 0 {
			if effi == 1 {
				str += fmt.Sprint(effi)
			}
			break
		}
		str += fmt.Sprint("x", toUpper(exp), " + ")
	}
	return str
}

func getOrder(p, k int) (int, []int) {
	n := p
	bases := make([]int, k)
	bases[0] = 1
	for i := 1; i < k; i++ {
		bases[i] = p * bases[i-1]
		n *= p
	}
	return n, bases
}

func checkIrreducible(p, k int, irreducible []int) bool {
	// let x from 0 to p-1, for Poly = x^k + [0,p)x^(k-1) + [0,p)x^(k-2) + [0,p)x + [0,p)
	// check Poly!=0 from all x
	for x := 0; x < p; x++ {
		pows := make([]int, k+1)
		pows[0] = 1
		for i := 1; i < k+1; i++ {
			pows[i] = (x * pows[i-1]) % p
		}
		n := decimal(irreducible, pows)
		//fmt.Println(x, pows, n)
		if n%p == 0 {
			return false
		}
	}
	return true
}

func decimal(effis, bases []int) int {
	n := 0
	for i := range bases {
		n += effis[i] * bases[i]
	}
	return n
}

type set map[int]struct{}

func (s set) contains(i int) bool {
	_, ok := s[i]
	return ok
}

func (s set) add(i int) {
	s[i] = struct{}{}
}

func (s set) size() int {
	return len(s)
}

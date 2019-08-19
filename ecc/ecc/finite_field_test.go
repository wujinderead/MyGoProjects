package ecc

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	upper = []rune{'⁰', '¹', '²', '³', '⁴', '⁵', '⁶', '⁷', '⁸', '⁹'}
)

func testFiniteField(p, k int, irreducible []int) {
	fmt.Println(string(upper))
	// check if it is an irreducible polynomial
	// a candidate polynomial is x^k + [0,p)x^(k-1) + [0,p)x^(k-2) + ··· + [0,p)x + [0,p)
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
	testFiniteField(5, 3, []int{3, 3, 0, 1})    // GF(5³), irreducible polynomial x³ + 3x¹ + 3
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

/*
g5=GF(5)
g53=GF(5**3)
print g53.polynomial()   # z³+3z+3

e5=EllipticCurve(g5, [2, 4])  # Elliptic Curve defined by y²=x³+2x+4 over GF(5)
print 'e5.order =', e5.order()   # 7
e5.points()  # O, (0, 2), (0, 3), (2, 1), (2, 4), (4, 1), (4, 4)]

e53=EllipticCurve(g53, [2, 4])  # Elliptic Curve defined by y²=x³+2x+4 over GF(5³)
print 'e53.order =', e53.order()   # 112

# get two points on curve e53, polynomial and vector form
p1=e53.points()[12]      # (3z+1, 2z²+2z) = ((1, 3, 0), (0, 2, 2))
p2=e53.points()[64]      # (2z²+2z+4, 4z²+3z+2) = ((4, 2, 2), (2, 3, 4))
print 'p1x=', g53.vector_space()(p1[0]), 'p1y=', g53.vector_space()(p1[1])
print 'p2x=', g53.vector_space()(p2[0]), 'p2y=', g53.vector_space()(p2[1])

# test equality
p1[1]**2 == p1[0]**3 + 2*p1[0] + 4   # (2z²+2z)² = (3z+1)³+2(3z+1)+4 mod z³+3z+3
p2[1]**2 == p2[0]**3 + 2*p2[0] + 4
*/

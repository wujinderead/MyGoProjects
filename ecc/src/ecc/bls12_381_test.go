package ecc

import (
	"fmt"
	"math/big"
	"testing"
)

/*
# sage script for bls12-381 field
x = -0xd201000000010000   # generator
t = x + 1
q = 1/3 * (x - 1)**2*(x**4 - x**2 + 1) + x
r = x**4 - x**2 + 1
n = q + 1 - t
cofactor1 = int((x - 1)**2 / 3)
cofactor2 = int((x**8 - 4*x**7 + 5*x**6 - 4*x**4 + 6*x**3 - 4*x**2 - 4*x + 13) / 9)

b = 4
F = GF(q)
E = EllipticCurve(F, [0, b])
print "E.order==n ?", E.order()==n
print "E.order/r == cofactor1 ?", int(n/r)==cofactor1

# extend to Fq²
R.<T> = PolynomialRing(F)
non_residue = -1
j = 1
quadratic_non_residue = u+j
F2.<u> = F.extension(T^2-non_residue,'u')             # u²+1
E2 = EllipticCurve(F2, [0,b*quadratic_non_residue])   # E'(Fq²)/(u²+1), E: y² = x³ + 4(u+1)
print "E2.order/r == cofactor2 ?", int(E2.order()/r)==cofactor2

F12_equation = (T^6 - j)^2 - non_residue
F12.<w> = F.extension(F12_equation)                   # w¹² - 2w⁶ + 2
E12 = EllipticCurve(F12, [0,b])                       # E(Fq¹²)/(w¹²-2w⁶+2), E: y² = x³ + 4

*/

var (
	// bls12-381 parameters
	// use generator z to generate 2 prime q and r
	// q = (z - 1)² ((z⁴ - z² + 1) / 3) + z
	// r = (z⁴ - z² + 1)
	// satisfies r | n, r | (q^k-1). in this case, smallest k is 12, i.e., r | (q¹² - 1)
	z, _ = new(big.Int).SetString("-d201000000010000", 16)
	q, _ = new(big.Int).SetString("1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab", 16)
	r, _ = new(big.Int).SetString("73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001", 16)

	// the order of elliptic curve over Fq is also parameterized as n = q+1-t, t=z+1, so n=q-z
	n, _ = new(big.Int).SetString("1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb15400008c0000000000aaab", 16)

	// group G1 is defined on E(Fq): y² = x³ + 4
	// g1 (g1x, g1y) is a point of subgroup with order r, i.e., r*g1 = O.
	g1x, _ = new(big.Int).SetString("17f1d3a73197d7942695638c4fa9ac0fc3688c4f9774b905a14e3a3f171bac586c55e83ff97a1aeffb3af00adb22c6bb", 16)
	g1y, _ = new(big.Int).SetString("8b3f481e3aaa0f1a09e30ed741d8ae4fcf5e095d5d00af600db18cb2c04b3edd03cc744a2888ae40caa232946c5e7e1", 16)

	// group G2 is defined on E'(Fq²/(u²+1)) : y² = x³ + 4(u + 1)
	// g2 (x1*u+x0, y1*u+y0) is a point of subgroup with order r, i.e., r*g2 = O.
	// the bilinear map is G1 × G2 -> Gt
	n2, _ = new(big.Int).SetString("2a437a4b8c35fc74bd278eaa22f25e9e2dc90e50e7046b466e59e49349e8bd050a62cfd16ddca6ef53149330978ef0137697386bf984315744a2d5eb3dd4d213f2484c55b94474ab096de2c62640b2643116b1e2788e6a8b2a9fffe1c7238e5", 16)
	x1, _ = new(big.Int).SetString("13e02b6052719f607dacd3a088274f65596bd0d09920b61ab5da61bbdc7f5049334cf11213945d57e5ac7d055d042b7e", 16)
	x0, _ = new(big.Int).SetString("24aa2b2f08f0a91260805272dc51051c6e47ad4fa403b02b4510b647ae3d1770bac0326a805bbefd48056c8c121bdb8", 16)
	y1, _ = new(big.Int).SetString("606c4a02ea734cc32acd2b02bc28b99cb3e287e85a763af267492ab572e99ab3f370d275cec1da1aaa9075ff05f79be", 16)
	y0, _ = new(big.Int).SetString("ce5d527727d6e118cc9cdc6da2e351aadfd9baa8cbdd3a76d429a695160d12c923ac9cc3baca289e193548608b82801", 16)
)

func TestBls12Parameters(t *testing.T) {
	fmt.Println("q =", q.BitLen(), new(big.Int).Mod(q, FOUR), q.String())
	fmt.Println("r =", r.BitLen(), new(big.Int).Mod(r, FOUR), r.String())
	fmt.Println("n =", n.BitLen(), n.String())
	qsubn := new(big.Int).Sub(q, n)
	fmt.Println("z =", z.BitLen(), z.String(), "q-n == z:", qsubn.Cmp(z) == 0)
	fq := FpCurve{head: nil, P: q, A: ZERO, B: FOUR, X: nil, Y: nil, Order: n}
	ndivr := new(big.Int).Div(n, r)
	nmodr := new(big.Int).Mod(n, r)
	fmt.Println("n/r =", ndivr, "mod:", nmodr) // n mod r = 0, i.e., r | n
	qq := new(big.Int).Set(q)
	for i := 2; i <= 12; i++ {
		qq.Mul(qq, q)
		qq.Mod(qq, r)
		fmt.Println("q"+toUpper(i), "mod r =", qq) // q¹² mod r = 1, i.e., r | (q¹² - 1)
	}

	g1 := &EcPoint{g1x, g1y}
	fmt.Println("g1 on curve:", fq.IsOnCurve(g1))
	fmt.Println("g1 * r =", fq.ScalaMult(g1, r.Bytes())) // G1*r=0, i.e., G1 is in subgroup of order r

	co1 := new(big.Int).Sub(z, ONE)
	co1.Mul(co1, co1)
	co1.Div(co1, THREE)
	// (z-1)²/3 = n/r, i.e., the cofactor of r order subgroup is (z-1)²/3
	fmt.Println("cofator1:", co1, co1.Cmp(ndivr))

	// (z⁸-4z⁷+5z⁶-4z⁴+6z³-4z²-4z+13)/9 = n2/r, i.e. the cofactor of r order subgroup in E'(Fq²).
	co2 := new(big.Int).SetInt64(13)
	zz := new(big.Int).Lsh(z, 2)   // 4z
	co2.Sub(co2, zz)               // -4z+13
	zz.Mul(zz, z)                  // 4z²
	co2.Sub(co2, zz)               // -4z²-4z+13
	zz.Mul(zz, z)                  // 4z³
	tmp := new(big.Int).Rsh(zz, 1) // 2z³
	co2.Add(co2, tmp)
	co2.Add(co2, zz) // 6z³-4z²-4z+13
	zz.Mul(zz, z)    // 4z⁴
	co2.Sub(co2, zz) // -4z⁴+6z³-4z²-4z+13
	zz.Mul(zz, z)
	zz.Mul(zz, z)  // 4z⁶
	tmp.Rsh(zz, 2) // z⁶
	co2.Add(co2, tmp)
	co2.Add(co2, zz) // 5z⁶-4z⁴+6z³-4z²-4z+13
	zz.Mul(zz, z)    // 4z⁷
	co2.Sub(co2, zz) // -4z⁷+5z⁶-4z⁴+6z³-4z²-4z+13
	zz.Mul(zz, z)
	zz.Rsh(zz, 2)    // z⁸
	co2.Add(co2, zz) // z⁸-4z⁷+5z⁶-4z⁴+6z³-4z²-4z+13
	tmp.SetInt64(9)
	co2.Div(co2, tmp) // (z⁸-4z⁷+5z⁶-4z⁴+6z³-4z²-4z+13)/9
	n2divr := new(big.Int).Div(n2, r)
	fmt.Println("cofator2:", co2, co2.Cmp(n2divr))
}

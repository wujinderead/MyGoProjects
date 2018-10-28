package cipher

import (
	"testing"
	"crypto/dsa"
	"crypto/rand"
	"fmt"
	"crypto/sha256"
	"io"
	"math/big"
)

func Test_dsaSign(t *testing.T) {
	parameters := new(dsa.Parameters)
	err := dsa.GenerateParameters(parameters, rand.Reader, dsa.L2048N256)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	priv := new(dsa.PrivateKey)
	priv.Parameters = *parameters
	err = dsa.GenerateKey(priv, rand.Reader)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Printf("p: %4d, %x\n", priv.P.BitLen(), priv.P.Bytes())
	fmt.Printf("q: %4d, %x\n", priv.Q.BitLen(), priv.Q.Bytes())
	fmt.Printf("g: %4d, %x\n", priv.G.BitLen(), priv.G.Bytes())
	fmt.Printf("x: %4d, %x\n", priv.X.BitLen(), priv.X.Bytes())
	fmt.Printf("y: %4d, %x\n", priv.Y.BitLen(), priv.Y.Bytes())

	hasher := sha256.New()
	hasher.Write([]byte("my name is van. i'm a artist, a performance artist."))
	hash := hasher.Sum([]byte{})
	fmt.Printf("h: %4d, %x\n", 256, hash)

	r, s, err := dsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Printf("r: %4d, %x\n", r.BitLen(), r.Bytes())
	fmt.Printf("s: %4d, %x\n", s.BitLen(), s.Bytes())
}

func Test_dsaSignFixed(t *testing.T) {
	parameters := new(dsa.Parameters)
	err := dsa.GenerateParameters(parameters, rand.Reader, dsa.L2048N256)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	priv := new(dsa.PrivateKey)
	priv.P, _ = new(big.Int).SetString("a521c4c4ad779187570677cdf42c21ba34b9a45a562867f5465d8ea925c5fdcc4aa2d7fffc508a772c217507f461268fbdfe2694625b9c06ce735c2cd99e12347c993db0fd6c2282c70601b41a49a8bd5e9f8c36317a4f0e701b8e2fd84265f44524ed061857adbcad39a11065ad278b1817f159d71a87ee899cb85da386779ec2dcf7ca79322b2a951c88c70d7c9742a04d5d30062da3a4febadcd0be2fcb0956925af0e59cbd6262b6114e6c517580d62424ebe9d8b601f81df019d0fac3db0419a1dc0ae147f63f1cf6ee98ba0f34d15100a5f3f0f75bbfaf593bf6e404658edb3022affda27e940592dd0d33794cc9f11cfdecc46f4a5630d283a13c2181", 16)
	priv.Q, _ = new(big.Int).SetString("d452142c5311dbf26ce8c7ba7d9d855df33e37b005662aabe05d91b0150d1241", 16)
	priv.G, _ = new(big.Int).SetString("2513eab991902ad8d8228c366f7c661293ff03628b48bbf3361a780769aca2ffae29e768adc0612ee1c1247d9707efce039094746a2ef8f381f24f1d99fbec2cc71b28661827c3d07670111f1c92b14085f98475360023a9c865daf811643cf07df7e27d306fc398461a1883c13189a59a5d7aca24f3c5137bb7bc57c5b0e6ef82ec7258a7e05d5be8632c84ed2a587396670da4e029335d9f4dac6f2bb74c588e814265550a56e596b4452f9200a17c82d5bd76fc676e6c9e9d6f7a6d15584bbd46cfeb4d5f9d6e69c736864e03e5b7df6b61b0f5bfe046b3bc16ab0b7e6b5a56cb6e164b651a4a83a3193ec02b9cbceb50ed2e6ac868156e98a1064e765194", 16)
	priv.X, _ = new(big.Int).SetString("2c39c9646245f4638a6288870ff303ef190801924057dcb69c8b292f0c4a047b", 16)
	priv.Y, _ = new(big.Int).SetString("2548176b43b9ad7d2d6aede44962a699b12783ee985f4ea67c2b44285418552c919c4250f61f6624efea44e4a720550451d20651898dbc34dfc3829a12d0ca51ca18353971155398559dfd402ec5b1832141ef89ece4b3b55e32c8e05ca10f91b76443e7704cf3c39c07640ca540b66b27dd0b0fbbc974a032152d08370add84874249ff723a0a20b8636d3e48b0a8685e15394379a55a88ceb5c55acf8d6d96b7092eb32282ff8b4cd6b3490e80abb3363aabe29cd4f92fc72e8f874013981ed04e4ec88f601c6059e55d082d184532827e987aebf2aaf0f27dd4cf0cc95fb12edb2f5d9ca251b20f956500a484cdb51929c2c70b1ca9eabd75a148fbee47ec", 16)

	fmt.Printf("p: %4d, %x\n", priv.P.BitLen(), priv.P.Bytes())
	fmt.Printf("q: %4d, %x\n", priv.Q.BitLen(), priv.Q.Bytes())
	fmt.Printf("g: %4d, %x\n", priv.G.BitLen(), priv.G.Bytes())
	fmt.Printf("x: %4d, %x\n", priv.X.BitLen(), priv.X.Bytes())
	fmt.Printf("y: %4d, %x\n", priv.Y.BitLen(), priv.Y.Bytes())

	hasher := sha256.New()
	hasher.Write([]byte("my name is van. i'm a artist, a performance artist."))
	hash := hasher.Sum([]byte{})
	fmt.Printf("h: %4d, %x\n", 256, hash)

	r, s, err := Sign(rand.Reader, priv, hash)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Printf("r: %4d, %x\n", r.BitLen(), r.Bytes())
	fmt.Printf("s: %4d, %x\n", s.BitLen(), s.Bytes())

	verified := Verify(&priv.PublicKey, hash, r, s)
	fmt.Println("verified ? ", verified)
}

func Sign(rand io.Reader, priv *dsa.PrivateKey, hash []byte) (r, s *big.Int, err error) {
	// FIPS 186-3, section 4.6

	n := priv.Q.BitLen()
	if priv.Q.Sign() <= 0 || priv.P.Sign() <= 0 || priv.G.Sign() <= 0 || priv.X.Sign() <= 0 || n&7 != 0 {
		err = dsa.ErrInvalidPublicKey
		return
	}
	n >>= 3

	var attempts int
	for attempts = 10; attempts > 0; attempts-- {
		k := new(big.Int)
		buf := make([]byte, n)
		for {
			_, err = io.ReadFull(rand, buf)
			if err != nil {
				return
			}
			k.SetBytes(buf)
			// priv.Q must be >= 128 because the test above
			// requires it to be > 0 and that
			//    ceil(log_2(Q)) mod 8 = 0
			// Thus this loop will quickly terminate.
			if k.Sign() > 0 && k.Cmp(priv.Q) < 0 {
				break
			}
		}

		kInv := fermatInverse(k, priv.Q)

		r = new(big.Int).Exp(priv.G, k, priv.P)
		r.Mod(r, priv.Q)

		if r.Sign() == 0 {
			continue
		}

		z := k.SetBytes(hash)

		s = new(big.Int).Mul(priv.X, r)
		s.Add(s, z)
		s.Mod(s, priv.Q)
		s.Mul(s, kInv)
		s.Mod(s, priv.Q)

		if s.Sign() != 0 {
			fmt.Printf("x: %4d, %x\n", k.BitLen(), k.Bytes())
			break
		}
	}

	// Only degenerate private keys will require more than a handful of
	// attempts.
	if attempts == 0 {
		return nil, nil, dsa.ErrInvalidPublicKey
	}

	return
}

// Verify verifies the signature in r, s of hash using the public key, pub. It
// reports whether the signature is valid.
//
// Note that FIPS 186-3 section 4.6 specifies that the hash should be truncated
// to the byte-length of the subgroup. This function does not perform that
// truncation itself.
func Verify(pub *dsa.PublicKey, hash []byte, r, s *big.Int) bool {
	// FIPS 186-3, section 4.7

	if pub.P.Sign() == 0 {
		return false
	}

	if r.Sign() < 1 || r.Cmp(pub.Q) >= 0 {
		return false
	}
	if s.Sign() < 1 || s.Cmp(pub.Q) >= 0 {
		return false
	}

	w := new(big.Int).ModInverse(s, pub.Q)

	n := pub.Q.BitLen()
	if n&7 != 0 {
		return false
	}
	z := new(big.Int).SetBytes(hash)

	u1 := new(big.Int).Mul(z, w)
	u1.Mod(u1, pub.Q)
	u2 := w.Mul(r, w)
	u2.Mod(u2, pub.Q)
	v := u1.Exp(pub.G, u1, pub.P)
	u2.Exp(pub.Y, u2, pub.P)
	v.Mul(v, u2)
	v.Mod(v, pub.P)
	v.Mod(v, pub.Q)

	return v.Cmp(r) == 0
}

func fermatInverse(k, P *big.Int) *big.Int {
	two := big.NewInt(2)
	pMinus2 := new(big.Int).Sub(P, two)
	return new(big.Int).Exp(k, pMinus2, P)
}
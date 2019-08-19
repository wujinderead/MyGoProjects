package ecc

import (
	"fmt"
	"math/big"
)

const (
	BN_BITS2 = 64
	BN_MASK2 = 0xffffffffffffffff
)

func gf2m_mul_1x1(r1, r0 *big.Word, a, b big.Word) {
	var h, l, s big.Word
	var tab = make([]big.Word, 16)
	var top3b = a >> 61
	var a1, a2, a4, a8 big.Word

	a1 = a & (0x1FFFFFFFFFFFFFFF)
	a2 = a1 << 1
	a4 = a2 << 1
	a8 = a4 << 1

	tab[0] = 0
	tab[1] = a1
	tab[2] = a2
	tab[3] = a1 ^ a2
	tab[4] = a4
	tab[5] = a1 ^ a4
	tab[6] = a2 ^ a4
	tab[7] = a1 ^ a2 ^ a4
	tab[8] = a8
	tab[9] = a1 ^ a8
	tab[10] = a2 ^ a8
	tab[11] = a1 ^ a2 ^ a8
	tab[12] = a4 ^ a8
	tab[13] = a1 ^ a4 ^ a8
	tab[14] = a2 ^ a4 ^ a8
	tab[15] = a1 ^ a2 ^ a4 ^ a8

	s = tab[b&0xF]
	l = s
	s = tab[b>>4&0xF]
	l ^= s << 4
	h = s >> 60
	s = tab[b>>8&0xF]
	l ^= s << 8
	h ^= s >> 56
	s = tab[b>>12&0xF]
	l ^= s << 12
	h ^= s >> 52
	s = tab[b>>16&0xF]
	l ^= s << 16
	h ^= s >> 48
	s = tab[b>>20&0xF]
	l ^= s << 20
	h ^= s >> 44
	s = tab[b>>24&0xF]
	l ^= s << 24
	h ^= s >> 40
	s = tab[b>>28&0xF]
	l ^= s << 28
	h ^= s >> 36
	s = tab[b>>32&0xF]
	l ^= s << 32
	h ^= s >> 32
	s = tab[b>>36&0xF]
	l ^= s << 36
	h ^= s >> 28
	s = tab[b>>40&0xF]
	l ^= s << 40
	h ^= s >> 24
	s = tab[b>>44&0xF]
	l ^= s << 44
	h ^= s >> 20
	s = tab[b>>48&0xF]
	l ^= s << 48
	h ^= s >> 16
	s = tab[b>>52&0xF]
	l ^= s << 52
	h ^= s >> 12
	s = tab[b>>56&0xF]
	l ^= s << 56
	h ^= s >> 8
	s = tab[b>>60]
	l ^= s << 60
	h ^= s >> 4

	/* compensate for the top three bits of a */

	if (top3b & 01) != 0 {
		l ^= b << 61
		h ^= b >> 3
	}
	if (top3b & 02) != 0 {
		l ^= b << 62
		h ^= b >> 2
	}
	if (top3b & 04) != 0 {
		l ^= b << 63
		h ^= b >> 1
	}

	*r1 = h
	*r0 = l
}

func gf2m_mul_2x2(r []big.Word, a1, a0, b1, b0 big.Word) {
	var m1, m0 big.Word
	/* r[3] = h1, r[2] = h0 r[1] = l1 r[0] = l0 */
	gf2m_mul_1x1(&r[3], &r[2], a1, b1)
	gf2m_mul_1x1(&r[1], &r[0], a0, b0)
	gf2m_mul_1x1(&m1, &m0, a0^a1, b0^b1)
	/* Correction on m1 ^= l1 ^ h1 m0 ^= l0 ^ h0 */
	r[2] ^= m1 ^ r[1] ^ r[3]            /* h0 ^= m1 ^ l1 ^ h1; */
	r[1] = r[3] ^ r[2] ^ r[0] ^ m1 ^ m0 /* l1 ^= l0 ^ h0 ^ m0; */
}

func bn_gf2m_mul(a, b []big.Word) []big.Word {
	var zlen, i, j, k int
	var x1, x0, y1, y0 big.Word
	var zz = make([]big.Word, 4)

	zlen = len(a) + len(b) + 2
	s := make([]big.Word, zlen)

	for j = 0; j < len(b); j += 2 {
		y0 = b[j]
		if (j + 1) == len(b) {
			y1 = 0
		} else {
			y1 = b[j+1]
		}
		for i = 0; i < len(a); i += 2 {
			x0 = a[i]
			if (i + 1) == len(a) {
				x1 = 0
			} else {
				x1 = a[i+1]
			}
			gf2m_mul_2x2(zz, x1, x0, y1, y0)
			for k = 0; k < 4; k++ {
				s[i+j+k] ^= big.Word(zz[k])
			}
		}
	}
	// the underlying big.Word[] for big.Int is little-endian
	return s
}

func bn_gf2m_sqr(a []big.Word) []big.Word {
	s := make([]big.Word, len(a)<<1)
	for i := len(a) - 1; i >= 0; i-- {
		s[(i<<1)+1] = sqr1(a[i])
		s[i<<1] = sqr0(a[i])
	}
	return s
}

func sqr_nibble(w big.Word) big.Word {
	return ((w & 8) << 3) | ((w & 4) << 2) | ((w & 2) << 1) | (w & 1)
}

func sqr1(w big.Word) big.Word {
	return sqr_nibble(w>>60)<<56 |
		sqr_nibble(w>>56)<<48 |
		sqr_nibble(w>>52)<<40 |
		sqr_nibble(w>>48)<<32 |
		sqr_nibble(w>>44)<<24 |
		sqr_nibble(w>>40)<<16 |
		sqr_nibble(w>>36)<<8 |
		sqr_nibble(w>>32)
}

func sqr0(w big.Word) big.Word {
	return sqr_nibble(w>>28)<<56 |
		sqr_nibble(w>>24)<<48 |
		sqr_nibble(w>>20)<<40 |
		sqr_nibble(w>>16)<<32 |
		sqr_nibble(w>>12)<<24 |
		sqr_nibble(w>>8)<<16 |
		sqr_nibble(w>>4)<<8 |
		sqr_nibble(w)
}

func bn_gf2m_poly2arr(a []big.Word) []int {
	var i, j, k int
	var mask big.Word
	p := make([]int, 6)

	for i = len(a) - 1; i >= 0; i-- {
		if a[i] == 0 {
			/* skip word if a->d[i] == 0 */
			continue
		}
		mask = 1 << (BN_BITS2 - 1)
		for j = BN_BITS2 - 1; j >= 0; j-- {
			if (a[i] & mask) != 0 {
				p[k] = BN_BITS2*i + j
				k++
			}
			mask >>= 1
		}
	}
	p[k] = -1
	return p
}

func bn_gf2m_mod_arr(a []big.Word, p []int) []big.Word {
	var j, k int
	var n, dN, d0, d1 int
	var zz big.Word

	z := make([]big.Word, len(a))
	copy(z, a)

	/* start reduction */
	dN = p[0] / BN_BITS2
	for j = len(z) - 1; j > dN; {
		zz = z[j]
		if z[j] == 0 {
			j--
			continue
		}
		z[j] = 0

		for k = 1; p[k] != 0; k++ {
			/* reducing component t^p[k] */
			n = p[0] - p[k]
			d0 = n % BN_BITS2
			d1 = BN_BITS2 - d0
			n /= BN_BITS2
			z[j-n] ^= zz >> uint(d0)
			if d0 != 0 {
				z[j-n-1] ^= zz << uint(d1)
			}
		}

		/* reducing component t^0 */
		n = dN
		d0 = p[0] % BN_BITS2
		d1 = BN_BITS2 - d0
		z[j-n] ^= zz >> uint(d0)
		if d0 != 0 {
			z[j-n-1] ^= zz << uint(d1)
		}
	}

	/* final round of reduction */
	for j == dN {

		d0 = p[0] % BN_BITS2
		zz = z[dN] >> uint(d0)
		if zz == 0 {
			break
		}
		d1 = BN_BITS2 - d0

		/* clear up the top d1 bits */
		if d0 != 0 {
			z[dN] = (z[dN] << uint(d1)) >> uint(d1)
		} else {
			z[dN] = 0
		}
		z[0] ^= zz /* reduction t^0 component */

		for k = 1; p[k] != 0; k++ {
			var tmp_ulong big.Word

			/* reducing component t^p[k] */
			n = p[k] / BN_BITS2
			d0 = p[k] % BN_BITS2
			d1 = BN_BITS2 - d0
			z[n] ^= zz << uint(d0)
			tmp_ulong = zz >> uint(d1)
			if d0 != 0 && tmp_ulong != 0 {
				z[n+1] ^= tmp_ulong
			}
		}
	}
	return z
}

func bn_gf2m_mod_arr_self(z []big.Word, p []int) {
	var j, k int
	var n, dN, d0, d1 int
	var zz big.Word

	/* start reduction */
	dN = p[0] / BN_BITS2
	for j = len(z) - 1; j > dN; {
		zz = z[j]
		if z[j] == 0 {
			j--
			continue
		}
		z[j] = 0

		for k = 1; p[k] != 0; k++ {
			/* reducing component t^p[k] */
			n = p[0] - p[k]
			d0 = n % BN_BITS2
			d1 = BN_BITS2 - d0
			n /= BN_BITS2
			z[j-n] ^= zz >> uint(d0)
			if d0 != 0 {
				z[j-n-1] ^= zz << uint(d1)
			}
		}

		/* reducing component t^0 */
		n = dN
		d0 = p[0] % BN_BITS2
		d1 = BN_BITS2 - d0
		z[j-n] ^= zz >> uint(d0)
		if d0 != 0 {
			z[j-n-1] ^= zz << uint(d1)
		}
	}

	/* final round of reduction */
	for j == dN {

		d0 = p[0] % BN_BITS2
		zz = z[dN] >> uint(d0)
		if zz == 0 {
			break
		}
		d1 = BN_BITS2 - d0

		/* clear up the top d1 bits */
		if d0 != 0 {
			z[dN] = (z[dN] << uint(d1)) >> uint(d1)
		} else {
			z[dN] = 0
		}
		z[0] ^= zz /* reduction t^0 component */

		for k = 1; p[k] != 0; k++ {
			var tmp_ulong big.Word

			/* reducing component t^p[k] */
			n = p[k] / BN_BITS2
			d0 = p[k] % BN_BITS2
			d1 = BN_BITS2 - d0
			z[n] ^= zz << uint(d0)
			tmp_ulong = zz >> uint(d1)
			if d0 != 0 && tmp_ulong != 0 {
				z[n+1] ^= tmp_ulong
			}
		}
	}
}

func bn_gf2m_mod(a, p []big.Word) []big.Word {
	arr := bn_gf2m_poly2arr(p)
	return bn_gf2m_mod_arr(a, arr)
}

func bn_num_bits_word(ll big.Word) int {
	if ll == 0 {
		return 0
	}
	var x, mask uint
	var bits uint = 1

	l := uint(ll)
	x = l >> 32
	mask = (0 - x) & BN_MASK2
	mask = 0 - (mask >> (BN_BITS2 - 1))
	bits += 32 & mask
	l ^= (x ^ l) & mask

	x = l >> 16
	mask = (0 - x) & BN_MASK2
	mask = 0 - (mask >> (BN_BITS2 - 1))
	bits += 16 & mask
	l ^= (x ^ l) & mask

	x = l >> 8
	mask = (0 - x) & BN_MASK2
	mask = 0 - (mask >> (BN_BITS2 - 1))
	bits += 8 & mask
	l ^= (x ^ l) & mask

	x = l >> 4
	mask = (0 - x) & BN_MASK2
	mask = 0 - (mask >> (BN_BITS2 - 1))
	bits += 4 & mask
	l ^= (x ^ l) & mask

	x = l >> 2
	mask = (0 - x) & BN_MASK2
	mask = 0 - (mask >> (BN_BITS2 - 1))
	bits += 2 & mask
	l ^= (x ^ l) & mask

	x = l >> 1
	mask = (0 - x) & BN_MASK2
	mask = 0 - (mask >> (BN_BITS2 - 1))
	bits += 1 & mask

	return int(bits)
}

func bn_num_bits(a []big.Word) int {
	i := len(a) - 1
	if bn_is_zero(a) {
		return 0
	}
	return (i * BN_BITS2) + bn_num_bits_word(a[i])
}

func bn_is_zero(a []big.Word) bool {
	for i := range a {
		if a[i] != 0 {
			return false
		}
	}
	return true
}

func bn_wexpand(a []big.Word, length int) []big.Word {
	ret := make([]big.Word, length)
	copy(ret, a)
	return ret
}

func bn_gf2m_mod_inv_vartime(a, p []big.Word) []big.Word {
	u := bn_gf2m_mod(a, p)        // u = a mod p
	v := make([]big.Word, len(p)) // v = p
	copy(v, p)
	var i int
	ubits := bn_num_bits(u)
	vbits := bn_num_bits(v) /* v is copy of p */
	top := len(p)

	u = bn_wexpand(u, top)
	b := make([]big.Word, top)
	b[0] = 1
	c := make([]big.Word, top)
	for true {
		for ubits != 0 && (u[0]&1) == 0 {
			var u0, u1, b0, b1, mask big.Word

			u0 = u[0]
			b0 = b[0]
			mask = big.Word(0 - (b0 & 1))
			b0 ^= p[0] & mask
			for i = 0; i < top-1; i++ {
				u1 = u[i+1]
				u[i] = ((u0 >> 1) | (u1 << (BN_BITS2 - 1))) & BN_MASK2
				u0 = u1
				b1 = b[i+1] ^ (p[i+1] & mask)
				b[i] = ((b0 >> 1) | (b1 << (BN_BITS2 - 1))) & BN_MASK2
				b0 = b1
			}
			u[i] = u0 >> 1
			b[i] = b0 >> 1
			ubits--
		}

		if ubits <= BN_BITS2 {
			if u[0] == 0 {
				fmt.Println("u0 =0 !")
				break
			}
			if u[0] == 1 {
				break
			}
		}

		if ubits < vbits {
			ubits, vbits = vbits, ubits
			u, v = v, u
			b, c = c, b
		}
		for i = 0; i < top; i++ {
			u[i] ^= v[i]
			b[i] ^= c[i]
		}
		if ubits == vbits {
			var ul big.Word
			utop := (ubits - 1) / BN_BITS2

			ul = u[utop]
			for ul == 0 && utop != 0 {
				utop--
				ul = u[utop]
			}
			ubits = utop*BN_BITS2 + bn_num_bits_word(ul)
		}
	}
	return b
}

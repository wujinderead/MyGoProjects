package main

import "fmt"

// https://leetcode.com/problems/sum-of-k-mirror-numbers/

// A k-mirror number is a positive integer without leading zeros that reads the same both forward
// and backward in base-10 as well as in base-k.
// For example, 9 is a 2-mirror number. The representation of 9 in base-10 and base-2 are 9 and
// 1001 respectively, which read the same both forward and backward.
// On the contrary, 4 is not a 2-mirror number. The representation of 4 in base-2 is 100, which
// does not read the same both forward and backward.
// Given the base k and the number n, return the sum of the n smallest k-mirror numbers.
// Example 1:
//   Input: k = 2, n = 5
//   Output: 25
//   Explanation:
//     The 5 smallest 2-mirror numbers and their representations in base-2 are listed as follows:
//       base-10    base-2
//         1          1
//         3          11
//         5          101
//         7          111
//         9          1001
//     Their sum = 1 + 3 + 5 + 7 + 9 = 25.
// Example 2:
//   Input: k = 3, n = 7
//   Output: 499
//   Explanation:
//     The 7 smallest 3-mirror numbers are and their representations in base-3 are listed as follows:
//       base-10    base-3
//         1          1
//         2          2
//         4          11
//         8          22
//         121        11111
//         151        12121
//         212        21212
//     Their sum = 1 + 2 + 4 + 8 + 121 + 151 + 212 = 499.
// Example 3:
//   Input: k = 7, n = 17
//   Output: 20379000
//   Explanation: The 17 smallest 7-mirror numbers are:
//     1, 2, 3, 4, 5, 6, 8, 121, 171, 242, 292, 16561, 65656, 2137312, 4602064, 6597956, 6958596
// Constraints:
//   2 <= k <= 9
//   1 <= n <= 30

// generate all base-10 palindromes and check if its base-k mirrored
func kMirror(k int, n int) int64 {
	sum := 0
	s := 1
	d := 0
	buf := make([]int, 50)
	ind := 0
	isMirror := true
outer:
	for {
		for num := s; num < s*10; num++ { // num from 100 to 900
			// say num=123, make a palindrome 12321, d=2
			x := num / 10 // 12, the left part
			xx := num     // 123, the palindrome
			for dd := 0; dd < d; dd++ {
				xx = xx*10 + x%10
				x = x / 10
			}
			// check if palindrome is k-mirror
			xxx := xx
			ind = 0
			for xx > 0 {
				buf[ind] = xx % k
				xx = xx / k
				ind++
			}
			isMirror = true
			for l, r := 0, ind-1; l < r; l, r = l+1, r-1 {
				if buf[l] != buf[r] {
					isMirror = false
					break
				}
			}
			if isMirror {
				n--
				sum += xxx
				fmt.Println(xxx, buf[:ind])
				if n == 0 {
					break outer
				}
			}
		}
		d++
		for num := s; num < s*10; num++ { // num from 100 to 900
			// say num=123, make a palindrome 123321, d=3
			x := num
			xx := num // 123, the palindrome
			for dd := 0; dd < d; dd++ {
				xx = xx*10 + x%10
				x = x / 10
			}
			// check if palindrome is k-mirror
			xxx := xx
			ind = 0
			for xx > 0 {
				buf[ind] = xx % k
				xx = xx / k
				ind++
			}
			isMirror = true
			for l, r := 0, ind-1; l < r; l, r = l+1, r-1 {
				if buf[l] != buf[r] {
					isMirror = false
					break
				}
			}
			if isMirror {
				n--
				sum += xxx
				fmt.Println(xxx, buf[:ind])
				if n == 0 {
					break outer
				}
			}
		}
		s = s * 10
	}
	return int64(sum)
}

func main() {
	for _, v := range []struct {
		k, n, ans int
	}{
		{2, 5, 25},
		{3, 7, 499},
		{7, 17, 20379000},
		// the first 30 9-mirror numbers in base-10 and base-9:
		// 1 1
		// 2 2
		// 3 3
		// 4 4
		// 5 5
		// 6 6
		// 7 7
		// 8 8
		// 191  232
		// 282  343
		// 373  454
		// 464  565
		// 555  676
		// 646  787
		// 656  808
		//  6886  10401
		// 25752  38283
		// 27472  41614
		// 42324  64046
		// 50605  76367
		//  626626 1154511
		// 1540451 2807082
		// 1713171 3201023
		// 1721271 3213123
		// 1828281 3385833
		// 1877781 3471743
		// 1885881 3483843
		// 2401042 4458544
		// 2434342 4520254
		// 2442442 4532354
		{9, 30, 18627530},
		// the first 30 7-mirror numbers in base-10 and base-7:
		// 1  1
		// 2  2
		// 3  3
		// 4  4
		// 5  5
		// 6  6
		// 8  11
		// 121 232
		// 171 333
		// 242 464
		// 292 565
		// 16561 66166
		// 65656 362263
		// 2137312 24111142
		// 4602064 54055045
		// 6597956 110040011
		// 6958596 113101311
		// 9470749 143333341
		// 61255216 1342442431
		// 230474032 5465665645
		// 466828664 14365656341
		// 485494584 15013431051
		// 638828836 21554645512
		// 657494756 22202420222
		// 858474858 30162626103
		// 25699499652 1566600066651
		// 40130703104 2620322230262
		// 45862226854 3212336332123
		// 61454945416 4303624263034
		// 64454545446 4441146411444
		{7, 30, 241030621167},
	} {
		fmt.Println(kMirror(v.k, v.n), v.ans)
	}
}

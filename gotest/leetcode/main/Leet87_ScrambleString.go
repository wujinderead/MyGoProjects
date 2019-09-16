package main

import "fmt"

// https://leetcode.com/problems/scramble-string/

// Given a string s1, we may represent it as a binary tree by
// partitioning it to two non-empty substrings recursively.
//
// Below is one possible representation of s1 = "great":
//
//     great
//    /    \
//   gr    eat
//  / \    /  \
// g   r  e   at
//            / \
//           a   t
// if we swap the "g" and "r", then swap "e" and "at", then swap "a" and "t",
// it produces a scrambled string "rgtae":
//
//     rgtae
//    /    \
//   rg    tae
//  / \    /  \
// r   g  ta  e
//        / \
//       t   a
// We say that "rgtae" is a scrambled string of "great".
// Given two strings s1 and s2 of the same length, determine if s2 is a scrambled string of s1.
// Example 1:
//   Input: s1 = "great", s2 = "rgeat"
//   Output: true
// Example 2:
//   Input: s1 = "abcde", s2 = "caebd"
//   Output: false

const alphabetSize = 128 // if only lowercase letter

// if s1[i...j] can scramble to s2[i...j], denote as canScramble(s1, s2),
// then there must be k (i≤k≤j), that makes
// canScramble(s1[i...k], s2[i...k]) and canScramble(s1[k...j], s2[k...j]),
// or canScramble(s1[i...k], s2[k...j]) and canScramble(s1[k...j], s2[i...k])
// the time complexity is 2^N or N! (not sure, but exponential anyway)
func isScramble(s1 string, s2 string) bool {
	alphabet := make([]int, alphabetSize)
	for i := 0; i < len(s1); i++ {
		alphabet[s1[i]]++
		alphabet[s2[i]]--
	}
	for _, v := range alphabet {
		if v != 0 {
			return false // different alphabet statistic, must not scramble
		}
	}
	// any permutation of a string with length≤3 can always scramble.
	// however for string length≥4, e.g. "1234", of its 24 permutations,
	// only 22 can scramble, with "2413" and "3142" can not.
	if len(s1) <= 3 {
		return true
	}
	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[len(s1)-i:len(s1)]) && isScramble(s1[i:], s2[:len(s1)-i]) {
			return true
		}
	}
	return false
}

// use dp to make the complexity polynomial.
// let F(i, j, k) denote that s1[i...i+k] and s2[j...j+k] can scramble,
// i.e., substring of s1 start at i, s2 start at j, with length k. then for 1≤m<k,
// F(i, j, k) = ( F(i, j, m) and F(i+m, j+m, k-m) ) or ( F(i, j+k-m, m) and F(i+m, j, k-m) )。
// base case is F(i, j, 1) = s1[i]==s2[j]. finally return F(0, 0, len(s1)).
// time complexity, O(n⁴)
func isScrambleDp(s1 string, s2 string) bool {
	l := len(s1)
	F := make([]bool, l*l*(l+1))
	for k := 1; k <= l; k++ {
		for i := 0; i <= l-k; i++ {
			for j := 0; j <= l-k; j++ {
				if k == 1 { // base case
					set3D(F, l, l+1, i, j, 1, s1[i] == s2[j])
					continue
				} // k>1, partition to check
				for m := 1; m < k; m++ {
					v := (get3D(F, l, l+1, i, j, m) && get3D(F, l, l+1, i+m, j+m, k-m)) ||
						(get3D(F, l, l+1, i, j+k-m, m) && get3D(F, l, l+1, i+m, j, k-m))
					if v {
						set3D(F, l, l+1, i, j, k, v)
					}
				}
			}
		}
	}
	return get3D(F, l, l+1, 0, 0, l)
}

func get3D(arr []bool, b, c, i, j, k int) bool {
	// like c/c++, int arr[a][b][c], return arr[i][j][k]
	return arr[i*b*c+j*c+k]
}

func set3D(arr []bool, b, c, i, j, k int, v bool) {
	arr[i*b*c+j*c+k] = v
}

func main() {
	fmt.Println(isScramble("abc", "cab"))
	fmt.Println(isScramble("1234", "2143"))
	fmt.Println(isScramble("1234", "2413"))
	fmt.Println(isScrambleDp("a", "b"))
	fmt.Println(isScrambleDp("ab", "ba"))
	fmt.Println(isScrambleDp("abc", "cab"))
	fmt.Println(isScrambleDp("1234", "2143"))
	fmt.Println(isScrambleDp("1234", "2413"))
}

package main

import "fmt"

// https://leetcode.com/problems/number-of-ways-of-cutting-a-pizza/

// Given a rectangular pizza represented as a rows x cols matrix containing the
// following characters: 'A' (an apple) and '.' (empty cell) and given the integer k.
// You have to cut the pizza into k pieces using k-1 cuts.
// For each cut you choose the direction: vertical or horizontal, then you choose a cut
// position at the cell boundary and cut the pizza into two pieces. If you cut the pizza
// vertically, give the left part of the pizza to a person. If you cut the pizza horizontally,
// give the upper part of the pizza to a person. Give the last piece of pizza to the last person.
// Return the number of ways of cutting the pizza such that each piece contains at least one
// apple. Since the answer can be a huge number, return this modulo 10^9 + 7.
// Example 1:
//   Input: pizza = ["A..","AAA","..."], k = 3
//   Output: 3
//   Explanation: The figure above shows the three ways to cut the pizza. Note that
//     pieces must contain at least one apple.
//       A..               A..               A│..      .│.
//       ───  ->  A│AA     ───  ->  AA│A     A│AA  ->  A│A
//       AAA      .│..     AAA      ..│.     .│..      .│.
//       ...            ,  ...           ,
// Example 2:
//   Input: pizza = ["A..","AA.","..."], k = 3
//   Output: 1
// Example 3:
//   Input: pizza = ["A..","A..","..."], k = 1
//   Output: 1
// Constraints:
//    1 <= rows, cols <= 50
//    rows == pizza.length
//    cols == pizza[i].length
//    1 <= k <= 10
//    pizza consists of characters 'A' and '.' only.

// to cut a pizza with left-upper point at p[i][j] into k part,
// denote number of ways as dp(i, j, k)
// if we do vertical cut, dp(i, j, k) += dp(i, j+x, k-1)
// if we do horizontal cut, dp(i, j, k) += dp(i+x, j, k-1)
// space O(mnk), time O(mnk*(m+n))
func ways(p []string, kp int) int {
	prime := 1000000007

	// count the number of apples in the sub-pizza with left-upper point at p[i][j]
	m, n := len(p), len(p[0])
	count := make([][]int, m+1)
	for i := range count {
		count[i] = make([]int, n+1)
	}
	for i:=m-1; i>=0; i-- {
		for j:=n-1; j>=0; j-- {
			count[i][j] = count[i+1][j]+count[i][j+1]-count[i+1][j+1]
			if p[i][j] == 'A' {
				count[i][j]++
			}
		}
	}

	// let p[k][i][j] be the ways to cut sub-pizza with left-upper point at p[i][j] into k pieces
	dp := make([][][]int, kp+1)
	for k := range dp {
		dp[k] = make([][]int, m)
		for i := range dp[k] {
			dp[k][i] = make([]int, n)
		}
	}
	for i:=m-1; i>=0; i-- {
		for j:=n-1; j>=0; j-- {
			if count[i][j] == 0 {
				continue
			}
			dp[1][i][j] = 1   // if count[i][j]>0, we can always cut it into 1 piece
			for k:=2; k<=kp; k++ {
				for nj:=j+1; nj<n; nj++ {   // vertical cut
					// e.g., to cut in to 2 pieces,
					// the first 2 cuts is 'count[i][j] == count[i][nj]' situation
					// the last cut is 'count[i][nj] == 0' situation
					//  ...A..  -> .│..A..   ..│.A..  ...│A..  ...A│..
					//  ..A...     .│.A...   ..│A...  ..A│...  ..A.│..
					if count[i][j] == count[i][nj] {   // no apple on left piece after cut
						continue
					}
					if count[i][nj] == 0 {             // no apple on right piece after cut
						break
					}
					dp[k][i][j] += dp[k-1][i][nj]
					dp[k][i][j] %= prime
				}
				for ni:=i+1; ni<m; ni++ {   // horizontal cut
					if count[i][j] == count[ni][j] {   // no apple on upper piece after cut
						continue
					}
					if count[ni][j] ==0 {              // no apple on lower piece after cut
						break
					}
					dp[k][i][j] += dp[k-1][ni][j]
					dp[k][i][j] %= prime
				}
				if dp[k][i][j] == 0 {  // if we can't cut it into k pieces, don't try k+1
					break
				}
			}
		}
	}
	return dp[kp][0][0]
}

func main() {
	fmt.Println(ways([]string{"A..","AAA","..."}, 3))
	fmt.Println(ways([]string{"A..","AA.","..."}, 3))
	fmt.Println(ways([]string{"A..","A..","..."}, 1))
	fmt.Println(ways([]string{".A..A","A.A..","A.AA.","AAAA.","A.AA."}, 5))
	fmt.Println(ways([]string{"..A.A.AAA...AAAAAA.AA..A..A.A......A.AAA.AAAAAA.AA","A.AA.A.....AA..AA.AA.A....AAA.A........AAAAA.A.AA.","A..AA.AAA..AAAAAAAA..AA...A..A...A..AAA...AAAA..AA","....A.A.AA.AA.AA...A.AA.AAA...A....AA.......A..AA.","AAA....AA.A.A.AAA...A..A....A..AAAA...A.A.A.AAAA..","....AA..A.AA..A.A...A.A..AAAA..AAAA.A.AA..AAA...AA","A..A.AA.AA.A.A.AA..A.A..A.A.AAA....AAAAA.A.AA..A.A",".AA.A...AAAAA.A..A....A...A.AAAA.AA..A.AA.AAAA.AA.","A.AA.AAAA.....AA..AAA..AAAAAAA...AA.A..A.AAAAA.A..","A.A...A.A...A..A...A.AAAA.A..A....A..AA.AAA.AA.AA.",".A.A.A....AAA..AAA...A.AA..AAAAAAA.....AA....A....","..AAAAAA..A..A...AA.A..A.AA......A.AA....A.A.AAAA.","...A.AA.AAA.AA....A..AAAA...A..AAA.AAAA.A.....AA.A","A.AAAAA..A...AAAAAAAA.AAA.....A.AAA.AA.A..A.A.A...","A.A.AA...A.A.AA...A.AA.AA....AA...AA.A..A.AA....AA","AA.A..A.AA..AAAAA...A..AAAAA.AA..AA.AA.A..AAAAA..A","...AA....AAAA.A...AA....AAAAA.A.AAAA.A.AA..AA..AAA","..AAAA..AA..A.AA.A.A.AA...A...AAAAAAA..A.AAA..AA.A","AA....AA....AA.A......AAA...A...A.AA.A.AA.A.A.AA.A","A.AAAA..AA..A..AAA.AAA.A....AAA.....A..A.AA.A.A...","..AA...AAAAA.A.A......AA...A..AAA.AA..A.A.A.AA..A.",".......AA..AA.AAA.A....A...A.AA..A.A..AAAAAAA.AA.A",".A.AAA.AA..A.A.A.A.A.AA...AAAA.A.A.AA..A...A.AAA..","A..AAAAA.A..A..A.A..AA..A...AAA.AA.A.A.AAA..A.AA..","A.AAA.A.AAAAA....AA..A.AAA.A..AA...AA..A.A.A.AA.AA",".A..AAAA.A.A.A.A.......AAAA.AA...AA..AAA..A...A.AA","A.A.A.A..A...AA..A.AAA..AAAAA.AA.A.A.A..AA.A.A....","A..A..A.A.AA.A....A...A......A.AA.AAA..A.AA...AA..",".....A..A...A.A...A..A.AA.A...AA..AAA...AA..A.AAA.","A...AA..A..AA.A.A.AAA..AA..AAA...AAA..AAA.AAAAA...","AA...AAA.AAA...AAAA..A...A..A...AA...A..AA.A...A..","A.AA..AAAA.AA.AAA.A.AA.A..AAAAA.A...A.A...A.AA....","A.......AA....AA..AAA.AAAAAAA.A.AA..A.A.AA....AA..",".A.A...AA..AA...AA.AAAA.....A..A..A.AA.A.AA...A.AA","..AA.AA.AA..A...AA.AA.AAAAAA.....A.AA..AA......A..","AAA..AA...A....A....AA.AA.AA.A.A.A..AA.AA..AAA.AAA","..AAA.AAA.A.AA.....AAA.A.AA.AAAAA..AA..AA.........",".AA..A......A.A.AAA.AAAA...A.AAAA...AAA.AAAA.....A","AAAAAAA.AA..A....AAAA.A..AA.A....AA.A...A.A....A..",".A.A.AA..A.AA.....A.A...A.A..A...AAA..A..AA..A.AAA","AAAA....A...A.AA..AAA..A.AAA..AA.........AA.AAA.A.","......AAAA..A.AAA.A..AAA...AAAAA...A.AA..A.A.AA.A.","AA......A.AAAAAAAA..A.AAA...A.A....A.AAA.AA.A.AAA.",".A.A....A.AAA..A..AA........A.AAAA.AAA.AA....A..AA",".AA.A...AA.AAA.A....A.A...A........A.AAA......A...","..AAA....A.A...A.AA..AAA.AAAAA....AAAAA..AA.AAAA..","..A.AAA.AA..A.AA.A...A.AA....AAA.A.....AAA...A...A",".AA.AA...A....A.AA.A..A..AAA.A.A.AA.......A.A...A.","...A...A.AA.A..AAAAA...AA..A.A..AAA.AA...AA...A.A.","..AAA..A.A..A..A..AA..AA...A..AA.AAAAA.A....A..A.A"}, 8))
}
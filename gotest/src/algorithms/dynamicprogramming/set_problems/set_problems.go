package set_problems

import "fmt"

// https://www.geeksforgeeks.org/coin-change-dp-7/
// Given a value N, if we want to make change for N cents,
// and we have infinite supply of each of S = [S1, S2, .. , Sm] valued coins,
// return how many ways can we make the change.
// O(MN) time and space, since table is updated line by line, it can be reduced to O(N).
func coinChange(coins []int, N int) int {
	// let m represents last m coins, n is the temporary target,
	// ways(m, n) represents the number of ways to make n cents using last m coins,
	// then ways(m, n) = ways(m-1, n) + ways(m, n-coins[M-m])
	// base case is ways(x, 0) = 1, ways(1, n) = (1 if n can divide last coin, else 0)
	M := len(coins)
	ways := make([][]int, M+1)
	for i := range ways {
		ways[i] = make([]int, N+1)
	}
	for i := 1; i <= M; i++ {
		ways[i][0] = 1 // set ways[1...M, 0] = 1
	}
	for j := 1; j <= N; j++ {
		// set ways[1, j=1...N] = 1 if j can divide the last one coin
		if j%coins[M-1] == 0 {
			ways[1][j] = 1
		}
	}
	for i := 2; i <= M; i++ {
		for j := 1; j <= N; j++ {
			ci := M - i       // current coin index
			v := ways[i-1][j] // ways[i, j] = ways[i-1, j] + ways[i, j-coins[ci]]
			if j-coins[ci] >= 0 {
				v += ways[i][j-coins[ci]]
			}
			ways[i][j] = v
		}
	}
	return ways[M][N]
}

// https://www.geeksforgeeks.org/tile-stacking-problem/
// from tiles of size 1...M, select N tiles that makes the sequence non-decreasing.
// can only use a single tile K times. return the number of ways to get sequence.
func tileStacking(M, N, K int) {
	// let m be tile sizes m...M, n be the number to select,
	// ways[m, n] is the ways to select n tile from tile size N...M
	// ways[m, n] = ways[m+1, n] + ways[m, n-1]
	// base case: ways[M-1, 1] = 1, ways[x, 0] = 1
	// we need to check if we overuses the number
	// todo complicated
}

// given set S={s0, ..., Sm} and target N, return if there is a subset
// that the sum of subset is equal to target N.
func subsetSum(set []int, N int) bool {
	// let m be the first m elements, n be the target,
	// has[m, n] = has[m-1, n] || has[m-1, n-set[m]]
	// base case: has[x, 0] = true, has[0, x] = (set[0]==x)
	M := len(set)
	has := make([][]bool, M)
	for i := range has {
		has[i] = make([]bool, N+1)
	}
	for i := 0; i < M; i++ {
		has[i][0] = true
	}
	has[0][set[0]] = true
	for i := 1; i < M; i++ {
		for j := 1; j <= N; j++ {
			v := has[i-1][j]
			if j-set[i] >= 0 {
				v = v || has[i-1][j-set[i]]
			}
			has[i][j] = v
		}
	}
	fmt.Println(has)
	return has[M-1][N]
}

// given set S={s0, ..., Sm} and target N, return all subsets that with sum N.
func perfectSum(set []int, N int) [][]int {
	// first get the has 2d array
	M := len(set)
	has := make([][]bool, M)
	for i := range has {
		has[i] = make([]bool, N+1)
	}
	for i := 0; i < M; i++ {
		has[i][0] = true
	}
	has[0][set[0]] = true
	for i := 1; i < M; i++ {
		for j := 1; j <= N; j++ {
			v := has[i-1][j]
			if j-set[i] >= 0 {
				v = v || has[i-1][j-set[i]]
			}
			has[i][j] = v
		}
	}
	// then use backtracking method to get all perfect sums
	subs := make([][]int, 0)
	sub := make([]int, M)
	perfectHelper(&subs, sub, 0, set, has, M-1, N)
	return subs
}

func perfectHelper(subs *[][]int, sub []int, index int, set []int, has [][]bool, i, j int) {
	if i == 0 && j != 0 { // i==0, j!=0, has[0][j]==true, set[0] should be included in subset
		sub[index] = set[i]
		cursub := make([]int, index+1) // include set[i], index need increment
		copy(cursub, sub[:index+1])
		*subs = append(*subs, cursub)
		return
	}
	if j == 0 { // j==0, find a subset
		cursub := make([]int, index)
		copy(cursub, sub[:index])
		*subs = append(*subs, cursub)
		return
	}
	if j-set[i] >= 0 && has[i-1][j-set[i]] { // include set[i], index need increment
		sub[index] = set[i]
		perfectHelper(subs, sub, index+1, set, has, i-1, j-set[i])
	}
	if has[i-1][j] { // not include set[i], no increment index
		perfectHelper(subs, sub, index, set, has, i-1, j)
	}
}

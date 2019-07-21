package set_problems

import (
	"fmt"
	"math"
)

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

// get the minimum operations for matrix chain multiplication
// for example, suppose A is 10 × 30, B is 30 × 5, C is 5 × 60 matrix. then,
// (AB)C = (10×30×5) + (10×5×60) = 1500 + 3000 = 4500 operations
// A(BC) = (30×5×60) + (10×30×60) = 9000 + 18000 = 27000 operations.
func matrixChainMultiplication(mats []int) int {
	// let min(i, j) be the minimum operation of matrix[i...j]
	// then min(i, j) = Min(min(i,i+1)+min(i+1,j), min(i,i+2)+min(i+2,j), ...)
	// time O(n³), space O(N²)
	N := len(mats)
	sum := make([][]int, N-1) // store min cost
	for i := 0; i < N-1; i++ {
		sum[i] = make([]int, N-1)
	}
	for i := 0; i < N-1; i++ {
		sum[i][i] = 0 // base case sum[i][i]=0, i.e., the diagonal is all zero
	}
	for sub := 1; sub < N-1; sub++ { // the matrix is updated diagonally
		for i := 0; i+sub < N-1; i++ {
			j := i + sub
			// calculate sum[i][j]
			sum[i][j] = math.MaxInt64
			for k := i; k < j; k++ { // get min sum
				// for matrix[a...b], row number is mats[a], col number is mats[b+1]
				cost := sum[i][k] + sum[k+1][j] + mats[i]*mats[k+1]*mats[j+1]
				if cost < sum[i][j] {
					sum[i][j] = cost
				}
			}
		}
	}
	return sum[0][N-2]
}

// 01 knapsack problem
// Given weights and values of n items, put these items in a knapsack
// of capacity W to get the maximum total value in the knapsack.
func knapsack01(weight, value []int, W int) int {
	// denote mv(k, w) be the max value of k items with limit weight w.
	// for last item with index k, if we include it, value = max(k-1, w).
	// if include it, value = value[k]+max(k-1, w-weight[k]), thus:
	// mv(k, w) = max( mv(k-1, w), value[k]+mv(k-1, w-weight[k]) )
	// base case: mv(0, w>=weight[0]) = value[0], mv(x, 0)=0
	// time complexity O(nW), space O(nW) can be reduced to O(W)
	mv := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		mv[i] = make([]int, W+1)
	}
	for i := 0; i < len(weight); i++ {
		mv[i][0] = 0
	}
	for j := weight[0]; j <= W; j++ { // use only first item, so mv[0][j>=weight[0]]=value[0]
		mv[0][j] = value[0]
	}
	for i := 1; i < len(weight); i++ {
		for j := 1; j <= W; j++ {
			mv[i][j] = mv[i-1][j]
			if j-weight[i] >= 0 && mv[i][j] < value[i]+mv[i-1][j-weight[i]] {
				mv[i][j] = value[i] + mv[i-1][j-weight[i]]
			}
		}
	}
	return mv[len(weight)-1][W]
}

package main

import "fmt"

// https://leetcode.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/

// Given a m x n matrix mat and an integer threshold. Return the maximum side-length
// of a square with a sum less than or equal to threshold or return 0 if there
// is no such square.
// Example 1:
//   Input: mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
//   Output: 2
//     1,1,3,2,4,3,2
//     1,1,3,2,4,3,2
//     1,1,3,2,4,3,2
//   Explanation: The maximum side length of square with sum less than 4 is 2 as shown.
// Example 2:
//   Input: mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]], threshold = 1
//   Output: 0
// Example 3:
//   Input: mat = [[1,1,1,1],[1,0,0,0],[1,0,0,0],[1,0,0,0]], threshold = 6
//   Output: 3
// Example 4:
//   Input: mat = [[18,70],[61,1],[25,85],[14,40],[11,96],[97,96],[63,45]], threshold = 40184
//   Output: 2
// Constraints:
//   1 <= m, n <= 300
//   m == mat.length
//   n == mat[i].length
//   0 <= mat[i][j] <= 10000
//   0 <= threshold <= 10^5

func maxSideLength(mat [][]int, threshold int) int {
	// acc[i][j]: the accumulated sum right and bottom of point (i, j)
	m, n := len(mat), len(mat[0])
    acc := make([][]int, m+1)
    for i := range acc {
    	acc[i] = make([]int, n+1)
	}
	for i:=m-1; i>=0; i-- {
		for j:=n-1; j>=0; j-- {
			acc[i][j] = mat[i][j]+acc[i+1][j]+acc[i][j+1]-acc[i+1][j+1]
		}
	}
	// binary search for low=0 and high=min(m,n). for each mid,
	// check if there exists one square with side length <= threshold
	// time is m*n*log(min(m,n))
	lo, hi := 1, m+1
	if m>n {
		hi = n+1
	}
	for lo<hi {
		mid := (lo+hi)/2
		exist := false
		outer: for i:=0; i+mid<=m; i++ {
			for j:=0; j+mid<=n; j++ {
				area := acc[i][j]-acc[i+mid][j]-acc[i][j+mid]+acc[i+mid][j+mid]
				if area<=threshold {
					exist = true
					break outer
				}
			}
		}
		// we want the max value of that CAN. so we search the first CAN'T. the candidate is 1 to n,
		// so we let lo=1, hi=n+1. the loop end with lo=hi with first that CAN'T. so lo-1 is max value that CAN.
		if exist {
			lo = mid+1
		} else {
			hi = mid
		}
	}
	return lo-1
}

func main() {
	fmt.Println(maxSideLength([][]int{{1,1,3,2,4,3,2},{1,1,3,2,4,3,2},{1,1,3,2,4,3,2}}, 4), 2)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 1), 0)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 2), 1)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 3), 1)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 7), 1)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 8), 2)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 9), 2)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 31), 3)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 32), 4)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 33), 4)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 49), 4)
	fmt.Println(maxSideLength([][]int{{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2},{2,2,2,2,2}}, 50), 5)
	fmt.Println(maxSideLength([][]int{{1,1,1,1},{1,0,0,0},{1,0,0,0},{1,0,0,0}}, 6), 3)
	fmt.Println(maxSideLength([][]int{{18,70},{61,1},{25,85},{14,40},{11,96},{97,96},{63,45}}, 40184), 2)
}
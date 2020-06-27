package main

import "fmt"

// Given a matrix, and a target, return the number of non-empty submatrices that sum to target. 
// A submatrix x1, y1, x2, y2 is the set of all cells matrix[x][y] with x1 ≤ x ≤ x2 and y1 ≤ y ≤ y2. 
// Two submatrices (x1, y1, x2, y2) and (x1', y1', x2', y2') are different if they have some 
// coordinate that is different: for example, if x1 != x1'. 
// Example 1: 
//   Input: matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
//   Output: 4
//   Explanation: The four 1x1 submatrices that only contain 0.
// Example 2: 
//   Input: matrix = [[1,-1],[-1,1]], target = 0
//   Output: 5
//   Explanation: The two 1x2 submatrices, plus the two 2x1 submatrices, plus the 2x2 submatrix.
// Note: 
//   1 <= matrix.length <= 300 
//   1 <= matrix[0].length <= 300 
//   -1000 <= matrix[i] <= 1000 
//   -10^8 <= target <= 10^8 

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	allcount := 0

	// suffix[i][j] is the sum of matrix[i:][j:]
	suffix := make([][]int, m+1)
	for i := range suffix {
		suffix[i] = make([]int, n+1)
	}
	for i:=m-1; i>=0; i-- {
		for j:=n-1; j>=0; j-- {
			suffix[i][j] = matrix[i][j] + suffix[i+1][j] + suffix[i][j+1] - suffix[i+1][j+1]
		}
	}

	// main logic, O(n*n*m)
	mapp := make(map[int]int, m+1)
	for i:=0; i<n; i++ {
		for j:=i+1; j<=n; j++ {      // choose every two-column pairs
			for r:=0; r<m; r++ {     // for each row
				v := suffix[r][i]-suffix[r][j]  // sum(matrix[r:][i:j])
				if v==target {
					allcount++
				}
				// if there are some x that make sum(matrix([x:][i:j])==v+target, 
				// and we have sum(matrix[r:][i:j])=v, we will get sum(matrix([x:r][i:j])==target, 
				// so we add the number of such x to allcount
				if vv, ok := mapp[v+target]; ok {
					allcount += vv
				}
				mapp[v] = mapp[v]+1
			}
			for k := range mapp {   // clear the map
				delete(mapp, k)
			}
		}
	}
	return allcount
}

func main() {
	fmt.Println(numSubmatrixSumTarget([][]int{{0,1,0},{1,1,1},{0,1,0}}, 0), 4)
	fmt.Println(numSubmatrixSumTarget([][]int{{1,-1},{-1,1}}, 0), 5)
	fmt.Println(numSubmatrixSumTarget([][]int{{1,5},{2,4},{6,2}}, 6), 4)
}
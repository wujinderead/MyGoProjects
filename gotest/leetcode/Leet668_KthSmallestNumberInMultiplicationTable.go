package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/

// Nearly every one have used the Multiplication Table. But could you find out
// the k-th smallest number quickly from the multiplication table?
// Given the height m and the length n of a m * n Multiplication Table, and a
// positive integer k, you need to return the k-th smallest number in this table.
// Example 1:
//   Input: m = 3, n = 3, k = 5
//   Output:
//   Explanation:
//     The Multiplication Table:
//       1 2 3
//       2 4 6
//       3 6 9
//     The 5-th smallest number is 3 (1, 2, 2, 3, 3).
// Example 2:
//   Input: m = 2, n = 3, k = 6
//   Output:
//   Explanation:
//     The Multiplication Table:
//       1 2 3
//       2 4 6
//     The 6-th smallest number is 6 (1, 2, 2, 3, 4, 6).
// Note:
//   The m and n will be in the range [1, 30000].
//   The k will be in the range [1, m * n]

func findKthNumber(m int, n int, k int) int {
	if m>n {
		m, n = n, m   // let n larger
	}
    lo, hi := 1, m*n
    for lo<hi {
    	mid := (lo+hi)/2
		// count how many numbers <= mid
		count := (mid/n)*n
		for i:=mid/n; i<m; i++ {
			j := mid/(i+1)
			if j==0 {
				break
			}
			count += j
		}
		// binary search
		if count>=k {
			hi = mid
		} else {
			lo = mid+1
		}
	}
	return hi
}

func main() {
	fmt.Println(findKthNumber(30000,30000, 30000*30000/2-17))
	fmt.Println(findKthNumber(30000,30000, 30000*30000))
	for _, v := range [][2]int{{1,1}, {1,2}, {3,1}, {3,3}, {3,4}, {5,4}} {
		m, n := v[0], v[1]
		arr := make([]int, m*n)
		for i:=0; i<m; i++ {
			for j:=0; j<n; j++ {
				arr[i*n+j] = (i+1)*(j+1)
			}
		}
		sort.Sort(sort.IntSlice(arr))
		fmt.Println(m, n, arr)
		for i:=1; i<=m*n; i++ {
			fmt.Println(findKthNumber(m, n, i), arr[i-1])
		}
		fmt.Println()
	}
}
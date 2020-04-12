package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/

// There is a one-dimensional garden on the x-axis. The garden starts at the point 0 and
// ends at the point n. (i.e The length of the garden is n).
// There are n + 1 taps located at points [0, 1, ..., n] in the garden.
// Given an integer n and an integer array ranges of length n + 1 where ranges[i] (0-indexed)
// means the i-th tap can water the area [i - ranges[i], i + ranges[i]] if it was open.
// Return the minimum number of taps that should be open to water the whole garden,
// If the garden cannot be watered return -1.
// Example 1:
//   Input: n = 5, ranges = [3,4,1,1,0,0]
//   Output: 1
//   Explanation: The tap at point 0 can cover the interval [-3,3]
//  3          +-+-+
//  2        +-+-+
//  1      +-+-+-+-+-+
//  0      +-+-+-+
//         +-+-+-+-+-+
//         0 1 2 3 4 5
//     The tap at point 1 can cover the interval [-3,5]
//     The tap at point 2 can cover the interval [1,3]
//     The tap at point 3 can cover the interval [2,4]
//     The tap at point 4 can cover the interval [4,4]
//     The tap at point 5 can cover the interval [5,5]
//     Opening Only the second tap will water the whole garden [0,5]
// Example 2:
//   Input: n = 3, ranges = [0,0,0,0]
//   Output: -1
//   Explanation: Even if you activate all the four taps you cannot water the whole garden.
// Example 3:
//   Input: n = 7, ranges = [1,2,1,0,2,1,0,1]
//   Output: 3
// Example 4:
//   Input: n = 8, ranges = [4,0,0,0,0,0,0,0,4]
//   Output: 2
// Example 5:
//   Input: n = 8, ranges = [4,0,0,0,4,0,0,0,4]
//   Output: 1
// Constraints:
//   1 <= n <= 10^4
//   ranges.length == n + 1
//   0 <= ranges[i] <= 100

// O(n*n), can improved to O(nr), r is max(ranges[i])
func minTaps(n int, ranges []int) int {
	// let F(i, j) be the minimal number taps of taps[i:] to water ranges[j: end]
	// then F(i, j) has two candidate, we want the min value:
	//   1+F(i+1, j+X)  // use tap[i], the use tap[i+1:] to water range[j+X: end], X is the area tap[i] can water
	//   F(i+1, j)      // no use tap[i], use tap[i+1:] to water range[j: end]
	// n+1 taps (from 0 to n), n segment (0 from to n-1)
	// old[i] is use last tap to cover area i:end
	old, new := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ { // the tap n can water
		old[i] = -1
		if i >= n-ranges[n] {
			old[i] = 1
		}
	}
	for k := n - 1; k >= 0; k-- { // from tap n-1 to tap 0
		if ranges[k] == 0 {
			continue // ignore 0 tap
		}
		for i := 0; i < n; i++ {
			if i < k-ranges[k] || i >= k+ranges[k] {
				new[i] = old[i]
				continue
			}
			if i >= k-ranges[k] && i < k+ranges[k] {
				fi1 := -1
				if k+ranges[k] < n {
					fi1 = old[k+ranges[k]]
				} else {
					fi1 = 0
				}
				if fi1 != -1 {
					fi1 = 1 + fi1
				}
				if old[i] == -1 && fi1 == -1 { // both -1
					new[i] = -1
				} else if old[i] != -1 && fi1 != -1 { // neither -1
					new[i] = min(old[i], fi1)
				} else { // one -1
					new[i] = old[i] + fi1 + 1
				}
			}
		}
		old, new = new, old
	}
	return old[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
	fmt.Println(minTaps(3, []int{0, 0, 0, 0}))
	fmt.Println(minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1}))
	fmt.Println(minTaps(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4}))
	fmt.Println(minTaps(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4}))
}

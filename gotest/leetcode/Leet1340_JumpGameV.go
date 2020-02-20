package main

import "fmt"

// https://leetcode.com/problems/jump-game-v/

// Given an array of integers arr and an integer d. In one step you can jump from
// index i to index:
//   i + x where: i + x < arr.length and 0 < x <= d.
//   i - x where: i - x >= 0 and 0 < x <= d.
// In addition, you can only jump from index i to index j if arr[i] > arr[j] and
// arr[i] > arr[k] for all indices k between i and j
// (More formally min(i, j) < k < max(i, j)).
// You can choose any index of the array and start jumping. Return the maximum
// number of indices you can visit.
// Notice that you can not jump outside of the array at any time.
// Example 1:
//   Input: arr = [6,4,14,6,8,13,9,7,10,6,12], d = 2
//   Output: 4
//   Explanation:
//     You can start at index 10. You can jump 10 --> 8 --> 6 --> 7 as shown.
//     Note that if you start at index 6 you can only jump to index 7. You cannot jump
//     to index 5 because 13 > 9. You cannot jump to index 4 because index 5 is between
//     index 4 and 6 and 13 > 9.
//     Similarly You cannot jump from index 3 to index 2 or index 1.
//
// h:  6 4 14 6 8 13 9 7 10 6 12
//         ▇
//         ▇      ▇
//         ▇      ▇           ▇
//         ▇      ▇      ▇    ▇
//         ▇      ▇  ▇   ▇    ▇
//         ▇    ▇ ▇  ▇   ▇    ▇
//         ▇    ▇ ▇  ▇ ▇ ▇    ▇
//     ▇   ▇  ▇ ▇ ▇  ▇ ▇ ▇  ▇ ▇
//     ▇ ▇ ▇  ▇ ▇ ▇  ▇ ▇ ▇  ▇ ▇
// i:  0 1 2  3 4 5  6 7 8  9 10
// Example 2:
//   Input: arr = [3,3,3,3,3], d = 3
//   Output: 1
//   Explanation: You can start at any index. You always cannot jump to any index.
// Example 3:
//   Input: arr = [7,6,5,4,3,2,1], d = 1
//   Output: 7
//   Explanation: Start at index 0. You can visit all the indicies.
// Example 4:
//   Input: arr = [7,1,7,1,7,1], d = 2
//   Output: 2
// Example 5:
//   Input: arr = [66], d = 1
//   Output: 1
// Constraints:
//   1 <= arr.length <= 1000
//   1 <= arr[i] <= 10^5
//   1 <= d <= arr.length

func maxJumps(arr []int, d int) int {
	// let mjp[i] be the maximum indices can visit,
	// then mjp[i]=1+max(mjp[k, i can visit k])
	counted := make([]bool, len(arr))
	mjp := make([]int, len(arr))
	curmax := 0
	for i := 0; i < len(arr); i++ {
		if !counted[i] {
			count(arr, i, d, counted, mjp)
		}
		curmax = max(curmax, mjp[i])
	}
	return curmax
}

func count(arr []int, i, d int, counted []bool, mjp []int) {
	counted[i] = true
	curmax := 0
	for j := 1; j <= d; j++ { // to right
		if i+j >= len(arr) || arr[i] <= arr[i+j] { // can't visit, stop
			break
		}
		if !counted[i+j] {
			count(arr, i+j, d, counted, mjp)
		}
		curmax = max(curmax, mjp[i+j])
	}
	for j := 1; j <= d; j++ { // to left
		if i-j < 0 || arr[i] <= arr[i-j] { // can't visit, stop
			break
		}
		if !counted[i-j] {
			count(arr, i-j, d, counted, mjp)
		}
		curmax = max(curmax, mjp[i-j])
	}
	mjp[i] = curmax + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
	fmt.Println(maxJumps([]int{3, 3, 3, 3, 3}, 3))
	fmt.Println(maxJumps([]int{7, 6, 5, 4, 3, 2, 1}, 1))
	fmt.Println(maxJumps([]int{7, 1, 7, 1, 7, 1}, 2))
	fmt.Println(maxJumps([]int{66}, 1))
}

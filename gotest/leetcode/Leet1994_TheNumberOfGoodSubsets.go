package main

import "fmt"

// https://leetcode.com/problems/the-number-of-good-subsets/

// You are given an integer array nums. We call a subset of nums good if its product can be
// represented as a product of one or more distinct prime numbers.
// For example, if nums = [1, 2, 3, 4]:
//   [2, 3], [1, 2, 3], and [1, 3] are good subsets with products 6 = 2*3, 6 = 2*3, and 3 = 3 respectively.
//   [1, 4] and [4] are not good subsets with products 4 = 2*2 and 4 = 2*2 respectively.
// Return the number of different good subsets in nums modulo 10^9 + 7.
// A subset of nums is any array that can be obtained by deleting some (possibly none or all)
// elements from nums. Two subsets are different if and only if the chosen indices to delete are different.
// Example 1:
//   Input: nums = [1,2,3,4]
//   Output: 6
//   Explanation: The good subsets are:
//     - [1,2]: product is 2, which is the product of distinct prime 2.
//     - [1,2,3]: product is 6, which is the product of distinct primes 2 and 3.
//     - [1,3]: product is 3, which is the product of distinct prime 3.
//     - [2]: product is 2, which is the product of distinct prime 2.
//     - [2,3]: product is 6, which is the product of distinct primes 2 and 3.
//     - [3]: product is 3, which is the product of distinct prime 3.
// Example 2:
//   Input: nums = [4,2,3,15]
//   Output: 5
//   Explanation: The good subsets are:
//     - [2]: product is 2, which is the product of distinct prime 2.
//     - [2,3]: product is 6, which is the product of distinct primes 2 and 3.
//     - [2,15]: product is 30, which is the product of distinct primes 2, 3, and 5.
//     - [3]: product is 3, which is the product of distinct prime 3.
//     - [15]: product is 15, which is the product of distinct primes 3 and 5.
// Constraints:
//   1 <= nums.length <= 10^5
//   1 <= nums[i] <= 30

// use bitmask to represent primes. e.g., 2=0010, 3=0100, 5=1000;
// then their product can be represented as bitwise-OR of their masks: 2*3=0110, 2*3*5=1110...
// also, we can use bitwise-AND to check if two products contains distinct primes.
// for this problem, there are 10 primes, with less than 2^10 distinct products.
// so time complexity is O(30*1024), space is O(1024)
func numberOfGoodSubsets(nums []int) int {
	const mod = int(1e9 + 7)
	// mask interested numbers
	interest := [31]int{
		1:  1,
		2:  1 << 1,
		3:  1 << 2,
		5:  1 << 3,
		7:  1 << 4,
		11: 1 << 5,
		13: 1 << 6,
		17: 1 << 7,
		19: 1 << 8,
		23: 1 << 9,
		29: 1 << 10,
	}
	interest[6] = interest[2] + interest[3]
	interest[10] = interest[2] + interest[5]
	interest[14] = interest[2] + interest[7]
	interest[15] = interest[3] + interest[5]
	interest[21] = interest[3] + interest[7]
	interest[22] = interest[2] + interest[11]
	interest[26] = interest[2] + interest[13]
	interest[30] = interest[2] + interest[3] + interest[5]

	// count occurrence
	count := [31]int{}
	for _, v := range nums {
		if interest[v] > 0 {
			count[v]++
		}
	}

	// get products count
	// could also use [1024]int{}, as there are only 10 primes, their possible distinct products are less than 2^10
	products := make(map[int]int)
	for i := 2; i <= 30; i++ {
		if interest[i] == 0 || count[i] == 0 {
			continue
		}
		ic := count[i]
		mask := interest[i]
		for k, v := range products {
			if k&mask == 0 { // k&mask==0, means k and mask are co-prime, we can add k*mask to the result
				products[k|mask] = products[k|mask] + (v*ic)%mod
			}
		}
		products[mask] = products[mask] + ic
	}

	// sum, then multiple 2^number_of_one
	sum := 0
	for _, v := range products {
		sum = (sum + v) % mod
	}
	for i := 0; i < count[1]; i++ {
		sum = (sum * 2) % mod
	}
	return sum
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{4, 2, 3, 15}, 5},
		{[]int{2, 2, 3, 3, 6, 6, 7, 7, 7, 10, 10, 10}, 79},
		{[]int{1, 1, 2, 2, 3, 3, 6, 6, 7, 7, 7, 10, 10, 10}, 316},
		{[]int{1, 1, 2, 2}, 8},
		{[]int{1, 1, 4}, 0},
		{[]int{4}, 0},
	} {
		fmt.Println(numberOfGoodSubsets(v.nums), v.ans)
	}
}

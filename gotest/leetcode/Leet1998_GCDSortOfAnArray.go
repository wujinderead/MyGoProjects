package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/gcd-sort-of-an-array/

// You are given an integer array nums, and you can perform the following operation any number
// of times on nums:
// Swap the positions of two elements nums[i] and nums[j] if gcd(nums[i], nums[j]) > 1 where
// gcd(nums[i], nums[j]) is the greatest common divisor of nums[i] and nums[j].
// Return true if it is possible to sort nums in non-decreasing order using the above swap method,
// or false otherwise.
// Example 1:
//   Input: nums = [7,21,3]
//   Output: true
//   Explanation: We can sort [7,21,3] by performing the following operations:
//     - Swap 7 and 21 because gcd(7,21) = 7. nums = [21,7,3]
//     - Swap 21 and 3 because gcd(21,3) = 3. nums = [3,7,21]
// Example 2:
//   Input: nums = [5,2,6,2]
//   Output: false
//   Explanation: It is impossible to sort the array because 5 cannot be swapped with any other element.
// Example 3:
//   Input: nums = [10,5,9,3,15]
//   Output: true
//   We can sort [10,5,9,3,15] by performing the following operations:
//     - Swap 10 and 15 because gcd(10,15) = 5. nums = [15,5,9,3,10]
//     - Swap 15 and 3 because gcd(15,3) = 3. nums = [3,5,9,15,10]
//     - Swap 10 and 15 because gcd(10,15) = 5. nums = [3,5,9,10,15]
// Constraints:
//   1 <= nums.length <= 3 * 10^4
//   2 <= nums[i] <= 105

// union-find to group all prime factors, e.g.,
//   6=2*3, {2,3} in the same set
//   35=5*7, {5,7} in the same set
//   21=3*7, so {2,3,5,7} in the same set
//   so any number that has a factor is 2,3,5,7 are in this set
//     if we need swap 6 and 35, we just check their "smallest prime factor" which is 2 and 5,
//     and (2,5) is in same set, so 6 and 35 are swappable
// use "sieve of Eratosthenes" to find the smallest prime factor and all prime factors for a range of numbers
func gcdSort(nums []int) bool {
	spf := make([]int, 100001)
	for i := 0; i < len(spf); i++ {
		spf[i] = i
	}
	for i := 2; i < len(spf); i++ {
		if spf[i] == i { // i is a prime
			// mark i*2, i*3, i*4 ... is non-prime, however no need to mark these,
			// just need to start at i*i, mark i*i, i*(i+1), i*(i+2) as non-prime
			// with their smallest prime factor as i
			for j := i; i*j < len(spf); j++ {
				if spf[i*j] > i {
					spf[i*j] = i
				}
			}
		}
	}

	// union find all primes factors, for each number in nums
	root := make([]int, 100001)
	for _, num := range nums {
		// merge all other factors with spf; e.g., if 30=2*3*5, merge (2,3), merge (2,5); this will group (2,3,5)
		s := spf[num]
		for num%s == 0 {
			num = num / s
		}
		for num > 1 {
			ss := spf[num]
			for num%ss == 0 {
				num = num / ss
			}
			// merge s and ss
			rs := getRoot(root, s)
			rss := getRoot(root, ss)
			if rs != rss { // not same set, merge it
				root[rss] = rs
			}
		}
	}

	// check the answer
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Sort(sort.IntSlice(sorted))
	for i := range sorted {
		// check if nums[i] and sorted[i] are swappable, just need to check if their spf are in the same set
		ra := getRoot(root, spf[nums[i]])
		rb := getRoot(root, spf[sorted[i]])
		if ra != rb {
			return false
		}
	}
	return true
}

func getRoot(root []int, index int) int {
	if root[index] != 0 {
		x := getRoot(root, root[index])
		root[index] = x
		return x
	}
	return index
}

func main() {
	for _, v := range []struct {
		n   []int
		ans bool
	}{
		{[]int{7, 21, 3}, true},
		{[]int{5, 2, 6, 2}, false},
		{[]int{10, 5, 9, 3, 30}, true},
		{[]int{7 * 7 * 11 * 11 * 13, 2 * 2 * 3 * 3 * 5 * 5}, false},
		{[]int{7 * 7 * 11 * 11 * 13, 2 * 2 * 3 * 3 * 5 * 5 * 7}, true},
		{[]int{7 * 7 * 11 * 11 * 13, 2 * 2 * 3 * 3 * 5 * 5, 5 * 13}, true},
		{[]int{128 * 13, 2 * 2 * 3 * 3 * 5 * 5}, true},
		{[]int{128 * 13, 3 * 3 * 5 * 5}, false},
		{[]int{128 * 13, 3 * 3 * 5 * 5, 2 * 3}, true},
	} {
		fmt.Println(gcdSort(v.n), v.ans)
	}
}

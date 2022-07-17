package main

import "fmt"

// https://leetcode.com/problems/maximum-and-sum-of-array/

// You are given an integer array nums of length n and an integer numSlots such that 2 * numSlots >= n.
// There are numSlots slots numbered from 1 to numSlots.
// You have to place all n integers into the slots such that each slot contains at most two numbers.
// The AND sum of a given placement is the sum of the bitwise AND of every number with its respective
// slot number.
// For example, the AND sum of placing the numbers [1, 3] into slot 1 and [4, 6] into slot 2 is equal
// to (1 AND 1) + (3 AND 1) + (4 AND 2) + (6 AND 2) = 1 + 1 + 0 + 2 = 4.
// Return the maximum possible AND sum of nums given numSlots slots.
// Example 1:
//   Input: nums = [1,2,3,4,5,6], numSlots = 3
//   Output: 9
//   Explanation: One possible placement is [1, 4] into slot 1, [2, 6] into slot 2,
//     and [3, 5] into slot 3.
//     This gives the maximum AND sum of (1 AND 1) + (4 AND 1) + (2 AND 2) + (6 AND 2)
//     + (3 AND 3) + (5 AND 3) = 1 + 0 + 2 + 2 + 3 + 1 = 9.
// Example 2:
//   Input: nums = [1,3,10,4,7,1], numSlots = 9
//   Output: 24
//   Explanation: One possible placement is [1, 1] into slot 1, [3] into slot 3, [4]
//     into slot 4, [7] into slot 7, and [10] into slot 9.
//     This gives the maximum AND sum of (1 AND 1) + (1 AND 1) + (3 AND 3) + (4 AND 4)
//     + (7 AND 7) + (10 AND 9) = 1 + 1 + 3 + 4 + 7 + 8 = 24.
//     Note that slots 2, 5, 6, and 8 are empty which is permitted.
// Constraints:
//   n == nums.length
//   1 <= numSlots <= 9
//   1 <= n <= 2 * numSlots
//   1 <= nums[i] <= 15

func maximumANDSum(nums []int, numSlots int) int {
	// we have 2*numSlots slots, number from 1 to 2*numSlots, slots[i]'s value is (i+1)/2
	// let dp[i][mask] is the answer to put masked nums to slots[:i]
	// then for each j in mask, we can put nums[j] to slots[i],
	// we got a candidate (nums[j] & ((i+1)/2)) + dp[i-1][mask-(1<<j)]
	// or we ignore slots[i], we have dp[i-1][mask]
	old := make([]int, 1<<len(nums)) // initial state: empty slot
	cur := make([]int, 1<<len(nums))
	for i := 1; i <= 2*numSlots; i++ { // from slot 1 to 2*numSlots
		for mask := 1; mask < (1 << len(nums)); mask++ { // NOTE: here has some redundant: when i=1, we only need 1-bit mask; when i=2, we only need (<=2)-bit mask
			// cur[mask] is initialized as dp[i-1][mask]
			cur[mask] = old[mask]
			for j := 0; j < len(nums); j++ {
				if mask&(1<<j) > 0 {
					// put nums[j] to slot i
					cand := (nums[j] & ((i + 1) / 2)) + old[mask-(1<<j)]
					if cand > cur[mask] {
						cur[mask] = cand
					}
				}
			}
		}
		old, cur = cur, old
	}
	return old[(1<<len(nums))-1]
}

// recursive
func maximumANDSum1(nums []int, numSlots int) int {
	mem := make([][]int, 2*numSlots+1)
	for i := range mem {
		mem[i] = make([]int, 1<<len(nums))
	}
	return dp(nums, mem, 2*numSlots, (1<<len(nums))-1)
}

func dp(nums []int, mem [][]int, slot, mask int) int {
	if slot == 0 || bitCount(mask) > slot {
		return 0
	}
	if mem[slot][mask] > 0 {
		return mem[slot][mask]
	}
	ans := dp(nums, mem, slot-1, mask)
	for j := 0; j < len(nums); j++ {
		if mask&(1<<j) > 0 {
			// put nums[j] to slot i
			cand := (nums[j] & ((slot + 1) / 2)) + dp(nums, mem, slot-1, mask-(1<<j))
			if cand > ans {
				ans = cand
			}
		}
	}
	mem[slot][mask] = ans
	return ans
}

func bitCount(x int) int {
	x = x - ((x >> 1) & 0x55555555)
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0F0F0F0F
	x = x + (x >> 8)
	x = x + (x >> 16)
	return x & 0x0000003F
}

func main() {
	for _, v := range []struct {
		n       []int
		ns, ans int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3, 9},
		{[]int{3}, 1, 1},
		{[]int{2, 3}, 1, 1},
		{[]int{1, 3, 10, 4, 7, 1}, 9, 24},
		{[]int{7, 6, 13, 13, 13, 6, 3, 12, 6, 4, 10, 3, 2}, 7, 54},
	} {
		fmt.Println(maximumANDSum(v.n, v.ns), maximumANDSum1(v.n, v.ns), v.ans)
	}
}

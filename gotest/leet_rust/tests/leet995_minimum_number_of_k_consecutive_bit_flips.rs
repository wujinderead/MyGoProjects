// https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/

// You are given a binary array nums and an integer k.
// A k-bit flip is choosing a subarray of length k from nums and simultaneously changing every 0
// in the subarray to 1, and every 1 in the subarray to 0.
// Return the minimum number of k-bit flips required so that there is no 0 in the array. If it is
// not possible, return -1.
// A subarray is a contiguous part of an array.
// Example 1:
//   Input: nums = [0,1,0], k = 1
//   Output: 2
//   Explanation: Flip nums[0], then flip nums[2].
// Example 2:
//   Input: nums = [1,1,0], k = 2
//   Output: -1
//   Explanation: No matter how we flip subarrays of size 2, we cannot make the array become [1,1,1].
// Example 3:
//   Input: nums = [0,0,0,1,0,1,1,0], k = 3
//   Output: 3
//   Explanation:
//     Flip nums[0],nums[1],nums[2]: nums becomes [1,1,1,1,0,1,1,0]
//     Flip nums[4],nums[5],nums[6]: nums becomes [1,1,1,1,1,0,0,0]
//     Flip nums[5],nums[6],nums[7]: nums becomes [1,1,1,1,1,1,1,1]
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= k <= nums.length

mod _minimum_number_of_k_consecutive_bit_flips {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i32,
    }

    impl Solution {
        // space O(n); if use sliding window, space O(k); if modify original array, space O(1)
        pub fn min_k_bit_flips(nums: Vec<i32>, k: i32) -> i32 {
            let mut flip = vec![0; nums.len()+1];
            let mut prefix = 0;  // use prefix to mark a range flipped
            let mut ans = 0;
            let kk = k as usize;
            for i in 0..nums.len() {
                prefix += flip[i];
                if prefix % 2 == nums[i] {  // nums[i]==0, need flip, nums in interval [i, i+k)
                    if nums.len()-i < kk {  // remain length < k, can't make all 1
                        return -1;
                    }
                    // mark nums in interval [i, i+k) flipped
                    flip[i] += 1;
                    prefix += 1;  // because flipped incremented, prefix also
                    flip[i+kk] -= 1;
                    ans += 1;  // increment answer
                }
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![0,1,0],
                k: 1,
                ans: 2,
            },
            Solution {
                nums: vec![1,1,0],
                k: 2,
                ans: -1,
            },
            Solution {
                nums: vec![0,0,0,1,0,1,1,0],
                k: 3,
                ans: 3,
            },
            Solution {
                nums: vec![0,1,1],
                k: 2,
                ans: -1,
            },
        ];
        for i in &testcases {
            let ans = Solution::min_k_bit_flips(i.nums.clone(), i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
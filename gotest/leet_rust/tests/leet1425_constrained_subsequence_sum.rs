// https://leetcode.com/problems/constrained-subsequence-sum/

// Given an integer array nums and an integer k, return the maximum sum of a non-empty subsequence
// of that array such that for every two consecutive integers in the subsequence, nums[i] and nums[j],
// where i < j, the condition j - i <= k is satisfied.
// A subsequence of an array is obtained by deleting some number of elements (can be zero) from the
// array, leaving the remaining elements in their original order.
// Example 1:
//   Input: nums = [10,2,-10,5,20], k = 2
//   Output: 37
//   Explanation: The subsequence is [10, 2, 5, 20].
// Example 2:
//   Input: nums = [-1,-2,-3], k = 1
//   Output: -1
//   Explanation: The subsequence must be non-empty, so we choose the largest number.
// Example 3:
//   Input: nums = [10,-2,-10,-5,20], k = 2
//   Output: 23
//   Explanation: The subsequence is [10, -2, -5, 20].
// Constraints:
//   1 <= k <= nums.length <= 10⁵
//   -10⁴ <= nums[i] <= 10⁴

mod _constrained_subsequence_sum {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i32,
    }

    // sliding window: let dp[i] be the maximum result you can get if the last element is nums[i],
    // then dp[i] = max(0, dp[i-k], ..., dp[i-1]) + nums[i]
    // use deque for sliding window minimal
    use std::collections::VecDeque;
    impl Solution {
        pub fn constrained_subset_sum(nums: Vec<i32>, k: i32) -> i32 {
            let mut ans = nums[0];
            let mut dp = vec![0; nums.len()];
            let k = k as usize;
            let mut deque = VecDeque::new();
            deque.push_back(0);
            dp[0] = nums[0];
            for i in 1..nums.len() {
                // get dp[i] = max(0, dp[i-k], ..., dp[i-1]) + nums[i]
                dp[i] = 0.max(dp[*deque.front().unwrap()]) + nums[i];
                ans = ans.max(dp[i]);
                // add dp[i] to queue
                if i-*deque.front().unwrap() >= k {
                    deque.pop_front();
                }
                while deque.len() > 0 && dp[*deque.back().unwrap()] <= dp[i] {
                    deque.pop_back();
                }
                deque.push_back(i);
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![10,2,-10,5,20],
                k: 2,
                ans: 37,
            },
            Solution {
                nums: vec![-1,-2,-3],
                k: 1,
                ans: -1,
            },
            Solution {
                nums: vec![10,-2,-10,-5,20],
                k: 2,
                ans: 23,
            },
        ];
        for i in testcases {
            let ans = Solution::constrained_subset_sum(i.nums, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/

// Given an integer array nums and an integer k, return the length of the shortest non-empty
// subarray of nums with a sum of at least k. If there is no such subarray, return -1.
// A subarray is a contiguous part of an array.
// Example 1: 
//   Input: nums = [1], k = 1
//   Output: 1
// Example 2: 
//   Input: nums = [1,2], k = 4
//   Output: -1
// Example 3: 
//   Input: nums = [2,-1,2], k = 3
//   Output: 3
// Constraints:
//   1 <= nums.length <= 10⁵
//   -10⁵ <= nums[i] <= 10⁵
//   1 <= k <= 10⁹

mod _shortest_subarray_with_sum_at_least_k {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i32,
    }

    use std::collections::VecDeque;
    impl Solution {
        pub fn shortest_subarray(nums: Vec<i32>, k: i32) -> i32 {
            let mut prefix = vec![0; nums.len()+1];
            for i in 0..nums.len() {
                prefix[i+1] = prefix[i] + nums[i] as i64;
            }
            let mut ans = nums.len()+1;
            let mut queue = VecDeque::new();
            queue.push_back(0); // 0 means the empty prefix
            for i in 1..prefix.len() {
                // monotonic stack: make prefix increasing
                while queue.len() > 0 && prefix[*queue.back().unwrap()] >= prefix[i] {
                    queue.pop_back();
                }
                queue.push_back(i);
                // from queue's front, check if prefix[i] - prefix[front] >= k
                // if so, pop front and update ans.
                while prefix[i] - prefix[*queue.front().unwrap()] >= k as i64 {
                    ans = ans.min(i-*queue.front().unwrap());
                    queue.pop_front();
                }
            }
            return if ans == nums.len()+1 { -1 } else { ans as i32 };
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                k: 1,
                nums: vec![1],
                ans: 1,
            },
            Solution {
                k: 4,
                nums: vec![1,2],
                ans: -1,
            },
            Solution {
                k: 3,
                nums: vec![2,-1,2],
                ans: 3,
            },
            Solution {
                k: 1000000000,
                nums: vec![-100000; 100000],
                ans: -1,
            },
        ];
        for i in testcases {
            let ans = Solution::shortest_subarray(i.nums, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
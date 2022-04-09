// https://leetcode.com/problems/jump-game-vi/

// You are given a 0-indexed integer array nums and an integer k.
// You are initially standing at index 0. In one move, you can jump at most k steps forward without
// going outside the boundaries of the array. That is, you can jump from index i to any index in the
// range [i + 1, min(n - 1, i + k)] inclusive.
// You want to reach the last index of the array (index n - 1). Your score is the sum of all nums[j]
// for each index j you visited in the array.
// Return the maximum score you can get.
// Example 1:
//   Input: nums = [1,-1,-2,4,-7,3], k = 2
//   Output: 7
//   Explanation: You can choose your jumps forming the subsequence [1,-1,4,3] (underlined above).
//     The sum is 7.
// Example 2:
//   Input: nums = [10,-5,-2,4,0,3], k = 3
//   Output: 17
//   Explanation: You can choose your jumps forming the subsequence [10,4,3] (underlined above).
//     The sum is 17.
// Example 3:
//   Input: nums = [1,-5,-20,4,-1,3,-6,-3], k = 2
//   Output: 0
// Constraints:
//   1 <= nums.length, k <= 10⁵
//   -10⁴ <= nums[i] <= 10⁴

mod _jump_game_v_i {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i32,
    }

    // sliding window max
    use std::collections::VecDeque;
    impl Solution {
        pub fn max_result(nums: Vec<i32>, k: i32) -> i32 {
            let mut queue = VecDeque::new();
            let mut dp = vec![0; nums.len()];
            queue.push_back(0);
            dp[0] = nums[0];
            // let dp[i] be the max score when reach i, then dp[i] = max(dp[i-k], ..., dp[i-1])+nums[i]
            for i in 1..nums.len() {
                let max = dp[*queue.front().unwrap()];  // queue front is the max index
                dp[i] = max + nums[i];
                // to slide window, remove dp[i-k] and insert dp[i]; pop queue if queue front is dp[i-k]
                if *queue.front().unwrap() as i32 == (i as i32 - k) {
                    queue.pop_front();
                }
                // insert dp[i]
                while queue.len() > 0 && dp[*queue.back().unwrap()] < dp[i] {
                    queue.pop_back();
                }
                queue.push_back(i);
            }
            return dp[nums.len()-1];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,-1,-2,4,-7,3],
                k: 2,
                ans: 7,
            },
            Solution {
                nums: vec![10,-5,-2,4,0,3],
                k: 3,
                ans: 17,
            },
            Solution {
                nums: vec![1,-5,-20,4,-1,3,-6,-3],
                k: 2,
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::max_result(i.nums, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
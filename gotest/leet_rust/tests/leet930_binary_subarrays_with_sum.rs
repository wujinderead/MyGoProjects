// https://leetcode.com/problems/binary-subarrays-with-sum/

// Given a binary array nums and an integer goal, return the number of non-empty subarrays with a
// sum goal.
// A subarray is a contiguous part of the array.
// Example 1
//   Input: nums = [1,0,1,0,1], goal = 2
//   Output: 4
//   Explanation: The 4 subarrays are bolded and underlined below:
//     [1,0,1,0,1]
//     [1,0,1,0,1]
//     [1,0,1,0,1]
//     [1,0,1,0,1]
// Example 2:
//   Input: nums = [0,0,0,0,0], goal = 0
//   Output: 15
// Constraints:
//   1 <= nums.length <= 3 * 10⁴
//   nums[i] is either 0 or 1.
//   0 <= goal <= nums.length

mod _binary_subarrays_with_sum {
    struct Solution{
        nums: Vec<i32>,
        goal: i32,
        ans: i32,
    }

    // sliding window need extra space, while sliding window don't
    // easier solution: (number of subarrays that sum at most goal) - (number of subarrays that sum at most goal-1)
    impl Solution {
        pub fn num_subarrays_with_sum(nums: Vec<i32>, goal: i32) -> i32 {
            let mut ans = 0;
            // special case: goal == 0
            if goal == 0 {
                let mut zero = 0;
                for i in 0..nums.len() {
                    if nums[i] == 0 {
                        zero += 1;
                    } else {
                        ans += zero*(zero+1)/2;
                        zero = 0;
                    }
                }
                ans += zero*(zero+1)/2;
                return ans;
            }
            // goal > 0
            let mut start = 0;
            let mut i = 0;
            let mut sum = 0;
            while i < nums.len() {
                sum += nums[i];
                if sum == goal {
                    let pi = i;
                    while i+1 < nums.len() && nums[i+1] == 0 {
                        i += 1;
                    }
                    let ps = start;
                    while nums[start] != 1 {
                        start += 1;
                    }
                    ans += ((start-ps+1) * (i-pi+1)) as i32;
                    start += 1;
                    sum = goal-1;
                }
                i += 1;
            }
            return ans as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,0,1,0,1],
                goal: 2,
                ans: 4,
            },
            Solution {
                nums: vec![0,0,0,0,0],
                goal: 0,
                ans: 15,
            },
            Solution {
                nums: vec![0,0,0,1,0,0],
                goal: 0,
                ans: 9,
            },
        ];
        for i in testcases {
            let ans = Solution::num_subarrays_with_sum(i.nums, i.goal);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
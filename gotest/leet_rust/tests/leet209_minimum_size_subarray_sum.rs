// https://leetcode.com/problems/minimum-size-subarray-sum/

// Given an array of positive integers nums and a positive integer target, return the minimal length
// of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than
// or equal to target. If there is no such subarray, return 0 instead.
// Example 1:
//   Input: target = 7, nums = [2,3,1,2,4,3]
//   Output: 2
//   Explanation: The subarray [4,3] has the minimal length under the problem constraint.
// Example 2:
//   Input: target = 4, nums = [1,4,4]
//   Output: 1
// Example 3:
//   Input: target = 11, nums = [1,1,1,1,1,1,1,1]
//   Output: 0
// Constraints:
//   1 <= target <= 10⁹
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁵

mod _minimum_size_subarray_sum {
    struct Solution{
        target: i32,
        nums: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
            const MAX: usize = 1e9 as usize;
            let (mut ans, mut start, mut sum) = (MAX, 0, 0);
            for i in 0..nums.len() {
                sum += nums[i];
                while sum >= target {
                    ans = ans.min(i-start+1);
                    sum -= nums[start];
                    start += 1;
                }
            }
            return (ans % MAX) as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                target: 7,
                nums: vec![2,3,1,2,4,3],
                ans: 2,
            },
            Solution {
                target: 4,
                nums: vec![1,4,4],
                ans: 1,
            },
            Solution {
                target: 11,
                nums: vec![1,1,1,1,1,1,1],
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::min_sub_array_len(i.target, i.nums);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
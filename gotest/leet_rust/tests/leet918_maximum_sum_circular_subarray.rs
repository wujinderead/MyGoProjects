// https://leetcode.com/problems/maximum-sum-circular-subarray/

// Given a circular integer array nums of length n, return the maximum possible sum of a non-empty
// subarray of nums.
// A circular array means the end of the array connects to the beginning of the array. Formally,
// the next element of nums[i] is nums[(i + 1) % n] and the previous element of nums[i] is
// nums[(i - 1 + n) % n].
// A subarray may only include each element of the fixed buffer nums at most once. Formally, for
// a subarray nums[i], nums[i + 1], ..., nums[j], there does not exist i <= k1, k2 <= j with
// k1 % n == k2 % n.
// Example 1:
//   Input: nums = [1,-2,3,-2]
//   Output: 3
//   Explanation: Subarray [3] has maximum sum 3.
// Example 2:
//   Input: nums = [5,-3,5]
//   Output: 10
//   Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10.
// Example 3:
//   Input: nums = [-3,-2,-3]
//   Output: -2
//   Explanation: Subarray [-2] has maximum sum -2.
// Constraints:
//   n == nums.length
//   1 <= n <= 3 * 10⁴
//   -3 * 10⁴ <= nums[i] <= 3 * 10⁴

mod _maximum_sum_circular_subarray {
    struct Solution{
        nums: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn max_subarray_sum_circular(nums: Vec<i32>) -> i32 {
            let (mut min, mut max) = (nums[0], nums[0]);
            let (mut p_min, mut p_max) = (nums[0], nums[0]);
            let mut sum = nums[0];
            for i in 1..nums.len() {
                sum += nums[i];  // current prefix sum
                min = min.min(sum.min(sum-p_max));  // min-sum subarray, current sum - minimal prefix
                max = max.max(sum.max(sum-p_min));  // max-sum subarray, current sum - maximal prefix
                p_min = p_min.min(sum);   // update min prefix
                p_max = p_max.max(sum);   // update max prefix
            }
            if max > 0 {  // max subarray, or sum - min subarray
                return max.max(sum-min);
            }
            // max<=0 means all numbers <= 0
            return max;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,-2,3,-2],
                ans: 3,
            },
            Solution {
                nums: vec![5,-3,5],
                ans: 10,
            },
            Solution {
                nums: vec![-3,-2,-3],
                ans: -2,
            },
            Solution {
                nums: vec![-2],
                ans: -2,
            }
        ];
        for i in testcases {
            let ans = Solution::max_subarray_sum_circular(i.nums);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
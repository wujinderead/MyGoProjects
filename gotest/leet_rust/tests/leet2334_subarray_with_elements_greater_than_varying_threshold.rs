// https://leetcode.com/problems/subarray-with-elements-greater-than-varying-threshold/

// You are given an integer array nums and an integer threshold.
// Find any subarray of nums of length k such that every element in the subarray is greater than
// threshold / k.
// Return the size of any such subarray. If there is no such subarray, return -1.
// A subarray is a contiguous non-empty sequence of elements within an array.
// Example 1:
//   Input: nums = [1,3,4,3,1], threshold = 6
//   Output: 3
//   Explanation: The subarray [3,4,3] has a size of 3, and every element is greater than 6 / 3 = 2.
//     Note that this is the only valid subarray.
// Example 2:
//   Input: nums = [6,5,6,5,8], threshold = 7
//   Output: 1
//   Explanation: The subarray [8] has a size of 1, and 8 > 7 / 1 = 7. So 1 is returned.
//     Note that the subarray [6,5] has a size of 2, and every element is greater than 7 / 2 = 3.5.
//     Similarly, the subarrays [6,5,6], [6,5,6,5], [6,5,6,5,8] also satisfy the given conditions.
//     Therefore, 2, 3, 4, or 5 may also be returned.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i], threshold <= 10⁹

mod _subarray_with_elements_greater_than_varying_threshold {
    struct Solution{
        nums: Vec<i32>,
        threshold: i32,
        ans: i32,
    }

    impl Solution {
        // monotonic stack
        pub fn valid_subarray_size(nums: Vec<i32>, threshold: i32) -> i32 {
            let mut stack = Vec::new();
            let mut left = vec![0; nums.len()];
            for i in 0..nums.len() {
                while stack.len() > 0 && nums[*stack.last().unwrap()] >= nums[i] {
                    stack.pop();
                }
                if stack.len() > 0 {
                    left[i] = *stack.last().unwrap() as i32;
                } else {
                    left[i] = -1;
                }
                stack.push(i);
            }
            stack.truncate(0);
            for i in (0..nums.len()).into_iter().rev() {
                while stack.len() > 0 && nums[*stack.last().unwrap()] >= nums[i] {
                    stack.pop();
                }
                let mut right = nums.len() as i32;
                if stack.len() > 0 {
                    right = *stack.last().unwrap() as i32;
                }
                stack.push(i);

                // check answer
                if (right-left[i]-1)*nums[i] > threshold {
                    return right-left[i]-1;
                }
            }
            return -1;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,3,4,3,1],
                threshold: 6,
                ans: 3,
            },
            Solution {
                nums: vec![6,5,6,5,8],
                threshold: 7,
                ans: 1,
            },
            Solution {
                nums: vec![vec![1], vec![100001; 10000]].concat(),
                threshold: 1000000000,
                ans: 10000,
            },
        ];
        for i in testcases {
            let ans = Solution::valid_subarray_size(i.nums, i.threshold);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
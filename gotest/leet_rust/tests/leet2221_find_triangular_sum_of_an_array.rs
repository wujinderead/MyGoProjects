// https://leetcode.com/problems/find-triangular-sum-of-an-array/

// You are given a 0-indexed integer array nums, where nums[i] is a digit between 0 and 9 (inclusive).
// The triangular sum of nums is the value of the only element present in nums after the following
// process terminates:
//   Let nums comprise of n elements. If n == 1, end the process. Otherwise,
//     create a new 0-indexed integer array newNums of length n - 1.
//   For each index i, where 0 <= i < n - 1, assign the value of newNums[i] as (
//     nums[i] + nums[i+1]) % 10, where % denotes modulo operator.
//   Replace the array nums with newNums.
//   Repeat the entire process starting from step 1.
// Return the triangular sum of nums. 
// Example 1:
//   Input: nums = [1,2,3,4,5]
//   Output: 8
//   Explanation:
//     The above diagram depicts the process from which we obtain the triangular sum of the array.
// Example 2:
//   Input: nums = [5]
//   Output: 5
//   Explanation:
//     Since there is only one element in nums, the triangular sum is the value of that element itself.
// Constraints:
//   1 <= nums.length <= 1000
//   0 <= nums[i] <= 9

mod _find_triangular_sum_of_an_array {
    struct Solution{
        nums: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        // can be o(n) use pascal triangle
        pub fn triangular_sum(mut nums: Vec<i32>) -> i32 {
            use std::mem;
            let mut new = vec![0; nums.len()];
            for i in (1..nums.len()).into_iter().rev() {
                for j in 0..i {
                    new[j] = (nums[j]+nums[j+1])%10;
                }
                mem::swap(&mut nums, &mut new);
            }
            return nums[0];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,2,3,4,5],
                ans: 8,
            },
            Solution {
                nums: vec![5],
                ans: 5,
            }
        ];
        for i in testcases {
            let ans = Solution::triangular_sum(i.nums);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
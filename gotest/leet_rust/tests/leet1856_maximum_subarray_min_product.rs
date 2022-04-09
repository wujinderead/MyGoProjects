// https://leetcode.com/problems/maximum-subarray-min-product/

// The min-product of an array is equal to the minimum value in the array multiplied by the
// array's sum.
// For example, the array [3,2,5] (minimum value is 2) has a min-product of 2*(3+2+5)=2*10=20.
// Given an array of integers nums, return the maximum min-product of any non-empty subarray of nums.
// Since the answer may be large, return it modulo 10⁹ + 7.
// Note that the min-product should be maximized before performing the modulo operation. Testcases
// are generated such that the maximum min-product without modulo will fit in a 64-bit signed integer.
// A subarray is a contiguous part of an array.
// Example 1:
//   Input: nums = [1,2,3,2]
//   Output: 14
//   Explanation: The maximum min-product is achieved with the subarray [2,3,2] (minimum value is 2).
//     2 * (2+3+2) = 2 * 7 = 14.
// Example 2:
//   Input: nums = [2,3,3,1,2]
//   Output: 18
//   Explanation: The maximum min-product is achieved with the subarray [3,3] (minimum value is 3).
//     3 * (3+3) = 3 * 6 = 18.
// Example 3:
//   Input: nums = [3,1,5,6,4,2]
//   Output: 60
//   Explanation: The maximum min-product is achieved with the subarray [5,6,4] (minimum value is 4).
//     4 * (5+6+4) = 4 * 15 = 60.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁷

mod _maximum_subarray_min_product {
    struct Solution{
        nums: Vec<i32>,
        ans: i32,
    }

    // similar to max histogram area problem
    impl Solution {
        pub fn max_sum_min_product(nums: Vec<i32>) -> i32 {
            let mut prefix = vec![0; nums.len()+1];
            for i in 1..=nums.len() {
                prefix[i] = (nums[i-1] as i64) + prefix[i-1];  // sum(nums[i...j])=prefix[j+1]-prefix[i]
            }
            // for each nums[i], find leftmost l that nums[l] < nums[i]
            let mut l = vec![0; nums.len()];
            let mut stack = Vec::with_capacity(nums.len()+1);
            stack.push(-1);
            for i in 0..nums.len() {
                while stack.len() > 1 && nums[*stack.last().unwrap() as usize] >= nums[i] {
                    stack.pop();
                }
                l[i] = *stack.last().unwrap();
                stack.push(i as i32);
            }
            // for each nums[i], find rightmost r that nums[r] < nums[i]
            let mut r = vec![0; nums.len()];
            stack.truncate(0);
            stack.push(nums.len() as i32);
            for i in (0..nums.len()).into_iter().rev() {
                while stack.len() > 1 && nums[*stack.last().unwrap() as usize] >= nums[i] {
                    stack.pop();
                }
                r[i] = *stack.last().unwrap();
                stack.push(i as i32);
            }
            // find max answer
            let mut ans = 0;
            for i in 0..nums.len() {
                // from nums[l[i]+1] to nums[r[i]-1], the min value is nums[i]
                ans = ans.max((nums[i] as i64)*(prefix[r[i] as usize] - prefix[(l[i]+1) as usize]));
            }
            return (ans % 1000000007) as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,2,3,2],
                ans: 14,
            },
            Solution {
                nums: vec![2,3,3,1,2],
                ans: 18,
            },
            Solution {
                nums: vec![3,1,5,6,4,2],
                ans: 60,
            },
            Solution {
                nums: vec![3],
                ans: 0,
            },
            Solution {
                nums: vec![2, 3],
                ans: 10,
            },
        ];
        for i in testcases {
            let ans = Solution::max_sum_min_product(i.nums);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
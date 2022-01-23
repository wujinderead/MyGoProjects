// https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together-ii/

// A swap is defined as taking two distinct positions in an array and swapping the values in them.
// A circular array is defined as an array where we consider the first element 
// and the last element to be adjacent.
// Given a binary circular array nums, return the minimum number of swaps required to group
// all 1's present in the array together at any location.
// Example 1:
//   Input: nums = [0,1,0,1,1,0,0]
//   Output: 1
//   Explanation: Here are a few of the ways to group all the 1's together:
//     [0,0,1,1,1,0,0] using 1 swap.
//     [0,1,1,1,0,0,0] using 1 swap.
//     [1,1,0,0,0,0,1] using 2 swaps (using the circular property of the array).
//     There is no way to group all 1's together with 0 swaps.
//     Thus, the minimum number of swaps required is 1.
// Example 2:
//   Input: nums = [0,1,1,1,0,0,1,1,0]
//   Output: 2
//   Explanation: Here are a few of the ways to group all the 1's together:
//     [1,1,1,0,0,0,0,1,1] using 2 swaps (using the circular property of the array).
//     [1,1,1,1,1,0,0,0,0] using 2 swaps.
//     There is no way to group all 1's together with 0 or 1 swaps.
//     Thus, the minimum number of swaps required is 2.
// Example 3:
//   Input: nums = [1,1,0,0,1]
//   Output: 0
//   Explanation: All the 1's are already grouped together due to the circular property of the array.
//     Thus, the minimum number of swaps required is 0.
// Constraints:
//   1 <= nums.length <= 10⁵
//   nums[i] is either 0 or 1.

mod _minimum_swaps_to_group_all_1s_together_i_i {
    struct Solution{
        nums: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn min_swaps(nums: Vec<i32>) -> i32 {
            let mut ones = 0;
            for i in 0..nums.len() {   // count how many 1's, let it be ones
                if nums[i] == 1 {
                    ones += 1;
                }
            }
            if ones <= 1 || ones >= nums.len()-1 {
                return 0;
            }
            let mut min = nums.len() as i32;
            let mut sum = 0;
            for i in 0..ones {   // in a segment of length = ones
                sum += nums[i];    // count how many ones
            }
            min = min.min(ones as i32 -sum); // ones - sum is what we need swap
            for i in 1..nums.len() {
                // sliding window, note: use % to handle index going to front
                sum = sum-nums[i-1]+nums[(i+ones-1)%nums.len()];
                min = min.min(ones as i32 -sum);
            }
            return min;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![0,1,0,1,1,0,0],
                ans: 1,
            },
            Solution {
                nums: vec![0,1,1,1,0,0,1,1,0],
                ans: 2,
            },
            Solution {
                nums: vec![1,1,0,0,1],
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::min_swaps(i.nums);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/minimum-average-difference/

// You are given a 0-indexed integer array nums of length n.
// The average difference of the index i is the absolute difference between the average of the first
// i + 1 elements of nums and the average of the last n - i - 1 elements. Both averages should be
// rounded down to the nearest integer.
// Return the index with the minimum average difference. If there are multiple such indices, return
// the smallest one.
// Note:
//   The absolute difference of two numbers is the absolute value of their difference.
//   The average of n elements is the sum of the n elements divided (integer division) by n.
//   The average of 0 elements is considered to be 0.
// Example 1:
//   Input: nums = [2,5,3,9,5,3]
//   Output: 3
//   Explanation:
//     - The average difference of index 0 is: |2 / 1 - (5 + 3 + 9 + 5 + 3) / 5| = |2/ 1 - 25 / 5| = |2 - 5| = 3.
//     - The average difference of index 1 is: |(2 + 5) / 2 - (3 + 9 + 5 + 3) / 4| =|7 / 2 - 20 / 4| = |3 - 5| = 2.
//     - The average difference of index 2 is: |(2 + 5 + 3) / 3 - (9 + 5 + 3) / 3| =|10 / 3 - 17 / 3| = |3 - 5| = 2.
//     - The average difference of index 3 is: |(2 + 5 + 3 + 9) / 4 - (5 + 3) / 2| =|19 / 4 - 8 / 2| = |4 - 4| = 0.
//     - The average difference of index 4 is: |(2 + 5 + 3 + 9 + 5) / 5 - 3 / 1| = |24 / 5 - 3 / 1| = |4 - 3| = 1.
//     - The average difference of index 5 is: |(2 + 5 + 3 + 9 + 5 + 3) / 6 - 0| = |27 / 6 - 0| = |4 - 0| = 4.
//     The average difference of index 3 is the minimum average difference so return 3.
// Example 2:
//   Input: nums = [0]
//   Output: 0
//   Explanation:
//     The only index is 0 so return 0.
//     The average difference of index 0 is: |0 / 1 - 0| = |0 - 0| = 0.
// Constraints:
//   1 <= nums.length <= 10⁵
//   0 <= nums[i] <= 10⁵

mod _minimum_average_difference {
    struct Solution{
        nums: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn minimum_average_difference(nums: Vec<i32>) -> i32 {
            let nums = nums.iter().map(|&s| s as i64).collect::<Vec<_>>();
            let mut sum = nums.iter().sum::<i64>();
            let mut min = 1e10 as i64;
            let mut left = 0;
            let mut ans = 0;
            for i in 0..nums.len()-1 {
                left += nums[i];
                sum -= nums[i];
                let abs = (left/(i as i64+1)-sum/((nums.len()-i-1) as i64)).abs();
                if abs < min {
                    min = abs;
                    ans = i as i32;
                }
            }
            left += nums[nums.len()-1];
            if left/(nums.len() as i64) < min {
                ans = nums.len() as i32-1;
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![2,5,3,9,5,3],
                ans: 3,
            },
            Solution {
                nums: vec![0],
                ans: 0,
            },
            Solution {
                nums: vec![4,2,0],
                ans: 2,
            },
            Solution {
                nums: vec![0,0,0,0,0],
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::minimum_average_difference(i.nums);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
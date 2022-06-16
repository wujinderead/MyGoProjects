// https://leetcode.com/problems/subarrays-with-k-different-integers/

// Given an integer array nums and an integer k, return the number of good subarrays of nums.
// A good array is an array where the number of different integers in that array is exactly k.
// For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.
// A subarray is a contiguous part of an array.
// Example 1:
//   Input: nums = [1,2,1,2,3], k = 2
//   Output: 7
//   Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1],
//     [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2]
// Example 2:
//   Input: nums = [1,2,1,3,4], k = 3
//   Output: 3
//   Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
// Constraints:
//   1 <= nums.length <= 2 * 10⁴
//   1 <= nums[i], k <= nums.length

mod _subarrays_with_k_different_integers {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i32,
    }

    use std::collections::HashMap;
    impl Solution {
        pub fn subarrays_with_k_distinct(nums: Vec<i32>, k: i32) -> i32 {
            return Solution::at_most(&nums, k) - Solution::at_most(&nums, k-1);
        }

        fn at_most(nums: &Vec<i32>, k: i32) -> i32 {
            let mut map = HashMap::new();
            let (mut ans, mut start, mut count) = (0, 0, 0);
            for i in 0..nums.len() {
                let x = map.entry(&nums[i]).or_insert(0);
                if *x == 0 {
                    count += 1;
                }
                *x += 1;
                while count > k {
                    let x = map.get_mut(&nums[start]).unwrap();
                    *x -= 1;
                    start += 1;
                    if *x == 0 {
                        count -= 1;
                    }
                }
                ans += i+1-start;
            }
            return ans as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,2,1,2,3],
                k: 2,
                ans: 7,
            },
            Solution {
                nums: vec![1,2,1,3,4],
                k: 3,
                ans: 3,
            },
            Solution {
                nums: vec![1,2,1,1,1,3,3,3,3,4],
                k: 1,
                ans: 19,
            },
        ];
        for i in testcases {
            let ans = Solution::subarrays_with_k_distinct(i.nums, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/partition-array-such-that-maximum-difference-is-k/

// You are given an integer array nums and an integer k. You may partition nums into one or more
// subsequences such that each element in nums appears in exactly one of the subsequences.
// Return the minimum number of subsequences needed such that the difference between the maximum
// and minimum values in each subsequence is at most k.
// A subsequence is a sequence that can be derived from another sequence by deleting some or no
// elements without changing the order of the remaining elements.
// Example 1:
//   Input: nums = [3,6,1,2,5], k = 2
//   Output: 2
//   Explanation:
//     We can partition nums into the two subsequences [3,1,2] and [6,5].
//     The difference between the maximum and minimum value in the first subsequence is 3 - 1 = 2.
//     The difference between the maximum and minimum value in the second subsequence is 6 - 5 = 1.
//     Since two subsequences were created, we return 2. It can be shown that 2 is the minimum
//     number of subsequences needed.
// Example 2:
//   Input: nums = [1,2,3], k = 1
//   Output: 2
//   Explanation:
//     We can partition nums into the two subsequences [1,2] and [3].
//     The difference between the maximum and minimum value in the first subsequence is 2 - 1 = 1.
//     The difference between the maximum and minimum value in the second subsequence is 3 - 3 = 0.
//     Since two subsequences were created, we return 2. Note that another optimal
//     solution is to partition nums into the two subsequences [1] and [2,3].
// Example 3:
//  Input: nums = [2,2,4,5], k = 0
//  Output: 3
//  Explanation:
//    We can partition nums into the three subsequences [2,2], [4], and [5].
//    The difference between the maximum and minimum value in the first subsequences is 2 - 2 = 0.
//    The difference between the maximum and minimum value in the second subsequences is 4 - 4 = 0.
//    The difference between the maximum and minimum value in the third subsequences is 5 - 5 = 0.
//    Since three subsequences were created, we return 3. It can be shown that 3 is
//    the minimum number of subsequences needed.
// Constraints:
//   1 <= nums.length <= 10⁵
//   0 <= nums[i] <= 10⁵
//   0 <= k <= 10⁵

mod _partition_array_such_that_maximum_difference_is_k {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i32,
    }

    impl Solution {
        pub fn partition_array(mut nums: Vec<i32>, k: i32) -> i32 {
            nums.sort();
            let mut i = 0;
            let mut count = 0;
            while i<nums.len() {
                count += 1;
                let upper = nums[i]+k;
                while i<nums.len() && nums[i] <= upper {
                    i += 1;
                }
            }
            return count;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![3,6,1,2,5],
                k: 2,
                ans: 2,
            },
            Solution {
                nums: vec![1,2,3],
                k: 1,
                ans: 2,
            },
            Solution {
                nums: vec![2,2,4,5],
                k: 0,
                ans: 3,
            },
            Solution {
                nums: vec![1,2],
                k: 1,
                ans: 1,
            },
            Solution {
                nums: vec![1],
                k: 1,
                ans: 1,
            },
        ];
        for i in testcases {
            let ans = Solution::partition_array(i.nums, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
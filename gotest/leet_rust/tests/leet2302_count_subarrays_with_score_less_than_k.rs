// https://leetcode.com/problems/count-subarrays-with-score-less-than-k/

// The score of an array is defined as the product of its sum and its length.
// For example, the score of [1, 2, 3, 4, 5] is (1 + 2 + 3 + 4 + 5) * 5 = 75.
// Given a positive integer array nums and an integer k, return the number of non-empty subarrays
// of nums whose score is strictly less than k.
// A subarray is a contiguous sequence of elements within an array.
// Example 1:
//   Input: nums = [2,1,4,3,5], k = 10
//   Output: 6
//   Explanation:
//     The 6 subarrays having scores less than 10 are:
//     - [2] with score 2 * 1 = 2.
//     - [1] with score 1 * 1 = 1.
//     - [4] with score 4 * 1 = 4.
//     - [3] with score 3 * 1 = 3.
//     - [5] with score 5 * 1 = 5.
//     - [2,1] with score (2 + 1) * 2 = 6.
//     Note that subarrays such as [1,4] and [4,3,5] are not considered because
//     their scores are 10 and 36 respectively, while we need scores strictly less than 10.
// Example 2:
//   Input: nums = [1,1,1], k = 5
//   Output: 5
//   Explanation:
//     Every subarray except [1,1,1] has a score less than 5.
//     [1,1,1] has a score (1 + 1 + 1) * 3 = 9, which is greater than 5.
//     Thus, there are 5 subarrays having scores less than 5.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁵
//   1 <= k <= 10¹⁵

mod _count_subarrays_with_score_less_than_k {
    struct Solution{
        nums: Vec<i32>,
        k: i64,
        ans: i32,
    }

    impl Solution {
        pub fn count_subarrays(nums: Vec<i32>, k: i64) -> i64 {
            let nums = nums.iter().map(|&s| s as i64).collect::<Vec<_>>();
            let mut ans = 0;
            let mut sum = 0;
            let mut i = 0;  // the start of valid subarray
            for (j, &v) in nums.iter().enumerate() {
                sum += v;
                while sum * ((j-i+1) as i64) >= k { // skip invalid start
                    sum -= nums[i];
                    i += 1;
                }
                ans += (j - i + 1) as i64;  // the number of valid subarray that ends at j
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![2,1,4,3,5],
                k: 10,
                ans: 6,
            },
            Solution {
                nums: vec![1,1,1],
                k: 5,
                ans: 5,
            },
        ];
        for i in testcases {
            let ans = Solution::count_subarrays(i.nums, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
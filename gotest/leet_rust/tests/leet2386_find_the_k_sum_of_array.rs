// https://leetcode.com/problems/find-the-k-sum-of-an-array/

// You are given an integer array nums and a positive integer k. You can choose
// any subsequence of the array and sum all of its elements together.
// We define the K-Sum of the array as the kᵗʰ largest subsequence sum that can
// be obtained (not necessarily distinct).
// Return the K-Sum of the array.
// A subsequence is an array that can be derived from another array by deleting
// some or no elements without changing the order of the remaining elements.
// Note that the empty subsequence is considered to have a sum of 0.
// Example 1:
//   Input: nums = [2,4,-2], k = 5
//   Output: 2
//   Explanation: All the possible subsequence sums that we can obtain are the
//     following sorted in decreasing order:
//     - 6, 4, 4, 2, 2, 0, 0, -2.
//     The 5-Sum of the array is 2.
// Example 2:
//   Input: nums = [1,-2,3,4,-10,12], k = 16
//   Output: 10
//   Explanation: The 16-Sum of the array is 10.
// Constraints:
//   n == nums.length
//   1 <= n <= 10⁵
//   -10⁹ <= nums[i] <= 10⁹
//   1 <= k <= min(2000, 2ⁿ)

mod _find_the_k_sum_of_array {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i64,
    }

    use std::collections::BinaryHeap;
    impl Solution {
        pub fn k_sum(mut nums: Vec<i32>, k: i32) -> i64 {
            let n = nums.len();
            // max subset: sum of positive numbers
            let mut ans = nums.iter().fold(0 as i64, |acc, &x| acc+0.max(x as i64));
            nums.sort_by_key(|&s| s.abs());  // sort by abs

            let mut heap = BinaryHeap::new();
            // reduce the max set by remove number with minimal abs()
            //   sum of current subset
            //   the index of subset's last removed number in sorted nums
            heap.push((ans-nums[0].abs() as i64, 0));

            let mut kk= k-1;
            while kk>0 {
                kk -= 1;
                let (s, i) = heap.pop().unwrap();  // pop current smallest-sum subset
                ans = s;
                // let subset {x, ..., y, nums[i]} be the current minimal set
                // then the next possible larger subsets are
                //   {x, ..., y, nums[i], nums[i+1]}, by add next larger number, or,
                //   or {x, ..., y, nums[i+1]}, by substitute last number.
                // push these 2 numbers to heap
                if i+1<n {
                    heap.push((s-nums[i+1].abs() as i64, i+1, ));
                    heap.push((s-nums[i+1].abs() as i64+nums[i].abs() as i64, i+1));
                }
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![2,4,-2],
                k: 5,
                ans: 2,
            },
            Solution {
                nums: vec![1,-2,3,4,-10,12],
                k: 16,
                ans: 10,
            },
        ];
        for i in testcases {
            let ans = Solution::k_sum(i.nums, i.k);
            println!("{}, {}", ans, i.ans);
        }
        for i in 1..=16 {
            let ans = Solution::k_sum(vec![1,2,4,8], i);
            println!("{}, {}", i, ans);
        }
    }
}
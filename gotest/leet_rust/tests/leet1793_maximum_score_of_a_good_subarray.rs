// https://leetcode.com/problems/maximum-score-of-a-good-subarray/

// You are given an array of integers nums (0-indexed) and an integer k.
// The score of a subarray (i, j) is defined as min(nums[i], nums[i+1], ...,nums[j]) * (j - i + 1).
// A good subarray is a subarray where i <= k <= j.
// Return the maximum possible score of a good subarray.
// Example 1:
//   Input: nums = [1,4,3,7,4,5], k = 3
//   Output: 15
//   Explanation: The optimal subarray is (1, 5) with a score of min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15.
// Example 2:
//   Input: nums = [5,5,4,5,4,1,1,1], k = 0
//   Output: 20
//   Explanation: The optimal subarray is (0, 4) with a score of min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 2 * 10⁴
//   0 <= k < nums.length

mod _maximum_score_of_a_good_subarray {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i32,
    }

    impl Solution {
        pub fn maximum_score(nums: Vec<i32>, k: i32) -> i32 {
            let (mut ans, mut min, mut l, mut r) = (nums[k as usize], nums[k as usize], k, k);
            while l>0 || r<(nums.len() as i32-1) {
                if l==0 {  // can only go right
                    r += 1;
                } else if r == (nums.len() as i32-1) {  // can only go left
                    l -= 1;
                } else if nums[(l-1) as usize] > nums[(r+1) as usize] { // can go both side, while left has larger value
                    l -= 1;
                } else {
                    r += 1;
                }
                min = min.min(nums[l as usize].min(nums[r as usize]));
                ans = ans.max(min*(r-l+1));
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,4,3,7,4,5],
                k: 3,
                ans: 15,
            },
            Solution {
                nums: vec![5,5,4,5,4,1,1,1],
                k: 0,
                ans: 20,
            },
        ];
        for i in testcases {
            let ans = Solution::maximum_score(i.nums, i.k);
            println!("{}, {}", ans, i.ans);
        }
    }
}
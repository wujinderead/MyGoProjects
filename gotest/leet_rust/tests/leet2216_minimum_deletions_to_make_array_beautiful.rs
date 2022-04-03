// https://leetcode.com/problems/minimum-deletions-to-make-array-beautiful/

// You are given a 0-indexed integer array nums. The array nums is beautiful if:
//   nums.length is even.
//   nums[i] != nums[i + 1] for all i % 2 == 0.
// Note that an empty array is considered beautiful.
// You can delete any number of elements from nums. When you delete an element, all the
// elements to the right of the deleted element will be shifted one unit to the left to
// fill the gap created and all the elements to the left of the deleted element will
// remain unchanged.
// Return the minimum number of elements to delete from nums to make it beautiful.
// Example 1:
//   Input: nums = [1,1,2,3,5]
//   Output: 1
//   Explanation: You can delete either nums[0] or nums[1] to make nums = [1,2,3,5]
//     which is beautiful. It can be proven you need at least 1 deletion to make nums beautiful.
// Example 2:
//   Input: nums = [1,1,2,2,3,3]
//   Output: 2
//   Explanation: You can delete nums[0] and nums[5] to make nums = [1,2,2,3]
//     which is beautiful. It can be proven you need at least 2 deletions to make num beautiful.
// Constraints:
//   1 <= nums.length <= 10⁵
//   0 <= nums[i] <= 10⁵

mod _minimum_deletions_to_make_array_beautiful {
    struct Solution{
        nums: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        // simple greedy solution
        pub fn min_deletion(nums: Vec<i32>) -> i32 {
            let mut has_prev = false;
            let mut prev = 0;
            let mut sum = 0;
            for i in 0..nums.len() {
                if !has_prev {
                    prev  = nums[i];
                    has_prev = true;
                    continue;
                }
                if prev == nums[i] {
                    sum += 1;
                } else {
                    has_prev = false;
                }
            }
            if has_prev {
                sum += 1;
            }
            return sum;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,1,2,3,5],
                ans: 1,
            },
            Solution {
                nums: vec![1,1,2,2,3,3],
                ans: 2,
            },
            Solution {
                nums: vec![2],
                ans: 1,
            },
            Solution {
                nums: vec![2,2],
                ans: 2,
            },
            Solution {
                nums: vec![2,2,2],
                ans: 3,
            },
            Solution {
                nums: vec![1,2],
                ans: 0,
            },
            Solution {
                nums: vec![1,2,2,3],
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::min_deletion(i.nums);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
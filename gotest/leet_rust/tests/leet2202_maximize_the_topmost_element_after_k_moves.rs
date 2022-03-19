// https://leetcode.com/problems/maximize-the-topmost-element-after-k-moves/

// You are given a 0-indexed integer array nums representing the contents of a pile,
// where nums[0] is the topmost element of the pile.
// In one move, you can perform either of the following:
//   If the pile is not empty, remove the topmost element of the pile.
//   If there are one or more removed elements, add any one of them back onto the pile.
//   This element becomes the new topmost element.
// You are also given an integer k, which denotes the total number of moves to be made.
// Return the maximum value of the topmost element of the pile possible after exactly k moves.
// In case it is not possible to obtain a non-empty pile after k moves, return -1.
// Example 1:
//   Input: nums = [5,2,2,4,0,6], k = 4
//   Output: 5
//   Explanation:
//     One of the ways we can end with 5 at the top of the pile after 4 moves is as follows:
//       - Step 1: Remove the topmost element = 5. The pile becomes [2,2,4,0,6].
//       - Step 2: Remove the topmost element = 2. The pile becomes [2,4,0,6].
//       - Step 3: Remove the topmost element = 2. The pile becomes [4,0,6].
//       - Step 4: Add 5 back onto the pile. The pile becomes [5,4,0,6].
//       Note that this is not the only way to end with 5 at the top of the pile.
//       It can be shown that 5 is the largest answer possible after 4 moves.
// Example 2:
//   Input: nums = [2], k = 1
//   Output: -1
//   Explanation:
//     In the first move, our only option is to pop the topmost element of the pile.
//     Since it is not possible to obtain a non-empty pile after one move, we return -1.
// Constraints:
//   1 <= nums.length <= 10⁵
//   0 <= nums[i], k <= 10⁹

mod _maximize_the_topmost_element_after_k_moves {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i32,
    }

    impl Solution {
        pub fn maximum_top(nums: Vec<i32>, k: i32) -> i32 {
            if nums.len() == 1 && k%2==1 {
                return -1;
            }
            let mut max = -1;
            let mut i = 0;
            let mut k = k;
            while i<nums.len() && k>=0 {
                if k != 1 {
                    max = max.max(nums[i]);
                }
                i += 1;
                k -= 1;
            }
            return max;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![5,2,2,4,0,6],
                k: 4,
                ans: 5,
            },
            Solution {
                nums: vec![2],
                k: 1,
                ans: -1,
            },
            Solution {
                nums: vec![3],
                k: 3,
                ans: -1,
            },
            Solution {
                nums: vec![35,43,23,86,23,45,84,2,18,83,79,28,54,81,12,94,14,0,0,29,94,12,13,1,48,85,22,95,24,5,73,10,96,97,72,41,52,1,91,3,20,22,41,98,70,20,52,48,91,84,16,30,27,35,69,33,67,18,4,53,86,78,26,83,13,96,29,15,34,80,16,49],
                k: 15,
                ans: 94,
            },
            Solution {
                nums: vec![73,63,62,16,95,92,93,52,89,36,75,79,67,60,42,93,93,74,94,73,35,86,96],
                k: 59,
                ans: 96,
            }
        ];
        for i in testcases {
            let ans = Solution::maximum_top(i.nums, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
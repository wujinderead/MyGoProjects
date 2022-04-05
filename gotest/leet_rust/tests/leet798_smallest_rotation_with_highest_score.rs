// https://leetcode.com/problems/smallest-rotation-with-highest-score/

// You are given an array nums. You can rotate it by a non-negative integer k so that the array
// becomes [nums[k], nums[k + 1], ... nums[nums.length - 1], nums[0], nums[1], ..., nums[k-1]].
// Afterward, any entries that are less than or equal to their index are worth one point.
// For example, if we have nums = [2,4,1,3,0], and we rotate by k = 2, it becomes [1,3,0,2,4].
// This is worth 3 points because 1 > 0 [no points], 3 > 1 [no points], 0 <= 2 [one point],
// 2 <= 3 [one point], 4 <= 4 [one point].
// Return the rotation index k that corresponds to the highest score we can achieve if we rotated
// nums by it. If there are multiple answers, return the smallest such index k.
// Example 1:
//   Input: nums = [2,3,1,4,0]
//   Output: 3
//   Explanation: Scores for each k are listed below:
//     k = 0,  nums = [2,3,1,4,0],    score 2
//     k = 1,  nums = [3,1,4,0,2],    score 3
//     k = 2,  nums = [1,4,0,2,3],    score 3
//     k = 3,  nums = [4,0,2,3,1],    score 4
//     k = 4,  nums = [0,2,3,1,4],    score 3
//     So we should choose k = 3, which has the highest score.
// Example 2:
//   Input: nums = [1,3,0,2,4]
//   Output: 0
//   Explanation: nums will always have 3 points no matter how it shifts.
//     So we will choose the smallest k, which is 0.
// Constraints:
//   1 <= nums.length <= 10⁵
//   0 <= nums[i] < nums.length

mod _smallest_rotation_with_highest_score {
    struct Solution{
        nums: Vec<i32>,
        ans: i32,
    }

    use std::collections::HashMap;
    impl Solution {
        // the score for k=0 does not matter, we concern how the score will change at certain k.
        // e.g. nums=[1,3,0,2,4], let's say we rotate to left,
        // for nums[0]=1, it has no score (but it doesn't matter), it got score when k=1, lose when k=5 (also 0)s
        // for nums[1]=3, it got score when k=2, lose when k=4
        // for nums[2]=0, ignore, it always has score
        // for nums[3]=2, it lose score when k=2, got when k=4
        // for nums[4]=4, it lose score when k=1, got when k=5 (also 0)
        // so we can see that, when k=1, there is 1 got and 1 lose. so the total score doesn't change.
        // it applies to other k.
        pub fn best_rotation(nums: Vec<i32>) -> i32 {
            let n = nums.len() as i32;
            let mut max_k = 0;
            let mut max_diff = 0;
            let mut lose = HashMap::new();
            // for nums[i]=v, it got score when k=i+1; it lose score when k=(i-v+1)%n
            for (i, &v) in nums.iter().enumerate() {
                let lose_k = (i as i32 - v + 1 + n) % n; // for nums[i]=v, it lose score when k=lose_k
                *lose.entry(lose_k).or_insert(0) += 1;
            }
            let mut diff = 0;
            for k in 1..n {
                // for each k, we always got 1 score, we just check how much score we lose for this k
                diff = diff + 1 - *lose.get(&k).unwrap_or(&0);
                if diff > max_diff {
                    max_diff = diff;
                    max_k = k;
                }
            }
            return max_k;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![2,3,1,4,0],
                ans: 3,
            },
            Solution {
                nums: vec![1,3,0,2,4],
                ans: 0,
            },
            Solution {
                nums: vec![6,2,8,3,5,2,4,3,7,6],
                ans: 9,
            }
        ];
        for i in testcases {
            let ans = Solution::best_rotation(i.nums);
            println!("{}, {}", ans, i.ans);
        }
    }
}
// https://leetcode.com/problems/largest-combination-with-bitwise-and-greater-than-zero/

// The bitwise AND of an array nums is the bitwise AND of all integers in nums.
// For example, for nums = [1, 5, 3], the bitwise AND is equal to 1 & 5 & 3 = 1.
// Also, for nums = [7], the bitwise AND is 7.
// You are given an array of positive integers candidates. Evaluate the bitwise 
// AND of every combination of numbers of candidates. Each number in candidates
// may only be used once in each combination.
// Return the size of the largest combination of candidates with a bitwise AND greater than 0.
// Example 1:
//   Input: candidates = [16,17,71,62,12,24,14]
//   Output: 4
//   Explanation: The combination [16,17,62,24] has a bitwise AND of 16 & 17 & 62 & 24 = 16 > 0.
//     The size of the combination is 4.
//     It can be shown that no combination with a size greater than 4 has a bitwise AND greater than 0.
//     Note that more than one combination may have the largest size.
//     For example, the combination [62,12,24,14] has a bitwise AND of 62 & 12 & 24 & 14 = 8 > 0.
// Example 2:
//   Input: candidates = [8,8]
//   Output: 2
//   Explanation: The largest combination [8,8] has a bitwise AND of 8 & 8 = 8 > 0.
//     The size of the combination is 2, so we return 2.
// Constraints:
//   1 <= candidates.length <= 10⁵
//   1 <= candidates[i] <= 10⁷

mod _largest_combination_with_bitwise_a_n_d_greater_than_zero {
    struct Solution{
        candidates: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        // count the number of 1's of every bit, return the maximal number
        pub fn largest_combination(candidates: Vec<i32>) -> i32 {
            let mut count = [0; 30];
            for c in candidates {
                for i in 0..30 {
                    if (1<<i) & c > 0 {
                        count[i] += 1;
                    }
                }
            }
            return *count.iter().max().unwrap();
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                candidates: vec![16,17,71,62,12,24,14],
                ans: 4,
            },
            Solution {
                candidates: vec![8,8],
                ans: 2,
            },
        ];
        for i in testcases {
            let ans = Solution::largest_combination(i.candidates);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
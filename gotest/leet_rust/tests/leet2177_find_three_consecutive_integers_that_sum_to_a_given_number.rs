// https://leetcode.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/

// Given an integer num, return three consecutive integers (as a sorted array) that sum to num.
// If num cannot be expressed as the sum of three consecutive integers, return an empty array.
// Example 1:
//   Input: num = 33
//   Output: [10,11,12]
//   Explanation: 33 can be expressed as 10 + 11 + 12 = 33.
//     10, 11, 12 are 3 consecutive integers, so we return [10, 11, 12].
// Example 2:
//   Input: num = 4
//   Output: []
//   Explanation: There is no way to express 4 as the sum of 3 consecutive integers.
// Constraints:
//   0 <= num <= 10¹⁵

mod _find_three_consecutive_integers_that_sum_to_a_given_number {
    struct Solution{
        num: i64,
        ans: Vec<i64>
    }

    impl Solution {
        pub fn sum_of_three(num: i64) -> Vec<i64> {
            if num % 3 != 0 {
                return vec![];
            }
            return vec![num/3-1, num/3, num/3+1];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                num: 33,
                ans: vec![10,11,12],
            },
            Solution {
                num: 4,
                ans: vec![],
            },
        ];
        for i in testcases {
            let ans = Solution::sum_of_three(i.num);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/smallest-value-of-the-rearranged-number/

// You are given an integer num. Rearrange the digits of num such that its value is minimized
// and it does not contain any leading zeros.
// Return the rearranged number with minimal value.
// Note that the sign of the number does not change after rearranging the digits.
// Example 1:
//   Input: num = 310
//   Output: 103
//   Explanation: The possible arrangements for the digits of 310 are 013, 031, 103, 130, 301, 310.
//     The arrangement with the smallest value that does not contain any leading zeros is 103.
// Example 2:
//   Input: num = -7605
//   Output: -7650
//   Explanation: Some possible arrangements for the digits of -7605 are -7650, -6705, -5076, -0567.
//   The arrangement with the smallest value that does not contain any leading zeros is -7650.
// Constraints:
//   -10¹⁵ <= num <= 10¹⁵

mod _smallest_value_of_the_rearranged_number {
    struct Solution{
        nums: i64,
        ans: i64,
    }

    impl Solution {
        pub fn smallest_number(num: i64) -> i64 {
            if num == 0 {
                return 0;
            }
            let mut ds = Vec::<i64>::new();
            let mut x = num.abs();
            while x != 0 {
                ds.push(x%10);
                x = x/10;
            }
            ds.sort();
            return if num > 0 {
                // swap first non-0 to front
                for (i, &v) in ds.iter().enumerate() {
                    if v != 0 {
                        ds.swap(0, i);
                        break;
                    }
                }
                let mut ans = 0;
                for &i in ds.iter() {
                    ans *= 10;
                    ans += i;
                }
                ans
            } else {
                ds.reverse();
                let mut ans = 0;
                for &i in ds.iter() {
                    ans *= 10;
                    ans += i;
                }
                -ans
            }
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: 310,
                ans: 103,
            },
            Solution {
                nums: -102,
                ans: -210,
            },
            Solution {
                nums: 200,
                ans: 200,
            },
            Solution {
                nums: -7605,
                ans: -7650,
            },
            Solution {
                nums: 0,
                ans: 0,
            },
            Solution {
                nums: -1,
                ans: -1,
            },
        ];
        for i in testcases {
            let ans = Solution::smallest_number(i.nums);
            println!("{}, {}", ans, i.ans);
        }
    } 
}

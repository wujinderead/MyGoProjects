// https://leetcode.com/problems/longest-binary-subsequence-less-than-or-equal-to-k/

// You are given a binary string s and a positive integer k.
// Return the length of the longest subsequence of s that makes up a binary
// number less than or equal to k.
// Note:
//   The subsequence can contain leading zeroes.
//   The empty string is considered to be equal to 0.
//   A subsequence is a string that can be derived from another string by
//     deleting some or no characters without changing the order of the remaining characters.
// Example 1:
//   Input: s = "1001010", k = 5
//   Output: 5
//   Explanation: The longest subsequence of s that makes up a binary number less
//     than or equal to 5 is "00010", as this number is equal to 2 in decimal.
//     Note that "00100" and "00101" are also possible, which are equal to 4 and 5
//     in decimal, respectively.
//     The length of this subsequence is 5, so 5 is returned.
// Example 2:
//   Input: s = "00101001", k = 1
//   Output: 6
//   Explanation: "000001" is the longest subsequence of s that makes up a binary
//     number less than or equal to 1, as this number is equal to 1 in decimal.
//     The length of this subsequence is 6, so 6 is returned.
// Constraints:
//   1 <= s.length <= 1000
//   s[i] is either '0' or '1'.
//   1 <= k <= 10⁹

mod _longest_binary_subsequence_less_than_or_equal_to_k {
    struct Solution{
        s: String,
        k: i32,
        ans: i32,
    }

    impl Solution {
        pub fn longest_subsequence(s: String, k: i32) -> i32 {
            let mut zeros = 0;
            for i in 0..s.len() {
                if s.as_bytes()[i] == b'0' {
                    zeros += 1;
                } else if i+30 >= s.len() && i32::from_str_radix(&s[i..], 2).unwrap() <= k {
                    return zeros + (s.len()-i) as i32;
                }
            }
            return zeros;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                s: "1001010".to_string(),
                k: 5,
                ans: 5,
            },
            Solution {
                s: "00101001".to_string(),
                k: 1,
                ans: 6,
            },
            Solution {
                s: "00000".to_string(),
                k: 1,
                ans: 5,
            },
            Solution {
                s: "00010".to_string(),
                k: 1,
                ans: 4,
            },
        ];
        for i in testcases {
            let ans = Solution::longest_subsequence(i.s, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
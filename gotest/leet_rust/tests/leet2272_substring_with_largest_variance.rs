// https://leetcode.com/problems/substring-with-largest-variance/

// The variance of a string is defined as the largest difference between the number of occurrences
// of any 2 characters present in the string. Note the two characters may or may not be the same.
// Given a string s consisting of lowercase English letters only, return the largest variance
// possible among all substrings of s.
// A substring is a contiguous sequence of characters within a string.
// Example 1:
//   Input: s = "aababbb"
//   Output: 3
//   Explanation:
//     All possible variances along with their respective substrings are listed below:
//     - Variance 0 for substrings "a", "aa", "ab", "abab", "aababb", "ba", "b", "bb", and "bbb".
//     - Variance 1 for substrings "aab", "aba", "abb", "aabab", "ababb", "aababbb", and "bab".
//     - Variance 2 for substrings "aaba", "ababbb", "abbb", and "babb".
//     - Variance 3 for substring "babbb".
//     Since the largest possible variance is 3, we return it.
// Example 2:
//   Input: s = "abcde"
//   Output: 0
//   Explanation:
//     No letter occurs more than once in s, so the variance of every substring is 0.
// Constraints:
//   1 <= s.length <= 10⁴
//   s consists of lowercase English letters.

mod _substring_with_largest_variance {
    struct Solution{
        s: String,
        ans: i32,
    }

    // Kadane algorithm
    impl Solution {
        pub fn largest_variance(s: String) -> i32 {
            let ss = s.as_bytes();
            let mut ans = 0;
            let mut count = [0; 26];
            for &c in ss {
                count[(c-b'a') as usize] += 1;
            }
            // for each char pair
            for i in 0..26 {
                for j in 0..26 {
                    if i == j || count[i] == 0 || count[j] == 0 {
                        continue;
                    }
                    let (a, b) = (b'a'+ i as u8, b'a' + j as u8);
                    let (mut remain_a, mut cur_a, mut cur_b) = (count[i], 0, 0);
                    for &c in ss {
                        if c != a && c != b {
                            continue
                        }
                        if c == b {
                            cur_b += 1;
                        } else {
                            cur_a += 1;
                            remain_a -= 1;
                        };
                        if cur_a > 0 {
                            println!("{} {}", cur_b, cur_a);
                            ans = ans.max(cur_b-cur_a);
                        }
                        if cur_b < cur_a && remain_a > 0 {
                            cur_a = 0;
                            cur_b = 0;
                        }
                    }
                }
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            /*Solution {
                s: "aababbb".to_string(),
                ans: 3,
            },
            Solution {
                s: "abcde".to_string(),
                ans: 0,
            },
            Solution {
                s: "lripaa".to_string(),
                ans: 1,
            },*/
            Solution {
                s: "abb".to_string(),
                ans: 1,
            }
        ];
        for i in testcases {
            let ans = Solution::largest_variance(i.s);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
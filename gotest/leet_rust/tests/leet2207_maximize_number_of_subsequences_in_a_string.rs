// https://leetcode.com/problems/maximize-number-of-subsequences-in-a-string/

// You are given a 0-indexed string text and another 0-indexed string pattern of length 2,
// both of which consist of only lowercase English letters.
// You can add either pattern[0] or pattern[1] anywhere in text exactly once. 
// Note that the character can be added even at the beginning or at the end of text.
// Return the maximum number of times pattern can occur as a subsequence of the modified text.
// A subsequence is a string that can be derived from another string by deleting some or
// no characters without changing the order of the remaining characters.
// Example 1: 
//   Input: text = "abdcdbc", pattern = "ac"
//   Output: 4
//   Explanation:
//     If we add pattern[0] = 'a' in between text[1] and text[2], we get "abadcdbc".
//     Now, the number of times "ac" occurs as a subsequence is 4.
//     Some other strings which have 4 subsequences "ac" after adding a character to
//     text are "aabdcdbc" and "abdacdbc".
//     However, strings such as "abdcadbc", "abdccdbc", and "abdcdbcc", although
//     obtainable, have only 3 subsequences "ac" and are thus suboptimal.
//     It can be shown that it is not possible to get more than 4 subsequences "ac"
//     by adding only one character.
// Example 2:
//   Input: text = "aabb", pattern = "ab"
//   Output: 6
//   Explanation:
//     Some of the strings which can be obtained from text and have 6 subsequences
//     "ab" are "aaabb", "aaabb", and "aabbb".
// Constraints:
//   1 <= text.length <= 10⁵
//   pattern.length == 2
//   text and pattern consist only of lowercase English letters.

mod _maximize_number_of_subsequences_in_a_string {
    struct Solution{
        text: String,
        pattern: String,
        ans: i32,
    }

    impl Solution {
        pub fn maximum_subsequence_count(text: String, pattern: String) -> i64 {
            let (text, pattern) = (text.into_bytes(), pattern.as_bytes());
            let (mut sum, mut p0, mut p1) = (0, 0, 0);
            for c in text {
                if c == pattern[1] {
                    sum += p0;
                    p1 += 1;
                }
                if c == pattern[0] {  // double if to handler pattern[0]==pattern[1]
                    p0 += 1;
                }
            }
            return sum + p0.max(p1);
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                text: "abdcdbc".to_string(),
                pattern: "ac".to_string(),
                ans: 4,
            },
            Solution {
                text: "aabb".to_string(),
                pattern: "ab".to_string(),
                ans: 6,
            },
            Solution {
                text: "aaccb".to_string(),
                pattern: "bb".to_string(),
                ans: 1,
            },
            Solution {
                text: "aacc".to_string(),
                pattern: "bb".to_string(),
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::maximum_subsequence_count(i.text, i.pattern);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
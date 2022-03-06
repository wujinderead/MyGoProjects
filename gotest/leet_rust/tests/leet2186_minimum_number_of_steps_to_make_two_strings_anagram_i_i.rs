// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii/

// You are given two strings s and t. In one step, you can append any character to either s or t.
// Return the minimum number of steps to make s and t anagrams of each other.
// An anagram of a string is a string that contains the same characters with a different
// (or the same) ordering.
// Example 1:
//   Input: s = "leetcode", t = "coats"
//   Output: 7
//   Explanation:
//     - In 2 steps, we can append the letters in "as" onto s = "leetcode", forming s = "leetcodeas".
//     - In 5 steps, we can append the letters in "leede" onto t = "coats", forming t = "coatsleede".
//     "leetcodeas" and "coatsleede" are now anagrams of each other.
//     We used a total of 2 + 5 = 7 steps.
//     It can be shown that there is no way to make them anagrams of each other with less than 7 steps.
// Example 2:
//   Input: s = "night", t = "thing"
//   Output: 0
//   Explanation: The given strings are already anagrams of each other.
//     Thus, we do not need any further steps.
// Constraints:
//   1 <= s.length, t.length <= 2 * 10⁵
//   s and t consist of lowercase English letters.

mod _minimum_number_of_steps_to_make_two_strings_anagram_i_i {
    struct Solution{
        s: String,
        t: String,
        ans: i32,
    }

    impl Solution {
        pub fn min_steps(s: String, t: String) -> i32 {
            let (mut ss, mut tt) = ([0 as i32; 26], [0 as i32; 26]);
            for b in s.into_bytes() {
                ss[(b-b'a') as usize] += 1;
            }
            for b in t.into_bytes() {
                tt[(b-b'a') as usize] += 1;
            }
            let (mut sum, mut i) = (0, 0);
            while i<26 {
                sum += (ss[i]-tt[i]).abs();
                i += 1;
            }
            return sum;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                s: "leetcode".to_string(),
                t: "coats".to_string(),
                ans: 7,
            },
            Solution {
                s: "night".to_string(),
                t: "thing".to_string(),
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::min_steps(i.s, i.t);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
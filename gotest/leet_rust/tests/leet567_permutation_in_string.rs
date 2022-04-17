// https://leetcode.com/problems/permutation-in-string/

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.
// Example 1:
//   Input: s1 = "ab", s2 = "eidbaooo"
//   Output: true
//   Explanation: s2 contains one permutation of s1 ("ba").
// Example 2:
//   Input: s1 = "ab", s2 = "eidboaoo"
//   Output: false
// Constraints:
//   1 <= s1.length, s2.length <= 10⁴
//   s1 and s2 consist of lowercase English letters.

mod _permutation_in_string {
    struct Solution{
        s1: String,
        s2: String,
        ans: bool,
    }

    // count occurrence of characters in s1, check if  substring of s2 has same occurrence
    impl Solution {
        pub fn check_inclusion(s1: String, s2: String) -> bool {
            if s2.len() < s1.len() {
                return false;
            }
            let s1 = s1.as_bytes();
            let s2 = s2.as_bytes();
            // count occurrence of characters in s1
            let mut o1 = [0; 26];
            for c in s1 {
                o1[(*c-b'a') as usize] += 1;
            }
            // check s2
            let mut o2 = [0; 26];
            for i in 0..s1.len() {
                o2[(s2[i]-b'a') as usize] += 1;
            }
            if o1==o2 {
                return true;
            }
            // sliding window
            for i in s1.len()..s2.len() {
                o2[(s2[i]-b'a') as usize] += 1;
                o2[(s2[i-s1.len()]-b'a') as usize] -= 1;
                if o1 == o2 {
                    return true;
                }
            }
            return false;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                s1: "ab".to_string(),
                s2: "eidbaooo".to_string(),
                ans: true,
            },
            Solution {
                s1: "ab".to_string(),
                s2: "eidboaoo".to_string(),
                ans: false,
            },
        ];
        for i in testcases {
            let ans = Solution::check_inclusion(i.s1, i.s2);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
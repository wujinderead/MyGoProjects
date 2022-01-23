// https://leetcode.com/problems/count-words-obtained-after-adding-a-letter/

// You are given two 0-indexed arrays of strings startWords and targetWords.
// Each string consists of lowercase English letters only.
// For each string in targetWords, check if it is possible to choose a string from
// startWords and perform a conversion operation on it to be equal to that from targetWords.
// The conversion operation is described in the following two steps:
// Append any lowercase letter that is not present in the string to its end.
// For example, if the string is "abc", the letters 'd', 'e', or 'y' can be 
// added to it, but not 'a'. If 'd' is added, the resulting string will be "abcd".
// Rearrange the letters of the new string in any arbitrary order.
// For example, "abcd" can be rearranged to "acbd", "bacd", "cbda", and so on. 
// Note that it can also be rearranged to "abcd" itself.
// Return the number of strings in targetWords that can be obtained by 
// performing the operations on any string of startWords.
// Note that you will only be verifying if the string in targetWords can be 
// obtained from a string in startWords by performing the operations. The strings in
// startWords do not actually change during this process.
// Example 1:
//   Input: startWords = ["ant","act","tack"], targetWords = ["tack","act","acti"]
//   Output: 2
//   Explanation:
//     - In order to form targetWords[0] = "tack", we use startWords[1] = "act",
//       append 'k' to it, and rearrange "actk" to "tack".
//     - There is no string in startWords that can be used to obtain targetWords[1] = "act".
//       Note that "act" does exist in startWords, but we must append one letter to
//       the string before rearranging it.
//     - In order to form targetWords[2] = "acti", we use startWords[1] = "act",
//       append 'i' to it, and rearrange "acti" to "acti" itself.
// Example 2:
//   Input: startWords = ["ab","a"], targetWords = ["abc","abcd"]
//   Output: 1
//   Explanation:
//     - In order to form targetWords[0] = "abc", we use startWords[0] = "ab", add
//       'c' to it, and rearrange it to "abc".
//     - There is no string in startWords that can be used to obtain targetWords[1] = "abcd".
// Constraints:
//   1 <= startWords.length, targetWords.length <= 5 * 10⁴
//   1 <= startWords[i].length, targetWords[j].length <= 26
//   Each string of startWords and targetWords consists of lowercase English letters only.
//   No letter occurs more than once in any string of startWords or targetWords.

mod _count_words_obtained_after_adding_a_letter {
    use std::collections::HashSet;

    struct Solution{
        start_words: Vec<&'static str>,
        target_words: Vec<&'static str>,
        ans: i32,
    }

    impl Solution {
        pub fn word_count(start_words: Vec<String>, target_words: Vec<String>) -> i32 {
            let mut sset = HashSet::new();
            let mut count = 0;
            for i in 0..start_words.len() {
                let mask = string_to_mask(&start_words[i]);
                sset.insert(mask);
            }
            'outer: for i in 0..target_words.len() {
                let mask = string_to_mask(&target_words[i]);
                for j in 0..27 {
                    if mask & (1<<j) > 0 {
                        let smark = mask ^ (1<<j);
                        if sset.contains(&smark) {
                            count += 1;
                            continue 'outer;
                        }
                    }
                }
            }
            return count;
        }
    }

    fn string_to_mask(s: &String) -> i32 {
        let chs = s.as_bytes();
        let mut mask = 0;
        for j in 0..chs.len() {
            mask = mask | 1<<(chs[j]-b'a');
        }
        return mask;
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                start_words: vec!["ant","act","tack"],
                target_words: vec!["tack","act","acti"],
                ans: 2,
            },
            Solution {
                start_words: vec!["ab","a"],
                target_words: vec!["abc","abcd"],
                ans: 1,
            },
            Solution {
                start_words: vec!["g","vf","ylpuk","nyf","gdj","j","fyqzg","sizec"],
                target_words: vec!["r","am","jg","umhjo","fov","lujy","b","uz","y"],
                ans: 2,
            },
        ];
        for i in testcases {
            let sw= i.start_words.iter().map(|s| s.to_string()).collect();
            let tw= i.target_words.iter().map(|s| s.to_string()).collect();
            let ans = Solution::word_count(sw, tw);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
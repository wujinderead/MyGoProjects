// https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/

// You are given an array of strings words. Each element of words consists of
// two lowercase English letters.
// Create the longest possible palindrome by selecting some elements from words 
// and concatenating them in any order. Each element can be selected at most once.
// Return the length of the longest palindrome that you can create. If it is 
// impossible to create any palindrome, return 0.
// A palindrome is a string that reads the same forward and backward.
// Example 1:
//   Input: words = ["lc","cl","gg"]
//   Output: 6
//   Explanation: One longest palindrome is "lc" + "gg" + "cl" = "lcggcl", of length 6.
//     Note that "clgglc" is another longest palindrome that can be created.
// Example 2:
//   Input: words = ["ab","ty","yt","lc","cl","ab"]
//   Output: 8
//   Explanation: One longest palindrome is "ty" + "lc" + "cl" + "yt" = "tylcclyt", of length 8.
//     Note that "lcyttycl" is another longest palindrome that can be created.
// Example 3:
//   Input: words = ["cc","ll","xx"]
//   Output: 2
//   Explanation: One longest palindrome is "cc", of length 2.
//     Note that "ll" is another longest palindrome that can be created, and so is "xx".
// Constraints:
//   1 <= words.length <= 10⁵
//   words[i].length == 2
//   words[i] consists of lowercase English letters.

mod _longest_palindrome_by_concatenating_two_letter_words {
    use std::collections::HashMap;

    struct Solution{
        words: Vec<String>,
        ans: i32,
    }

    impl Solution {
        pub fn longest_palindrome(words: Vec<String>) -> i32 {
            let mut map = HashMap::new();
            // count words
            for word in words.iter() {
                let count = map.entry(word).or_insert(0);
                *count += 1;
            }
            let mut count = 0;
            let mut extra_pair = 0;
            let mut chars = vec![0; 2];
            for (key, value) in &map {
                let chs = key.as_bytes();
                if chs[0] == chs[1] {  // "aa"
                    count += 4*(value/2);
                    if value%2 == 1 {
                        extra_pair = 2;
                    }
                } else if chs[0] < chs[1] {  // "ab"
                    chars[0] = chs[1];
                    chars[1] = chs[0];
                    let newkey = String::from_utf8(chars.to_vec()).unwrap();
                    if let Some(&v) = map.get(&newkey) {
                        count += 4*v.min(*value);
                    }
                }
            }
            return count + extra_pair
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                words: vec!["lc".to_string(),"cl".to_string(),"gg".to_string()],
                ans: 6,
            },
            Solution {
                words: vec!["ab".to_string(),"ty".to_string(),"yt".to_string(),"lc".to_string(),"cl".to_string(),"ab".to_string()],
                ans: 8,
            },
            Solution {
                words: vec!["aa".to_string(),"xx".to_string(),"bb".to_string()],
                ans: 2,
            },
            Solution {
                words: vec!["aa".to_string(),"xx".to_string(),"bb".to_string(),"xx".to_string()],
                ans: 6,
            },
            Solution {
                words: vec!["aa".to_string(),"xx".to_string(),"bb".to_string(),"xx".to_string(),"xx".to_string()],
                ans: 6,
            },
        ];
        for i in testcases {
            let ans = Solution::longest_palindrome(i.words);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/swap-for-longest-repeated-character-substring/

// You are given a string text. You can swap two of the characters in the text.
// Return the length of the longest substring with repeated characters.
// Example 1:
//   Input: text = "ababa"
//   Output: 3
//   Explanation: We can swap the first 'b' with the last 'a', or the last 'b' with the first 'a'.
//     Then, the longest repeated character substring is "aaa" with length 3.
// Example 2:
//   Input: text = "aaabaaa"
//   Output: 6
//   Explanation: Swap 'b' with the last 'a' (or the first 'a'), and we get
//     longest repeated character substring "aaaaaa" with length 6.
// Example 3:
//   Input: text = "aaaaa"
//   Output: 5
//   Explanation: No need to swap, longest repeated character substring is "aaaaa"
//     with length is 5.
// Constraints:
//   1 <= text.length <= 2 * 10⁴
//   text consist of lowercase English characters only.

mod _swap_for_longest_repeated_character_substring {
    struct Solution{
        text: String,
        ans: i32,
    }

    impl Solution {
        pub fn max_rep_opt1(text: String) -> i32 {
            // first group same letters
            let text = text.as_bytes();
            let mut cur = text[0];
            let mut gs = vec![1];  // group size
            let mut gc = vec![text[0]];  // group char
            let mut max = 1;
            let mut cc = [0; 26];  // char count
            cc[(text[0]-b'a') as usize] += 1;
            for i in 1..text.len() {
                if text[i] == cur {
                    *gs.last_mut().unwrap() += 1;
                } else {
                    gs.push(1);
                    gc.push(text[i]);
                    cur = text[i];
                }
                cc[(text[i]-b'a') as usize] += 1;
            }
            // check each group
            for i in 0..gs.len() {
                if gs[i] != cc[(gc[i]-b'a') as usize] { // check if this group contains all gc[i]
                    max = max.max(gs[i]+1);
                } else {
                    max = max.max(gs[i]);
                }
                if gs[i]==1 && i>0 && i<gs.len()-1 && gc[i-1]==gc[i+1] {  // current group size 1, and adjacent groups have same character
                    if gs[i-1]+gs[i+1] != cc[(gc[i-1]-b'a') as usize] {  // if other group has same character
                        max = max.max(gs[i-1]+gs[i+1]+1);
                    } else {
                        max = max.max(gs[i-1]+gs[i+1]);
                    }
                }
            }
            return max;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                text: "ababa".to_string(),
                ans: 3,
            },
            Solution {
                text: "aaabaaa".to_string(),
                ans: 6,
            },
            Solution {
                text: "aaaaa".to_string(),
                ans: 5,
            },
        ];
        for i in testcases {
            let ans = Solution::max_rep_opt1(i.text);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
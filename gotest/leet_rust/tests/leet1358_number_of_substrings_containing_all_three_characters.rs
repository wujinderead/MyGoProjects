// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/

// Given a string s consisting only of characters a, b and c.
// Return the number of substrings containing at least one occurrence of all these characters
// a, b and c.
// Example 1:
//   Input: s = "abcabc"
//   Output: 10
//   Explanation: The substrings containing at least one occurrence of the
//     characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc",
//     "cab", "cabc" and "abc" (again).
// Example 2:
//   Input: s = "aaacb"
//   Output: 3
//   Explanation: The substrings containing at least one occurrence of the
//     characters a, b and c are "aaacb", "aacb" and "acb".
// Example 3:
//   Input: s = "abc"
//   Output: 1
// Constraints:
//   3 <= s.length <= 5 x 10^4
//   s only consists of a, b or c characters.

mod _number_of_substrings_containing_all_three_characters {
    struct Solution{
        s: String,
        ans: i32,
    }

    impl Solution {
        pub fn number_of_substrings(s: String) -> i32 {
            let (mut a, mut b, mut c) = (Option::<usize>::None, Option::<usize>::None, Option::<usize>::None);
            let s = s.as_bytes();
            let  mut count = 0;
            for i in 0..s.len() {
                if s[i] == b'a' {
                    if b.is_some() && c.is_some() {
                        count += b.unwrap().min(c.unwrap())+1;
                    }
                    a = Some(i);
                }
                if s[i] == b'b' {
                    if a.is_some() && c.is_some() {
                        count += a.unwrap().min(c.unwrap())+1;
                    }
                    b = Some(i);
                }
                if s[i] == b'c' {
                    if a.is_some() && b.is_some() {
                        count += a.unwrap().min(b.unwrap())+1;
                    }
                    c = Some(i);
                }
            }
            return count as i32;
        }

        pub fn number_of_substrings_two_pointer(s: String) -> i32 {
            let (mut i, mut j, mut f, mut count) = (0, 0, [0,0,0], 0);
            let s= s.as_bytes();
            while i<s.len() {
                f[(s[i]-b'a') as usize] += 1;
                while f[0]>0 && f[1]>0 && f[2]>0 {
                    f[(s[j]-b'a') as usize] -= 1;
                    j += 1;
                }
                count += j as i32;
                i += 1;
            }
            return count;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                s: "abcabc".to_string(),
                ans: 10,
            },
            Solution {
                s: "aaabc".to_string(),
                ans: 3,
            },
            Solution {
                s: "abc".to_string(),
                ans: 1,
            },
            Solution {
                s: "abbac".to_string(),
                ans: 3,
            },
            Solution {
                s: "abba".to_string(),
                ans: 0,
            }
        ];
        for i in &testcases {
            let ans = Solution::number_of_substrings(i.s.clone());
            println!("{}, {}", ans, i.ans);
        }
        for i in &testcases {
            let ans = Solution::number_of_substrings_two_pointer(i.s.clone());
            println!("{}, {}", ans, i.ans);
        }
    } 
}
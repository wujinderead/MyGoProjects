// https://leetcode.com/problems/construct-string-with-repeat-limit/

// You are given a string s and an integer repeatLimit. Construct a new string repeatLimitedString
// using the characters of s such that no letter appears more than repeatLimit times in a row.
// You do not have to use all characters from s.
// Return the lexicographically largest repeatLimitedString possible.
// A string a is lexicographically larger than a string b if in the first position where a and b
// differ, string a has a letter that appears later in the alphabet than the corresponding letter
// in b. If the first min(a.length, b.length) characters do not differ, then the longer string is
// the lexicographically larger one.
// Example 1:
//   Input: s = "cczazcc", repeatLimit = 3
//   Output: "zzcccac"
//   Explanation: We use all of the characters from s to construct the repeatLimitedString "zzcccac".
//     The letter 'a' appears at most 1 time in a row.
//     The letter 'c' appears at most 3 times in a row.
//     The letter 'z' appears at most 2 times in a row.
//     Hence, no letter appears more than repeatLimit times in a row and the string is a valid repeatLimitedString.
//     The string is the lexicographically largest repeatLimitedString possible so we return "zzcccac".
//     Note that the string "zzcccca" is lexicographically larger but the letter 'c'
//     appears more than 3 times in a row, so it is not a valid repeatLimitedString.
// Example 2:
//   Input: s = "aababab", repeatLimit = 2
//   Output: "bbabaa"
//   Explanation: We use only some of the characters from s to construct the repeatLimitedString "bbabaa".
//    The letter 'a' appears at most 2 times in a row.
//    The letter 'b' appears at most 2 times in a row.
//    Hence, no letter appears more than repeatLimit times in a row and the string is a valid repeatLimitedString.
//    The string is the lexicographically largest repeatLimitedString possible so we return "bbabaa".
//    Note that the string "bbabaaa" is lexicographically larger but the letter 'a'
//    appears more than 2 times in a row, so it is not a valid repeatLimitedString.
// Constraints:
//   1 <= repeatLimit <= s.length <= 10⁵
//   s consists of lowercase English letters.

mod _construct_string_with_repeat_limit {
    struct Solution{
        s: String,
        repeat_limit: i32,
        ans: String,
    }

    impl Solution {
        pub fn repeat_limited_string(s: String, repeat_limit: i32) -> String {
            let s_len = s.len();
            let mut chars = s.into_bytes();
            // sort chars in descending order
            chars.sort_by_key(|&b| u8::MAX-b);
            chars.push(0);

            let mut i = 0;
            let mut buf = Vec::<u8>::with_capacity(s_len);

            // while we have letter
            'outer: while i < s_len {
                let mut j = i;
                while chars[i] == chars[j] {
                    j += 1;
                }
                let mut jj = j;   // [i..jj] is same largest char, jj is next smaller char
                let mut same = 0;
                while i < j {  // for same chars
                    while same < repeat_limit && i < j {
                        buf.push(chars[i]);
                        same += 1;
                        i += 1;
                    }
                    if i == j {  // i == j, same char over
                        i = jj;
                        break;   // start from new char
                    } else {     // need append separator
                        same = 0;
                        if jj == s_len {
                            break 'outer;  // can't append a separator, return
                        }
                        buf.push(chars[jj]);
                        jj += 1;
                    }
                }
            }
            return unsafe { String::from_utf8_unchecked(buf) };
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                s: String::from("cczazcc"),
                repeat_limit: 3,
                ans: String::from("zzcccac"),
            },
            Solution {
                s: String::from("aababab"),
                repeat_limit: 2,
                ans: String::from("bbabaa"),
            },
            Solution {
                s: String::from("zzzzzcccbbbba"),
                repeat_limit: 2,
                ans: String::from("zzczzczcbbabb"),
            },
            Solution {
                s: String::from("robnsdvpuxbapuqgopqvxdrchivlifeepy"),
                repeat_limit: 2,
                ans: String::from("yxxvvuvusrrqqppopponliihgfeeddcbba"),
            },
        ];
        for i in testcases {
            let ans = Solution::repeat_limited_string(i.s, i.repeat_limit);
            println!("{}, {}", ans, i.ans);
        }
    }
}
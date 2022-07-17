// https://leetcode.com/problems/move-pieces-to-obtain-a-string/

// You are given two strings start and target, both of length n. Each string consists only of the
// characters 'L', 'R', and '_' where:
// The characters 'L' and 'R' represent pieces, where a piece 'L' can move to the left only if there
// is a blank space directly to its left, and a piece 'R' can move to the right only if there is a
// blank space directly to its right. The character '_' represents a blank space that can be
// occupied by any of the 'L' or 'R' pieces.
// Return true if it is possible to obtain the string target by moving the pieces of the string
// start any number of times. Otherwise, return false.
// Example 1:
//   Input: start = "_L__R__R_", target = "L______RR"
//   Output: true
//   Explanation: We can obtain the string target from start by doing the
//   following moves:
//     - Move the first piece one step to the left, start becomes equal to "L___R__R_".
//     - Move the last piece one step to the right, start becomes equal to "L___R___R".
//     - Move the second piece three steps to the right, start becomes equal to "L______RR".
//     Since it is possible to get the string target from start, we return true.
// Example 2:
//   Input: start = "R_L_", target = "__LR"
//   Output: false
//   Explanation: The 'R' piece in the string start can move one step to the right to obtain "_RL_".
//     After that, no pieces can move anymore, so it is impossible to obtain the string target from start.
// Example 3:
//   Input: start = "_R", target = "R_"
//   Output: false
//   Explanation: The piece in the string start can move only to the right, so it
//     is impossible to obtain the string target from start.
// Constraints:
//   n == start.length == target.length
//   1 <= n <= 10⁵
//   start and target consist of the characters 'L', 'R', and '_'.

mod _move_pieces_to_obtain_a_string {
    struct Solution{
        start: String,
        target: String,
        ans: bool,
    }

    impl Solution {
        pub fn can_change(start: String, target: String) -> bool {
            let (mut i, mut j) = (0, 0);
            let (s, t) = (start.as_bytes(), target.as_bytes());
            while i<s.len() {
                while i<s.len() && s[i] == b'_' {
                    i += 1;
                }
                if i == s.len() {  // no LR in start, break
                    break;
                }
                // s[i] = L or R, search target
                while j<s.len() && t[j] == b'_' {
                    j += 1;
                }
                if j==s.len() || s[i] != t[j] || (s[i] == b'R' && i>j) || (s[i] == b'L' && i<j) {
                    return false;
                }
                i += 1;
                j += 1;
            }
            while j<s.len() {  // find if there is extra LR in target
                if t[j] != b'_' {
                    return false;
                }
                j += 1;
            }
            return true;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                start: "_L__R__R_".to_string(),
                target: "L______RR".to_string(),
                ans: true,
            },
            Solution {
                start: "R_L_".to_string(),
                target: "__LR".to_string(),
                ans: false,
            },
            Solution {
                start: "_R".to_string(),
                target: "R_".to_string(),
                ans: false,
            },
        ];
        for i in testcases {
            let ans = Solution::can_change(i.start, i.target);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
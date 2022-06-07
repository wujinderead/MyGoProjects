// https://leetcode.com/problems/score-of-parentheses/

// Given a balanced parentheses string s, return the score of the string.
// The score of a balanced parentheses string is based on the following rule:
//   "()" has score 1.
//   AB has score A + B, where A and B are balanced parentheses strings.
//   (A) has score 2 * A, where A is a balanced parentheses string.
// Example 1:
//   Input: s = "()"
//   Output: 1
// Example 2:
//   Input: s = "(())"
//   Output: 2
// Example 3:
//   Input: s = "()()"
//   Output: 2
// Constraints:
//   2 <= s.length <= 50
//   s consists of only '(' and ')'.
//   s is a balanced parentheses string.

mod _score_of_parentheses {
    struct Solution{
        s: String,
        ans: i32,
    }

    // another solution: https://leetcode.com/problems/score-of-parentheses/discuss/1856699/C%2B%2B-BEATS-100-OMG!!!-(-%22-)-O(1)-Space-Explained
    // just check the depth of the inner '()' and sum them.
    impl Solution {
        pub fn score_of_parentheses(s: String) -> i32 {
            let mut i_stack = Vec::new();
            let mut v_stack = Vec::new();
            let s = s.as_bytes();
            for (i, &c) in s.iter().enumerate() {
                if c==b'(' {
                    i_stack.push(i);
                    continue;
                }
                // the char is ')'
                let lp = i_stack.pop().unwrap();
                if lp == i-1 {  // a single '()'
                    v_stack.push((lp, 1));
                    continue;
                }
                // else '(xxx)', pop those xxx from stack and sum it, then push the sum*2
                let mut sum = 0;
                while v_stack.len() > 0 && v_stack.last().unwrap().0 > lp {
                    sum += v_stack.pop().unwrap().1;
                }
                v_stack.push((lp, sum*2));
            }
            return v_stack.iter().fold(0, |acc, &x| acc + x.1);
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                s: "()".to_string(),
                ans: 1,
            },
            Solution {
                s: "(())".to_string(),
                ans: 2,
            },
            Solution {
                s: "()()".to_string(),
                ans: 2,
            },
            Solution {
                s: "(()())".to_string(),
                ans: 4,
            },
            Solution {
                s: "(()())()".to_string(),
                ans: 5,
            },
            Solution {
                s: "((()())())".to_string(),
                ans: 10,
            },
            Solution {
                s: "(((())))".to_string(),
                ans: 8,
            },
            Solution {
                s: "(((())))(()())".to_string(),
                ans: 12,
            },
            Solution {
                s: "(((()))(()()))".to_string(),
                ans: 16,
            },
            Solution {
                s: "(((()))(()())())()".to_string(),
                ans: 19,
            },
        ];
        for i in testcases {
            let ans = Solution::score_of_parentheses(i.s);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
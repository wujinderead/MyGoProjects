// https://leetcode.com/problems/check-if-there-is-a-valid-parentheses-string-path/

// A parentheses string is a non-empty string consisting only of '(' and ')'. It is valid if any of
// the following conditions is true:
//   It is ().
//   It can be written as AB (A concatenated with B), where A and B are valid parentheses strings.
//   It can be written as (A), where A is a valid parentheses string.
// You are given an m x n matrix of parentheses grid. A valid parentheses string path in the grid
// is a path satisfying all of the following conditions:
//   The path starts from the upper left cell (0, 0).
//   The path ends at the bottom-right cell (m - 1, n - 1).
//   The path only ever moves down or right.
//   The resulting parentheses string formed by the path is valid.
// Return true if there exists a valid parentheses string path in the grid. 
// Otherwise, return false.
// Example 1:
//   Input: grid = [["(","(","("],[")","(",")"],["(","(",")"],["(","(",")"]]
//   Output: true
//   Explanation: The above diagram shows two possible paths that form valid parentheses strings.
//     The first path shown results in the valid parentheses string "()(())".
//     The second path shown results in the valid parentheses string "((()))".
//     Note that there may be other valid parentheses string paths.
// Example 2:
//   Input: grid = [[")",")"],["(","("]]
//   Output: false
//   Explanation: The two possible paths form the parentheses strings "))(" and ")((".
//     Since neither of them are valid parentheses strings, we return false.
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 100
//   grid[i][j] is either '(' or ')'.

mod _check_if_there_is_a_valid_parentheses_string_path {
    struct Solution{
        grid: Vec<Vec<char>>,
        ans: bool,
    }

    // O(mn(m+n))
    // can be O(mn), since the set contains only consecutive values, we can only save the max and min value
    use std::collections::HashSet;
    impl Solution {
        pub fn has_valid_path(grid: Vec<Vec<char>>) -> bool {
            if grid[0][0] == ')' {
                return false;
            }
            let (m, n) = (grid.len(), grid[0].len());
            // for first row
            let mut dp = vec![HashSet::with_capacity(1); n];
            dp[0].insert(1);
            let mut prev = 1;
            for j in 1..n {
                if grid[0][j] == '(' {
                    prev += 1;
                } else {
                    prev -= 1;
                    if prev < 0 {
                        break;
                    }
                }
                dp[j].insert(prev);
            }
            // for each row down
            let mut prev = 1;
            for i in 1..m {
                // for dp[0]
                if prev >= 0 {
                    if grid[i][0] == '(' {
                        dp[0].remove(&prev);
                        prev += 1;
                        dp[0].insert(prev);
                    } else {
                        dp[0].remove(&prev);
                        prev -= 1;
                        if prev >= 0 {
                            dp[0].insert(prev);
                        }
                    }
                }
                for j in 1..n {
                    let d = if grid[i][j] == '(' { 1 } else { -1 };
                    dp[j] = dp[j].union(&dp[j-1]).map(|&s| s+d).collect();
                    dp[j].remove(&-1);
                }
            }
            return dp[n-1].contains(&0);
        }
    }
    
    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                grid: vec![['(','(','('],[')','(',')'],['(','(',')'],['(','(',')']].iter().map(|s| s.to_vec()).collect(),
                ans: true,
            },
            Solution {
                grid: vec![['(','(','('],[')','(',')'],[')','(',')'],['(','(',')']].iter().map(|s| s.to_vec()).collect(),
                ans: true,
            },
            Solution {
                grid: vec![[')',')'],['(','(']].iter().map(|s| s.to_vec()).collect(),
                ans: false,
            }
        ];
        for i in testcases {
            let ans = Solution::has_valid_path(i.grid);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
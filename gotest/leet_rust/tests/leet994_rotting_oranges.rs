// https://leetcode.com/problems/rotting-oranges/

// You are given an m x n grid where each cell can have one of three values:
//   0 representing an empty cell,
//   1 representing a fresh orange, or
//   2 representing a rotten orange.
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
// Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this
// is impossible, return -1.
// Example 1:
//   Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
//   Output: 4
// Example 2:
//   Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
//   Output: -1
//   Explanation: The orange in the bottom left corner (row 2, column 0) is never
//     rotten, because rotting only happens 4-directionally.
// Example 3:
//   Input: grid = [[0,2]]
//   Output: 0
//   Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 10
//   grid[i][j] is 0, 1, or 2.

mod _rotting_oranges {
    struct Solution{
        grid: Vec<Vec<i32>>,
        ans: i32,
    }

    use std::collections::VecDeque;
    impl Solution {
        pub fn oranges_rotting(mut grid: Vec<Vec<i32>>) -> i32 {
            let mut sum_good = 0;
            let mut rotten = VecDeque::new();
            for i in 0..grid.len() {
                for j in 0..grid[0].len() {
                    if grid[i][j] == 1 {
                        sum_good += 1;
                    }
                    if grid[i][j] == 2 {
                        rotten.push_back([i as i32, j as i32]);
                    }
                }
            }
            let mut step = 0;
            while sum_good > 0 && rotten.len() > 0 {
                let l = rotten.len();
                for _ in 0..l {  // bfs for current layer
                    let cur = rotten.pop_front().unwrap();
                    for (i, j) in [(cur[0]-1,cur[1]), (cur[0]+1,cur[1]), (cur[0],cur[1]-1), (cur[0],cur[1]+1)] {
                        if i>=0 && (i as usize)<grid.len() &&
                            j>=0 && (j as usize)<grid[0].len() &&
                            grid[i as usize][j as usize]==1 {
                            grid[i as usize][j as usize] = 2;
                            rotten.push_back([i, j]);
                            sum_good -= 1;
                        }
                    }
                }
                step += 1;
            }
            return match sum_good>0 {
                true => -1,
                false => step,
            };
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                grid: vec![vec![2,1,1],vec![1,1,0],vec![0,1,1]],
                ans: 4,
            },
            Solution {
                grid: vec![vec![2,1,1],vec![0,1,1],vec![1,0,1]],
                ans: -1,
            },
            Solution {
                grid: vec![vec![0,2]],
                ans: 0,
            }
        ];
        for i in testcases {
            let ans = Solution::oranges_rotting(i.grid);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
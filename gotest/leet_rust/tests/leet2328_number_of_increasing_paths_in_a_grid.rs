// https://leetcode.com/problems/number-of-increasing-paths-in-a-grid/

// You are given an m x n integer matrix grid, where you can move from a cell to any adjacent cell
// in all 4 directions.
// Return the number of strictly increasing paths in the grid such that you can start from any cell
// and end at any cell. Since the answer may be very large, return it modulo 10⁹ + 7.
// Two paths are considered different if they do not have exactly the same sequence of visited cells.
// Example 1:
//   Input: grid = [[1,1],[3,4]]
//   Output: 8
//   Explanation: The strictly increasing paths are:
//     - Paths with length 1: [1], [1], [3], [4].
//     - Paths with length 2: [1 -> 3], [1 -> 4], [3 -> 4].
//     - Paths with length 3: [1 -> 3 -> 4].
//     The total number of paths is 4 + 3 + 1 = 8.
// Example 2:
//   Input: grid = [[1],[2]]
//   Output: 3
//   Explanation: The strictly increasing paths are:
//     - Paths with length 1: [1], [2].
//     - Paths with length 2: [1 -> 2].
//     The total number of paths is 2 + 1 = 3.
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 1000
//   1 <= m * n <= 10⁵
//   1 <= grid[i][j] <= 10⁵

mod _number_of_increasing_paths_in_a_grid {
    struct Solution{
        grid: Vec<Vec<i32>>,
        ans: i32,
    }

    impl Solution {
        pub fn count_paths_sort(grid: Vec<Vec<i32>>) -> i32 {
            const P: i32 = 1e9 as i32 + 7;
            let (m, n) = (grid.len(), grid[0].len());
            let (mm, nn) = (m as i32, n as i32);
            // sort cells
            let mut cells = Vec::with_capacity(m*n);
            for i in 0..m {
                for j in 0..n  {
                    cells.push((grid[i][j], i, j));
                }
            }
            cells.sort_by_key(|s| s.0);

            // dp
            let mut dp = vec![vec![0; n]; m];
            for (cur, i, j) in cells {
                dp[i][j] = 1;
                for (di, dj) in [(1,0),(-1,0),(0,1),(0,-1)] {
                    let (ni, nj) = (i as i32 + di, j as i32 + dj);
                    if ni>=0 && ni<mm && nj>=0 && nj<nn {
                        let (ni, nj) = (ni as usize, nj as usize);
                        if grid[ni][nj] < cur {
                            dp[i][j] = (dp[i][j]+dp[ni][nj]) % P
                        }
                    }
                }
            }
            let mut ans = 0;
            for i in 0..m {
                for j in 0..n  {
                    ans = (ans+dp[i][j]) % P
                }
            }
            return ans;
        }

        // dfs + memo
        pub fn count_paths(grid: Vec<Vec<i32>>) -> i32 {
            const P: i32 = 1e9 as i32 + 7;
            let (m, n) = (grid.len(), grid[0].len());

            // dfs
            let mut dp = vec![vec![0; n]; m];
            let mut ans = 0;
            for i in 0..m {
                for j in 0..n  {
                    ans = (ans + Solution::visit(&grid, &mut dp, i, j)) % P
                }
            }
            return ans;
        }

        fn visit(grid: &Vec<Vec<i32>>, dp: &mut Vec<Vec<i32>>, i: usize, j :usize) -> i32 {
            if dp[i][j] > 0 {
                return dp[i][j];
            }

            const P: i32 = 1e9 as i32 + 7;
            let (m, n) = (grid.len(), grid[0].len());
            let (mm, nn) = (m as i32, n as i32);

            dp[i][j] = 1;
            for (di, dj) in [(1,0),(-1,0),(0,1),(0,-1)] {
                let (ni, nj) = (i as i32 + di, j as i32 + dj);
                if ni>=0 && ni<mm && nj>=0 && nj<nn {
                    let (ni, nj) = (ni as usize, nj as usize);
                    if grid[ni][nj] < grid[i][j] {
                        dp[i][j] = (dp[i][j] + Solution::visit(grid, dp, ni, nj)) % P;
                    }
                }
            }
            return dp[i][j];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                grid: vec![[1,1],[3,4]].iter().map(|s| s.to_vec()).collect(),
                ans: 8,
            },
            Solution {
                grid: vec![[1],[2]].iter().map(|s| s.to_vec()).collect(),
                ans: 3,
            },
            Solution {
                grid: vec![[1,1],[1,1]].iter().map(|s| s.to_vec()).collect(),
                ans: 4,
            },
        ];
        for i in testcases {
            let ans1 = Solution::count_paths(i.grid.clone());
            let ans2 = Solution::count_paths_sort(i.grid.clone());
            println!("{}, {}, {}", ans1, ans2, i.ans);
        }
    } 
}
// https://leetcode.com/problems/count-unguarded-cells-in-the-grid/

// You are given two integers m and n representing a 0-indexed m x n grid. You are also given two 2D
// integer arrays guards and walls where guards[i] = [rowi, coli] and walls[j] = [rowj, colj]
// represent the positions of the iᵗʰ guard and jᵗʰ wall respectively.
// A guard can see every cell in the four cardinal directions (north, east,south, or west) starting
// from their position unless obstructed by a wall or another guard. A cell is guarded if there is
// at least one guard that can see it.
// Return the number of unoccupied cells that are not guarded.
// Example 1:
//   Input: m = 4, n = 6, guards = [[0,0],[1,1],[2,3]], walls = [[0,1],[2,2],[1,4]]
//   Output: 7
//   Explanation: The guarded and unguarded cells are shown in red and green
//     respectively in the above diagram. There are a total of 7 unguarded cells, so we return 7.
// Example 2:
//   Input: m = 3, n = 3, guards = [[1,1]], walls = [[0,1],[1,0],[2,1],[1,2]]
//   Output: 4
//   Explanation: The unguarded cells are shown in green in the above diagram.
//     There are a total of 4 unguarded cells, so we return 4.
// Constraints:
//   1 <= m, n <= 10⁵
//   2 <= m * n <= 10⁵
//   1 <= guards.length, walls.length <= 5 * 10⁴
//   2 <= guards.length + walls.length <= m * n
//   guards[i].length == walls[j].length == 2
//   0 <= rowi, rowj < m
//   0 <= coli, colj < n
//   All the positions in guards and walls are unique.

mod _count_unguarded_cells_in_the_grid {
    struct Solution{
        m: i32,
        n: i32,
        guards: Vec<Vec<i32>>,
        walls: Vec<Vec<i32>>,
        ans: i32,
    }

    // just traverse each row from left to right, right to left,
    // each column from up to down, down to up
    impl Solution {
        pub fn count_unguarded(m: i32, n: i32, guards: Vec<Vec<i32>>, walls: Vec<Vec<i32>>) -> i32 {
            let mut grid = vec![vec![0; n as usize]; m as usize];
            for g in &guards {
                grid[g[0] as usize][g[1] as usize] = 1;
            }
            for w in &walls {
                grid[w[0] as usize][w[1] as usize] = 2;
            }
            let mut count = 0;
            for i in 0..grid.len() {
                let mut has = false;
                for j in 0..grid[0].len() {
                    if grid[i][j] == 0 && has {
                        grid[i][j] = -1;  // mark as guarded
                        count += 1;
                    } else if grid[i][j] > 0 {
                        has = grid[i][j]==1;
                    }
                }
                has = false;
                for j in (0..grid[0].len()).into_iter().rev() {
                    if grid[i][j] == 0 && has {
                        grid[i][j] = -1;  // mark as guarded
                        count += 1;
                    } else if grid[i][j] > 0 {
                        has = grid[i][j]==1;
                    }
                }
            }
            for j in 0..grid[0].len() {
                let mut has = false;
                for i in 0..grid.len() {
                    if grid[i][j] == 0 && has {
                        grid[i][j] = -1;  // mark as guarded
                        count += 1;
                    } else if grid[i][j] > 0 {
                        has = grid[i][j]==1;
                    }
                }
                has = false;
                for i in (0..grid.len()).into_iter().rev() {
                    if grid[i][j] == 0 && has {
                        grid[i][j] = -1;  // mark as guarded
                        count += 1;
                    } else if grid[i][j] > 0 {
                        has = grid[i][j]==1;
                    }
                }
            }
            return m*n- count- guards.len() as i32 - walls.len() as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                m: 4,
                n: 6,
                guards: vec![[0,0],[1,1],[2,3]].iter().map(|s| s.to_vec()).collect(),
                walls:vec![[0,1],[2,2],[1,4]].iter().map(|s| s.to_vec()).collect(),
                ans: 7,
            },
            Solution {
                m: 3,
                n: 3,
                guards: vec![[1,1]].iter().map(|s| s.to_vec()).collect(),
                walls:vec![[0,1],[1,0],[2,1],[1,2]].iter().map(|s| s.to_vec()).collect(),
                ans: 4,
            },
        ];
        for i in testcases {
            let ans = Solution::count_unguarded(i.m, i.n, i.guards, i.walls);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
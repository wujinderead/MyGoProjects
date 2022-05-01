// https://leetcode.com/problems/escape-the-spreading-fire/

// You are given a 0-indexed 2D integer array grid of size m x n which represents a field.
// Each cell has one of three values:
//   0 represents grass,
//   1 represents fire,
//   2 represents a wall that you and fire cannot pass through.
// You are situated in the top-left cell, (0, 0), and you want to travel to the safehouse at the
// bottom-right cell, (m - 1, n - 1). Every minute, you may move to an adjacent grass cell. After
// your move, every fire cell will spread to all adjacent cells that are not walls.
// Return the maximum number of minutes that you can stay in your initial position before moving
// while still safely reaching the safehouse. If this is impossible, return -1. If you can always
// reach the safehouse regardless of the minutes stayed, return 10⁹.
// Note that even if the fire spreads to the safehouse immediately after you have reached it,
// it will be counted as safely reaching the safehouse.
// A cell is adjacent to another cell if the former is directly north, east, south, or west of
// the latter (i.e., their sides are touching).
// Example 1:
//   Input: grid = [[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]
//   Output: 3
//   Explanation: The figure above shows the scenario where you stay in the initial position for 3 minutes.
//     You will still be able to safely reach the safehouse.
//     Staying for more than 3 minutes will not allow you to safely reach the safehouse.
// Example 2:
//   Input: grid = [[0,0,0,0],[0,1,2,0],[0,2,0,0]]
//   Output: -1
//   Explanation: The figure above shows the scenario where you immediately move towards the safehouse.
//     Fire will spread to any cell you move towards and it is impossible to safely reach the safehouse.
//     Thus, -1 is returned.
// Example 3:
//   Input: grid = [[0,0,0],[2,2,0],[1,2,0]]
//   Output: 1000000000
//   Explanation: The figure above shows the initial grid.
//     Notice that the fire is contained by walls and you will always be able to safely reach the safehouse.
//     Thus, 10⁹ is returned.
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   2 <= m, n <= 300
//   4 <= m * n <= 2 * 10⁴
//   grid[i][j] is either 0, 1, or 2.
//   grid[0][0] == grid[m - 1][n - 1] == 0

mod _escape_the_spreading_fire {
    struct Solution{
        grid: Vec<Vec<i32>>,
        ans: i32,
    }

    // we can only focus time difference on the bottom-right grid.
    // the reason is that on a valid path from top-left to bot-right,
    // the time difference of fire_time - people_time is non-increasing.
    use std::collections::VecDeque;
    impl Solution {
        pub fn maximum_minutes(grid: Vec<Vec<i32>>) -> i32 {
            let (m, n) = (grid.len() as i32, grid[0].len() as i32);
            let (mm, nn) = (grid.len(), grid[0].len());
            let mut f_time = vec![vec![-1; nn]; mm];
            let mut p_time = vec![vec![-1; nn]; mm];
            let mut fires = VecDeque::new();
            for i in 0..mm {
                for j in 0..nn {
                    if grid[i][j] == 1 {
                        fires.push_back([i as i32, j as i32]);
                    }
                }
            }
            // bfs to find when does a grid get fired
            let mut t = 1;
            while fires.len() > 0 {
                for _i in 0..fires.len() {
                    let cur = fires.pop_front().unwrap();
                    let (i, j) = (cur[0], cur[1]);
                    for (di, dj) in [(0,1), (0,-1), (-1,0), (1,0)] {
                        let (ni, nj) = (i+di, j+dj);
                        if ni>=0 && ni<m && nj>=0 && nj<n {
                            let (ni, nj) = (ni as usize, nj as usize);
                            if grid[ni][nj]==0 && f_time[ni][nj]==-1 {
                                f_time[ni][nj] = t;
                                fires.push_back([ni as i32, nj as i32]);
                            }
                        }
                    }
                }
                t += 1;
            }
            // check when does a person reach a grid
            let mut person = VecDeque::new();
            person.push_back([0, 0]);
            t = 1;
            while person.len() > 0 {
                for _i in 0..person.len() {
                    let cur = person.pop_front().unwrap();
                    let (i, j) = (cur[0], cur[1]);
                    for (di, dj) in [(0,1), (0,-1), (-1,0), (1,0)] {
                        let (ni, nj) = (i+di, j+dj);
                        if ni>=0 && ni<m && nj>=0 && nj<n {
                            let (ni, nj) = (ni as usize, nj as usize);
                            if grid[ni][nj]==0 && p_time[ni][nj]==-1 {
                                p_time[ni][nj] = t;
                                person.push_back([ni as i32, nj as i32]);
                            }
                        }
                    }
                }
                t += 1;
            }
            // check edge case
            if p_time[mm-1][nn-1] == -1 {  // person can't reach
                return -1;
            }
            if f_time[mm-1][nn-1] == -1 {  // fire can't reach
                return 1e9 as i32;
            }
            if f_time[mm-1][nn-1] < p_time[mm-1][nn-1] {
                return -1;
            }
            let diff = f_time[mm-1][nn-1] - p_time[mm-1][nn-1];
            // whether person are 'followed' by fire on both two pathes toward bot-right cell.
            if p_time[mm-2][nn-1] > -1 && p_time[mm-1][nn-2] > -1 &&
                (f_time[mm-2][nn-1] - p_time[mm-2][nn-1] > diff || f_time[mm-1][nn-2] - p_time[mm-1][nn-2] > diff) {
               return diff;
            }
            return diff-1;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                grid: vec![[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]].iter().map(|s| s.to_vec()).collect(),
                ans: 3,
            },
            Solution {
                grid: vec![[0,0,0,0],[0,1,2,0],[0,2,0,0]].iter().map(|s| s.to_vec()).collect(),
                ans: -1,
            },
            Solution {
                grid: vec![[0,0,0],[2,2,0],[1,2,0]].iter().map(|s| s.to_vec()).collect(),
                ans: 1e9 as i32,
            },
        ];
        for i in testcases {
            let ans = Solution::maximum_minutes(i.grid);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
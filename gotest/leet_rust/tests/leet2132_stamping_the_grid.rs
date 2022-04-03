// https://leetcode.com/problems/stamping-the-grid/

// You are given an m x n binary matrix grid where each cell is either 0 (empty) or 1 (occupied).
// You are then given stamps of size stampHeight x stampWidth. We want to fit the stamps such
// that they follow the given restrictions and requirements:
//   Cover all the empty cells.
//   Do not cover any of the occupied cells.
//   We can put as many stamps as we want.
//   Stamps can overlap with each other.
//   Stamps are not allowed to be rotated.
//   Stamps must stay completely inside the grid.
// Return true if it is possible to fit the stamps while following the given restrictions and
// requirements. Otherwise, return false.
// Example 1: 
//   Input: grid = [[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0]], stampHeight = 4, stampWidth = 3
//   Output: true
//   Explanation: We have two overlapping stamps (labeled 1 and 2 in the image)
//     that are able to cover all the empty cells.
// Example 2:
//   Input: grid = [[1,0,0,0],[0,1,0,0],[0,0,1,0],[0,0,0,1]], stampHeight = 2, stampWidth = 2
//   Output: false
//   Explanation: There is no way to fit the stamps onto all the empty cells
//     without the stamps going outside the grid.
// Constraints:
//   m == grid.length
//   n == grid[r].length
//   1 <= m, n <= 10⁵
//   1 <= m * n <= 2 * 10⁵
//   grid[r][c] is either 0 or 1.
//   1 <= stampHeight, stampWidth <= 10⁵

mod _stamping_the_grid {
    struct Solution{
        grid: Vec<Vec<i32>>,
        stamp_height: i32,
        stamp_width: i32,
        ans: bool,
    }

    impl Solution {
        pub fn possible_to_stamp(grid: Vec<Vec<i32>>, stamp_height: i32, stamp_width: i32) -> bool {
            let (m, n) = (grid.len(), grid[0].len());
            let (h, w) = (stamp_height as usize, stamp_width as usize);
            // get cumulative sum of the original grid
            let mut cum = vec![vec![0; n+1]; m+1];
            for i in (0..m).into_iter().rev() {
                for j in (0..n).into_iter().rev() {
                    cum[i][j] = cum[i+1][j] + cum[i][j+1] + grid[i][j] - cum[i+1][j+1];
                }
            }
            // good[i][j]==1 if (i,j) is the top-left corner of a valid stamp
            //println!("{:?}", cum);
            let mut good = vec![vec![0; n+1]; m+1];
            // check each empty node on grid
            for i in 0..m {
                for j in 0..n { // check cumulative sum of grid[i][j] to grid[i+h-1][j+w-1]
                    if grid[i][j] == 0 && i+h-1 < m && j+w-1 < n {
                        if cum[i][j] - cum[i+h][j] - cum[i][j+w] + cum[i+h][j+w] == 0 {
                            good[i][j] = 1;
                        }
                    }
                }
            }
            //println!("{:?}", good);
            // get cumulative sum of `good`
            for i in (0..m).into_iter().rev() {
                for j in (0..n).into_iter().rev() {
                    good[i][j] = good[i+1][j] + good[i][j+1] + good[i][j] - good[i+1][j+1];
                }
            }
            // for each empty node in original grid, if this node can be stamped,
            // there must exist 1s in good[i-h+1][j-w+1] to good[i][j] (here good is the original good)
            for i in 0..m {
                for j in 0..n {
                    if grid[i][j] != 0 {
                        continue;
                    }
                    let ni = if i<h-1 { 0 } else { i+1-h };
                    let nj = if j<w-1 { 0 } else { j+1-w };
                    if good[ni][nj] - good[i+1][nj] - good[ni][j+1] + good[i+1][j+1] == 0 {
                        return false; // all-0 among good[i-h+1][j-w+1] to good[i][j], can't stamp
                    }
                }
            }
            return true;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                grid: vec![vec![1,0,0,0],vec![1,0,0,0],vec![1,0,0,0],vec![1,0,0,0],vec![1,0,0,0]],
                stamp_height: 4,
                stamp_width: 3,
                ans: true,
            },
            Solution {
                grid: vec![vec![1,0,0,0],vec![0,1,0,0],vec![0,0,1,0],vec![0,0,0,1]],
                stamp_height: 2,
                stamp_width: 2,
                ans: false,
            }
        ];
        for i in testcases {
            let ans = Solution::possible_to_stamp(i.grid, i.stamp_height, i.stamp_width);
            println!("{}, {}", ans, i.ans);
        }
    } 
}

//[[4, 3, 2, 1, 0], [3, 3, 2, 1, 0], [2, 2, 2, 1, 0], [1, 1, 1, 1, 0], [0, 0, 0, 0, 0]]
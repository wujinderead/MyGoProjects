// https://leetcode.com/problems/path-with-minimum-effort/

// You are a hiker preparing for an upcoming hike. You are given heights, a 2D array of size rows
// x columns, where heights[row][col] represents the height of cell (row, col). You are situated
// in the top-left cell, (0, 0), and you hope to travel to the bottom-right cell,
// (rows-1, columns-1) (i.e., 0-indexed). You can move up, down, left, or right, and you wish to
// find a route that requires the minimum effort.
// A route's effort is the maximum absolute difference in heights between two consecutive cells
// of the route.
// Return the minimum effort required to travel from the top-left cell to the bottom-right cell.
// Example 1:
//   Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
//   Output: 2
//   Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in consecutive cells.
//     This is better than the route of [1,2,2,2,5], where the maximum absolute difference is 3.
// Example 2:
//   Input: heights = [[1,2,3],[3,8,4],[5,3,5]]
//   Output: 1
//   Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1
//     in consecutive cells, which is better than route [1,3,5,3,5].
// Example 3:
//   Input: heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
//   Output: 0
//   Explanation: This route does not require any effort.
// Constraints:
//   rows == heights.length
//   columns == heights[i].length
//   1 <= rows, columns <= 100
//   1 <= heights[i][j] <= 10⁶

mod _path_with_minimum_effort {
    struct Solution{
        heights: Vec<Vec<i32>>,
        ans: i32,
    }

    use std::collections::BinaryHeap;
    use std::ops::Neg;
    impl Solution {
        // dijkstra to find the shortest path
        pub fn minimum_effort_path(heights: Vec<Vec<i32>>) -> i32 {
            let (m, n) = (heights.len(), heights[0].len());
            let (mm, nn) = (m as i32, n as i32);
            let mut min = vec![vec![1<<30; n]; m];
            let mut heap = BinaryHeap::new();
            heap.push((0, 0, 0));
            while heap.len() > 0 {
                let (diff, x, y) = heap.pop().unwrap();
                let (xx, yy, diff) = (x as i32, y as i32, -diff);
                if min[x][y] <= diff {  // already find shortest path for (x, y), skip
                    continue;
                }
                min[x][y] = diff;
                if x == m-1 && y == n-1 {
                    return diff;
                }
                for (dx, dy) in [(1,0), (-1,0), (0,1), (0,-1)] {
                    if xx+dx>=0 && xx+dx<mm && yy+dy>=0 && yy+dy<nn {
                        let (nx, ny) = ((xx+dx) as usize, (yy+dy) as usize);
                        if min[nx][ny] == 1<<30 {  // new node is unvisited
                            let tmp = (heights[x][y] - heights[nx][ny]).abs();
                            // the distance to new node is the max of previous_diff and current_diff
                            heap.push((diff.max(tmp).neg(), nx, ny)); // negate cause we want min heap
                        }
                    }
                }
            }
            return -1;
        }

        // binary search: binary search the min diff,
        // check if we can reach target with given diff
        pub fn minimum_effort_path_binary_search(heights: Vec<Vec<i32>>) -> i32 {
            let (mut left, mut right) = (0, 999999);
            while left < right {
                let mid = (left+right)/2;
                let mut visited = vec![vec![false; heights[0].len()]; heights.len()];
                Solution::visit(0, 0, &heights, &mut visited, mid);
                if visited[heights.len()-1][heights[0].len()-1] { // can reach target
                    right = mid;
                } else {
                    left = mid+1;
                }
            }
            return left;
        }

        fn visit(x: usize, y: usize, heights: &Vec<Vec<i32>>, visited: &mut Vec<Vec<bool>>, diff: i32) {
            visited[x][y] = true;
            let (xx, yy) = (x as i32, y as i32);
            let (m, n) = (heights.len() as i32, heights[0].len() as i32);
            for (dx, dy) in [(1,0), (-1,0), (0,1), (0,-1)] {
                if xx+dx>=0 && xx+dx<m && yy+dy>=0 && yy+dy<n {
                    let (nx, ny) = ((xx+dx) as usize, (yy+dy) as usize);
                    if !visited[nx][ny] && (heights[x][y]-heights[nx][ny]).abs() <= diff{
                        Solution::visit(nx, ny, heights, visited, diff);
                    }
                }
            }
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                heights: vec![[1,2,2],[3,8,2],[5,3,5]].iter().map(|s| s.to_vec()).collect(),
                ans: 2,
            },
            Solution {
                heights: vec![[1,2,3],[3,8,4],[5,3,5]].iter().map(|s| s.to_vec()).collect(),
                ans: 1,
            },
            Solution {
                heights: vec![[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]].iter().map(|s| s.to_vec()).collect(),
                ans: 0,
            },
            Solution {
                heights: vec![[1]].iter().map(|s| s.to_vec()).collect(),
                ans: 0,
            },
            Solution {
                heights: vec![[1,2]].iter().map(|s| s.to_vec()).collect(),
                ans: 1,
            },
            Solution {
                heights: vec![[2],[1]].iter().map(|s| s.to_vec()).collect(),
                ans: 1,
            },
            Solution {
                heights: vec![[1],[1000000]].iter().map(|s| s.to_vec()).collect(),
                ans: 999999,
            },
        ];
        for i in testcases {
            let ans1 = Solution::minimum_effort_path(i.heights.clone());
            let ans2 = Solution::minimum_effort_path_binary_search(i.heights.clone());
            println!("{}, {}, {}", ans1, ans2, i.ans);
        }
    } 
}
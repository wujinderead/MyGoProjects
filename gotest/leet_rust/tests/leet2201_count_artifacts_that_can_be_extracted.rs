// https://leetcode.com/problems/count-artifacts-that-can-be-extracted/

// There is an n x n 0-indexed grid with some artifacts buried in it. You are given the integer
// n and a 0-indexed 2D integer array artifacts describing the positions of the rectangular
// artifacts where artifacts[i] = [r1i, c1i, r2i, c2i] denotes that the iᵗʰ artifact is buried
// in the subgrid where:
//   (r1i, c1i) is the coordinate of the top-left cell of the iᵗʰ artifact and
//   (r2i, c2i) is the coordinate of the bottom-right cell of the iᵗʰ artifact.
// You will excavate some cells of the grid and remove all the mud from them. 
// If the cell has a part of an artifact buried underneath, it will be uncovered.
// If all the parts of an artifact are uncovered, you can extract it.
// Given a 0-indexed 2D integer array dig where dig[i] = [ri, ci] indicates that you will
// excavate the cell (ri, ci), return the number of artifacts that you can extract.
// The test cases are generated such that:
//   No two artifacts overlap.
//   Each artifact only covers at most 4 cells.
//   The entries of dig are unique.
// Example 1:
//   Input: n = 2, artifacts = [[0,0,0,0],[0,1,1,1]], dig = [[0,0],[0,1]]
//   Output: 1
//   Explanation:
//     The different colors represent different artifacts.
//     Excavated cells are labeled with a 'D' in the grid.
//     There is 1 artifact that can be extracted, namely the red artifact.
//     The blue artifact has one part in cell (1,1) which remains uncovered,
//     so we cannot extract it. Thus, we return 1.
// Example 2:
//   Input: n = 2, artifacts = [[0,0,0,0],[0,1,1,1]], dig = [[0,0],[0,1],[1,1]]
//   Output: 2
//   Explanation: Both the red and blue artifacts have all parts uncovered
//     (labeled with a 'D') and can be extracted, so we return 2.
// Constraints:
//   1 <= n <= 1000
//   1 <= artifacts.length, dig.length <= min(n², 10⁵)
//   artifacts[i].length == 4
//   dig[i].length == 2
//   0 <= r1i, c1i, r2i, c2i, ri, ci <= n - 1
//   r1i <= r2i
//   c1i <= c2i
//   No two artifacts will overlap.
//   The number of cells covered by an artifact is at most 4.
//   The entries of dig are unique.

mod _count_artifacts_that_can_be_extracted {

    struct Solution{
        n: i32,
        artifacts: Vec<Vec<i32>>,
        dig: Vec<Vec<i32>>,
        ans: i32,
    }

    use std::collections::HashMap;
    impl Solution {
        pub fn dig_artifacts(n: i32, artifacts: Vec<Vec<i32>>, dig: Vec<Vec<i32>>) -> i32 {
            let mut sum = 0;
            let mut point_art = HashMap::new();
            let mut size_art = vec![0; artifacts.len()];
            for (ind, v) in artifacts.iter().enumerate() {
                for i in v[0]..=v[2] {
                    for j in v[1]..=v[3] {
                        point_art.insert(i*n+j, ind);
                    }
                }
                size_art[ind] = (v[2]-v[0]+1) * (v[3]-v[1]+1);
            }
            for d in dig {
                if let Some(&ind) = point_art.get(&(d[0]*n+d[1])) {
                    size_art[ind] -= 1;
                    if size_art[ind] == 0 {
                        sum += 1;
                    }
                }
            }
            return sum;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                n: 2,
                artifacts: vec![vec![0,0,0,0],vec![0,1,1,1]],
                dig: vec![vec![0,0],vec![0,1]],
                ans: 1,
            },
            Solution {
                n: 2,
                artifacts: vec![vec![0,0,0,0],vec![0,1,1,1]],
                dig: vec![vec![0,0],vec![0,1],vec![1,1]],
                ans: 2,
            },
        ];
        for i in testcases {
            let ans = Solution::dig_artifacts(i.n, i.artifacts, i.dig);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
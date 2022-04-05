// https://leetcode.com/problems/detonate-the-maximum-bombs/

// You are given a list of bombs. The range of a bomb is defined as the area where its effect
// can be felt. This area is in the shape of a circle with the center as the location of the bomb.
// The bombs are represented by a 0-indexed 2D integer array bombs where bombs[i] = [xi, yi, ri].
// xi and yi denote the X-coordinate and Y-coordinate of the location of the iᵗʰ bomb, whereas
// ri denotes the radius of its range.
// You may choose to detonate a single bomb. When a bomb is detonated, it will detonate all bombs
// that lie in its range. These bombs will further detonate the bombs that lie in their ranges.
// Given the list of bombs, return the maximum number of bombs that can be detonated if you
// are allowed to detonate only one bomb.
// Example 1:
//   Input: bombs = [[2,1,3],[6,1,4]]
//   Output: 2
//   Explanation:
//     The above figure shows the positions and ranges of the 2 bombs.
//     If we detonate the left bomb, the right bomb will not be affected.
//     But if we detonate the right bomb, both bombs will be detonated.
//     So the maximum bombs that can be detonated is max(1, 2) = 2.
// Example 2:
//   Input: bombs = [[1,1,5],[10,10,5]]
//   Output: 1
//   Explanation:
//     Detonating either bomb will not detonate the other bomb, so the maximum
//     number of bombs that can be detonated is 1.
// Example 3:
//   Input: bombs = [[1,2,3],[2,3,1],[3,4,2],[4,5,3],[5,6,4]]
//   Output: 5
//   Explanation:
//     The best bomb to detonate is bomb 0 because:
//     - Bomb 0 detonates bombs 1 and 2. The red circle denotes the range of bomb 0.
//     - Bomb 2 detonates bomb 3. The blue circle denotes the range of bomb 2.
//     - Bomb 3 detonates bomb 4. The green circle denotes the range of bomb 3.
//     Thus all 5 bombs are detonated.
// Constraints:
//   1 <= bombs.length <= 100
//   bombs[i].length == 3
//   1 <= xi, yi, ri <= 10⁵

mod _detonate_the_maximum_bombs {
    struct Solution{
        bombs: Vec<Vec<i32>>,
        ans: i32,
    }

    // note: the graph is a directed graph
    impl Solution {
        pub fn maximum_detonation(bombs: Vec<Vec<i32>>) -> i32 {
            let mut neighbor = vec![vec![]; bombs.len()];
            for i in 0..bombs.len() {
                for j in (i+1)..bombs.len() {
                    let a1 = (bombs[i][0]-bombs[j][0]) as i64;
                    let a2 = (bombs[i][1]-bombs[j][1]) as i64;
                    if a1*a1+a2*a2 <= (bombs[i][2] as i64)*(bombs[i][2] as i64) {
                        // bomb j in bomb i's range, i can detonate j
                        neighbor[i].push(j);
                    }
                    if a1*a1+a2*a2 <= (bombs[j][2] as i64)*(bombs[j][2] as i64) {
                        neighbor[j].push(i);
                    }
                }
            }
            let mut ans = 0;
            let mut i = 0;
            while i<bombs.len() {
                let mut number = 0;
                let mut visited = vec![false; bombs.len()];
                Solution::visit(i, &neighbor, &mut visited, &mut number);
                ans = ans.max(number);
                i += 1;
            }
            return ans;
        }

        fn visit(i: usize, neighbor: &Vec<Vec<usize>>, visited: &mut Vec<bool>, number: &mut i32) {
            *number += 1;
            visited[i] = true;
            for &n in &neighbor[i] {
                if !visited[n] {
                    Solution::visit(n, neighbor, visited, number);
                }
            }
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                bombs: vec![[2,1,3],[6,1,4]].iter().map(|s| s.to_vec()).collect(),
                ans: 2,
            },
            Solution {
                bombs: vec![[1,1,5],[10,10,5]].iter().map(|s| s.to_vec()).collect(),
                ans: 1,
            },
            Solution {
                bombs: vec![[1,2,3],[2,3,1],[3,4,2],[4,5,3],[5,6,4]].iter().map(|s| s.to_vec()).collect(),
                ans: 5,
            },
            Solution {
                bombs: vec![[1,1,100000],[100000,100000,1]].iter().map(|s| s.to_vec()).collect(),
                ans: 1,
            },
            Solution {
                bombs: vec![[54,95,4],[99,46,3],[29,21,3],[96,72,8],[49,43,3],[11,20,3],[2,57,1],[69,51,7],[97,1,10],[85,45,2],[38,47,1],[83,75,3],[65,59,3],[33,4,1],[32,10,2],[20,97,8],[35,37,3]].iter().map(|s| s.to_vec()).collect(),
                ans: 1,
            }
        ];
        for i in testcases {
            let ans = Solution::maximum_detonation(i.bombs);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
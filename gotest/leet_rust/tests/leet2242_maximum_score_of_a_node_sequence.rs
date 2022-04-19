// https://leetcode.com/problems/maximum-score-of-a-node-sequence/

// There is an undirected graph with n nodes, numbered from 0 to n - 1.
// You are given a 0-indexed integer array scores of length n where scores[i] denotes the score of
// node i. You are also given a 2D integer array edges where edges[i] = [ai, bi] denotes that there
// exists an undirected edge connecting nodes ai and bi.
// A node sequence is valid if it meets the following conditions:
//   There is an edge connecting every pair of adjacent nodes in the sequence.
//   No node appears more than once in the sequence.
// The score of a node sequence is defined as the sum of the scores of the nodes in the sequence.
// Return the maximum score of a valid node sequence with a length of 4. If no such sequence exists,
// return -1.
// Example 1:
//   Input: scores = [5,2,9,8,4], edges = [[0,1],[1,2],[2,3],[0,2],[1,3],[2,4]]
//   Output: 24
//   Explanation: The figure above shows the graph and the chosen node sequence [0,1,2,3].
//     The score of the node sequence is 5 + 2 + 9 + 8 = 24.
//     It can be shown that no other node sequence has a score of more than 24.
//     Note that the sequences [3,1,2,0] and [1,0,2,3] are also valid and have a score of 24.
//     The sequence [0,3,2,4] is not valid since no edge connects nodes 0 and 3.
// Example 2:
//   Input: scores = [9,20,6,4,11,12], edges = [[0,3],[5,3],[2,4],[1,3]]
//   Output: -1
//   Explanation: The figure above shows the graph.
//     There are no valid node sequences of length 4, so we return -1.
// Constraints:
//   n == scores.length
//   4 <= n <= 5 * 10⁴
//   1 <= scores[i] <= 10⁸
//   0 <= edges.length <= 5 * 10⁴
//   edges[i].length == 2
//   0 <= ai, bi <= n - 1
//   ai != bi
//   There are no duplicate edges.

mod _maximum_score_of_a_node_sequence {
    struct Solution{
        scores: Vec<i32>,
        edges: Vec<Vec<i32>>,
        ans: i32,
    }

    impl Solution {
        pub fn maximum_score(scores: Vec<i32>, edges: Vec<Vec<i32>>) -> i32 {
            // find the top-3 neighbors with larger score
            let mut neighbor = vec![Vec::<usize>::with_capacity(3); scores.len()];
            for e in &edges {
                let (s, t) = (e[0] as usize, e[1] as usize);
                if neighbor[s].len()<3 {
                    neighbor[s].push(t);
                } else {
                    let mut min = 0;
                    if scores[neighbor[s][1]] < scores[neighbor[s][min]] {
                        min = 1;
                    }
                    if scores[neighbor[s][2]] < scores[neighbor[s][min]] {
                        min = 2;
                    }
                    if scores[neighbor[s][min]] < scores[t] {
                        neighbor[s][min] = t;
                    }
                }
                if neighbor[t].len()<3 {
                    neighbor[t].push(s);
                } else {
                    let mut min = 0;
                    if scores[neighbor[t][1]] < scores[neighbor[t][min]] {
                        min = 1;
                    }
                    if scores[neighbor[t][2]] < scores[neighbor[t][min]] {
                        min = 2;
                    }
                    if scores[neighbor[t][min]] < scores[s] {
                        neighbor[t][min] = s;
                    }
                }
            }
            // find max-scored 4-node path
            let mut max = -1;
            for e in &edges {
                let (s, t) = (e[0] as usize, e[1] as usize);
                for &i in &neighbor[s] {
                    if i == t {
                        continue;
                    }
                    for &j in &neighbor[t] {
                        if s == j || i == j {
                            continue;
                        }
                        max = max.max(scores[s]+scores[t]+scores[i]+scores[j]);
                    }
                }
            }
            return max;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                scores: vec![5,2,9,8,4],
                edges: vec![[0,1],[1,2],[2,3],[0,2],[1,3],[2,4]].iter().map(|s| s.to_vec()).collect(),
                ans: 24,
            },
            Solution {
                scores: vec![9,20,6,4,11,12],
                edges: vec![[0,3],[5,3],[2,4],[1,3]].iter().map(|s| s.to_vec()).collect(),
                ans: -1,
            },
            Solution {
                scores: vec![10,7,3,6,4,9,11,8,3,5,4,12],
                edges: vec![[0,1],[0,2],[0,3],[0,4],[0,5],[0,6],[6,7],[6,8],[6,9],[6,10],[6,11]].iter().map(|s| s.to_vec()).collect(),
                ans: 42,
            }
        ];
        for i in testcases {
            let ans = Solution::maximum_score(i.scores, i.edges);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
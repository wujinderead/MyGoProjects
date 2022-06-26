// https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/

// You are given an integer n. There is an undirected graph with n nodes, numbered from 0 to n - 1.
// You are given a 2D integer array edges where edges[i] = [ai,bi] denotes that there exists an
// undirected edge connecting nodes ai and bi.
// Return the number of pairs of different nodes that are unreachable from each other.
// Example 1:
//   Input: n = 3, edges = [[0,1],[0,2],[1,2]]
//   Output: 0
//   Explanation: There are no pairs of nodes that are unreachable from each other.
//     Therefore, we return 0.
// Example 2:
//   Input: n = 7, edges = [[0,2],[0,5],[2,4],[1,6],[5,4]]
//   Output: 14
//   Explanation: There are 14 pairs of nodes that are unreachable from each other:
//     [[0,1],[0,3],[0,6],[1,2],[1,3],[1,4],[1,5],[2,3],[2,6],[3,4],[3,5],[3,6],[4,6],[5,6]].
//     Therefore, we return 14.
// Constraints:
//   1 <= n <= 10⁵
//   0 <= edges.length <= 2 * 10⁵
//   edges[i].length == 2
//   0 <= ai, bi < n
//   ai != bi
//   There are no repeated edges.

mod _count_unreachable_pairs_of_nodes_in_an_undirected_graph {
    struct Solution{
        n: i32,
        edges: Vec<Vec<i32>>,
        ans: i32,
    }

    // union find or dfs
    use std::collections::HashMap;
    impl Solution {
        pub fn count_pairs(n: i32, edges: Vec<Vec<i32>>) -> i64 {
            let mut root = vec![-1; n as usize];
            for pair in edges {
                let r1 = Solution::get_root(&mut root, pair[0]);
                let r2 = Solution::get_root(&mut root, pair[1]);
                if r1 == r2 {
                    continue;
                }
                root[r2 as usize] = r1;
            }
            let mut map = HashMap::new();
            for i in 0..n {
                let r = Solution::get_root(&mut root, i);
                *map.entry(r).or_insert(0) += 1;
            }
            let mut ans = n as i64 * (n-1) as i64 / 2;
            for (_k, v) in map {
                ans -= v*(v-1)/2;
            }
            return ans;
        }

        fn get_root(root: &mut Vec<i32>, ind: i32) -> i32 {
            if root[ind as usize] == -1 {
                return ind;
            }
            let x = Solution::get_root(root, root[ind as usize]);
            root[ind as usize] = x;
            return x;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                n: 3,
                edges: vec![[0,1],[0,2],[1,2]].iter().map(|s| s.to_vec()).collect(),
                ans: 0,
            },
            Solution {
                n: 7,
                edges: vec![[0,2],[0,5],[2,4],[1,6],[5,4]].iter().map(|s| s.to_vec()).collect(),
                ans: 14,
            },
        ];
        for i in testcases {
            let ans = Solution::count_pairs(i.n, i.edges);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/minimum-height-trees/

// A tree is an undirected graph in which any two vertices are connected by exactly one path.
// In other words, any connected graph without simple cycles is a tree.
// Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where edges[i] =
// [ai, bi] indicates that there is an undirected edge between the two nodes ai and bi in the tree,
// you can choose any node of the tree as the root. When you select a node x as the root, the result
// tree has height h. Among all possible rooted trees, those with minimum height (i.e. min(h)) are
// called minimum height trees (MHTs).
// Return a list of all MHTs' root labels. You can return the answer in any order.
// The height of a rooted tree is the number of edges on the longest downward path between the root
// and a leaf.
// Example 1:
//   Input: n = 4, edges = [[1,0],[1,2],[1,3]]
//   Output: [1]
//   Explanation: As shown, the height of the tree is 1 when the root is the node
//     with label 1 which is the only MHT.
// Example 2:
//   Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
//   Output: [3,4]
// Constraints:
//   1 <= n <= 2 * 10⁴
//   edges.length == n - 1
//   0 <= ai, bi < n
//   ai != bi
//   All the pairs (ai, bi) are distinct.
//   The given input is guaranteed to be a tree and there will be no repeated edges.

mod _minimum_height_trees {
    struct Solution{
        n: i32,
        edges: Vec<Vec<i32>>,
        ans: Vec<i32>,
    }

    use std::collections::VecDeque;
    impl Solution {
        // topological sort, remove leaves layer by layer,
        // then root of MHT is the last remained 1 or 2 nodes
        pub fn find_min_height_trees(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
            if n==1 {
                return vec![0];
            }
            let mut degree = vec![0; n as usize];
            let mut tree = vec![Vec::<usize>::new(); n as usize];
            for e in &edges {
                degree[e[0] as usize] += 1;
                degree[e[1] as usize] += 1;
                tree[e[0] as usize].push(e[1] as usize);
                tree[e[1] as usize].push(e[0] as usize);
            }
            let mut queue = VecDeque::new();
            let mut visited = 0;
            for i in 0..degree.len() {
                if degree[i] == 1 {
                    queue.push_back(i);
                }
            }
            while n-visited > 2 {
                let l = queue.len();
                visited += l as i32;
                for _ in 0..l {
                    let cur = queue.pop_front().unwrap();
                    for &n in &tree[cur] {
                        degree[n] -= 1;
                        if degree[n] == 1 {  // find a new leaf
                            queue.push_back(n);
                        }
                    }
                }
            }
            return queue.iter().map(|&s| s as i32).collect();
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                n: 4,
                edges: vec![[1,0],[1,2],[1,3]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![1],
            },
            Solution {
                n: 6,
                edges: vec![[3,0],[3,1],[3,2],[3,4],[5,4]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![3,4],
            },
            Solution {
                n: 7,
                edges: vec![[0,1],[1,2],[2,3],[3,4],[5,2],[6,5]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![2],
            },
            Solution {
                n: 8,
                edges: vec![[0,1],[1,2],[2,3],[3,4],[5,2],[6,5],[4,7]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![2,3],
            },
            Solution {
                n: 9,
                edges: vec![[0,1],[1,2],[2,3],[3,4],[5,2],[6,5],[4,7],[6,8]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![2],
            },
            Solution {
                n: 1,
                edges: vec![],
                ans: vec![0],
            },
            Solution {
                n: 3,
                edges: vec![[0,1],[0,2]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![0],
            },
            Solution {
                n: 2,
                edges: vec![[0,1]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![0,1],
            }
        ];
        for i in testcases {
            let ans1 = Solution::find_min_height_trees(i.n, i.edges.clone());
            let ans2 = Solution::find_min_height_trees_long_path(i.n, i.edges.clone());
            println!("{:?}, {:?}, {:?}", ans1, ans2, i.ans);
        }
    } 
}
// https://leetcode.com/problems/longest-path-with-different-adjacent-characters/

// You are given a tree (i.e. a connected, undirected graph that has no cycles) rooted at node 0
// consisting of n nodes numbered from 0 to n - 1. The tree is represented by a 0-indexed array
// parent of size n, where parent[i] is the parent of node i. Since node 0 is the root, parent[0] == -1.
// You are also given a string s of length n, where s[i] is the character assigned to node i.
// Return the length of the longest path in the tree such that no pair of adjacent nodes on the
// path have the same character assigned to them.
// Example 1:
//   Input: parent = [-1,0,0,1,1,2], s = "abacbe"
//   Output: 3
//   Explanation: The longest path where each two adjacent nodes have different
//     characters in the tree is the path: 0 -> 1 -> 3. The length of this path is 3, so 3 is returned.
//     It can be proven that there is no longer path that satisfies the conditions.
// Example 2:
//   Input: parent = [-1,0,0,0], s = "aabc"
//   Output: 3
//   Explanation: The longest path where each two adjacent nodes have different
//     characters is the path: 2 -> 0 -> 3. The length of this path is 3, so 3 is returned.
// Constraints:
//   n == parent.length == s.length
//   1 <= n <= 10⁵
//   0 <= parent[i] <= n - 1 for all i >= 1
//   parent[0] == -1
//   parent represents a valid tree.
//   s consists of only lowercase English letters.

extern crate core;

mod _longest_path_with_different_adjacent_characters {
    struct Solution{
        parent: Vec<i32>,
        s: String,
        ans: i32,
    }

    use std::collections::VecDeque;
    impl Solution {
        // bfs: from leaf to root; dfs: from root to leaf also works
        pub fn longest_path(parent: Vec<i32>, s: String) -> i32 {
            let mut max = 0;
            // count number of child
            let mut nc = vec![0; parent.len()];
            let s = s.as_bytes();
            for i in 1..parent.len() {
                nc[parent[i] as usize] += 1;
            }
            // add leaf to queue
            let mut queue = VecDeque::new();
            for i in 0..nc.len() {
                if nc[i] == 0 {
                    queue.push_back(i);
                }
            }
            // visit node in BFS manner: visit a node only if all its children visited
            let mut leng = vec![1; parent.len()]; // the longest path below current node
            while queue.len()>0 {
                let cur = queue.pop_front().unwrap();
                if cur == 0 {
                    max = max.max(leng[cur]);
                    break;
                }
                let p = parent[cur] as usize;
                nc[p] -= 1;
                if nc[p] == 0 {   // all child of p visited
                    queue.push_back(p);
                }
                if s[cur] == s[p] {  // same char, leng not updated
                    continue;
                }
                max = max.max(leng[cur]+leng[p]);  // update max path
                leng[p] = leng[p].max(leng[cur]+1); // update leng of p
            }
            return max;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                parent: vec![-1,0,0,1,1,2],
                s: "abacbe".to_string(),
                ans: 3,
            },
            Solution {
                parent: vec![-1,0,0,0],
                s: "aabc".to_string(),
                ans: 3,
            },
            Solution {
                parent: vec![-1],
                s: "z".to_string(),
                ans: 1,
            },
            Solution {
                parent: vec![-1,0,1],
                s: "abc".to_string(),
                ans: 3,
            },
        ];
        for i in testcases {
            let ans = Solution::longest_path(i.parent, i.s);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
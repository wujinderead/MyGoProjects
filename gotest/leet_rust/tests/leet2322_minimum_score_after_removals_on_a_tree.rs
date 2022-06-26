// https://leetcode.com/problems/minimum-score-after-removals-on-a-tree/

// There is an undirected connected tree with n nodes labeled from 0 to n - 1 and n - 1 edges.
// You are given a 0-indexed integer array nums of length n where nums[i] represents the value of
// the iᵗʰ node. You are also given a 2D integer array edges of length n - 1 where edges[i] =
// [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.
// Remove two distinct edges of the tree to form three connected components. For a pair of removed
// edges, the following steps are defined:
//   Get the XOR of all the values of the nodes for each of the three components respectively.
//   The difference between the largest XOR value and the smallest XOR value is the score of the pair.
// For example, say the three components have the node values: [4,5,7], [1,9], and [3,3,3]. The
// three XOR values are 4 ^ 5 ^ 7 = 6, 1 ^ 9 = 8, and 3 ^ 3 ^ 3 = 3. The largest XOR value is 8 and
// the smallest XOR value is 3. The score is then 8 - 3 = 5.
// Return the minimum score of any possible pair of edge removals on the given tree.+
// Example 1:
//   Input: nums = [1,5,5,4,11], edges = [[0,1],[1,2],[1,3],[3,4]]
//   Output: 9
//   Explanation: The diagram above shows a way to make a pair of removals.
//     - The 1ˢᵗ component has nodes [1,3,4] with values [5,4,11]. Its XOR value is 5 ^ 4 ^ 11 = 10.
//     - The 2ⁿᵈ component has node [0] with value [1]. Its XOR value is 1 = 1.
//     - The 3ʳᵈ component has node [2] with value [5]. Its XOR value is 5 = 5.
//     The score is the difference between the largest and smallest XOR value which is 10 - 1 = 9.
//     It can be shown that no other pair of removals will obtain a smaller score than 9.
// Example 2:
//   Input: nums = [5,5,2,4,4,2], edges = [[0,1],[1,2],[5,2],[4,3],[1,3]]
//   Output: 0
//   Explanation: The diagram above shows a way to make a pair of removals.
//     - The 1ˢᵗ component has nodes [3,4] with values [4,4]. Its XOR value is 4 ^ 4 = 0.
//     - The 2ⁿᵈ component has nodes [1,0] with values [5,5]. Its XOR value is 5 ^ 5 = 0.
//     - The 3ʳᵈ component has nodes [2,5] with values [2,2]. Its XOR value is 2 ^ 2 = 0.
//     The score is the difference between the largest and smallest XOR value which is 0 - 0 = 0.
//     We cannot obtain a smaller score than 0.
// Constraints:
//   n == nums.length
//   3 <= n <= 1000
//   1 <= nums[i] <= 10⁸
//   edges.length == n - 1
//   edges[i].length == 2
//   0 <= ai, bi < n
//   ai != bi
//   edges represents a valid tree.

mod _minimum_score_after_removals_on_a_tree {
    struct Solution{
        nums: Vec<i32>,
        edges: Vec<Vec<i32>>,
        ans: i32,
    }

    // https://leetcode.com/problems/minimum-score-after-removals-on-a-tree/discuss/2198665/Python-3-Explanation-with-pictures-BFS
    impl Solution {
        pub fn minimum_score(mut nums: Vec<i32>, edges: Vec<Vec<i32>>) -> i32 {
            // make the tree
            let mut tree = vec![vec![]; nums.len()];
            for e in &edges {
                tree[e[0] as usize].push(e[1] as usize);
                tree[e[1] as usize].push(e[0] as usize);
            }

            // descendants[i][j] = true, means j is descendant of i
            let mut descendants = vec![vec![false; nums.len()]; nums.len()];

            // dfs
            let mut visited = vec![false; nums.len()];
            for i in 0..nums.len() {
                if !visited[i] {
                    Solution::visit(i, &tree, &mut nums, &mut visited, &mut descendants);
                }
            }
            let all = nums[0];
            let mut ans = 1e9 as i32;

            // for each edge pair, remove these two edges
            for i in 0..edges.len() {
                for j in (i+1)..edges.len() {
                    let (mut a, b) = (edges[i][0] as usize, edges[i][1] as usize);
                    let (mut c, d) = (edges[j][0] as usize, edges[j][1] as usize);
                    if descendants[a][b] == true {  // make a child
                        a = b;
                    }
                    if descendants[c][d] == true {  // make c child
                        c = d;
                    }
                    // three situation
                    let mut x = [0; 3];
                    if descendants[a][c] {  // c is a's descendant
                        x[0] = nums[a] ^ nums[c];
                        x[1] = nums[c];
                        x[2] = all^nums[a];
                    } else if descendants[c][a] {  // a is c's descendant
                        x[0] = nums[c] ^ nums[a];
                        x[1] = nums[a];
                        x[2] = all^nums[c];
                    } else {   // a and c are independent
                        x[0] = nums[a];
                        x[1] = nums[c];
                        x[2] = all^nums[a]^nums[c];
                    }
                    ans = ans.min(x[0].max(x[1].max(x[2])) - x[0].min(x[1].min(x[2])));
                }
            }
            return ans;
        }

        fn visit(ind: usize, tree: &Vec<Vec<usize>>, nums: &mut Vec<i32>, visited: &mut Vec<bool>, descendants: &mut Vec<Vec<bool>>) {
            visited[ind] = true;
            for &next in &tree[ind] {
                if !visited[next as usize] {
                    descendants[ind][next] = true;
                    Solution::visit(next, tree, nums, visited, descendants);
                    nums[ind] ^= nums[next];
                    for j in 0..nums.len() {
                        descendants[ind][j] = descendants[ind][j] || descendants[next][j];
                    }
                }
            }
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,5,5,4,11],
                edges: vec![[0,1],[1,2],[1,3],[3,4]].iter().map(|s| s.to_vec()).collect(),
                ans: 9,
            },
            Solution {
                nums: vec![5,5,2,4,4,2],
                edges: vec![[0,1],[1,2],[5,2],[4,3],[1,3]].iter().map(|s| s.to_vec()).collect(),
                ans: 0,
            },
            Solution {
                nums: vec![2,7,13],
                edges: vec![[1,2],[1,0]].iter().map(|s| s.to_vec()).collect(),
                ans: 11,
            },
        ];
        for i in testcases {
            let ans = Solution::minimum_score(i.nums, i.edges);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/minimum-weighted-subgraph-with-the-required-paths/

// You are given an integer n denoting the number of nodes of a weighted directed graph.
// The nodes are numbered from 0 to n - 1.
// You are also given a 2D integer array edges where edges[i] = [fromi, toi, weighti] denotes
// that there exists a directed edge from fromi to toi with weight weighti.
// Lastly, you are given three distinct integers src1, src2, and dest denoting three distinct
// nodes of the graph.
// Return the minimum weight of a subgraph of the graph such that it is possible to reach dest
// from both src1 and src2 via a set of edges of this subgraph. In case such a subgraph does
// not exist, return -1.
// A subgraph is a graph whose vertices and edges are subsets of the original graph. The weight
// of a subgraph is the sum of weights of its constituent edges.
// Example 1:
//   Input: n = 6, edges = [[0,2,2],[0,5,6],[1,0,3],[1,4,5],[2,1,1],[2,3,3],[2,3,4],[3,4,2],[4,5,1]], src1 = 0, src2 = 1, dest = 5
//   Output: 9
//   Explanation:
//     The above figure represents the input graph.
//     The blue edges represent one of the subgraphs that yield the optimal answer.
//     Note that the subgraph [[1,0,3],[0,5,6]] also yields the optimal answer.
//     It is not possible to get a subgraph with less weight satisfying all the constraints.
// Example 2:
//   Input: n = 3, edges = [[0,1,1],[2,1,1]], src1 = 0, src2 = 1, dest = 2
//   Output: -1
//   Explanation:
//     The above figure represents the input graph.
//     It can be seen that there does not exist any path from node 1 to node 2,
//     hence there are no subgraphs satisfying all the constraints.
// Constraints:
//   3 <= n <= 10⁵
//   0 <= edges.length <= 10⁵
//   edges[i].length == 3
//   0 <= fromi, toi, src1, src2, dest <= n - 1
//   fromi != toi
//   src1, src2, and dest are pairwise distinct.
//   1 <= weight[i] <= 10⁵

mod _minimum_weighted_subgraph_with_the_required_paths {
    struct Solution{
        n: i32,
        edges: Vec<Vec<i32>>,
        src1: i32,
        src2: i32,
        dest: i32,
        ans: i64,
    }

    use std::collections::binary_heap::BinaryHeap;
    use std::cmp::Ordering;
    impl Solution {
        // three dijkstra processes, let any node x be the mediate node:
        // find the shortest path from src1 to x (run dijkstra with src1 as source node);
        // find the shortest path from src2 to x (run dijkstra with src2 as source node);
        // find the shortest path from x to dest (run dijkstra with dest as source node on edge-reversed graph);
        // then we need find the node x that make (shortest1[x] + shortest2[x] + shortest3[x]) minimal
        pub fn minimum_weight(n: i32, edges: Vec<Vec<i32>>, src1: i32, src2: i32, dest: i32) -> i64 {
            // make a graph
            let nn = n as usize;
            let mut graph = vec![Vec::<[i32; 2]>::new(); nn];
            let mut reverse = vec![Vec::<[i32; 2]>::new(); nn];
            for e in &edges {
                graph[e[0] as usize].push([e[1], e[2]]);
                reverse[e[1] as usize].push([e[0], e[2]]);
            }

            // run dijkstra
            const MAX: i64 = 1e11 as i64;
            let shortest1 = Solution::dijkstra(&graph, src1 as usize);
            let shortest2 = Solution::dijkstra(&graph, src2 as usize);
            let shortest3 = Solution::dijkstra(&reverse, dest as usize);
            if shortest1[dest as usize] == MAX || shortest2[dest as usize] == MAX {  // src1 or src2 can't reach dest
                return -1;
            }
            let mut min = MAX;
            for (i, &_) in shortest1.iter().enumerate() {
                min = min.min(shortest1[i] + shortest2[i] + shortest3[i]);
            }
            return min;
        }

        fn dijkstra(graph: &Vec<Vec<[i32; 2]>>, src: usize) -> Vec<i64> {
            // the state to be stored in heap for dijkstra
            #[derive(Copy, Clone, Eq, PartialEq)]
            struct State(usize, i64);  // node number, cumulative dist
            impl PartialOrd for State {
                fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
                    // default BinaryHeap is max-heap, reverse other and self to make a min-heap
                    Some(other.1.cmp(&self.1))
                }
            }
            impl Ord for State {
                fn cmp(&self, other: &Self) -> Ordering {
                    other.1.cmp(&self.1)
                }
            }

            // dijkstra to find shortest path from src to other nodes
            const MAX: i64 = 1e11 as i64;
            let mut shortest = vec![MAX; graph.len()];
            let mut heap = BinaryHeap::new();
            heap.push(State(src, 0));
            while heap.len() > 0 {
                let pop = heap.pop().unwrap();
                let (cur, dist) = (pop.0, pop.1);
                if shortest[cur] == MAX {  // include current node in dijkstra set
                    shortest[cur] = dist;
                    for next in &graph[cur] {
                        let (next_node, edge_len) = (next[0] as usize, next[1] as i64);
                        if shortest[next_node] == MAX {  // if neighbor not in set, add neighbor to heap
                            heap.push(State(next_node, dist+edge_len));
                        }
                    }
                }
            }
            return shortest;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                n: 6,
                edges: [[0,2,2],[0,5,6],[1,0,3],[1,4,5],[2,1,1],[2,3,3],[2,3,4],[3,4,2],[4,5,1]].iter().map(|v| v.to_vec()).collect(),
                src1: 1,
                src2: 0,
                dest: 5,
                ans: 9,
            },
            Solution {
                n: 3,
                edges: [[0,1,1],[2,1,1]].iter().map(|v| v.to_vec()).collect(),
                src1: 0,
                src2: 1,
                dest: 2,
                ans: -1,
            },
            Solution {
                n: 8,
                edges: [[4,7,24],[1,3,30],[4,0,31],[1,2,31],[1,5,18],[1,6,19],[4,6,25],[5,6,32],[0,6,50]].iter().map(|v| v.to_vec()).collect(),
                src1: 4,
                src2: 1,
                dest: 6,
                ans: 44,
            },
            Solution {
                n: 6,
                edges: [[0,2,10],[0,4,2],[1,4,2],[1,3,10],[3,5,10],[4,5,20],[2,5,10]].iter().map(|v| v.to_vec()).collect(),
                src1: 0,
                src2: 1,
                dest: 5,
                ans: 24,
            },
        ];
        for i in testcases {
            let ans = Solution::minimum_weight(i.n, i.edges, i.src1, i.src2, i.dest);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
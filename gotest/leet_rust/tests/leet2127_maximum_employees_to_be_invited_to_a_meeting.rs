// https://leetcode.com/problems/maximum-employees-to-be-invited-to-a-meeting/

// A company is organizing a meeting and has a list of n employees, waiting to be invited.
// They have arranged for a large circular table, capable of seating any number of employees.
// The employees are numbered from 0 to n - 1. Each employee has a favorite person and they will
// attend the meeting only if they can sit next to their favorite person at the table.
// The favorite person of an employee is not themself.
// Given a 0-indexed integer array favorite, where favorite[i] denotes the favorite person of
// the iᵗʰ employee, return the maximum number of employees that can be invited to the meeting.
// Example 1:
//   Input: favorite = [2,2,1,2]
//   Output: 3
//   Explanation:
//     The above figure shows how the company can invite employees 0, 1, and 2, and seat them at the round table.
//     All employees cannot be invited because employee 2 cannot sit beside employees 0, 1, and 3, simultaneously.
//     Note that the company can also invite employees 1, 2, and 3, and give them their desired seats.
//     The maximum number of employees that can be invited to the meeting is 3.
// Example 2:
//   Input: favorite = [1,2,0]
//   Output: 3
//   Explanation:
//     Each employee is the favorite person of at least one other employee, and the
//     only way the company can invite them is if they invite every employee.
//     The seating arrangement will be the same as that in the figure given in
//     example 1:
//     - Employee 0 will sit between employees 2 and 1.
//     - Employee 1 will sit between employees 0 and 2.
//     - Employee 2 will sit between employees 1 and 0.
//     The maximum number of employees that can be invited to the meeting is 3.
// Example 3:
//   Input: favorite = [3,0,1,4,1]
//   Output: 4
//   Explanation:
//     The above figure shows how the company will invite employees 0, 1, 3, and 4, and seat them at the round table.
//     Employee 2 cannot be invited because the two spots next to their favorite employee 0 are taken.
//     So the company leaves them out of the meeting.
//     The maximum number of employees that can be invited to the meeting is 4.
// Constraints:
//   n == favorite.length
//   2 <= n <= 10⁵
//   0 <= favorite[i] <= n - 1
//   favorite[i] != i

// we deem each person i as a vertex, and i->favorite[i] as a directed edge in a graph.
// then according to the description, for each i, there may be many in-edges, but there will
// only be one out-edge, so in each connected components, there will be exactly one circle.
// if the circle size > 2, then the valid vertices are all vertices in this circle.
// if the circle size = 2, then we want the deepest tree that connected to both side.
//   case1:
//     ABCD is a circle, they are valid; EF invalid
//         A -> B <- E
//         ^    |
//         |    v
//     F-> D <- C
//   case2:
//     AB is circle, so for the trees that connected to A (ACDG) and B (BEFH), we want the deepest.
//     so DCABEF is the answer for the left component.
//     and case2 can union with each other. so the final answer is DCABEF + IJK
//                      ->                                <-
//          D -> C- > A <- B <- E <- F             I -> J -> K
//                    ^         ^
//                    |         |
//                    G         H
mod _maximum_employees_to_be_invited_to_a_meeting {
    struct Solution{
        favorite: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn maximum_invitations(favorite: Vec<i32>) -> i32 {
            let mut max_circle = 0;
            let mut sum_pairs = 0;

            // make the graph, use favorite[i]->i as edge, to facilitate computation
            let mut graph: Vec<Vec<usize>> = vec![vec![]; favorite.len()];
            let mut pairs = Vec::new();
            let mut i = 0;
            while i<favorite.len() {
                let fi = favorite[i] as usize;
                if favorite[fi] as usize == i && i < fi {  // find pairs
                   pairs.push([i, fi])
                }
                graph[fi].push(i);
                i += 1;
            }

            let mut visited = vec![false; favorite.len()];
            // for case2:
            for p in pairs {
                // visit the tree rooted at p[0], find the height
                let d0 = visit_tree(p[0], p[1], &graph, &mut visited);
                // visit the tree rooted at p[1], find the height
                let d1 = visit_tree(p[1], p[0], &graph, &mut visited);
                sum_pairs += d0+d1;
            }

            // for case1:
            i = 0;
            let mut depth = vec![0; favorite.len()];
            while i < visited.len() {
                if !visited[i] {
                    // find the circle size for each connected components
                    let circle_size = visit_circle(i, i, &mut depth, &graph, &mut visited);
                    max_circle = max_circle.max(circle_size);
                }
                i += 1;
            }
            return max_circle.max(sum_pairs);
        }
    }

    fn visit_tree(p0: usize, p1: usize, graph: &Vec<Vec<usize>>, visited: &mut Vec<bool>) -> i32 {
        visited[p0] = true;
        let mut max = 0;
        for i in &graph[p0] {
            if *i != p1 {
                max = max.max(visit_tree(*i, p1, graph, visited))
            }
        }
        return max+1;
    }

    fn visit_circle(i: usize, prev: usize, depth: &mut Vec<i32>, graph: &Vec<Vec<usize>>, visited: &mut Vec<bool>) -> i32 {
        visited[i] = true;       // mark as visited
        depth[i] = depth[prev]+1;  // the depth for current vertex
        let mut ret = 0;
        for next in &graph[i] {
            if visited[*next] {
                if depth[*next] == 0 {  // visited vertex with depth=0, not visited in current round
                    continue
                }
                // found a visited node in current round, then circle size is the depth difference
                ret = depth[i] - depth[*next] + 1;
                break;
            }
            // visit unvisited child, if can find circle, break
            ret = visit_circle(*next, i, depth, graph, visited);
            if ret > 0 {
                break;
            }
        }
        depth[i] = 0;
        return ret;
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                favorite: vec![2,2,1,2],
                ans: 3,
            },
            Solution {
                favorite: vec![2,2,1,2,5,4,4,6,6,5,9,5,4],
                ans: 9,
            },
            Solution {
                favorite: vec![1,2,0],
                ans: 3,
            },
            Solution {
                favorite: vec![3,0,1,4,1],
                ans: 4,
            },
            Solution {
                favorite: vec![2,2,1,2,5,4,4,6,6,5,9,5,4,16,13,14,15,14,14,15,16],
                ans: 9,
            },
            Solution {
                favorite: vec![1,2,3,4,5,6,3,8,9,10,11,8],
                ans: 4,
            },
            Solution {
                favorite: vec![12,10,16,18,9,2,20,4,1,0,8,18,4,6,1,0,3,15,6,2,17],
                ans: 4,
            }
        ];
        for i in testcases {
            let ans = Solution::maximum_invitations(i.favorite);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/total-cost-to-hire-k-workers/

// You are given a 0-indexed integer array costs where costs[i] is the cost of hiring the iᵗʰ worker.
// You are also given two integers k and candidates. We want to hire exactly k workers according to
// the following rules:
//   You will run k sessions and hire exactly one worker in each session.
//   In each hiring session, choose the worker with the lowest cost from either the first candidates
//     workers or the last candidates workers. Break the tie by the smallest index.
// For example, if costs = [3,2,7,7,1,2] and candidates = 2, then in the first hiring session, we will
// choose the 4ᵗʰ worker because they have the lowest cost [3,2,7,7,1,2].
// In the second hiring session, we will choose 1ˢᵗ worker because they have the same lowest cost as
// 4ᵗʰ worker but they have the smallest index [3,2,7,7,2]. Please note that the indexing may be changed
// in the process.
// If there are fewer than candidates workers remaining, choose the worker with the lowest cost among
// them. Break the tie by the smallest index. A worker can only be chosen once.
// Return the total cost to hire exactly k workers.
// Example 1:
//   Input: costs = [17,12,10,2,7,2,11,20,8], k = 3, candidates = 4
//   Output: 11
//   Explanation: We hire 3 workers in total. The total cost is initially 0.
//     - In the first hiring round we choose the worker from [17,12,10,2,7,2,11,20,8].
//     The lowest cost is 2, and we break the tie by the smallest index, which is 3.
//     The total cost = 0 + 2 = 2.
//     - In the second hiring round we choose the worker from [17,12,10,7,2,11,20,8].
//     The lowest cost is 2 (index 4). The total cost = 2 + 2 = 4.
//     - In the third hiring round we choose the worker from [17,12,10,7,11,20,8]. The
//     lowest cost is 7 (index 3). The total cost = 4 + 7 = 11. Notice that the worker
//     with index 3 was common in the first and last four workers.
//     The total hiring cost is 11.
// Example 2:
//   Input: costs = [1,2,4,1], k = 3, candidates = 3
//   Output: 4
//   Explanation: We hire 3 workers in total. The total cost is initially 0.
//     - In the first hiring round we choose the worker from [1,2,4,1]. The lowest
//     cost is 1, and we break the tie by the smallest index, which is 0. The total cost =
//     0 + 1 = 1. Notice that workers with index 1 and 2 are common in the first and
//     last 3 workers.
//     - In the second hiring round we choose the worker from [2,4,1]. The lowest cost
//     is 1 (index 2). The total cost = 1 + 1 = 2.
//     - In the third hiring round there are less than three candidates. We choose the
//     worker from the remaining workers [2,4]. The lowest cost is 2 (index 0). The
//     total cost = 2 + 2 = 4.
//     The total hiring cost is 4.
// Constraints:
//   1 <= costs.length <= 10⁵
//   1 <= costs[i] <= 10⁵
//   1 <= k, candidates <= costs.length

mod _total_cost_to_hire_k_workers {
    struct Solution{
        costs: Vec<i32>,
        k: i32,
        candidates: i32,
        ans: i64,
    }

    use std::cmp::Reverse;
    use std::collections::BinaryHeap;
    impl Solution {
        pub fn total_cost(costs: Vec<i32>, mut k: i32, candidates: i32) -> i64 {
            let mut left = BinaryHeap::new();
            let mut right = BinaryHeap::new();
            let (mut i, mut j, mut cost) = (0, costs.len()-1, 0 as i64);
            let candidates = candidates as usize;
            while k>0 {
                while left.len()<candidates && i<=j {
                    left.push(Reverse(costs[i]));
                    i += 1;
                }
                while right.len()<candidates && j>=i {
                    right.push(Reverse(costs[j]));
                    j -= 1;
                }
                let pa = if left.len() == 0 { i32::MAX } else { left.peek().unwrap().0 };
                let pb = if right.len() == 0 { i32::MAX } else { right.peek().unwrap().0 };
                if pa <= pb {
                    left.pop();
                    cost += pa as i64;
                } else {
                    right.pop();
                    cost += pb as i64;
                }
                k -= 1;
            }
            return cost;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                costs: vec![17, 12, 10, 2, 7, 2, 11, 20, 8],
                k: 3,
                candidates: 4,
                ans: 11,
            },
            Solution {
                costs: vec![1, 17, 12, 10, 2, 7, 2, 11, 20, 8, 3],
                k: 3,
                candidates: 4,
                ans: 6,
            },
            Solution {
                costs: vec![1, 2, 4, 1],
                k: 3,
                candidates: 3,
                ans: 4,
            },
        ];
        for i in testcases {
            let ans = Solution::total_cost(i.costs, i.k, i.candidates);
            println!("{}, {}", ans, i.ans);
        }
    }
}
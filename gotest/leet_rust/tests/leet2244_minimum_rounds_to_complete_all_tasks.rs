// https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/

// You are given a 0-indexed integer array tasks, where tasks[i] represents the difficulty level
// of a task. In each round, you can complete either 2 or 3 tasks of the same difficulty level.
// Return the minimum rounds required to complete all the tasks, or -1 if it is not possible to
// complete all the tasks.
// Example 1:
//   Input: tasks = [2,2,3,3,2,4,4,4,4,4]
//   Output: 4
//   Explanation: To complete all the tasks, a possible plan is:
//     - In the first round, you complete 3 tasks of difficulty level 2.
//     - In the second round, you complete 2 tasks of difficulty level 3.
//     - In the third round, you complete 3 tasks of difficulty level 4.
//     - In the fourth round, you complete 2 tasks of difficulty level 4.
//     It can be shown that all the tasks cannot be completed in fewer than 4 rounds, so the answer is 4.
// Example 2:
//   Input: tasks = [2,3,3]
//   Output: -1
//   Explanation: There is only 1 task of difficulty level 2, but in each round,
//     you can only complete either 2 or 3 tasks of the same difficulty level. Hence,
//     you cannot complete all the tasks, and the answer is -1.
// Constraints:
//   1 <= tasks.length <= 10⁵
//   1 <= tasks[i] <= 10⁹

mod _minimum_rounds_to_complete_all_tasks {
    struct Solution{
        tasks: Vec<i32>,
        ans: i32,
    }

    use std::collections::HashMap;
    impl Solution {
        pub fn minimum_rounds(tasks: Vec<i32>) -> i32 {
            let mut map = HashMap::new();
            for t in &tasks {
                *map.entry(t).or_insert(0) += 1;
            }
            let mut sum = 0;
            for (&_k, &v) in map.iter() {
                if v==1 {
                    return -1;
                }
                sum += (v-1)/3+1;
            }
            return sum;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                tasks: vec![2,2,3,3,2,4,4,4,4,4],
                ans: 4,
            },
            Solution {
                tasks: vec![2,3,2],
                ans: -1,
            },
        ];
        for i in testcases {
            let ans = Solution::minimum_rounds(i.tasks);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/last-stone-weight-ii/

// You are given an array of integers stones where stones[i] is the weight of the iᵗʰ stone.
// We are playing a game with the stones. On each turn, we choose any two stones and smash them
// together. Suppose the stones have weights x and y with x <= y. The result of this smash is:
//   If x == y, both stones are destroyed, and
//   If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
// At the end of the game, there is at most one stone left.
// Return the smallest possible weight of the left stone. If there are no stones left, return 0.
// Example 1:
//   Input: stones = [2,7,4,1,8,1]
//   Output: 1
//   Explanation:
//     We can combine 2 and 4 to get 2, so the array converts to [2,7,1,8,1] then,
//     we can combine 7 and 8 to get 1, so the array converts to [2,1,1,1] then,
//     we can combine 2 and 1 to get 1, so the array converts to [1,1,1] then,
//     we can combine 1 and 1 to get 0, so the array converts to [1], then that's the optimal value.
// Example 2:
//   Input: stones = [31,26,33,21,40]
//   Output: 5
// Constraints:
//   1 <= stones.length <= 30
//   1 <= stones[i] <= 100

mod _last_stone_weight_i_i {
    struct Solution{
        stones: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        // the problem is to partition the stones to two sets, and minimise the difference of sum.
        // so it's actually a knapsack problem, to find a subset that most equal to sum/2
        // O(n*sum)
        pub fn last_stone_weight_ii(stones: Vec<i32>) -> i32 {
            let stones = stones.iter().map(|&s| s as usize).collect::<Vec<_>>();
            let s = stones.iter().sum::<usize>();
            // dp[i]=true, we can use stones[:j] to get sum i
            let mut dp = vec![false; s/2+1];
            dp[0] = true;
            for i in 0..stones.len() {
                for j in (stones[i]..=s/2).into_iter().rev() {
                    dp[j] = dp[j] || dp[j-stones[i]];
                }
            }
            let mut i = s/2;
            while dp[i] == false {
                i -= 1;
            }
            return (s-i-i) as i32;
        }

        // use hash set: O(2^n)
        pub fn last_stone_weight_ii_set(stones: Vec<i32>) -> i32 {
            use std::collections::HashSet;
            let sum = stones.iter().sum::<i32>();
            let mut set = HashSet::new();
            set.insert(0);
            for i in 0..stones.len() {
                let mut to_add = Vec::new();
                for &s in set.iter() {
                    if s+stones[i] <= sum/2 {
                        to_add.push(s+stones[i]);
                    }
                }
                for s in to_add {
                    set.insert(s);
                }
            }
            let mut i = sum/2;
            while !set.contains(&i) {
                i -= 1;
            }
            return sum-i-i;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                stones: vec![2,7,4,1,8,1],
                ans: 1,
            },
            Solution {
                stones: vec![31,26,33,21,40],
                ans: 5,
            },
            Solution {
                stones: vec![91],
                ans: 91,
            },
        ];
        for i in &testcases {
            let ans = Solution::last_stone_weight_ii(i.stones.clone());
            println!("{}, {}", ans, i.ans);
        }
        for i in &testcases {
            let ans = Solution::last_stone_weight_ii_set(i.stones.clone());
            println!("{}, {}", ans, i.ans);
        }
    } 
}
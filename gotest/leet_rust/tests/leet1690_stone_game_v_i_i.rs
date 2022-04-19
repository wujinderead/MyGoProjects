// https://leetcode.com/problems/stone-game-vii/

// Alice and Bob take turns playing a game, with Alice starting first.
// There are n stones arranged in a row. On each player's turn, they can remove either the leftmost
// stone or the rightmost stone from the row and receive points equal to the sum of the remaining
// stones' values in the row. The winner is the one with the higher score when there are no stones
// left to remove.
// Bob found that he will always lose this game (poor Bob, he always loses), so he decided to
// minimize the score's difference. Alice's goal is to maximize the difference in the score.
// Given an array of integers stones where stones[i] represents the value of the iᵗʰ stone from the
// left, return the difference in Alice and Bob's score if they both play optimally.
// Example 1:
//   Input: stones = [5,3,1,4,2]
//   Output: 6
//   Explanation:
//     - Alice removes 2 and gets 5 + 3 + 1 + 4 = 13 points. Alice = 13, Bob = 0,stones = [5,3,1,4].
//     - Bob removes 5 and gets 3 + 1 + 4 = 8 points. Alice = 13, Bob = 8, stones = [3,1,4].
//     - Alice removes 3 and gets 1 + 4 = 5 points. Alice = 18, Bob = 8, stones = [1,4].
//     - Bob removes 1 and gets 4 points. Alice = 18, Bob = 12, stones = [4].
//     - Alice removes 4 and gets 0 points. Alice = 18, Bob = 12, stones = [].
//     The score difference is 18 - 12 = 6.
// Example 2:
//   Input: stones = [7,90,5,1,100,10,10,2]
//   Output: 122
// Constraints:
//   n == stones.length
//   2 <= n <= 1000
//   1 <= stones[i] <= 1000

mod _stone_game_v_i_i {
    struct Solution{
        stones: Vec<i32>,
        ans: i32,
    }

    use std::collections::HashMap;
    impl Solution {
        pub fn stone_game_vii(stones: Vec<i32>) -> i32 {
            // get prefix
            let mut prefix = vec![0; stones.len()+1];
            for i in 1..=stones.len() {
                prefix[i] = prefix[i-1]+stones[i-1]; // sum(stones[i...j])=prefix[j+1]-prefix[i]
            }
            // dp
            let mut map = HashMap::<(usize, usize), i32>::new();
            for i in 0..stones.len() { // initial: dp(i,i)=0
                map.insert((i, i), 0);
            }
            for diff in 1..stones.len() {
                let mut i = 0;
                while i+diff<stones.len() {
                    let j = i+diff;
                    // dp(i,j) = max(sum(stones[i+1...j])-dp(i+1,j), sum(stones[i...j-1])-dp(i,j-1))
                    let cand1 = prefix[j+1]-prefix[i+1]-*map.get(&(i+1,j)).unwrap();
                    let cand2 = prefix[j]-prefix[i]-*map.get(&(i,j-1)).unwrap();
                    map.insert((i,j), cand1.max(cand2));
                    i += 1;
                }
            }
            return *map.get(&(0,stones.len()-1)).unwrap();
        }

        pub fn stone_game_vii_1d(stones: Vec<i32>) -> i32 {
            // get prefix
            let mut prefix = vec![0; stones.len()+1];
            for i in 1..=stones.len() {
                prefix[i] = prefix[i-1]+stones[i-1]; // sum(stones[i...j])=prefix[j+1]-prefix[i]
            }
            // dp
            let mut dp = vec![0; stones.len()]; // initial: dp(i,i)=0
            for diff in 1..stones.len() {
                let mut i = 0;
                while i+diff<stones.len() {
                    let j = i+diff;
                    // dp(i,j) = max(sum(stones[i+1...j])-dp(i+1,j), sum(stones[i...j-1])-dp(i,j-1))
                    let cand2 = prefix[j]-prefix[i]-dp[i];
                    let cand1 = prefix[j+1]-prefix[i+1]-dp[i+1];
                    dp[i] = cand1.max(cand2);
                    i += 1;
                }
            }
            return dp[0];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                stones: vec![5,3,1,4,2],
                ans: 6,
            },
            Solution {
                stones: vec![7,90,5,1,100,10,10,2],
                ans: 122,
            },
        ];
        for i in &testcases {
            let ans = Solution::stone_game_vii(i.stones.clone());
            println!("{}, {}", ans, i.ans);
        }
        for i in &testcases {
            let ans = Solution::stone_game_vii_1d(i.stones.clone());
            println!("{}, {}", ans, i.ans);
        }
    } 
}
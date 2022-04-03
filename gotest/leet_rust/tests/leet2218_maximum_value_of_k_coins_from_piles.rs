// https://leetcode.com/problems/maximum-value-of-k-coins-from-piles/

// There are n piles of coins on a table. Each pile consists of a positive number of coins
// of assorted denominations.
// In one move, you can choose any coin on top of any pile, remove it, and add it to your wallet.
// Given a list piles, where piles[i] is a list of integers denoting the composition of the
// iᵗʰ pile from top to bottom, and a positive integer k, return the maximum total value of
// coins you can have in your wallet if you choose exactly k coins optimally.
// Example 1:
//   Input: piles = [[1,100,3],[7,8,9]], k = 2
//   Output: 101
//   Explanation:
//     The above diagram shows the different ways we can choose k coins.
//     The maximum total we can obtain is 101.
// Example 2:
//   Input: piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7
//   Output: 706
//   Explanation:
//     The maximum total can be obtained if we choose all coins from the last pile.
// Constraints:
//   n == piles.length
//   1 <= n <= 1000
//   1 <= piles[i][j] <= 10⁵
//   1 <= k <= sum(piles[i].length) <= 2000

mod _maximum_value_of_k_coins_from_piles {
    struct Solution{
        piles: Vec<Vec<i32>>,
        k: i32,
        ans: i32,
    }

    impl Solution {
        // dp: dp[i][k] max benefit to choose k coins in piles[...i]
        pub fn max_value_of_coins(mut piles: Vec<Vec<i32>>, k: i32) -> i32 {
            // accumulate each pile
            for i in 0..piles.len() {
                for j in 1..piles[i].len() {
                    piles[i][j] += piles[i][j-1];  // new piles[i][j] = sum(piles[i][0...j])
                }
            }
            let kk = k as usize;
            let old = &mut vec![0; kk+1]; // we can choose max k coins
            let new = &mut vec![0; kk+1];
            // dp[0]: for piles 0
            for i in 1..=kk.min(piles[0].len()) {
                old[i] = piles[0][i-1];
            }
            // dp: for piles i, dp[i][k] = dp[i-1][k-x] + sum(piles[i][..x])
            let mut sum = piles[0].len();
            for i in 1..piles.len() {
                sum += piles[i].len();  // sum of coins in piles[0...i]
                let maxk = kk.min(sum); // we can choose maxk coins from piles[0...i]
                for j in 1..=maxk {  // dp[i][j] choose j coins in piles[0...i]
                    new[j] = old[j];  // no choose piles[i]
                    for x in 1..=j.min(piles[i].len()) {  // choose x coins in piles[i]
                        new[j] = new[j].max(piles[i][x-1] + old[j-x])  // choose j-x coins in piles[0...i-1]
                    }
                }
                std::mem::swap(old, new);
            }
            return old[kk];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                piles: vec![vec![1,100,3],vec![7,8,9]],
                k: 2,
                ans: 101,
            },
            Solution {
                piles: vec![vec![100],vec![100],vec![100],vec![100],vec![100],vec![100],vec![1,1,1,1,1,1,700]],
                k: 7,
                ans: 706,
            }
        ];
        for i in testcases {
            let ans = Solution::max_value_of_coins(i.piles, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
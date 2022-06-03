// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/

// You are given an integer array prices where prices[i] is the price of a given stock on the
// iᵗʰ day, and an integer k.
// Find the maximum profit you can achieve. You may complete at most k transactions.
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock
// before you buy again).
// Example 1:
//   Input: k = 2, prices = [2,4,1]
//   Output: 2
//   Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
// Example 2:
//   Input: k = 2, prices = [3,2,6,5,0,3]
//   Output: 7
//   Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
//     Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
// Constraints:
//   0 <= k <= 100
//   0 <= prices.length <= 1000
//   0 <= prices[i] <= 1000

mod _best_time_to_buy_and_sell_stock_i_v {
    struct Solution{
        k: i32,
        prices: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn max_profit(k: i32, prices: Vec<i32>) -> i32 {
            if prices.len() < 2 {
                return 0;
            }
            let k = (prices.len()/2).min(k as usize);
            let mut buy = vec![100000; k];
            let mut sell = vec![0; k];
            for i in 0..prices.len() {
                buy[0] = buy[0].min(prices[i]);
                sell[0] = sell[0].max(prices[i]-buy[0]);
                for j in 1..k {
                    buy[j] = buy[j].min(prices[i]-sell[j-1]);
                    sell[j] = sell[j].max(prices[i]-buy[j]);
                }
            }
            return sell[k-1];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                k: 2,
                prices: vec![2,4,1],
                ans: 2,
            },
            Solution {
                k: 2,
                prices: vec![3,2,6,5,0,3],
                ans: 7,
            },
            Solution {
                k: 1,
                prices: vec![3,2,6,5,0,3],
                ans: 4,
            },
            Solution {
                k: 3,
                prices: vec![3,2,6,5,0,3],
                ans: 7,
            },
        ];
        for i in testcases {
            let ans = Solution::max_profit(i.k, i.prices);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

// You are given an array prices where prices[i] is the price of a given stock on the iᵗʰ day.
// Find the maximum profit you can achieve. You may complete as many transactions as you like
// (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:
// After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).
// Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock
// before you buy again).
// Example 1:
//   Input: prices = [1,2,3,0,2]
//   Output: 3
//   Explanation: transactions = [buy, sell, cooldown, buy, sell]
// Example 2:
//   Input: prices = [1]
//   Output: 0
// Constraints:
//   1 <= prices.length <= 5000
//   0 <= prices[i] <= 1000

mod _best_time_to_buy_and_sell_stock_with_cooldown {
    struct Solution{
        prices: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn max_profit(prices: Vec<i32>) -> i32 {
            if prices.len() < 2 {
                return 0;
            }
            let mut buy = -(prices[0].min(prices[1]));
            let mut sell0 = 0;
            let mut sell1 = 0.max(prices[1]+buy);
            for i in 2..prices.len() {
                // b[x] must depend on s[x-2], but s[x] can be the max of s[x-1] or x[x-2]
                if i%2 == 0 {
                    // b[2]               s[0]
                    buy = buy.max(sell0-prices[i]);
                    // s[2]  s[0]                                       s[1]
                    sell0 = sell0.max(prices[i]+buy).max(sell1);
                } else {
                    // b[3]               s[1]
                    buy = buy.max(sell1-prices[i]);
                    // s[3]  s[1]                                       s[2]
                    sell1 = sell1.max(prices[i]+buy).max(sell0);
                }
            }
            return sell0.max(sell1);
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                prices: vec![1,2,3,0,2],
                ans: 3,
            },
            Solution {
                prices: vec![1],
                ans: 0,
            },
            Solution {
                prices: vec![2,1,4],
                ans: 3,
            },
            Solution {
                prices: vec![6,1,6,4,3,0,2],
                ans: 7,
            }
        ];
        for i in testcases {
            let ans = Solution::max_profit(i.prices);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

// You are given an integer array prices where prices[i] is the price of a given stock on the iᵗʰ day.
// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of
// the stock at any time. However, you can buy it then immediately sell it on the same day.
// Find and return the maximum profit you can achieve.
// Example 1:
//   Input: prices = [7,1,5,3,6,4]
//   Output: 7
//   Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
//     Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
//     Total profit is 4 + 3 = 7.
// Example 2:
//   Input: prices = [1,2,3,4,5]
//   Output: 4
//   Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
//     Total profit is 4.
// Example 3:
//   Input: prices = [7,6,4,3,1]
//   Output: 0
//   Explanation: There is no way to make a positive profit, so we never buy the
//     stock to achieve the maximum profit of 0.
// Constraints:
//   1 <= prices.length <= 3 * 10⁴
//   0 <= prices[i] <= 10⁴

mod _best_time_to_buy_and_sell_stock_i_i {
    struct Solution{
        prices: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        // sum every increase
        pub fn max_profit1(prices: Vec<i32>) -> i32 {
            let mut sum = 0;
            let mut prev = prices[0];
            for i in 1..prices.len() {
                if prices[i] >= prev {
                    sum += prices[i]-prev;
                }
                prev = prices[i];
            }
            return sum;
        }

        // dp solution
        // for prices[0...i]
        // buy is the max money you can have when last action is buy
        // sell is the max money you can have when last action is sell
        pub fn max_profit2(prices: Vec<i32>) -> i32 {
            let (mut buy, mut sell) = (-prices[0], 0);
            for i in 1..prices.len() {
                sell = sell.max(prices[i]+buy);
                buy = buy.max(sell-prices[i]);
            }
            return sell;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                prices: vec![7,1,5,3,6,4],
                ans: 7,
            },
            Solution {
                prices: vec![1,2,3,4,5],
                ans: 4,
            },
            Solution {
                prices: vec![7,6,4,3,1],
                ans: 0,
            },
        ];
        for i in testcases {
            let ans1 = Solution::max_profit1(i.prices.clone());
            let ans2 = Solution::max_profit2(i.prices.clone());
            println!("{}, {}, {}", ans1, ans2, i.ans);
        }
    } 
}
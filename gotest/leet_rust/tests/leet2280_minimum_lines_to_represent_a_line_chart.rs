// https://leetcode.com/problems/minimum-lines-to-represent-a-line-chart/

// You are given a 2D integer array stockPrices where stockPrices[i] = [dayi, pricei] indicates the
// price of the stock on day dayi is pricei. A line chart is created from the array by plotting the
// points on an XY plane with the X-axis representing the day and the Y-axis representing the price
// and connecting adjacent points. One such example is shown below:
// Return the minimum number of lines needed to represent the line chart.
// Example 1:
//   Input: stockPrices = [[1,7],[2,6],[3,5],[4,4],[5,4],[6,3],[7,2],[8,1]]
//   Output: 3
//   Explanation:
//     The diagram above represents the input, with the X-axis representing the day and Y-axis representing the price.
//     The following 3 lines can be drawn to represent the line chart:
//     - Line 1 (in red) from (1,7) to (4,4) passing through (1,7), (2,6), (3,5), and (4,4).
//     - Line 2 (in blue) from (4,4) to (5,4).
//     - Line 3 (in green) from (5,4) to (8,1) passing through (5,4), (6,3), (7,2), and (8,1).
//     It can be shown that it is not possible to represent the line chart using less than 3 lines.
// Example 2:
//   Input: stockPrices = [[3,4],[1,2],[7,8],[2,3]]
//   Output: 1
//   Explanation:
//     As shown in the diagram above, the line chart can be represented with a single line.
// Constraints:
//   1 <= stockPrices.length <= 10⁵
//   stockPrices[i].length == 2
//   1 <= dayi, pricei <= 10⁹
//   All dayi are distinct.

mod _minimum_lines_to_represent_a_line_chart {
    struct Solution{
        stock_prices: Vec<Vec<i32>>,
        ans: i32,
    }

    impl Solution {
        pub fn minimum_lines(mut stock_prices: Vec<Vec<i32>>) -> i32 {
            if stock_prices.len() == 1 {
                return 0;
            }
            stock_prices.sort_by_key(|s| s[0]);
            let mut ans = 1;
            let mut dx = stock_prices[1][0]-stock_prices[0][0];
            let mut dy = stock_prices[1][1]-stock_prices[0][1];
            for i in 2..stock_prices.len() {
                let cx = stock_prices[i][0]-stock_prices[i-1][0];
                let cy = stock_prices[i][1]-stock_prices[i-1][1];
                if (dx as i64)*(cy as i64) != (cx as i64)*(dy as i64) {
                    dx = cx;
                    dy = cy;
                    ans += 1;
                }
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                stock_prices: vec![[1,7],[2,6],[3,5],[4,4],[5,4],[6,3],[7,2],[8,1]].iter().map(|s| s.to_vec()).collect(),
                ans: 3,
            },
            Solution {
                stock_prices: vec![[3,4],[1,2],[7,8],[2,3]].iter().map(|s| s.to_vec()).collect(),
                ans: 1,
            },
            Solution {
                stock_prices: vec![[1,1],[500000000,499999999],[1000000000,999999998]].iter().map(|s| s.to_vec()).collect(),
                ans: 2,
            }
        ];
        for i in testcases {
            let ans = Solution::minimum_lines(i.stock_prices);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
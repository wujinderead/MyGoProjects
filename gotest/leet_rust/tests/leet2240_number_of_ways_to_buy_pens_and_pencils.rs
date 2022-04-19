// https://leetcode.com/problems/number-of-ways-to-buy-pens-and-pencils/

// You are given an integer total indicating the amount of money you have. You are also given two
// integers cost1 and cost2 indicating the price of a pen and pencil respectively. You can spend
// part or all of your money to buy multiple quantities (or none) of each kind of writing utensil.
// Return the number of distinct ways you can buy some number of pens and pencils.
// Example 1:
//   Input: total = 20, cost1 = 10, cost2 = 5
//   Output: 9
//   Explanation: The price of a pen is 10 and the price of a pencil is 5.
//     - If you buy 0 pens, you can buy 0, 1, 2, 3, or 4 pencils.
//     - If you buy 1 pen, you can buy 0, 1, or 2 pencils.
//     - If you buy 2 pens, you cannot buy any pencils.
//     The total number of ways to buy pens and pencils is 5 + 3 + 1 = 9.
// Example 2:
//   Input: total = 5, cost1 = 10, cost2 = 10
//   Output: 1
//   Explanation: The price of both pens and pencils are 10, which cost more than
//     total, so you cannot buy any writing utensils. Therefore, there is only 1 way:
//     buy 0 pens and 0 pencils.
// Constraints:
//   1 <= total, cost1, cost2 <= 10⁶

mod _number_of_ways_to_buy_pens_and_pencils {
    struct Solution{
        total: i32,
        cost1: i32,
        cost2: i32,
        ans: i64,
    }

    impl Solution {
        pub fn ways_to_buy_pens_pencils(total: i32, cost1: i32, cost2: i32) -> i64 {
            let mut sum = 0;
            for i in 0..=total/cost1 {
                let remain = total-cost1*i;
                sum += (remain/cost2) as i64+1;
            }
            return sum;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                total: 20,
                cost1: 10,
                cost2: 5,
                ans: 9,
            },
            Solution {
                total: 5,
                cost1: 10,
                cost2: 10,
                ans: 1,
            },
            Solution {
                total: 100,
                cost1: 1,
                cost2: 1,
                ans: 5151,
            },
        ];
        for i in testcases {
            let ans = Solution::ways_to_buy_pens_pencils(i.total, i.cost1, i.cost2);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
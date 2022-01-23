// https://leetcode.com/problems/earliest-possible-day-of-full-bloom/

// You have n flower seeds. Every seed must be planted first before it can begin to grow,
// then bloom. Planting a seed takes time and so does the growth of a seed. You are given
// two 0-indexed integer arrays plantTime and growTime, of length n each: plantTime[i] is
// the number of full days it takes you to plant the iᵗʰ seed.
// Every day, you can work on planting exactly one seed.
//   You do not have to work on planting the same seed on consecutive days,
//     but the planting of a seed is not complete until you have worked plantTime[i]
//     days on planting it in total.
//   growTime[i] is the number of full days it takes the iᵗʰ seed to grow after being
//     completely planted.
//   After the last day of its growth, the flower blooms and stays bloomed forever.
//   From the beginning of day 0, you can plant the seeds in any order.
// Return the earliest possible day where all seeds are blooming.
// Example 1:
//   Input: plantTime = [1,4,3], growTime = [2,3,1]
//   Output: 9
//   Explanation: The grayed out pots represent planting days, colored pots
//     represent growing days, and the flower represents the day it blooms.
//     One optimal way is:
//     On day 0, plant the 0ᵗʰ seed. The seed grows for 2 full days and blooms on day 3.
//     On days 1, 2, 3, and 4, plant the 1ˢᵗ seed. The seed grows for 3 full days and blooms on day 8.
//     On days 5, 6, and 7, plant the 2ⁿᵈ seed. The seed grows for 1 full day and blooms on day 9.
//     Thus, on day 9, all the seeds are blooming.
// Example 2:
//   Input: plantTime = [1,2,3,2], growTime = [2,1,2,1]
//   Output: 9
//   Explanation: The grayed out pots represent planting days, colored pots
//     represent growing days, and the flower represents the day it blooms.
//     One optimal way is:
//     On day 1, plant the 0ᵗʰ seed. The seed grows for 2 full days and blooms on day 4.
//     On days 0 and 3, plant the 1ˢᵗ seed. The seed grows for 1 full day and blooms on day 5.
//     On days 2, 4, and 5, plant the 2ⁿᵈ seed. The seed grows for 2 full days and blooms on day 8.
//     On days 6 and 7, plant the 3ʳᵈ seed. The seed grows for 1 full day and blooms on day 9.
//     Thus, on day 9, all the seeds are blooming.
// Example 3:
//   Input: plantTime = [1], growTime = [1]
//   Output: 2
//   Explanation: On day 0, plant the 0ᵗʰ seed. The seed grows for 1 full day and blooms on day 2.
//     Thus, on day 2, all the seeds are blooming.
// Constraints:
//   n == plantTime.length == growTime.length
//   1 <= n <= 10⁵
//   1 <= plantTime[i], growTime[i] <= 10⁴

mod _earliest_possible_day_of_full_bloom {
    struct Solution{
        plant_time: Vec<i32>,
        grow_time: Vec<i32>,
        ans: i32,
    }

    // say the plant time is [p0,p1,p2], grow time is [g0,g1,g2].
    // the answer is actually the maximal value of the following values:
    //   p0       + g0
    //   p0+p1    + g1
    //   p0+p1+p2 + g2
    // we want the maximal value minimal.
    // since the left part is increasing, so we want right part decreasing, i.e., g0>g1>g2.
    // but is it the optimal answer? e.g, what if we switch p0 and p1?
    // the answer change from max(p0+g0, p0+p1+g1) to max(p1+g1, p1+p0+g0).
    // we have pre-condition: g0>g1, then:
    // if g0>p1+g1,
    // max(p0+g0, p0+p1+g1)    max(p1+g1, p1+p0+g0)
    //       ||                       ||
    //      p0+g0         <          p1+p0+g0
    // the answer get increasing (worse).
    // if g0<p1+g1,
    // max(p0+g0, p0+p1+g1)    max(p1+g1, p1+p0+g0)
    //       ||                       ||
    //     p0+p1+g1      <           p1+p0+g0
    // the answer is also get increasing (worse).
    // so the optimal way is to plant the flowers by decreasing grow_time.
    impl Solution {
        pub fn earliest_full_bloom(plant_time: Vec<i32>, grow_time: Vec<i32>) -> i32 {
            let mut pairs = vec![[0, 0]; plant_time.len()];
            for i in 0..pairs.len() {
                pairs[i] = [plant_time[i], grow_time[i]];
            }
            pairs.sort_by_key(|p| -p[1]);
            let mut max = 0;
            let mut left = 0;
            for i in 0..pairs.len() {
                left += pairs[i][0];
                max = max.max(left+pairs[i][1])
            }
            return max
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                plant_time: vec![1,4,3],
                grow_time: vec![2,3,1],
                ans: 9,
            },
            Solution {
                plant_time: vec![1,2,3,2],
                grow_time: vec![2,1,2,1],
                ans: 9,
            },
            Solution {
                plant_time: vec![1],
                grow_time: vec![1],
                ans: 2,
            },
            Solution {
                plant_time: vec![3,11,29,4,4,26,26,12,13,10,30,19,27,2,10],
                grow_time: vec![10,13,22,17,18,15,21,11,24,14,18,23,1,30,6],
                ans: 227,
            },
        ];
        for i in testcases {
            let ans = Solution::earliest_full_bloom(i.plant_time, i.grow_time);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
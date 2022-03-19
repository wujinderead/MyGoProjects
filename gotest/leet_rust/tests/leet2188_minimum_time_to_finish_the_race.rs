// https://leetcode.com/problems/minimum-time-to-finish-the-race/

// You are given a 0-indexed 2D integer array tires where tires[i] = [fi, ri] indicates that the
// iᵗʰ tire can finish its xᵗʰ successive lap in fi * ri⁽ˣ⁻¹⁾ seconds.
// For example, if fi = 3 and ri = 2, then the tire would finish its 1ˢᵗ lap in 3 seconds,
// its 2ⁿᵈ lap in 3 * 2 = 6 seconds, its 3ʳᵈ lap in 3 * 2² = 12 seconds, etc.
// You are also given an integer changeTime and an integer numLaps.
// The race consists of numLaps laps and you may start the race with any tire. You have an unlimited
// supply of each tire and after every lap, you may change to any given tire (including the current
// tire type) if you wait changeTime seconds.
// Return the minimum time to finish the race.
// Example 1:
//   Input: tires = [[2,3],[3,4]], changeTime = 5, numLaps = 4
//   Output: 21
//   Explanation:
//     Lap 1: Start with tire 0 and finish the lap in 2 seconds.
//     Lap 2: Continue with tire 0 and finish the lap in 2 * 3 = 6 seconds.
//     Lap 3: Change tires to a new tire 0 for 5 seconds and then finish the lap in another 2 seconds.
//     Lap 4: Continue with tire 0 and finish the lap in 2 * 3 = 6 seconds.
//     Total time = 2 + 6 + 5 + 2 + 6 = 21 seconds.
//     The minimum time to complete the race is 21 seconds.
// Example 2:
//   Input: tires = [[1,10],[2,2],[3,4]], changeTime = 6, numLaps = 5
//   Output: 25
//   Explanation:
//     Lap 1: Start with tire 1 and finish the lap in 2 seconds.
//     Lap 2: Continue with tire 1 and finish the lap in 2 * 2 = 4 seconds.
//     Lap 3: Change tires to a new tire 1 for 6 seconds and then finish the lap in another 2 seconds.
//     Lap 4: Continue with tire 1 and finish the lap in 2 * 2 = 4 seconds.
//     Lap 5: Change tires to tire 0 for 6 seconds then finish the lap in another 1 second.
//     Total time = 2 + 4 + 6 + 2 + 4 + 6 + 1 = 25 seconds.
//     The minimum time to complete the race is 25 seconds.
// Constraints:
//   1 <= tires.length <= 10⁵
//   tires[i].length == 2
//   1 <= fi, changeTime <= 10⁵
//   2 <= ri <= 10⁵
//   1 <= numLaps <= 1000

mod _minimum_time_to_finish_the_race {
    struct Solution{
        tires: Vec<Vec<i32>>,
        change_time: i32,
        num_laps: i32,
        ans: i32,
    }

    // sort tires, for fi<fj, must have ti>tj
    use std::cmp::Ordering::{Greater, Less};
    impl Solution {
        pub fn minimum_finish_time(mut tires: Vec<Vec<i32>>, change_time: i32, num_laps: i32) -> i32 {
            // sort tires by fi then by ri
            tires.sort_by(|a, b| {
                if a[0] < b[0] || (a[0] == b[0] && a[1] < b[1]) {
                    return Less;
                }
                return Greater;
            });
            // filter usable tires, for fi<fj, must have ti>tj
            let mut sorted = vec![&tires[0]];
            for v in &tires {
                let &last = sorted.last().unwrap();
                if v[0] > last[0] && v[1] < last[1] {
                    sorted.push(v);
                }
            }

            // for each tire, we must change after 20 laps; so we want to
            // find the minimal time to finish each n_lap if we don't change tire
            let mut dp = vec![i32::MAX; num_laps as usize + 1];
            for &tire in &sorted {
                let (mut f, mut sum, t) = (tire[0], tire[0], tire[1]);
                dp[1] = dp[1].min(f);
                for i in 2..=num_laps.min(20) as usize {
                    f *= t;
                    sum += f;
                    if sum >= 1e5 as i32 {
                        break;
                    }
                    dp[i] = dp[i].min(sum);
                }
            }

            // check the minimal time to finish num_laps
            // dp[i] = min(dp[i-x] + dp[x] + change_time)   1<=x<=20
            // it means (time to finish i-x laps) + change_tire_time + (time to finish x laps)
            for i in 2..=num_laps as usize {
                for j in 1..=20.min(i-1) {
                    dp[i] = dp[i].min(dp[j]+dp[i-j]+change_time);
                }
            }
            return dp[num_laps as usize];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution{
                tires: vec![vec![2,3],vec![3,4]],
                change_time: 5,
                num_laps: 4,
                ans: 21,
            },
            Solution{
                tires: vec![vec![1,10],vec![2,2],vec![3,4]],
                change_time: 6,
                num_laps: 5,
                ans: 25,
            },
            Solution{
                tires: (vec![[3,4],[84,2],[63,8],[72,8],[82,7],[83,6],[23,2],[77,5],[51,10],[28,2],[47,9],[8,3],[48,3],[56,3],[8,10],[66,6],[92,9],[44,6],[23,5],[5,6],[86,9],[13,10],[91,3],[2,2],[8,4],[67,8],[63,6],[52,5],[42,10],[3,9],[66,5],[35,10],[63,6],[65,6],[22,8],[40,9],[43,4],[73,9],[81,5],[32,2],[30,5],[80,9],[50,4],[35,4],[52,7],[11,5],[7,8],[68,3],[54,8],[49,8]]).iter().map(|v| v.to_vec()).collect(),
                change_time: 90,
                num_laps: 87,
                ans: 2526,
            },
            Solution{
                tires: vec![vec![99,7]],
                change_time: 85,
                num_laps: 95,
                ans: 17395,
            },
        ];
        for i in testcases {
            let ans = Solution::minimum_finish_time(i.tires, i.change_time, i.num_laps);
            println!("{}, {}", ans, i.ans);
        }
    } 
}

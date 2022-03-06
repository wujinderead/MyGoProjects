// https://leetcode.com/problems/minimum-time-to-complete-trips/

// You are given an array time where time[i] denotes the time taken by the iᵗʰ bus to complete
// one trip.
// Each bus can make multiple trips successively; that is, the next trip can start immediately
// after completing the current trip. Also, each bus operates independently; that is, the trips
// of one bus do not influence the trips of any other bus.
// You are also given an integer totalTrips, which denotes the number of trips all buses should make
// in total. Return the minimum time required for all buses to complete at least totalTrips trips.
// Example 1:
//   Input: time = [1,2,3], totalTrips = 5
//   Output: 3
//   Explanation:
//     - At time t = 1, the number of trips completed by each bus are [1,0,0].
//       The total number of trips completed is 1 + 0 + 0 = 1.
//     - At time t = 2, the number of trips completed by each bus are [2,1,0].
//       The total number of trips completed is 2 + 1 + 0 = 3.
//     - At time t = 3, the number of trips completed by each bus are [3,1,1].
//       The total number of trips completed is 3 + 1 + 1 = 5.
//     So the minimum time needed for all buses to complete at least 5 trips is 3.
// Example 2:
//   Input: time = [2], totalTrips = 1
//   Output: 2
//   Explanation:
//     There is only one bus, and it will complete its first trip at t = 2.
//     So the minimum time needed to complete 1 trip is 2.
// Constraints:
//   1 <= time.length <= 10⁵
//   1 <= time[i], totalTrips <= 10⁷

mod _minimum_time_to_complete_trips {
    struct Solution{
        time: Vec<i32>,
        total_trips: i32,
        ans: i64,
    }

    impl Solution {
        // binary search: find minimal n that sum(n/time[i]) >= total_trips
        pub fn minimum_time(time: Vec<i32>, total_trips: i32) -> i64 {
            let (mut left, mut right): (i64, i64) = (1, 1e14 as i64);
            while left < right {
                let mid = (left+right)/2;
                let mut sum = 0;
                for &t in &time {
                    sum += mid/(t as i64);
                }
                if sum < (total_trips as i64) {
                    left = mid+1;
                } else {
                    right = mid;
                }
            }
            return left;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution{
                time: vec![1,2,3],
                total_trips: 5,
                ans: 3,
            },
            Solution{
                time: vec![2],
                total_trips: 1,
                ans: 2,
            },
            Solution{
                time: vec![1; 100000],
                total_trips: 10000000,
                ans: 100,
            },
            Solution{
                time: vec![100000000],
                total_trips: 100000000,
                ans: 1e14 as i64,
            }
        ];
        for i in testcases {
            let ans = Solution::minimum_time(i.time, i.total_trips);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
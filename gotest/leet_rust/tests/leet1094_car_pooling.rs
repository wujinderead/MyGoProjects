// https://leetcode.com/problems/car-pooling/

// There is a car with capacity empty seats. The vehicle only drives east
// (i.e., it cannot turn around and drive west).
// You are given the integer capacity and an array trips where trip[i] =
// [numPassengersi, fromi, toi] indicates that the iᵗʰ trip has numPassengersi passengers
// and the locations to pick them up and drop them off are fromi and toi respectively.
// The locations are given as the number of kilometers due east from the car's initial location.
// Return true if it is possible to pick up and drop off all passengers for all the given trips,
// or false otherwise.
// Example 1:
//   Input: trips = [[2,1,5],[3,3,7]], capacity = 4
//   Output: false
// Example 2:
//   Input: trips = [[2,1,5],[3,3,7]], capacity = 5
//   Output: true
// Constraints:
//   1 <= trips.length <= 1000
//   trips[i].length == 3
//   1 <= numPassengersi <= 100
//   0 <= fromi < toi <= 1000
//   1 <= capacity <= 10⁵

mod _car_pooling {
    struct Solution{
        trips: Vec<Vec<i32>>,
        capacity: i32,
        ans: bool,
    }

    impl Solution {
        pub fn car_pooling(trips: Vec<Vec<i32>>, capacity: i32) -> bool {
            let mut ts: Vec<(i32, i32, i32)> = Vec::with_capacity(trips.len()*2);
            for item in trips.iter() {
                ts.push((item[1], item[0], 1));  // 1 for pick
                ts.push((item[2], item[0], 0));  // 0 for drop
            }
            ts.sort_by_key(|k| k.0 * 1000 + k.2);  // by position, if equal, down first
            let mut cap = capacity;
            for &item in ts.iter() {
                if item.2 == 1 {  // pick
                    if item.1 > cap {  // full, can't pick
                        return false;
                    }
                    cap -= item.1  // decrease capacity
                } else {  // drop
                    cap += item.1  // increase capacity
                }
            }
            return true
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                trips: vec![vec![2,1,5],vec![3,3,7]],
                capacity: 4,
                ans: false,
            },
            Solution {
                trips: vec![vec![2,1,5],vec![3,3,7]],
                capacity: 5,
                ans: true,
            },
            Solution {
                trips: vec![vec![2,1,5],vec![4,5,7],vec![1,5,7]],
                capacity: 5,
                ans: true,
            },
        ];
        for i in testcases {
            let ans = Solution::car_pooling(i.trips, i.capacity);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/car-fleet/

// There are n cars going to the same destination along a one-lane road. The destination is target
// miles away.
// You are given two integer array position and speed, both of length n, where position[i] is the
// position of the iᵗʰ car and speed[i] is the speed of the iᵗʰ car (in miles per hour).
// A car can never pass another car ahead of it, but it can catch up to it and drive bumper to
// bumper at the same speed. The faster car will slow down to match the slower car's speed. The
// distance between these two cars is ignored (i.e., they are assumed to have the same position).
// A car fleet is some non-empty set of cars driving at the same position and same speed. Note that
// a single car is also a car fleet.
// If a car catches up to a car fleet right at the destination point, it will still be considered
// as one car fleet.
// Return the number of car fleets that will arrive at the destination.
// Example 1:
//   Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
//   Output: 3
//   Explanation:
//     The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12.
//     The car starting at 0 does not catch up to any other car, so it is a fleet by itself.
//     The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6.
//     The fleet moves at speed 1 until it reaches target.
//     Note that no other cars meet these fleets before the destination, so the answer is 3.
// Example 2:
//   Input: target = 10, position = [3], speed = [3]
//   Output: 1
//   Explanation: There is only one car, hence there is only one fleet.
// Example 3:
//   Input: target = 100, position = [0,2,4], speed = [4,2,1]
//   Output: 1
//   Explanation:
//     The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting each other at 4.
//     The fleet moves at speed 2.
//     Then, the fleet (speed 2) and the car starting at 4 (speed 1) become one fleet,
//     meeting each other at 6. The fleet moves at speed 1 until it reaches target.
// Constraints:
//   n == position.length == speed.length
//   1 <= n <= 10⁵
//   0 < target <= 10⁶
//   0 <= position[i] < target
//   All the values of position are unique.
//   0 < speed[i] <= 10⁶

mod _car_fleet {
    struct Solution{
        target: i32,
        position: Vec<i32>,
        speed: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        // any faster car will block by a slower car, and form a fleet.
        // the fleet's time to target depends on the slowest car in the fleet
        pub fn car_fleet(target: i32, position: Vec<i32>, speed: Vec<i32>) -> i32 {
            let mut tuples = vec![[0; 2]; position.len()];
            for i in 0..position.len() {
                tuples[i] = [position[i], speed[i]];
            }
            tuples.sort_by_key(|t| t[0]);  // sort by position
            let mut cur_max = 0 as f32;
            let target = target as f32;
            let mut count = 0;
            for i in (0..tuples.len()).into_iter().rev() {  // from last car to first
                let time = (target-tuples[i][0] as f32) / tuples[i][1] as f32; // time to reach target
                if time > cur_max {
                    cur_max = time;
                    count += 1;
                }
            }
            return count;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                target: 12,
                position: vec![10,8,0,5,3],
                speed: vec![2,4,1,1,3],
                ans: 3,
            },
            Solution {
                target: 10,
                position: vec![3],
                speed: vec![3],
                ans: 1,
            },
            Solution {
                target: 100,
                position: vec![0,2,4],
                speed: vec![4,2,1],
                ans: 1,
            },
            Solution {
                target: 13,
                position: vec![10,2,5,7,4,6,11],
                speed: vec![7,5,10,5,9,4,1],
                ans: 2,
            }
        ];
        for i in testcases {
            let ans = Solution::car_fleet(i.target, i.position, i.speed);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/maximum-consecutive-floors-without-special-floors/

// Alice manages a company and has rented some floors of a building as office space. Alice has
// decided some of these floors should be special floors, used for relaxation only.
// You are given two integers bottom and top, which denote that Alice has rented all the floors
// from bottom to top (inclusive). You are also given the integer array special, where special[i]
// denotes a special floor that Alice has designated for relaxation.
// Return the maximum number of consecutive floors without a special floor.
// Example 1:
//   Input: bottom = 2, top = 9, special = [4,6]
//   Output: 3
//   Explanation: The following are the ranges (inclusive) of consecutive floors without a special floor:
//     - (2, 3) with a total amount of 2 floors.
//     - (5, 5) with a total amount of 1 floor.
//     - (7, 9) with a total amount of 3 floors.
//     Therefore, we return the maximum number which is 3 floors.
// Example 2:
//   Input: bottom = 6, top = 8, special = [7,6,8]
//   Output: 0
//   Explanation: Every floor rented is a special floor, so we return 0.
// Constraints:
//   1 <= special.length <= 10⁵
//   1 <= bottom <= special[i] <= top <= 10⁹
//   All the values of special are unique.

mod _maximum_consecutive_floors_without_special_floors {
    struct Solution{
        bottom: i32,
        top: i32,
        special: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn max_consecutive(bottom: i32, top: i32, mut special: Vec<i32>) -> i32 {
            let mut max = 0;
            special.sort();
            for i in 1..special.len() {
                max = max.max(special[i]-special[i-1]-1);
            }
            max = max.max(special[0]-bottom);
            max = max.max(top-special[special.len()-1]);
            return max;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                bottom: 2,
                top: 9,
                special: vec![4,6],
                ans: 3,
            },
            Solution {
                bottom: 6,
                top: 8,
                special: vec![7,6,8],
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::max_consecutive(i.bottom, i.top, i.special);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
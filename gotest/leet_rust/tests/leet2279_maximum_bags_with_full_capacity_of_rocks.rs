// https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/

// You have n bags numbered from 0 to n - 1. You are given two 0-indexed integer arrays capacity
// and rocks. The iᵗʰ bag can hold a maximum of capacity[i] rocks and currently contains rocks[i]
// rocks. You are also given an integer additionalRocks, the number of additional rocks you can
// place in any of the bags.
// Return the maximum number of bags that could have full capacity after 
// placing the additional rocks in some bags.
// Example 1:
//   Input: capacity = [2,3,4,5], rocks = [1,2,4,4], additionalRocks = 2
//   Output: 3
//   Explanation:
//     Place 1 rock in bag 0 and 1 rock in bag 1.
//     The number of rocks in each bag are now [2,3,4,4].
//     Bags 0, 1, and 2 have full capacity.
//     There are 3 bags at full capacity, so we return 3.
//     It can be shown that it is not possible to have more than 3 bags at full capacity.
//     Note that there may be other ways of placing the rocks that result in an answer of 3.
// Example 2:
//   Input: capacity = [10,2,2], rocks = [2,2,0], additionalRocks = 100
//   Output: 3
//   Explanation:
//     Place 8 rocks in bag 0 and 2 rocks in bag 2.
//     The number of rocks in each bag are now [10,2,2].
//     Bags 0, 1, and 2 have full capacity.
//     There are 3 bags at full capacity, so we return 3.
//     It can be shown that it is not possible to have more than 3 bags at full capacity.
//     Note that we did not use all of the additional rocks.
// Constraints:
//   n == capacity.length == rocks.length
//   1 <= n <= 5 * 10⁴
//   1 <= capacity[i] <= 10⁹
//   0 <= rocks[i] <= capacity[i]
//   1 <= additionalRocks <= 10⁹

mod _maximum_bags_with_full_capacity_of_rocks {
    struct Solution{
        capacity: Vec<i32>,
        rocks: Vec<i32>,
        additional_rocks: i32,
        ans: i32,
    }

    impl Solution {
        pub fn maximum_bags(capacity: Vec<i32>, rocks: Vec<i32>, mut additional_rocks: i32) -> i32 {
            let mut remain = Vec::with_capacity(capacity.len());
            let mut ans = 0;
            for i in 0..capacity.len() {
                if capacity[i] == rocks[i] {
                    ans += 1;
                    continue;
                }
                remain.push(capacity[i]-rocks[i]);
            }
            remain.sort();
            for i in 0..remain.len() {
                if additional_rocks<remain[i] {
                    break;
                } else {
                    ans += 1;
                    additional_rocks -= remain[i];
                }
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                capacity: vec![2,3,4,5],
                rocks: vec![1,2,4,4],
                additional_rocks: 2,
                ans: 3,
            },
            Solution {
                capacity: vec![10,2,2],
                rocks: vec![2,2,0],
                additional_rocks: 100,
                ans: 3,
            },
        ];
        for i in testcases {
            let ans = Solution::maximum_bags(i.capacity, i.rocks, i.additional_rocks);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
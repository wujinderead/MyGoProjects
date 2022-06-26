// https://leetcode.com/problems/count-number-of-ways-to-place-houses/

// There is a street with n * 2 plots, where there are n plots on each side of the street. The plots
// on each side are numbered from 1 to n. On each plot, a house can be placed.
// Return the number of ways houses can be placed such that no two houses are adjacent to each other
// on the same side of the street. Since the answer may be very large, return it modulo 10⁹ + 7.
// Note that if a house is placed on the iᵗʰ plot on one side of the street, a house can also be
// placed on the iᵗʰ plot on the other side of the street.
// Example 1:
//   Input: n = 1
//   Output: 4
//   Explanation:
//     Possible arrangements:
//     1. All plots are empty.
//     2. A house is placed on one side of the street.
//     3. A house is placed on the other side of the street.
//     4. Two houses are placed, one on each side of the street.
// Example 2:
//   Input: n = 2
//   Output: 9
//   Explanation: The 9 possible arrangements are shown in the diagram above.
// Constraints:
//   1 <= n <= 10⁴

mod _count_number_of_ways_to_place_houses {
    struct Solution{
        n: i32,
        ans: i32,
    }

    impl Solution {
        pub fn count_house_placements(n: i32) -> i32 {
            const MOD: i32 = 1e9 as i32 + 7;
            // one: sequence end with 1 (means house); zero: sequence end with 0 (means empty)
            let (mut one, mut zero) = (1, 1);
            for _i in 1..n {
                let tmp = zero;
                zero = (zero+one) % MOD;
                one = tmp;
            }
            return (((zero+one) as i64) * ((zero+one) as i64) % (MOD as i64)) as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                n: 1,
                ans: 4,
            },
            Solution {
                n: 2,
                ans: 9,
            },
            Solution {
                n: 3,
                ans: 25,
            },
            Solution {
                n: 7,
                ans: 1156,
            }
        ];
        for i in testcases {
            let ans = Solution::count_house_placements(i.n);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
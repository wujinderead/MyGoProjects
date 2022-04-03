// https://leetcode.com/problems/number-of-ways-to-select-buildings/

// You are given a 0-indexed binary string s which represents the types of buildings along
// a street where:
//   s[i] = '0' denotes that the iᵗʰ building is an office and
//   s[i] = '1' denotes that the iᵗʰ building is a restaurant.
// As a city official, you would like to select 3 buildings for random inspection. However,
// to ensure variety, no two consecutive buildings out of the selected buildings can be of
// the same type.
// For example, given s = "001101", we cannot select the 1ˢᵗ, 3ʳᵈ, and 5ᵗʰ buildings as that
// would form "011" which is not allowed due to having two consecutive buildings of the same type.
// Return the number of valid ways to select 3 buildings.
// Example 1:
//   Input: s = "001101"
//   Output: 6
//   Explanation:
//     The following sets of indices selected are valid:
//     - [0,2,4] from "001101" forms "010"
//     - [0,3,4] from "001101" forms "010"
//     - [1,2,4] from "001101" forms "010"
//     - [1,3,4] from "001101" forms "010"
//     - [2,4,5] from "001101" forms "101"
//     - [3,4,5] from "001101" forms "101"
//     No other selection is valid. Thus, there are 6 total ways.
// Example 2:
//   Input: s = "11100"
//   Output: 0
//   Explanation: It can be shown that there are no valid selections.
// Constraints:
//   3 <= s.length <= 10⁵
//   s[i] is either '0' or '1'.

mod _number_of_ways_to_select_buildings {
    struct Solution{
        s: String,
        ans: i32,
    }

    impl Solution {
        pub fn number_of_ways(s: String) -> i64 {
            let ss = s.as_bytes();
            let (mut sum, mut n0, mut n1, mut n10, mut n01) = (0, 0, 0, 0, 0);
            for &i in ss {
                if i==b'0' {
                    sum += n01;
                    n10 += n1;
                    n0 += 1;
                } else {
                    sum += n10;
                    n01 += n0;
                    n1 += 1;
                }
            }
            return sum;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                s: "001101".to_string(),
                ans: 6,
            },
            Solution {
                s: "11100".to_string(),
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::number_of_ways(i.s);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
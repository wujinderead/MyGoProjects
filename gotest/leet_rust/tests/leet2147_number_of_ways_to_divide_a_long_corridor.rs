// https://leetcode.com/problems/number-of-ways-to-divide-a-long-corridor/

// Along a long library corridor, there is a line of seats and decorative plants.
// You are given a 0-indexed string corridor of length n consisting of letters 
// 'S' and 'P' where each 'S' represents a seat and each 'P' represents a plant.
// One room divider has already been installed to the left of index 0, and another
// to the right of index n - 1. Additional room dividers can be installed. For
// each position between indices i - 1 and i (1 <= i <= n - 1), at most one divider
// can be installed.
// Divide the corridor into non-overlapping sections, where each section has exactly two
// seats with any number of plants. There may be multiple ways to perform the division.
// Two ways are different if there is a position with a room divider installed in the first
// way but not in the second way.
// Return the number of ways to divide the corridor. Since the answer may be 
// very large, return it modulo 10⁹ + 7. If there is no way, return 0.
// Example 1:
//   Input: corridor = "SSPPSPS"
//   Output: 3
//   Explanation: There are 3 different ways to divide the corridor.
//     The black bars in the above image indicate the two room dividers already installed.
//     Note that in each of the ways, each section has exactly two seats.
// Example 2:
//   Input: corridor = "PPSPSP"
//   Output: 1
//   Explanation: There is only 1 way to divide the corridor, by not installing any additional dividers.
//     Installing any would create some section that does not have exactly two seats.
// Example 3:
//   Input: corridor = "S"
//   Output: 0
//   Explanation: There is no way to divide the corridor because there will always
//     be a section that does not have exactly two seats.
// Constraints:
//   n == corridor.length
//   1 <= n <= 10⁵
//   corridor[i] is either 'S' or 'P'.

mod _number_of_ways_to_divide_a_long_corridor {
    struct Solution{
        corr: String,
        ans: i32,
    }

    impl Solution {
        pub fn number_of_ways(corridor: String) -> i32 {
            const MOD: i64 = 1e9 as i64 + 7;
            let mut ns = 0;
            let bytes = corridor.as_bytes();
            for &b in bytes {
                if b == b'S' {
                    ns += 1;
                }
            }

            // special case: seat number is odd, or seat number <= 2
            if ns % 2 == 1 {
                return 0;
            }
            if ns <= 2 {
                return ns/2;
            }

            // just how many Ps between the 2nd and 3rd, 4th and 5th...
            // say there are k Ps between 2nd and 3rd seats, there are k+1 ways to split.
            // just multiply these values.
            let mut ans = 1;
            let mut i: usize = 0;
            while i<corridor.len() && bytes[i] != b'S' {
                i +=1;
            }
            i += 1;
            while i<corridor.len() {
                while i<corridor.len() && bytes[i] != b'S' {
                    i += 1;
                }
                let pi = i as i64;
                i += 1;
                while i<corridor.len() && bytes[i] != b'S' {
                    i += 1;
                }
                if i < corridor.len() {
                    ans = (ans * (i as i64 - pi)) % MOD;
                }
                i += 1;
            }
            return ans as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                corr: String::from("SSPPSPS"),
                ans: 3,
            },
            Solution {
                corr: String::from("PPSPSP"),
                ans: 1,
            },
            Solution {
                corr: String::from("S"),
                ans: 0,
            },
            Solution {
                corr: String::from("PPPPPPPSPPPSPPPPSPPPSPPPPPSPPPSPPSPPSPPPPPSPSPPPPPSPPSPPPPPSPPSPPSPPPSPPPPSPPPPSPPPPPSPSPPPPSPSPPPSPPPPSPPPPPSPSPPSPPPPSPPSPPSPPSPPPSPPSPSPPSSSS"),
                ans: 18335643,
            }
        ];
        for i in testcases {
            let ans = Solution::number_of_ways(i.corr);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
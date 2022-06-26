// https://leetcode.com/problems/number-of-distinct-roll-sequences/

// You are given an integer n. You roll a fair 6-sided dice n times. Determine the total number of
// distinct sequences of rolls possible such that the following conditions are satisfied:
//   The greatest common divisor of any adjacent values in the sequence is equal to 1.
//   There is at least a gap of 2 rolls between equal valued rolls. More formally,
//     if the value of the iᵗʰ roll is equal to the value of the jᵗʰ roll, then abs(i - j) > 2.
// Return the total number of distinct sequences possible. Since the answer may 
// be very large, return it modulo 10⁹ + 7.
// Two sequences are considered distinct if at least one element is different.
// Example 1:
//   Input: n = 4
//   Output: 184
//   Explanation: Some of the possible sequences are (1, 2, 3, 4), (6, 1, 2, 3), (1, 2, 3, 1), etc.
//     Some invalid sequences are (1, 2, 1, 3), (1, 2, 3, 6).
//     (1, 2, 1, 3) is invalid since the first and third roll have an equal value
//     and abs(1 - 3) = 2 (i and j are 1-indexed).
//     (1, 2, 3, 6) is invalid since the greatest common divisor of 3 and 6 = 3.
//     There are a total of 184 distinct sequences possible, so we return 184.
// Example 2:
//   Input: n = 2
//   Output: 22
//   Explanation: Some of the possible sequences are (1, 2), (2, 1), (3, 2).
//     Some invalid sequences are (3, 6), (2, 4) since the greatest common divisor is not equal to 1.
//     There are a total of 22 distinct sequences possible, so we return 22.
// Constraints:
//   1 <= n <= 10⁴

mod _number_of_distinct_roll_sequences {
    struct Solution{
        n: i32,
        ans: i32,
    }

    // current value k depends on previous value i, j.
    // for example i=1, j=2, k can be [3,5]. so new23 += old12, new25 += old12
    impl Solution {
        pub fn distinct_sequences(n: i32) -> i32 {
            const MOD: i32 = 1e9 as i32 + 7;
            if n==1 {
                return 6;
            }
            let mut old = vec![0; 66];
            let next = vec![
                vec![],
                vec![2,3,4,5,6],
                vec![1,3,5],
                vec![1,2,4,5],
                vec![1,3,5],
                vec![1,2,3,4,6],
                vec![1,5],
            ];
            for i in 1..=6 {
                for &j in &next[i] {
                    old[i*10+j] = 1;
                }
            }
            let mut new = vec![0; 66];
            for _v in 2..n {
                for i in 1..=6 {
                    for &j in &next[i] {
                        for &k in &next[j] {
                            if k != i {
                                new[j*10+k] += old[i*10+j];
                                new[j*10+k] %= MOD;
                            }
                        }
                    }
                }
                for i in 0..old.len() {
                    old[i] = new[i];
                    new[i] = 0;
                }
            }
            let mut ans = 0;
            for i in 0..old.len() {
                ans = (ans+old[i]) % MOD;
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                n: 4,
                ans: 184,
            },
            Solution {
                n: 2,
                ans: 22,
            }
        ];
        for i in testcases {
            let ans = Solution::distinct_sequences(i.n);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
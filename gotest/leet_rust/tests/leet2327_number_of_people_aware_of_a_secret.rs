// https://leetcode.com/problems/number-of-people-aware-of-a-secret/

// On day 1, one person discovers a secret.
// You are given an integer delay, which means that each person will share the secret with a new
// person every day, starting from delay days after discovering the secret. You are also given an
// integer forget, which means that each person will forget the secret forget days after discovering
// it. A person cannot share the secret on the same day they forgot it, or on any day afterwards.
// Given an integer n, return the number of people who know the secret at the end of day n. Since
// the answer may be very large, return it modulo 10⁹ + 7.
// Example 1:
//   Input: n = 6, delay = 2, forget = 4
//   Output: 5
//   Explanation:
//     Day 1: Suppose the first person is named A. (1 person)
//     Day 2: A is the only person who knows the secret. (1 person)
//     Day 3: A shares the secret with a new person, B. (2 people)
//     Day 4: A shares the secret with a new person, C. (3 people)
//     Day 5: A forgets the secret, and B shares the secret with a new person, D. (3 people)
//     Day 6: B shares the secret with E, and C shares the secret with F. (5 people)
// Example 2:
//   Input: n = 4, delay = 1, forget = 3
//   Output: 6
//   Explanation:
//     Day 1: The first person is named A. (1 person)
//     Day 2: A shares the secret with B. (2 people)
//     Day 3: A and B share the secret with 2 new people, C and D. (4 people)
//     Day 4: A forgets the secret. B, C, and D share the secret with 3 new people. (6 people)
// Constraints:
//   2 <= n <= 1000
//   1 <= delay < forget <= n

mod _number_of_people_aware_of_a_secret {
    struct Solution{
        n: i32, delay: i32, forget: i32,
        ans: i32,
    }

    // can be O(n)
    impl Solution {
        // let dp[i] be the number of person that firstly knows the secret,
        // then if dp[i] = x, then dp[i+delay..i+forget] += x
        pub fn people_aware_of_secret(n: i32, delay: i32, forget: i32) -> i32 {
            const P: i32 = 1e9 as i32 + 7;
            let mut dp = vec![0; n as usize+1];
            dp[1] = 1;
            for i in 1..n {
                for j in delay..forget {
                    if i+j <= n {
                        dp[(i+j) as usize] += dp[i as usize];
                        dp[(i+j) as usize] %= P;
                    }
                }
            }

            // get answer
            let mut ans = 0;
            for i in (n-forget+1)..=n {
                ans = (ans+dp[i as usize]) % P;
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                n: 6, delay: 2, forget: 4,
                ans: 5,
            },
            Solution {
                n: 4, delay: 1, forget: 3,
                ans: 6,
            },
        ];
        for i in testcases {
            let ans = Solution::people_aware_of_secret(i.n, i.delay, i.forget);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
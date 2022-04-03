// https://leetcode.com/problems/find-good-days-to-rob-the-bank/

// You and a gang of thieves are planning on robbing a bank. You are given a 0-indexed integer
// array security, where security[i] is the number of guards on duty on the iᵗʰ day. The days
// are numbered starting from 0. You are also given an integer time.
// The iᵗʰ day is a good day to rob the bank if:
//   There are at least time days before and after the iᵗʰ day,
//   The number of guards at the bank for the time days before i are non-increasing, and
//   The number of guards at the bank for the time days after i are non-decreasing.
// More formally, this means day i is a good day to rob the bank if and only if
// security[i - time] >= security[i - time + 1] >= ... >= security[i] <= ... <=
// security[i + time - 1] <= security[i + time].
// Return a list of all days (0-indexed) that are good days to rob the bank. The order that
// the days are returned in does not matter.
// Example 1:
//   Input: security = [5,3,3,3,5,6,2], time = 2
//   Output: [2,3]
//   Explanation:
//     On day 2, we have security[0] >= security[1] >= security[2] <= security[3] <= security[4].
//     On day 3, we have security[1] >= security[2] >= security[3] <= security[4] <=security[5].
//     No other days satisfy this condition, so days 2 and 3 are the only good days to rob the bank.
// Example 2:
//   Input: security = [1,1,1,1,1], time = 0
//   Output: [0,1,2,3,4]
//   Explanation:
//     Since time equals 0, every day is a good day to rob the bank, so return every day.
// Example 3:
//   Input: security = [1,2,3,4,5,6], time = 2
//   Output: []
//   Explanation:
//     No day has 2 days before it that have a non-increasing number of guards.
//     Thus, no day is a good day to rob the bank, so return an empty list.
// Constraints:
//   1 <= security.length <= 10⁵
//   0 <= security[i], time <= 10⁵

mod _find_good_days_to_rob_the_bank {
    struct Solution{
        security: Vec<i32>,
        time: i32,
        ans: Vec<i32>,
    }

    impl Solution {
        pub fn good_days_to_rob_bank(security: Vec<i32>, time: i32) -> Vec<i32> {
            let (mut dec, mut inc) = (vec![1; security.len()], vec![1 ;security.len()]);
            let mut i = 1;
            while i<security.len() {
                if security[i-1] >= security[i] {
                    dec[i] = dec[i-1]+1;
                }
                i += 1;
            }
            let mut i = (security.len()-2) as i32;
            while i>=0 {
                let ii = i as usize;
                if security[ii] <= security[ii+1] {
                    inc[ii] = inc[ii+1]+1;
                }
                i -= 1;
            }
            let mut i = time;
            let mut ans = Vec::new();
            while i<(security.len() as i32)-time {
                let ii = i as usize;
                if dec[ii] >= (time+1) && inc[ii] >= (time+1) {
                    ans.push(i);
                }
                i += 1;
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                security: vec![5,3,3,3,5,6,2],
                time: 2,
                ans: vec![2,3],
            },
            Solution {
                security: vec![1,1,1,1,1],
                time: 0,
                ans: vec![0,1,2,3,4],
            },
            Solution {
                security: vec![1,2,3,4,5,6],
                time: 2,
                ans: vec![],
            },
        ];
        for i in testcases {
            let ans = Solution::good_days_to_rob_bank(i.security, i.time);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
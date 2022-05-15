// https://leetcode.com/problems/count-number-of-texts/

// Alice is texting Bob using her phone. The mapping of digits to letters is shown in the
// figure below.
// In order to add a letter, Alice has to press the key of the corresponding digit i times,
// where i is the position of the letter in the key.
// For example, to add the letter 's', Alice has to press '7' four times. 
// Similarly, to add the letter 'k', Alice has to press '5' twice.
// Note that the digits '0' and '1' do not map to any letters, so Alice does not use them.
// However, due to an error in transmission, Bob did not receive Alice's text 
// message but received a string of pressed keys instead.
// For example, when Alice sent the message "bob", Bob received the string "2266622".
// Given a string pressedKeys representing the string received by Bob, return 
// the total number of possible text messages Alice could have sent.
// Since the answer may be very large, return it modulo 10⁹ + 7.
// Example 1:
//   Input: pressedKeys = "22233"
//   Output: 8
//   Explanation:
//     The possible text messages Alice could have sent are: "aaadd", "abdd", "badd", "cdd", "aaae", "abe", "bae", and "ce".
//     Since there are 8 possible messages, we return 8.
// Example 2:
//   Input: pressedKeys = "222222222222222222222222222222222222"
//   Output: 82876089
//   Explanation:
//     There are 2082876103 possible text messages Alice could have sent.
//     Since we need to return the answer modulo 10⁹ + 7, we return 2082876103 % (10⁹ + 7) = 82876089.
// Constraints:
//   1 <= pressedKeys.length <= 10⁵
//   pressedKeys only consists of digits from '2' - '9'.

mod _count_number_of_texts {
    struct Solution{
        s: String,
        ans: i32,
    }

    impl Solution {
        pub fn count_texts(pressed_keys: String) -> i32 {
            const P: i64 = 1e9 as i64 + 7;
            let map = vec![0, 0, 0, 0, 0, 0, 0, 1, 0, 1];
            let mut s3 = vec![0; (pressed_keys.len()+1).max(4)];
            let mut s4 = vec![0; (pressed_keys.len()+1).max(5)];
            s3[1] = 1;
            s3[2] = 2;
            s3[3] = 4;  // for 2,3,4,5,6,8, at most 3 chars together
            s4[1] = 1;
            s4[2] = 2;
            s4[3] = 4;
            s4[4] = 8;  // for 7,9, at most 4 chars together
            for i in 4..(pressed_keys.len()+1) {
                s3[i] = (s3[i-1]+s3[i-2]+s3[i-3]) % P;
            }
            for i in 5..(pressed_keys.len()+1) {
                s4[i] = (s4[i-1]+s4[i-2]+s4[i-3]+s4[i-4]) % P;
            }
            let s = vec![s3, s4];
            let mut ans = 1;
            let mut count = 1;
            let bytes = pressed_keys.as_bytes();
            for i in 1..pressed_keys.len() {
                if bytes[i] == bytes[i-1] {
                    count += 1;
                } else {  // for each group of characters, multiple their results
                    ans = (ans*s[map[(bytes[i-1]-b'0') as usize]][count as usize]) % P;
                    count = 1;
                }
            }
            ans = (ans*s[map[(bytes[bytes.len()-1]-b'0') as usize]][count as usize]) % P;
            return ans as i32;
        }

        // no need to group
        pub fn count_texts_dp(pressed_keys: String) -> i32 {
            const P: i32 = 1e9 as i32 + 7;
            let mut dp = vec![0; pressed_keys.len()+1];
            dp[0] = 1;
            let s = pressed_keys.as_bytes();
            for i in 1..s.len()+1 {  // dp[i] the result for string with length i
                dp[i] = dp[i-1];
                if i>=2 && s[i-1] == s[i-2] {
                    dp[i] = (dp[i-1] + dp[i-2]) % P;
                    if i>=3 && s[i-1] == s[i-3] {
                        dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % P;
                        if i>=4 && s[i-1] == s[i-4] && (s[i-1] == b'7' || s[i-1] == b'9') {
                            dp[i] = (dp[i-1] + dp[i-2] + dp[i-3] + dp[i-4]) % P;
                        }
                    }
                }
            }
            return dp[s.len()];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                s: "22233".to_string(),
                ans: 8,
            },
            Solution {
                s: "222222222222222222222222222222222222".to_string(),
                ans: 82876089,
            },
            Solution {
                s: "2".to_string(),
                ans: 1,
            },
            Solution {
                s: "33".to_string(),
                ans: 2,
            },
            Solution {
                s: "55555555999977779555".to_string(),
                ans: 20736,
            }
        ];
        for i in &testcases {
            let ans = Solution::count_texts(i.s.clone());
            println!("{}, {}", ans, i.ans);
        }
        for i in &testcases {
            let ans = Solution::count_texts_dp(i.s.clone());
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/

// You are given a 0-indexed binary string s which represents a sequence of train cars.
// s[i] = '0' denotes that the iᵗʰ car does not contain illegal goods and s[i] = '1' denotes
// that the iᵗʰ car does contain illegal goods.
// As the train conductor, you would like to get rid of all the cars containing illegal goods.
// You can do any of the following three operations any number of times:
//   Remove a train car from the left end (i.e., remove s[0]) which takes 1 unit of time.
//   Remove a train car from the right end (i.e., remove s[s.length - 1]) which takes 1 unit of time.
//   Remove a train car from anywhere in the sequence which takes 2 units of time.
// Return the minimum time to remove all the cars containing illegal goods.
// Note that an empty sequence of cars is considered to have no cars containing illegal goods.
// Example 1:
//   Input: s = "1100101"
//   Output: 5
//   Explanation:
//     One way to remove all the cars containing illegal goods from the sequence is to
//       - remove a car from the left end 2 times. Time taken is 2 * 1 = 2.
//       - remove a car from the right end. Time taken is 1.
//       - remove the car containing illegal goods found in the middle. Time taken is 2.
//     This obtains a total time of 2 + 1 + 2 = 5.
//     An alternative way is to
//       - remove a car from the left end 2 times. Time taken is 2 * 1 = 2.
//       - remove a car from the right end 3 times. Time taken is 3 * 1 = 3.
//     This also obtains a total time of 2 + 3 = 5.
//     5 is the minimum time taken to remove all the cars containing illegal goods.
//     There are no other ways to remove them with less time.
// Example 2:
//   Input: s = "0010"
//   Output: 2
//   Explanation:
//     One way to remove all the cars containing illegal goods from the sequence is to
//       - remove a car from the left end 3 times. Time taken is 3 * 1 = 3.
//     This obtains a total time of 3.
//     Another way to remove all the cars containing illegal goods from the sequence is to
//       - remove the car containing illegal goods found in the middle. Time taken is 2.
//     This obtains a total time of 2.
//     Another way to remove all the cars containing illegal goods from the sequence is to
//       - remove a car from the right end 2 times. Time taken is 2 * 1 = 2.
//     This obtains a total time of 2.
//       2 is the minimum time taken to remove all the cars containing illegal goods.
//     There are no other ways to remove them with less time.
// Constraints:
//   1 <= s.length <= 2 * 10⁵
//   s[i] is either '0' or '1'.

// lee215's solution:
// just compute left part cost, make right part cost as n-1-i (i.e., remove all tight part)
// the reason is: if we have a better choice than n-1-i for right part,
// it means that we delete some 1 in the middle, which will be covered in computing left.
// for (int i = 0; i < n; ++i) {
//     left = Math.min(left + (s.charAt(i) - '0') * 2, i + 1);
//     res = Math.min(res, left + n - 1 - i);
// }
mod _minimum_time_to_remove_all_cars_containing_illegal_goods {
    struct Solution{
        s: String,
        ans: i32,
    }

    impl Solution {
        pub fn minimum_time(s: String) -> i32 {
            let bytes = s.clone().into_bytes();
            if s.len() == 1 {
                return (bytes[0] - b'0') as i32;
            }
            let (mut left, mut right) = (vec![0; s.len()], vec![0; s.len()]);
            // left to right
            if bytes[0] == b'1' {
                left[0] = 1;
            }
            let mut i = 1 as i32;
            while i < bytes.len() as i32 {
                let ii = i as usize;
                if bytes[ii] == b'1' {
                    left[ii] = (left[ii-1]+2).min(i+1);
                } else {
                    left[ii] = left[ii-1];
                }
                i += 1;
            }
            // right to left
            if bytes[s.len()-1] == b'1' {
                right[s.len()-1] = 1;
            }
            i = s.len() as i32 - 2;
            while i >= 0 {
                let ii = i as usize;
                if bytes[ii] == b'1' {
                    right[ii] = (right[ii+1]+2).min(s.len() as i32 - i);
                } else {
                    right[ii] = right[ii+1];
                }
                i -= 1;
            }

            // the minimal cost by plus left with right
            i = 0;
            let mut min = s.len() as i32;
            while i<(s.len()-1) as i32 {
                let ii = i as usize;
                min = min.min(left[ii]+right[ii+1]);
                i += 1;
            }
            return min;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                s: "1100101".to_string(),
                ans: 5,
            },
            Solution {
                s: "0100".to_string(),
                ans: 2,
            },
        ];
        for i in testcases {
            let ans = Solution::minimum_time(i.s);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
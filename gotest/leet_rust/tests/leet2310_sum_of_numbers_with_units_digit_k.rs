// https://leetcode.com/problems/sum-of-numbers-with-units-digit-k/

// Given two integers num and k, consider a set of positive integers with the following properties:
//   The units digit of each integer is k.
//   The sum of the integers is num.
// Return the minimum possible size of such a set, or -1 if no such set exists.
// Note:
//   The set can contain multiple instances of the same integer, and the sum of
//   an empty set is considered 0.
//   The units digit of a number is the rightmost digit of the number.
// Example 1:
//   Input: num = 58, k = 9
//   Output: 2
//   Explanation:
//     One valid set is [9,49], as the sum is 58 and each integer has a units digit of 9.
//     Another valid set is [19,39].
//     It can be shown that 2 is the minimum possible size of a valid set.
// Example 2:
//   Input: num = 37, k = 2
//   Output: -1
//   Explanation: It is not possible to obtain a sum of 37 using only integers that have a units digit of 2.
// Example 3:
//   Input: num = 0, k = 7
//   Output: 0
//   Explanation: The sum of an empty set is considered 0.
// Constraints:
//   0 <= num <= 3000
//   0 <= k <= 9

mod _sum_of_numbers_with_units_digit_k {
    struct Solution{
        num: i32,
        k: i32,
        ans: i32,
    }

    impl Solution {
        pub fn minimum_numbers(num: i32, k: i32) -> i32 {
            if num == 0 {
                return 0;
            }
            for i in 1..=10 {
                if (i*k)%10 == num%10 && i*k <= num {
                    return i;
                }
            }
            return -1;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution{
                num: 58,
                k: 9,
                ans: 2,
            },
            Solution{
                num: 37,
                k: 2,
                ans: -1,
            },
            Solution{
                num: 10,
                k: 0,
                ans: 1,
            },
            Solution{
                num: 70,
                k: 7,
                ans: 10,
            },
            Solution{
                num: 80,
                k: 7,
                ans: 10,
            },
        ];
        for i in testcases {
            let ans = Solution::minimum_numbers(i.num, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
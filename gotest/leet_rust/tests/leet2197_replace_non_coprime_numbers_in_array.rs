// https://leetcode.com/problems/replace-non-coprime-numbers-in-array/

// You are given an array of integers nums. Perform the following steps:
//   Find any two adjacent numbers in nums that are non-coprime.
//   If no such numbers are found, stop the process.
//   Otherwise, delete the two numbers and replace them with their LCM (Least Common Multiple).
// Repeat this process as long as you keep finding two adjacent non-coprime numbers.
// Return the final modified array. It can be shown that replacing adjacent non-coprime numbers
// in any arbitrary order will lead to the same result.
// The test cases are generated such that the values in the final array are 
// less than or equal to 10⁸.
// Two values x and y are non-coprime if GCD(x, y) > 1 where GCD(x, y) is the 
// Greatest Common Divisor of x and y.
// Example 1:
//   Input: nums = [6,4,3,2,7,6,2]
//   Output: [12,7,6]
//   Explanation:
//     - (6, 4) are non-coprime with LCM(6, 4) = 12. Now, nums = [12,3,2,7,6,2].
//     - (12, 3) are non-coprime with LCM(12, 3) = 12. Now, nums = [12,2,7,6,2].
//     - (12, 2) are non-coprime with LCM(12, 2) = 12. Now, nums = [12,7,6,2].
//     - (6, 2) are non-coprime with LCM(6, 2) = 6. Now, nums = [12,7,6].
//     There are no more adjacent non-coprime numbers in nums.
//     Thus, the final modified array is [12,7,6].
//     Note that there are other ways to obtain the same resultant array.
// Example 2:
//   Input: nums = [2,2,1,1,3,3,3]
//   Output: [2,1,1,3]
//   Explanation:
//     - (3, 3) are non-coprime with LCM(3, 3) = 3. Now, nums = [2,2,1,1,3,3].
//     - (3, 3) are non-coprime with LCM(3, 3) = 3. Now, nums = [2,2,1,1,3].
//     - (2, 2) are non-coprime with LCM(2, 2) = 2. Now, nums = [2,1,1,3].
//     There are no more adjacent non-coprime numbers in nums.
//     Thus, the final modified array is [2,1,1,3].
//     Note that there are other ways to obtain the same resultant array.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁵
//   The test cases are generated such that the values in the final array are
//     less than or equal to 10⁸.

mod _replace_non_coprime_numbers_in_array {
    struct Solution{
        nums: Vec<i32>,
        ans: Vec<i32>,
    }

    // just use stack
    impl Solution {
        pub fn replace_non_coprimes(nums: Vec<i32>) -> Vec<i32> {
            let mut stack = Vec::new();
            for nn in nums {
                let mut n = nn;
                while stack.len() > 0 {
                    let x = stack.pop().unwrap();
                    let gcd = Solution::gcd(x, n);
                    if gcd == 1 {
                        stack.push(x);
                        break;
                    }
                    n = n/gcd*x;
                }
                stack.push(n);
            }
            return stack;
        }

        fn gcd(a: i32, b: i32) -> i32 {
            if a==0 {
                return b;
            }
            return Solution::gcd(b%a, a);
        }
    }
    
    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![6,4,3,2,7,6,2],
                ans: vec![12,7,6],
            },
            Solution {
                nums: vec![2,2,1,1,3,3,3],
                ans: vec![2,1,1,3],
            },
        ];
        for i in testcases {
            let ans = Solution::replace_non_coprimes(i.nums);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
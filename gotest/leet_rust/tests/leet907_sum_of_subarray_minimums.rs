// https://leetcode.com/problems/sum-of-subarray-minimums/

// Given an array of integers arr, find the sum of min(b), where b ranges over every (contiguous)
// subarray of arr. Since the answer may be large, return the answer modulo 10⁹ + 7.
// Example 1:
//   Input: arr = [3,1,2,4]
//   Output: 17
//   Explanation:
//     Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].
//     Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.
//     Sum is 17.
// Example 2:
//   Input: arr = [11,81,94,43,3]
//   Output: 444
// Constraints:
//   1 <= arr.length <= 3 * 10⁴
//   1 <= arr[i] <= 3 * 10⁴

mod _sum_of_subarray_minimums {
    struct Solution{
        arr: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn sum_subarray_mins(arr: Vec<i32>) -> i32 {
            const P: i64 = 1e9 as i64 + 7;
            // find the last value that > arr[i] from i to left
            let mut left = vec![0; arr.len()];
            let mut stack = Vec::new();
            for i in 0..arr.len() {
                while stack.len() > 0 && arr[*stack.last().unwrap()] > arr[i] {
                    stack.pop();
                }
                left[i] = if stack.len() == 0 { 0 } else { *stack.last().unwrap()+1 };
                stack.push(i);
            }
            // find the last value that >= arr[i] from i to right
            let mut ans = 0;
            stack.truncate(0);
            for i in (0..arr.len()).into_iter().rev() {
                while stack.len() > 0 && arr[*stack.last().unwrap()] >= arr[i] {
                    stack.pop();
                }
                let right = if stack.len() == 0 { arr.len()-1 } else { *stack.last().unwrap()-1 };
                stack.push(i);
                //println!("{} {} {} {:?}", arr[i], left[i], right, &arr[left[i]..right+1]);
                // in arr[left...right], arr[i] is the min, so there are
                // (i-left+1)*(right-i+1) subarrays that contains arr[i]
                let c = ((i-left[i]+1) as i64 * (right-i+1) as i64) % P;
                ans = (ans + c * arr[i] as i64) % P;
            }
            return ans as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                arr: vec![3,1,2,4],
                ans: 17,
            },
            Solution {
                arr: vec![1,1,1,1],
                ans: 10,
            },
            Solution {
                arr: vec![1,1,2,1],
                ans: 11,
            },
            Solution {
                arr: vec![11,81,94,43,3],
                ans: 444,
            },
        ];
        for i in testcases {
            let ans = Solution::sum_subarray_mins(i.arr);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
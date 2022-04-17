// https://leetcode.com/problems/longest-turbulent-subarray/

// Given an integer array arr, return the length of a maximum size turbulent subarray of arr.
// A subarray is turbulent if the comparison sign flips between each adjacent pair of elements
// in the subarray.
// More formally, a subarray [arr[i], arr[i + 1], ..., arr[j]] of arr is said to be turbulent
// if and only if:
//   For i <= k < j:
//   arr[k] > arr[k + 1] when k is odd, and
//   arr[k] < arr[k + 1] when k is even.
// Or, for i <= k < j:
//   arr[k] > arr[k + 1] when k is even, and
//   arr[k] < arr[k + 1] when k is odd.
// Example 1:
//   Input: arr = [9,4,2,10,7,8,8,1,9]
//   Output: 5
//   Explanation: arr[1] > arr[2] < arr[3] > arr[4] < arr[5]
// Example 2:
//   Input: arr = [4,8,12,16]
//   Output: 2
// Example 3:
//   Input: arr = [100]
//   Output: 1
// Constraints:
//   1 <= arr.length <= 4 * 10⁴
//   0 <= arr[i] <= 10⁹

mod _longest_turbulent_subarray {
    struct Solution{
        arr: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn max_turbulence_size(arr: Vec<i32>) -> i32 {
            let mut cur = 1;
            let mut max = 1;
            for i in 1..arr.len() {
                if arr[i] == arr[i-1] {
                    cur = 1;
                } else if cur == 1 || (arr[i-2]>arr[i-1] && arr[i-1]<arr[i]) || (arr[i-2]<arr[i-1] && arr[i-1]>arr[i]) {
                    cur += 1;
                    max = max.max(cur);
                } else {
                    cur = 2;
                }
            }
            return max;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                arr: vec![9,4,2,10,7,8,8,1,9],
                ans: 5,
            },
            Solution {
                arr: vec![4,8,12,16],
                ans: 2,
            },
            Solution {
                arr: vec![100],
                ans: 1,
            },
            Solution {
                arr: vec![1,1,2],
                ans: 2,
            },
            Solution {
                arr: vec![1,1,2,2,3,4],
                ans: 2,
            },
        ];
        for i in testcases {
            let ans = Solution::max_turbulence_size(i.arr);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
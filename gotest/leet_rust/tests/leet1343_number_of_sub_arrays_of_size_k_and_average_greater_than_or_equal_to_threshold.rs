// https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/

// Given an array of integers arr and two integers k and threshold, return the number of sub-arrays
// of size k and average greater than or equal to threshold.
// Example 1:
//   Input: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
//   Output: 3
//   Explanation: Sub-arrays [2,5,5],[5,5,5] and [5,5,8] have averages 4, 5 and 6 respectively.
//     All other sub-arrays of size 3 have averages less than 4 (the threshold).
// Example 2:
//   Input: arr = [11,13,17,23,29,31,7,5,2,3], k = 3, threshold = 5
//   Output: 6
//   Explanation: The first 6 sub-arrays of size 3 have averages greater than 5.
//     Note that averages are not integers.
// Constraints:
//   1 <= arr.length <= 10⁵
//   1 <= arr[i] <= 10⁴
//   1 <= k <= arr.length
//   0 <= threshold <= 10⁴

mod _number_of_sub_arrays_of_size_k_and_average_greater_than_or_equal_to_threshold {
    struct Solution{
        arr: Vec<i32>,
        k: i32,
        threshold: i32,
        ans: i32,
    }

    impl Solution {
        pub fn num_of_subarrays(arr: Vec<i32>, k: i32, threshold: i32) -> i32 {
            let mut sum = 0;
            let mut count = 0;
            let kk= k as usize;
            for i in 0..kk {
                sum += arr[i];
            }
            if sum >= threshold*k {
                count += 1;
            }
            for i in kk..arr.len() {
                sum = sum+arr[i]-arr[i-kk];
                if sum >= threshold*k {
                    count += 1;
                }
            }
            return count;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                arr: vec![2,2,2,2,5,5,5,8],
                k: 3,
                threshold: 4,
                ans: 3,
            },
            Solution {
                arr: vec![11,13,17,23,29,31,7,5,2,3],
                k: 3,
                threshold: 5,
                ans: 6,
            },
        ];
        for i in testcases {
            let ans = Solution::num_of_subarrays(i.arr, i.k, i.threshold);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
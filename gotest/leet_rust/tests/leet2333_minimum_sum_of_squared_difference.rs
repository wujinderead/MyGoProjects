// https://leetcode.com/problems/minimum-sum-of-squared-difference/

// You are given two positive 0-indexed integer arrays nums1 and nums2, both of length n.
// The sum of squared difference of arrays nums1 and nums2 is defined as the sum of
// (nums1[i] - nums2[i])² for each 0 <= i < n.
// You are also given two positive integers k1 and k2. You can modify any of the elements of nums1
// by +1 or -1 at most k1 times. Similarly, you can modify any of the elements of nums2 by +1 or -1
// at most k2 times.
// Return the minimum sum of squared difference after modifying array nums1 at most k1 times and
// modifying array nums2 at most k2 times.
// Note: You are allowed to modify the array elements to become negative integers.
// Example 1:
//   Input: nums1 = [1,2,3,4], nums2 = [2,10,20,19], k1 = 0, k2 = 0
//   Output: 579
//   Explanation: The elements in nums1 and nums2 cannot be modified because k1 = 0 and k2 = 0.
//     The sum of square difference will be: (1 - 2)² + (2 - 10)² + (3 - 20)² + (4 -19)² = 579.
// Example 2:
//   Input: nums1 = [1,4,10,12], nums2 = [5,8,6,9], k1 = 1, k2 = 1
//   Output: 43
//   Explanation: One way to obtain the minimum sum of square difference is:
//     - Increase nums1[0] once.
//     - Increase nums2[2] once.
//     The minimum of the sum of square difference will be:
//     (2 - 5)2 + (4 - 8)2 + (10 - 7)2 + (12 - 9)² = 43.
//     Note that, there are other ways to obtain the minimum of the sum of square
//     difference, but there is no way to obtain a sum smaller than 43.
// Constraints:
//   n == nums1.length == nums2.length
//   1 <= n <= 10⁵
//   0 <= nums1[i], nums2[i] <= 10⁵
//   0 <= k1, k2 <= 10⁹

mod _minimum_sum_of_squared_difference {
    struct Solution{
        nums1: Vec<i32>,
        nums2: Vec<i32>,
        k1: i32,
        k2: i32,
        ans: i32,
    }

    impl Solution {
        pub fn min_sum_square_diff(nums1: Vec<i32>, nums2: Vec<i32>, k1: i32, k2: i32) -> i64 {
            let mut k = k1+k2;
            let mut diff = vec![0; 100001];
            let mut max = 0;
            // count diff to bucket
            for i in 0..nums1.len() {
                let d = (nums1[i]-nums2[i]).abs();
                max = max.max(d);
                diff[d as usize] += 1;
            }
            let mut ans = 0;
            while max > 0 {
                let mut m = max as usize;
                if k==0 || diff[m] > k {
                    diff[m] -= k;
                    diff[m-1] += k;
                    while m > 0 {
                        ans += (m as i64)*(m as i64)*(diff[m] as i64);
                        m -= 1;
                    }
                    return ans;
                }
                diff[m-1] += diff[m]; // decrease current diff
                k -= diff[m];
                max -= 1;
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums1: vec![1,2,3,4],
                nums2: vec![2,10,20,19],
                k1: 0,
                k2: 0,
                ans: 579,
            },
            Solution {
                nums1: vec![1,4,10,12],
                nums2: vec![5,8,6,9],
                k1: 1,
                k2: 1,
                ans: 43,
            },
            Solution {
                nums1: vec![3,3,3,3],
                nums2: vec![3,3,3,3],
                k1: 10,
                k2: 10,
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::min_sum_square_diff(i.nums1, i.nums2, i.k1, i.k2);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
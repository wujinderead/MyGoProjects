// https://leetcode.com/problems/maximum-length-of-repeated-subarray/

// Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears
// in both arrays.
// Example 1:
//   Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
//   Output: 3
//   Explanation: The repeated subarray with maximum length is [3,2,1].
// Example 2:
//   Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
//   Output: 5
// Constraints:
//   1 <= nums1.length, nums2.length <= 1000
//   1 <= nums1[i], nums2[i] <= 100

mod _maximum_length_of_repeated_subarray {
    struct Solution{
        nums1: Vec<i32>,
        nums2: Vec<i32>,
        ans: i32,
    }

    // just compare two arrays with index diff
    impl Solution {
        pub fn find_length(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
            let mut max = 0;
            for diff in 0..nums1.len()-1 {
                let mut cur = 0;
                for i in 0..nums1.len()-diff {
                    if i+diff>=nums2.len() {
                        break;
                    }
                    if nums1[i] == nums2[i+diff] {
                        cur += 1;
                        max = max.max(cur);
                    } else {
                        cur = 0;
                    }
                }
                cur = 0;
                for i in 0..nums1.len()-diff {
                    if i>=nums2.len() {
                        break;
                    }
                    if nums1[i+diff] == nums2[i] {
                        cur += 1;
                        max = max.max(cur);
                    } else {
                        cur = 0;
                    }
                }
            }
            return max;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums1: vec![1,2,3,2,1],
                nums2: vec![3,2,1,4,7],
                ans: 3,
            },
            Solution {
                nums1: vec![0,0,0,0,0],
                nums2: vec![0,0,0,0,0],
                ans: 5,
            },
            Solution {
                nums1: vec![1,2,3,2,1],
                nums2: vec![3,2,1,4],
                ans: 3,
            }
        ];
        for i in testcases {
            let ans = Solution::find_length(i.nums1, i.nums2);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
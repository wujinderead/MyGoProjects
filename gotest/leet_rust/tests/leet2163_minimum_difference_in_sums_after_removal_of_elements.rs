// https://leetcode.com/problems/minimum-difference-in-sums-after-removal-of-elements/

// You are given a 0-indexed integer array nums consisting of 3 * n elements.
// You are allowed to remove any subsequence of elements of size exactly n from nums.
// The remaining 2 * n elements will be divided into two equal parts:
//   The first n elements belonging to the first part and their sum is sumfirst.
//   The next n elements belonging to the second part and their sum is sumsecond.
// The difference in sums of the two parts is denoted as sumfirst - sumsecond.
// For example, if sumfirst = 3 and sumsecond = 2, their difference is 1. 
// Similarly, if sumfirst = 2 and sumsecond = 3, their difference is -1.
// Return the minimum difference possible between the sums of the two parts after the
// removal of n elements.
// Example 1:
//   Input: nums = [3,1,2]
//   Output: -1
//   Explanation: Here, nums has 3 elements, so n = 1.
//     Thus we have to remove 1 element from nums and divide the array into two equal parts.
//     - If we remove nums[0] = 3, the array will be [1,2]. The difference in sums
//       of the two parts will be 1 - 2 = -1.
//     - If we remove nums[1] = 1, the array will be [3,2]. The difference in sums
//       of the two parts will be 3 - 2 = 1.
//     - If we remove nums[2] = 2, the array will be [3,1]. The difference in sums
//       of the two parts will be 3 - 1 = 2.
//     The minimum difference between sums of the two parts is min(-1,1,2) = -1.
// Example 2:
//   Input: nums = [7,9,5,8,1,3]
//   Output: 1
//   Explanation: Here n = 2. So we must remove 2 elements and divide the remaining array into
//     two parts containing two elements each.
//     If we remove nums[2] = 5 and nums[3] = 8, the resultant array will be [7,9,1,3].
//     The difference in sums will be (7+9) - (1+3) = 12.
//     To obtain the minimum difference, we should remove nums[1] = 9 and nums[4] = 1.
//     The resultant array becomes [7,5,8,3]. The difference in sums of the two parts is (7+5) - (8+3) = 1.
//     It can be shown that it is not possible to obtain a difference smaller than 1.
// Constraints:
//   nums.length == 3 * n
//   1 <= n <= 10⁵
//   1 <= nums[i] <= 10⁵

mod _minimum_difference_in_sums_after_removal_of_elements {
    struct Solution{
        nums: Vec<i32>,
        ans: i32,
    }

    // for each i in n<=i<=2n, find the top-n minimal numbers in nums[:i],
    // and the top-n maximal numbers in nums[i:], and get the minimal difference.
    use std::collections::BinaryHeap;
    impl Solution {
        pub fn minimum_difference(ns: Vec<i32>) -> i64 {
            let n = ns.len()/3;
            let nums = ns.iter().map(|&x| x as i64).collect::<Vec<_>>();

            // find the sum(top-n minimal) in nums[:i], n<=i<=2n
            let mut l_heap = BinaryHeap::from(nums[..n].to_vec());
            let mut l_min = Vec::<i64>::with_capacity(n+1);
            l_min.push(nums[..n].iter().sum());
            let mut i = n;
            while i<2*n {
                if nums[i] < *l_heap.peek().unwrap() {
                    let popped = l_heap.pop().unwrap();
                    l_heap.push(nums[i]);
                    l_min.push(l_min[l_min.len()-1]-popped+nums[i]);
                } else {
                    l_min.push(l_min[l_min.len()-1]);
                }
                i += 1;
            }

            // find the sum(top-n maximal) in nums[i:], n<=i<=2n
            // convert to -x to make a min-heap
            let mut r_heap = BinaryHeap::from(nums[2*n..].iter().map(|&x| -x).collect::<Vec<_>>());
            let mut r_max = Vec::<i64>::with_capacity(n+1);
            r_max.push(nums[2*n..].iter().sum());
            i = 2*n-1;
            while i >= n {
                if -nums[i] < *r_heap.peek().unwrap() {
                    let popped = -r_heap.pop().unwrap();
                    r_heap.push(-nums[i]);
                    r_max.push(r_max[r_max.len()-1]-popped+nums[i]);
                } else {
                    r_max.push(r_max[r_max.len()-1]);
                }
                i -= 1;
            }
            r_max.reverse();  // need reverse because i is from 2n to n

            // get minimal difference
            let mut min = l_min[0] - r_max[0];
            for (i, _b) in l_min.iter().enumerate() {
                min = min.min(l_min[i]-r_max[i]);
            }
            return min as i64;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![3,1,2],
                ans: -1,
            },
            Solution {
                nums: vec![7,9,5,8,1,3],
                ans: 1,
            },
        ];
        for i in testcases {
            let ans = Solution::minimum_difference(i.nums);
            println!("{}, {}", ans, i.ans);
        }
    }
}
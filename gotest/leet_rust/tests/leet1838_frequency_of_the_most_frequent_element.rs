// https://leetcode.com/problems/frequency-of-the-most-frequent-element/

// The frequency of an element is the number of times it occurs in an array.
// You are given an integer array nums and an integer k. In one operation, you can choose an index
// of nums and increment the element at that index by 1.
// Return the maximum possible frequency of an element after performing at most k operations.
// Example 1:
//   Input: nums = [1,2,4], k = 5
//   Output: 3
//   Explanation: Increment the first element three times and the second element
//     two times to make nums = [4,4,4]. 4 has a frequency of 3.
// Example 2:
//   Input: nums = [1,4,8,13], k = 5
//   Output: 2
//   Explanation: There are multiple optimal solutions:
//     - Increment the first element three times to make nums = [4,4,8,13]. 4 has a frequency of 2.
//     - Increment the second element four times to make nums = [1,8,8,13]. 8 has a frequency of 2.
//     - Increment the third element five times to make nums = [1,4,13,13]. 13 has a frequency of 2.
// Example 3:
//   Input: nums = [3,9,6], k = 2
//   Output: 1
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁵
//   1 <= k <= 10⁵

mod _frequency_of_the_most_frequent_element {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i32,
    }

    impl Solution {
        // binary search
        pub fn max_frequency(nums: Vec<i32>, k: i32) -> i32 {
            let mut nums = nums.iter().map(|&s| s as i64).collect::<Vec<_>>();
            nums.sort();
            let mut prefix = vec![0; nums.len()+1];
            for i in 1..=nums.len() {
                prefix[i] = nums[i-1]+prefix[i-1];  // sum(nums[i...j]) = prefix[j+1]-prefix[i]
            }
            let (mut left, mut right) = (1, nums.len() as i64);
            while left < right {
                let mid = (left+right+1)/2;
                // find the minimal cost to make a `mid`-sized subarray equal
                // i.e.: the cost to make nums[i...mid-1] to equal nums[mid-1]
                //       the cost to make nums[i+1...mid] to equal nums[mid]  ...
                let mut min = 1e10 as i64;
                for i in 0..(nums.len()-mid as usize+1) {
                    min = min.min(nums[i+mid as usize-1]*mid-(prefix[i+mid as usize]-prefix[i]));
                }
                if k as i64 >= min {  // k can make a `mid`-sized subarray equal, make a larger mid
                    left = mid;
                } else {
                    right = mid-1;
                }
            }
            return left as i32;
        }

        // sliding window:
        // say nums[i...j] is the window that we can make numbers in window to equal nums[j],
        // which means sum(nums[i...j])+k >= (j-i+1)*nums[j]
        // if we append a new nums[j+1] to window, we need pop some value to find a new i'
        // so that sum(nums[i'...j+1])+k >= (j+1-i'+1)*nums[j+1]
        pub fn max_frequency_sliding_window(nums: Vec<i32>, k: i32) -> i32 {
            let mut nums = nums.iter().map(|&s| s as i64).collect::<Vec<_>>();
            nums.sort();
            let mut sum = nums[0];
            let mut ans = 1;
            let mut i = 0;
            for j in 1..nums.len() {
                sum += nums[j];
                while (sum+k as i64) < (j as i64 - i as i64 + 1) * nums[j] {
                    sum -= nums[i];
                    i += 1;
                }
                ans = ans.max(j-i+1);
            }
            return ans as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,2,4],
                k: 5,
                ans: 3,
            },
            Solution {
                nums: vec![1,2,4],
                k: 2,
                ans: 2,
            },
            Solution {
                nums: vec![1,4,8,13],
                k: 5,
                ans: 2,
            },
            Solution {
                nums: vec![3,9,6],
                k: 2,
                ans: 1,
            },
            Solution {
                nums: vec![1,2,3,4],
                k: 6,
                ans: 4,
            },
            Solution {
                nums: vec![1,2,3,4],
                k: 5,
                ans: 3,
            },
            Solution {
                nums: vec![9930,9923,9983,9997,9934,9952,9945,9914,9985,9982,9970,9932,9985,9902,9975,9990,9922,9990,9994,9937,9996,9964,9943,9963,9911,9925,9935,9945,9933,9916,9930,9938,10000,9916,9911,9959,9957,9907,9913,9916,9993,9930,9975,9924,9988,9923,9910,9925,9977,9981,9927,9930,9927,9925,9923,9904,9928,9928,9986,9903,9985,9954,9938,9911,9952,9974,9926,9920,9972,9983,9973,9917,9995,9973,9977,9947,9936,9975,9954,9932,9964,9972,9935,9946,9966],
                k: 3056,
                ans: 73,
            }
        ];
        for i in &testcases {
            let ans = Solution::max_frequency(i.nums.clone(), i.k);
            println!("{}, {}", ans, i.ans);
        }
        for i in &testcases {
            let ans = Solution::max_frequency_sliding_window(i.nums.clone(), i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
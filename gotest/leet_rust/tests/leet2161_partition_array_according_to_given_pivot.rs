// https://leetcode.com/problems/partition-array-according-to-given-pivot/

// You are given a 0-indexed integer array nums and an integer pivot. Rearrange nums such that
// the following conditions are satisfied:
//   Every element less than pivot appears before every element greater than pivot.
//   Every element equal to pivot appears in between the elements less than and greater than pivot.
//   The relative order of the elements less than pivot and the elements greater than pivot is maintained.
// More formally, consider every pi, pj where pi is the new position of the iᵗʰ element and
// pj is the new position of the jᵗʰ element. For elements less than pivot, if i < j and
// nums[i] < pivot and nums[j] < pivot, then pi < pj. Similarly for elements greater than pivot,
// if i < j and nums[i] > pivot and nums[j] > pivot, then pi < pj.
// Return nums after the rearrangement.
// Example 1:
//   Input: nums = [9,12,5,10,14,3,10], pivot = 10
//   Output: [9,5,3,10,10,12,14]
//   Explanation:
//     The elements 9, 5, and 3 are less than the pivot so they are on the left side of the array.
//     The elements 12 and 14 are greater than the pivot so they are on the right side of the array.
//     The relative ordering of the elements less than and greater than pivot is also maintained.
//     [9, 5, 3] and [12, 14] are the respective orderings.
// Example 2:
//   Input: nums = [-3,4,3,2], pivot = 2
//   Output: [-3,2,4,3]
//   Explanation:
//     The element -3 is less than the pivot so it is on the left side of the array.
//     The elements 4 and 3 are greater than the pivot so they are on the right side of the array.
//     The relative ordering of the elements less than and greater than pivot is also maintained.
//     [-3] and [4, 3] are the respective orderings.
// Constraints:
//   1 <= nums.length <= 10⁵
//   -10⁶ <= nums[i] <= 10⁶
//   pivot equals to an element of nums.

mod _partition_array_according_to_given_pivot {
    struct Solution{
        nums: Vec<i32>,
        pivot: i32,
        ans: Vec<i32>,
    }

    impl Solution {
        pub fn pivot_array(nums: Vec<i32>, pivot: i32) -> Vec<i32> {
            let mut ans = vec![0; nums.len()];
            let (mut low, mut high, mut np) = (0, nums.len()-1, 0);
            for &v in nums.iter() {
                if v < pivot {
                    ans[low] = v;
                    low += 1;
                } else if v > pivot {
                    ans[high] = v;
                    high -= 1;
                } else {
                    np += 1;
                }
            }
            while np > 0 {
                ans[low] = pivot;
                low += 1;
                np -= 1;
            }
            high = nums.len()-1;
            while low < high {
                let tmp = ans[low];
                ans[low] = ans[high];
                ans[high] = tmp;
                low += 1;
                high -= 1;
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![9,12,5,10,14,3,10],
                pivot: 10,
                ans: vec![9,5,3,10,10,12,14],
            },
            Solution {
                nums: vec![-3,4,3,2],
                pivot: 2,
                ans: vec![-3,2,4,3],
            },
        ];
        for i in testcases {
            let ans = Solution::pivot_array(i.nums, i.pivot);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
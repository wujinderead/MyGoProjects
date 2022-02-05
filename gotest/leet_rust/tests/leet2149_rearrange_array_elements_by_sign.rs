// https://leetcode.com/problems/rearrange-array-elements-by-sign/

// You are given a 0-indexed integer array nums of even length consisting of an
// equal number of positive and negative integers.
// You should rearrange the elements of nums such that the modified array 
// follows the given conditions:
//   Every consecutive pair of integers have opposite signs.
//   For all integers with the same sign, the order in which they were present in nums is preserved.
//   The rearranged array begins with a positive integer.
// Return the modified array after rearranging the elements to satisfy the aforementioned conditions.
// Example 1:
//   Input: nums = [3,1,-2,-5,2,-4]
//   Output: [3,-2,1,-5,2,-4]
//   Explanation:
//     The positive integers in nums are [3,1,2]. The negative integers are [-2,-5,-4].
//     The only possible way to rearrange them such that they satisfy all conditions is [3,-2,1,-5,2,-4].
//     Other ways such as [1,-2,2,-5,3,-4], [3,1,2,-2,-5,-4], [-2,3,-5,1,-4,2] are
//     incorrect because they do not satisfy one or more conditions.
// Example 2:
//   Input: nums = [-1,1]
//   Output: [1,-1]
//   Explanation:
//     1 is the only positive integer and -1 the only negative integer in nums.
//     So nums is rearranged to [1,-1].
// Constraints:
//   2 <= nums.length <= 2 * 10⁵
//   nums.length is even
//   1 <= |nums[i]| <= 10⁵
//   nums consists of equal number of positive and negative integers.

mod _rearrange_array_elements_by_sign {
    struct Solution{
        nums: Vec<i32>,
        ans: Vec<i32>,
    }

    impl Solution {
        pub fn rearrange_array(nums: Vec<i32>) -> Vec<i32> {
            let mut ans = vec![0; nums.len()];
            let (mut pi, mut ni) = (0, 1);
            for &i in nums.iter() {
                if i > 0 {
                    ans[pi] = i;
                    pi += 2;
                } else {
                    ans[ni] = i;
                    ni += 2;
                }
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![3,1,-2,-5,2,-4],
                ans: vec![3,-2,1,-5,2,-4],
            },
            Solution {
                nums: vec![-1,1],
                ans: vec![1,-1],
            },
        ];
        for i in testcases {
            let ans = Solution::rearrange_array(i.nums);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
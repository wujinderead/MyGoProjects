// https://leetcode.com/problems/find-all-lonely-numbers-in-the-array/

// You are given an integer array nums. A number x is lonely when it appears only once,
// and no adjacent numbers (i.e. x + 1 and x - 1) appear in the array.
// Return all lonely numbers in nums. You may return the answer in any order.
// Example 1:
//   Input: nums = [10,6,5,8]
//   Output: [10,8]
//   Explanation:
//     - 10 is a lonely number since it appears exactly once and 9 and 11 does not appear in nums.
//     - 8 is a lonely number since it appears exactly once and 7 and 9 does not appear in nums.
//     - 5 is not a lonely number since 6 appears in nums and vice versa.
//     Hence, the lonely numbers in nums are [10, 8].
//     Note that [8, 10] may also be returned.
// Example 2:
//   Input: nums = [1,3,5,3]
//   Output: [1,5]
//   Explanation:
//     - 1 is a lonely number since it appears exactly once and 0 and 2 does not appear in nums.
//     - 5 is a lonely number since it appears exactly once and 4 and 6 does not appear in nums.
//     - 3 is not a lonely number since it appears twice.
//     Hence, the lonely numbers in nums are [1, 5].
//     Note that [5, 1] may also be returned.
// Constraints:
//   1 <= nums.length <= 10⁵
//   0 <= nums[i] <= 10⁶

mod _find_all_lonely_numbers_in_the_array {
    use std::collections::HashMap;

    struct Solution{
        nums: Vec<i32>,
        ans: Vec<i32>,
    }

    impl Solution {
        pub fn find_lonely(nums: Vec<i32>) -> Vec<i32> {
            let mut ans = Vec::new();
            let mut map = HashMap::new();
            for &i in nums.iter() {
                *map.entry(i).or_insert(0) += 1;
            }
            for (&k, &_v) in map.iter() {
                if map.get(&k)==Some(&1) && map.get(&(k-1)).is_none() && map.get(&(k+1)).is_none() {
                    ans.push(k);
                }
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![10,6,5,8],
                ans: vec![10,8],
            },
            Solution {
                nums: vec![1,3,5,3],
                ans: vec![1,5],
            },
        ];
        for i in testcases {
            let ans = Solution::find_lonely(i.nums);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/replace-elements-in-an-array/

// You are given a 0-indexed array nums that consists of n distinct positive integers. Apply m
// operations to this array, where in the iᵗʰ operation you replace the number operations[i][0]
// with operations[i][1].
// It is guaranteed that in the iᵗʰ operation:
//   operations[i][0] exists in nums.
//   operations[i][1] does not exist in nums.
// Return the array obtained after applying all the operations.
// Example 1:
//   Input: nums = [1,2,4,6], operations = [[1,3],[4,7],[6,1]]
//   Output: [3,2,7,1]
//   Explanation: We perform the following operations on nums:
//     - Replace the number 1 with 3. nums becomes [3,2,4,6].
//     - Replace the number 4 with 7. nums becomes [3,2,7,6].
//     - Replace the number 6 with 1. nums becomes [3,2,7,1].
//     We return the final array [3,2,7,1].
// Example 2:
//   Input: nums = [1,2], operations = [[1,3],[2,1],[3,2]]
//   Output: [2,1]
//   Explanation: We perform the following operations to nums:
//     - Replace the number 1 with 3. nums becomes [3,2].
//     - Replace the number 2 with 1. nums becomes [3,1].
//     - Replace the number 3 with 2. nums becomes [2,1].
//     We return the array [2,1].
// Constraints:
//   n == nums.length
//   m == operations.length
//   1 <= n, m <= 10⁵
//   All the values of nums are distinct.
//   operations[i].length == 2
//   1 <= nums[i], operations[i][0], operations[i][1] <= 10⁶
//   operations[i][0] will exist in nums when applying the iᵗʰ operation.
//   operations[i][1] will not exist in nums when applying the iᵗʰ operation.

mod _replace_elements_in_an_array {
    struct Solution{
        nums: Vec<i32>,
        operations: Vec<Vec<i32>>,
        ans: Vec<i32>,
    }

    use std::collections::HashMap;
    impl Solution {
        pub fn array_change(mut nums: Vec<i32>, operations: Vec<Vec<i32>>) -> Vec<i32> {
            let mut map = HashMap::new();
            for (i, &v) in nums.iter().enumerate() {
                map.insert(v, i);
            }
            for s in operations {
                let &i = map.get(&s[0]).unwrap();
                nums[i] = s[1];
                map.insert(s[1], i);
                map.remove(&s[0]);
            }
            return nums;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,2,4,6],
                operations: vec![[1,3],[4,7],[6,1]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![3,2,7,1],
            },
            Solution {
                nums: vec![1,2],
                operations: vec![[1,3],[2,1],[3,2]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![2,1],
            },
        ];
        for i in testcases {
            let ans = Solution::array_change(i.nums, i.operations);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
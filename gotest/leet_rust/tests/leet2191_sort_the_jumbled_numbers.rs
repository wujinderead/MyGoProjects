// https://leetcode.com/problems/sort-the-jumbled-numbers/

// You are given a 0-indexed integer array mapping which represents the mapping rule of a shuffled
// decimal system. mapping[i] = j means digit i should be mapped to digit j in this system.
// The mapped value of an integer is the new integer obtained by replacing each occurrence of
// digit i in the integer with mapping[i] for all 0 <= i <= 9.
// You are also given another integer array nums. Return the array nums sorted in non-decreasing
// order based on the mapped values of its elements.
// Notes:
// Elements with the same mapped values should appear in the same relative order as in the input.
// The elements of nums should only be sorted based on their mapped values and not be replaced by them.
// Example 1:
//   Input: mapping = [8,9,4,0,2,1,3,5,7,6], nums = [991,338,38]
//   Output: [338,38,991]
//   Explanation:
//     Map the number 991 as follows:
//     1. mapping[9] = 6, so all occurrences of the digit 9 will become 6.
//     2. mapping[1] = 9, so all occurrences of the digit 1 will become 9.
//     Therefore, the mapped value of 991 is 669.
//     338 maps to 007, or 7 after removing the leading zeros.
//     38 maps to 07, which is also 7 after removing leading zeros.
//     Since 338 and 38 share the same mapped value, they should remain in the same
//     relative order, so 338 comes before 38.
//     Thus, the sorted array is [338,38,991].
// Example 2:
//   Input: mapping = [0,1,2,3,4,5,6,7,8,9], nums = [789,456,123]
//   Output: [123,456,789]
//   Explanation: 789 maps to 789, 456 maps to 456, and 123 maps to 123.
//     Thus, the sorted array is [123,456,789].
// Constraints:
//   mapping.length == 10
//   0 <= mapping[i] <= 9
//   All the values of mapping[i] are unique.
//   1 <= nums.length <= 3 * 10⁴
//   0 <= nums[i] < 10⁹

mod _sort_the_jumbled_numbers {

    struct Solution{
        mapping: Vec<i32>,
        nums: Vec<i32>,
        ans: Vec<i32>,
    }

    use std::cmp::Ordering::{Greater, Less};
    impl Solution {
        pub fn sort_jumbled(mapping: Vec<i32>, nums: Vec<i32>) -> Vec<i32> {
            let mut mapped = nums.iter().enumerate().map(|(idx, &x)| {
                if x == 0 {
                    return (mapping[0], idx);
                }
                let mut x = x;
                let mut m = 0;
                let mut i = 1;
                while x > 0 {
                    let c = x%10;
                    m += i*mapping[c as usize];
                    x = x/10;
                    i *= 10;
                }
                return (m, idx);  // mapped value and index pair
            }).collect::<Vec<_>>();
            mapped.sort_by(|&a, &b| {
               if a.0 < b.0 || (a.0 == b.0 && a.1 < b.1) {
                   return Less;
               }
               return Greater;
            });
            return mapped.iter().map(|&x| nums[x.1]).collect::<Vec<_>>();
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                mapping: vec![8,9,4,0,2,1,3,5,7,6],
                nums: vec![991,338,38],
                ans: vec![338,38,991],
            },
            Solution {
                mapping: vec![0,1,2,3,4,5,6,7,8,9],
                nums: vec![789,456,123],
                ans: vec![123,456,789],
            },
            Solution {
                mapping: vec![9,8,7,6,5,4,3,2,1,0],
                nums: vec![0,1,2,3,4,5,6,7,8,9],
                ans: vec![9,8,7,6,5,4,3,2,1,0],
            },
        ];
        for i in testcases {
            let ans = Solution::sort_jumbled(i.mapping, i.nums);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
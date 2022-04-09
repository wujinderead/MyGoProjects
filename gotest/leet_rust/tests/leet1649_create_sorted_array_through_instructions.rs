// https://leetcode.com/problems/create-sorted-array-through-instructions/

// Given an integer array instructions, you are asked to create a sorted array from the elements
// in instructions. You start with an empty container nums. For each element from left to right
// in instructions, insert it into nums. The cost of each insertion is the minimum of the following:
//   The number of elements currently in nums that are strictly less than instructions[i].
//   The number of elements currently in nums that are strictly greater than instructions[i].
// For example, if inserting element 3 into nums = [1,2,3,5], the cost of insertion is min(2, 1)
// (elements 1 and 2 are less than 3, element 5 is greater than 3) and nums will become [1,2,3,3,5].
// Return the total cost to insert all elements from instructions into nums. Since the answer may be
// large, return it modulo 10⁹ + 7
// Example 1:
//   Input: instructions = [1,5,6,2]
//   Output: 1
//   Explanation: Begin with nums = [].
//     Insert 1 with cost min(0, 0) = 0, now nums = [1].
//     Insert 5 with cost min(1, 0) = 0, now nums = [1,5].
//     Insert 6 with cost min(2, 0) = 0, now nums = [1,5,6].
//     Insert 2 with cost min(1, 2) = 1, now nums = [1,2,5,6].
//     The total cost is 0 + 0 + 0 + 1 = 1.
// Example 2:
//   Input: instructions = [1,2,3,6,5,4]
//   Output: 3
//   Explanation: Begin with nums = [].
//     Insert 1 with cost min(0, 0) = 0, now nums = [1].
//     Insert 2 with cost min(1, 0) = 0, now nums = [1,2].
//     Insert 3 with cost min(2, 0) = 0, now nums = [1,2,3].
//     Insert 6 with cost min(3, 0) = 0, now nums = [1,2,3,6].
//     Insert 5 with cost min(3, 1) = 1, now nums = [1,2,3,5,6].
//     Insert 4 with cost min(3, 2) = 2, now nums = [1,2,3,4,5,6].
//     The total cost is 0 + 0 + 0 + 0 + 1 + 2 = 3.
// Example 3:
//   Input: instructions = [1,3,3,3,2,4,2,1,2]
//   Output: 4
//   Explanation: Begin with nums = [].
//   Insert 1 with cost min(0, 0) = 0, now nums = [1].
//   Insert 3 with cost min(1, 0) = 0, now nums = [1,3].
//   Insert 3 with cost min(1, 0) = 0, now nums = [1,3,3].
//   Insert 3 with cost min(1, 0) = 0, now nums = [1,3,3,3].
//   Insert 2 with cost min(1, 3) = 1, now nums = [1,2,3,3,3].
//   Insert 4 with cost min(5, 0) = 0, now nums = [1,2,3,3,3,4].
//   Insert 2 with cost min(1, 4) = 1, now nums = [1,2,2,3,3,3,4].
//   Insert 1 with cost min(0, 6) = 0, now nums = [1,1,2,2,3,3,3,4].
//   Insert 2 with cost min(2, 4) = 2, now nums = [1,1,2,2,2,3,3,3,4].
//   The total cost is 0 + 0 + 0 + 0 + 1 + 0 + 1 + 0 + 2 = 4.
// Constraints:
//   1 <= instructions.length <= 10⁵
//   1 <= instructions[i] <= 10⁵

mod _create_sorted_array_through_instructions {
    struct Solution{
        instructions: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        pub fn create_sorted_array(instructions: Vec<i32>) -> i32 {
            let mut less = instructions.iter().enumerate().map(|(i, &s) | (s, 0, i)).collect::<Vec<_>>();
            let mut greater = instructions.iter().enumerate().map(|(i, &s) | (s, 0, i)).collect::<Vec<_>>();
            let mut buf = Vec::<(i32, i32, usize)>::with_capacity(instructions.len());
            Solution::merge_sort(&mut less, &mut buf, 0, instructions.len()-1, |a: i32, b: i32| a < b);
            Solution::merge_sort(&mut greater, &mut buf, 0, instructions.len()-1, |a: i32, b: i32| a > b);
            // restore the index order
            let mut min = vec![0; instructions.len()];
            let mut sum = 0;
            for t in less {
                min[t.2] = t.1;
            }
            for t in greater {
                sum += min[t.2].min(t.1);
                sum %= 1e9 as i32 + 7;
            }
            return sum;
        }

        fn merge_sort(nums: &mut Vec<(i32, i32, usize)>, buf: &mut Vec<(i32, i32, usize)>, start: usize, end: usize, f: fn(i32, i32) -> bool) {
            if start == end {
                return;
            }
            let mid = (start+end)/2;
            Solution::merge_sort(nums, buf, start, mid, f);
            Solution::merge_sort(nums, buf, mid+1, end, f);
            buf.clear();
            let (mut i, mut j) = (start, mid+1);
            while i <= mid && j<= end {
                if f(nums[i].0, nums[j].0) {
                    buf.push(nums[i]);
                    i += 1;
                } else {
                    buf.push((nums[j].0, nums[j].1+(i-start) as i32, nums[j].2));
                    j += 1;
                }
            }
            while i <= mid {
                buf.push(nums[i]);
                i += 1;
            }
            while j <= end {
                buf.push((nums[j].0, nums[j].1+(i-start) as i32, nums[j].2));
                j += 1;
            }
            nums.as_mut_slice()[start..=end].copy_from_slice(buf.as_slice());
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                instructions: vec![1,5,6,2],
                ans: 1,
            },
            Solution {
                instructions: vec![1,2,3,6,5,4],
                ans: 3,
            },
            Solution {
                instructions: vec![1,3,3,3,2,4,2,1,2],
                ans: 4,
            },
        ];
        for i in testcases {
            let ans = Solution::create_sorted_array(i.instructions);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
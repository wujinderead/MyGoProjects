// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/

// You have k lists of sorted integers in non-decreasing order. Find the smallest range that
// includes at least one number from each of the k lists.
// We define the range [a, b] is smaller than range [c, d] if b - a < d - c or a < c if b - a == d - c.
// Example 1:
//   Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
//   Output: [20,24]
//   Explanation:
//     List 1: [4, 10, 15, 24, 26], 24 is in range [20,24].
//     List 2: [0, 9, 12, 20], 20 is in range [20,24].
//     List 3: [5, 18, 22, 30], 22 is in range [20,24].
// Example 2:
//   Input: nums = [[1,2,3],[1,2,3],[1,2,3]]
//   Output: [1,1]
// Constraints:
//   nums.length == k
//   1 <= k <= 3500
//   1 <= nums[i].length <= 50
//   -10⁵ <= nums[i][j] <= 10⁵
//   nums[i] is sorted in non-decreasing order.

mod _smallest_range_covering_elements_from_k_lists {
    struct Solution{
        nums: Vec<Vec<i32>>,
        ans: Vec<i32>,
    }

    impl Solution {
        pub fn smallest_range(nums: Vec<Vec<i32>>) -> Vec<i32> {
            // sort all list: can use merge k sorted list
            let all = nums.iter().fold(0, |acc, x| acc + x.len());
            let mut ns = Vec::<(i32, usize)>::with_capacity(all);
            nums.iter().enumerate().for_each(|(i, v)|
                v.iter().for_each(|&vv| ns.push((vv, i))));
            ns.sort_by_key(|s| s.0);
            // initial interval: [min_value, max_value]
            let (mut l, mut r) = (ns.first().unwrap().0, ns.last().unwrap().0);
            let mut ls = 0;             // how many list covered
            let nl = nums.len() as i32; // list number
            let mut cs = vec![0; nums.len()]; // occurrence of each list
            // sliding window
            let (mut i, mut j) = (0, 0);
            while i < ns.len() {
                while ls < nl && i<ns.len() {
                    // include ns[i] in the window
                    if cs[ns[i].1] == 0 {  // new list
                        ls += 1;
                    }
                    cs[ns[i].1] += 1;
                    i += 1;
                }
                while ls == nl {  // [j,i-1] is a valid interval that cover all lists
                    if ns[i-1].0 - ns[j].0 < r-l {
                        r = ns[i-1].0;
                        l = ns[j].0;
                    }
                    // exclude ns[j]
                    if cs[ns[j].1] == 1 {
                        ls -= 1;
                    }
                    cs[ns[j].1] -= 1;
                    j += 1;
                }
            }
            return vec![l, r];
        }

        // for k lists, keep a heap contains the smallest value of each list. each time we pop the
        // minimal value, we compare it to the maximal value of the heap
        pub fn smallest_range_heap(nums: Vec<Vec<i32>>) -> Vec<i32> {
            use std::cmp::Reverse;
            use std::collections::BinaryHeap;
            let mut heap = BinaryHeap::new();
            let mut max = -100000;
            let (mut l, mut r) = (-100000, 100000);
            for (i, ns) in nums.iter().enumerate() {
                heap.push(Reverse((ns[0], i, 0)));
                max = max.max(ns[0]);
            }
            while heap.len() == nums.len() {
                let Reverse((min, l_ind, n_ind)) = heap.pop().unwrap();
                if max-min < r-l {
                    r = max;
                    l = min;
                }
                if n_ind+1 < nums[l_ind].len() { // push next value of the list
                    heap.push(Reverse((nums[l_ind][n_ind+1], l_ind, n_ind+1)));
                    max = max.max(nums[l_ind][n_ind+1]);
                }
            }
            return vec![l,r];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![vec![4,10,15,24,26],vec![0,9,12,20],vec![5,18,22,30]],
                ans: vec![20,24],
            },
            Solution {
                nums: vec![vec![1,2,3],vec![1,2,3],vec![1,2,3]],
                ans: vec![1,1],
            },
            Solution {
                nums: vec![vec![1,2,3]],
                ans: vec![1,1],
            },
            Solution {
                nums: vec![vec![1],vec![6],vec![4]],
                ans: vec![1,6],
            },
            Solution {
                nums: vec![vec![1,6,10],vec![2,6,9],vec![6,7,12]],
                ans: vec![6,6],
            },
        ];
        for i in &testcases {
            let ans = Solution::smallest_range(i.nums.clone());
            println!("{:?}, {:?}", ans, i.ans);
        }
        for i in &testcases {
            let ans = Solution::smallest_range_heap(i.nums.clone());
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/maximum-segment-sum-after-removals/

// You are given two 0-indexed integer arrays nums and removeQueries, both of length n.
// For the iᵗʰ query, the element in nums at the index removeQueries[i] is removed,
// splitting nums into different segments.
// A segment is a contiguous sequence of positive integers in nums. A segment sum is the sum
// of every element in a segment.
// Return an integer array answer, of length n, where answer[i] is the maximum segment sum after
// applying the iᵗʰ removal.
// Note: The same index will not be removed more than once.
// Example 1:
//   Input: nums = [1,2,5,6,1], removeQueries = [0,3,2,4,1]
//   Output: [14,7,2,2,0]
//   Explanation: Using 0 to indicate a removed element, the answer is as follows:
//     Query 1: Remove the 0th element, nums becomes [0,2,5,6,1] and the maximum
//     segment sum is 14 for segment [2,5,6,1].
//     Query 2: Remove the 3rd element, nums becomes [0,2,5,0,1] and the maximum
//     segment sum is 7 for segment [2,5].
//     Query 3: Remove the 2nd element, nums becomes [0,2,0,0,1] and the maximum
//     segment sum is 2 for segment [2].
//     Query 4: Remove the 4th element, nums becomes [0,2,0,0,0] and the maximum
//     segment sum is 2 for segment [2].
//     Query 5: Remove the 1st element, nums becomes [0,0,0,0,0] and the maximum
//     segment sum is 0, since there are no segments.
//     Finally, we return [14,7,2,2,0].
// Example 2:
//   Input: nums = [3,2,11,1], removeQueries = [3,2,1,0]
//   Output: [16,5,3,0]
//   Explanation: Using 0 to indicate a removed element, the answer is as follows:
//     Query 1: Remove the 3rd element, nums becomes [3,2,11,0] and the maximum
//     segment sum is 16 for segment [3,2,11].
//     Query 2: Remove the 2nd element, nums becomes [3,2,0,0] and the maximum segment
//     sum is 5 for segment [3,2].
//     Query 3: Remove the 1st element, nums becomes [3,0,0,0] and the maximum segment
//     sum is 3 for segment [3].
//     Query 4: Remove the 0th element, nums becomes [0,0,0,0] and the maximum segment
//     sum is 0, since there are no segments.
//     Finally, we return [16,5,3,0].
// Constraints:
//   n == nums.length == removeQueries.length
//   1 <= n <= 10⁵
//   1 <= nums[i] <= 10⁹
//   0 <= removeQueries[i] < n
//   All the values of removeQueries are unique.

mod _maximum_segment_sum_after_removals {
    struct Solution{
        nums: Vec<i32>,
        remove_queries: Vec<i32>,
        ans: Vec<i64>,
    }

    use std::collections::{BinaryHeap, HashSet, BTreeSet};
    use std::ops::Bound::{Excluded, Unbounded};

    // use tree map to perform split
    impl Solution {
        pub fn maximum_segment_sum(nums: Vec<i32>, remove_queries: Vec<i32>) -> Vec<i64> {
            let n = nums.len();
            // prefix sum for range query
            let mut prefix = vec![0 as i64; n+1];
            for (i, &v) in nums.iter().enumerate() {
                prefix[i+1] = prefix[i] + v as i64; // sum(a[i..j]) = prefix[j+1]-prefix[i]
            }
            let mut ans = vec![0; n];
            let mut set = HashSet::new();
            let mut heap = BinaryHeap::new();
            let mut tree = BTreeSet::new();
            set.insert((-1, n as i32));
            tree.insert(-1);
            tree.insert(n as i32);
            for (i, &v) in remove_queries.iter().enumerate() {
                // get current boundary of remove index
                let &upper = tree.range((Excluded(&v), Unbounded)).next().unwrap();
                let &lower = tree.range((Unbounded, Excluded(&v))).rev().next().unwrap();
                tree.insert(v);
                set.remove(&(lower+1, upper-1));
                if upper-v > 1 {
                    set.insert((v+1, upper-1));
                    heap.push((prefix[upper as usize]-prefix[v as usize+1], v+1, upper-1));
                }
                if v-lower > 1 {
                    set.insert((lower+1, v-1));
                    heap.push((prefix[v as usize]-prefix[(lower+1) as usize], lower+1, v-1));
                }
                while heap.len()>0 {
                    let (val, l, r) = *heap.peek().unwrap();
                    if set.contains(&(l, r)) {
                        ans[i] = val;
                        break;
                    } else {
                        heap.pop();
                    }
                }
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,2,5,6,1],
                remove_queries: vec![0,3,2,4,1],
                ans: vec![14,7,2,2,0],
            },
            Solution {
                nums: vec![3,2,11,1],
                remove_queries: vec![3,2,1,0],
                ans: vec![16,5,3,0],
            },
        ];
        for i in testcases {
            let ans = Solution::maximum_segment_sum(i.nums, i.remove_queries);
            println!("{:?}, {:?}", ans, i.ans);
        }
    }
}
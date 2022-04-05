// https://leetcode.com/problems/maximum-candies-allocated-to-k-children/

// You are given a 0-indexed integer array candies. Each element in the array denotes a pile
// of candies of size candies[i]. You can divide each pile into any number of sub piles, but
// you cannot merge two piles together.
// You are also given an integer k. You should allocate piles of candies to k children such
// that each child gets the same number of candies. Each child can take at most one pile of
// candies and some piles of candies may go unused.
// Return the maximum number of candies each child can get.
// Example 1:
//   Input: candies = [5,8,6], k = 3
//   Output: 5
//   Explanation: We can divide candies[1] into 2 piles of size 5 and 3, and
//    candies[2] into 2 piles of size 5 and 1. We now have five piles of candies of sizes 5,
//    5, 3, 5, and 1. We can allocate the 3 piles of size 5 to 3 children. It can be
//    proven that each child cannot receive more than 5 candies.
// Example 2:
//   Input: candies = [2,5], k = 11
//   Output: 0
//   Explanation: There are 11 children but only 7 candies in total, so it is
//     impossible to ensure each child receives at least one candy. Thus, each child gets
//     no candy and the answer is 0.
// Constraints:
//   1 <= candies.length <= 10⁵
//   1 <= candies[i] <= 10⁷
//   1 <= k <= 10¹²

mod _maximum_candies_allocated_to_k_children {
    struct Solution{
        candies: Vec<i32>,
        k: i64,
        ans: i32,
    }

    impl Solution {
        pub fn maximum_candies(candies: Vec<i32>, k: i64) -> i32 {
            let candies = candies.iter().map(|&s| s as i64).collect::<Vec<_>>();
            let (mut left, mut right) = (1, 1e12 as i64);
            // for example, left=5 and right=6
            // if 5 can, 6 can't:
            //   left=5, right=6, mid=5,
            //     then left=6
            //   left=6, right=6, mid=6,
            //     then right=5
            //   return 5
            // if 5 can, 6 can:
            //   left=5, right=6, mid=5,
            //     then left=6
            //   left=6, right=6, mid=6,
            //     then left=7
            //   return 6
            while left <= right {
                let mid = (left+right)/2;
                let mut sum = 0;
                for &c in &candies {
                    sum += c/mid;
                }
                if sum < k {
                    right = mid-1;
                } else {
                    left = mid+1;
                }
            }
            return right as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                candies: vec![5,8,6],
                k: 3,
                ans: 5,
            },
            Solution {
                candies: vec![2,5],
                k: 11,
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::maximum_candies(i.candies, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
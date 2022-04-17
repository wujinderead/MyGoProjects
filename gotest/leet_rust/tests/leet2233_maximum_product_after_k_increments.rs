// https://leetcode.com/problems/maximum-product-after-k-increments/

// You are given an array of non-negative integers nums and an integer k. In one operation, you may
// choose any element from nums and increment it by 1.
// Return the maximum product of nums after at most k operations. Since the answer may be very
// large, return it modulo 10⁹ + 7.
// Example 1:
//   Input: nums = [0,4], k = 5
//   Output: 20
//   Explanation: Increment the first number 5 times.
//     Now nums = [5, 4], with a product of 5 * 4 = 20.
//     It can be shown that 20 is maximum product possible, so we return 20.
//     Note that there may be other ways to increment nums to have the maximum product.
// Example 2:
//   Input: nums = [6,3,3,2], k = 2
//   Output: 216
//   Explanation: Increment the second number 1 time and increment the fourth number 1 time.
//     Now nums = [6, 4, 3, 3], with a product of 6 * 4 * 3 * 3 = 216.
//     It can be shown that 216 is maximum product possible, so we return 216.
//     Note that there may be other ways to increment nums to have the maximum product.
// Constraints:
//   1 <= nums.length, k <= 10⁵
//   0 <= nums[i] <= 10⁶

mod _maximum_product_after_k_increments {
    struct Solution{
        nums: Vec<i32>,
        k: i32,
        ans: i32,
    }

    // always increment the lowest number
    impl Solution {
        pub fn maximum_product(nums: Vec<i32>, k: i32) -> i32 {
            // sort the array
            let mut ns = nums.iter().map(|&s| s as i64).collect::<Vec<_>>();
            let mut kk = k as i64;
            ns.sort();
            let mut pre = ns[0];
            let mut i = 1;
            while i < nums.len() {
                if ns[i] * (i as i64) - pre > kk {
                    break;
                }
                pre += ns[i];
                i += 1;
            }
            // ns[:i] can be added as ns[i-1]
            kk -= ns[i-1] * (i as i64 -1) - (pre-ns[i-1]);
            let x = ns[i-1];
            // still remains k increments, these increment need be parted averagely
            for j in 0..i {
                ns[j] = x + kk / (i as i64);
            }
            for j in 0..(kk as usize)%i {
                ns[j] += 1;
            }
            // get final result
            let mut ans = 1;
            for x in ns {
                ans = (ans * x) % 1000000007;
            }
            return ans as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![0,4],
                k: 5,
                ans: 20,
            },
            Solution {
                nums: vec![6,3,3,2],
                k: 2,
                ans: 216,
            },
            Solution {
                nums: vec![1,5,7],
                k: 2,
                ans: 105,
            },
            Solution {
                nums: vec![1,1,5,7],
                k: 3,
                ans: 210,
            }
        ];
        for i in testcases {
            let ans = Solution::maximum_product(i.nums, i.k);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
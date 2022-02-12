// https://leetcode.com/problems/largest-divisible-subset/

// Given a set of distinct positive integers nums, return the largest subset answer such that
// every pair (answer[i], answer[j]) of elements in this subset satisfies:
//   answer[i] % answer[j] == 0, or
//   answer[j] % answer[i] == 0
// If there are multiple solutions, return any of them.
// Example 1:
//   Input: nums = [1,2,3]
//   Output: [1,2]
//   Explanation: [1,3] is also accepted.
// Example 2:
//   Input: nums = [1,2,4,8]
//   Output: [1,2,4,8]
// Constraints:
//   1 <= nums.length <= 1000
//   1 <= nums[i] <= 2 * 10⁹
//   All the integers in nums are unique.

mod _largest_divisible_subset {
    struct Solution{
        nums: Vec<i32>,
        ans: Vec<i32>,
    }

    impl Solution {
        pub fn largest_divisible_subset(mut nums: Vec<i32>) -> Vec<i32> {
            let (mut ns, mut prev) = (vec![0; nums.len()], vec![0; nums.len()]);
            nums.sort();
            let mut ans = Vec::new();
            let mut i = 0;
            let mut maxi = 0;
            while i < nums.len() {
                ns[i] = 1;
                prev[i] = i;
                let mut j = 0;
                while j < i {
                    if nums[i] %  nums[j] == 0 {
                        if ns[j] + 1 > ns[i] {
                            ns[i] = ns[j] + 1;
                            prev[i] = j;
                        }
                    }
                    j += 1;
                }
                if ns[i] > ns[maxi] {
                    maxi = i;
                }
                i += 1;
            }
            while prev[maxi] != maxi {
                ans.push(nums[maxi]);
                maxi = prev[maxi];
            }
            ans.push(nums[maxi]);
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![1,2,3],
                ans: vec![1,2],
            },
            Solution {
                nums: vec![1,2,4,8],
                ans: vec![1,2,4,8],
            },
            Solution {
                nums: vec![3,5,7],
                ans: vec![3],
            },
            Solution {
                nums: vec![1,2,3,4,6,8],
                ans: vec![1,2,4,8],
            },
        ];
        for i in testcases {
            let ans = Solution::largest_divisible_subset(i.nums);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
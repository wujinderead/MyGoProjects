// https://leetcode.com/problems/minimum-operations-to-halve-array-sum/

// You are given an array nums of positive integers. In one operation, you can
// choose any number from nums and reduce it to exactly half the number. (Note that
// you may choose this reduced number in future operations.)
// Return the minimum number of operations to reduce the sum of nums by at least half.
// Example 1:
//   Input: nums = [5,19,8,1]
//   Output: 3
//   Explanation: The initial sum of nums is equal to 5 + 19 + 8 + 1 = 33.
//     The following is one of the ways to reduce the sum by at least half:
//     Pick the number 19 and reduce it to 9.5.
//     Pick the number 9.5 and reduce it to 4.75.
//     Pick the number 8 and reduce it to 4.
//     The final array is [5, 4.75, 4, 1] with a total sum of 5 + 4.75 + 4 + 1 = 14.75.
//     The sum of nums has been reduced by 33 - 14.75 = 18.25, which is at least half of the initial sum, 18.25 >= 33/2 = 16.5.
//     Overall, 3 operations were used so we return 3.
//     It can be shown that we cannot reduce the sum by at least half in less than 3 operations.
// Example 2:
//   Input: nums = [3,8,20]
//   Output: 3
//   Explanation: The initial sum of nums is equal to 3 + 8 + 20 = 31.
//     The following is one of the ways to reduce the sum by at least half:
//     Pick the number 20 and reduce it to 10.
//     Pick the number 10 and reduce it to 5.
//     Pick the number 3 and reduce it to 1.5.
//     The final array is [1.5, 8, 5] with a total sum of 1.5 + 8 + 5 = 14.5.
//     The sum of nums has been reduced by 31 - 14.5 = 16.5, which is at least half of the initial sum, 16.5 >= 31/2 = 16.5.
//     Overall, 3 operations were used so we return 3.
//     It can be shown that we cannot reduce the sum by at least half in less than 3 operations.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁷

mod _minimum_operations_to_halve_array_sum {
    struct Solution{
        nums: Vec<i32>,
        ans: i32,
    }

    use std::cmp::Ordering;
    use std::collections::BinaryHeap;

    // max-heap to always halve the max value, since f64 is not Ord, we need wrap f64 as NonNaN
    impl Solution {
        pub fn halve_array(nums: Vec<i32>) -> i32 {
            #[derive(PartialEq)]
            struct NonNan(f64);
            impl Eq for NonNan {}
            impl PartialOrd for NonNan {
                fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
                    self.0.partial_cmp(&other.0)
                }
            }
            impl Ord for NonNan {
                fn cmp(&self, other: &NonNan) -> Ordering {
                    self.partial_cmp(other).unwrap()
                }
            }
            let floats = nums.iter().map(|&x| x as f64).collect::<Vec<_>>();
            let mut cur_sum = floats.iter().sum::<f64>();
            let target = cur_sum / 2.0;
            let mut heap = BinaryHeap::new();
            floats.iter().for_each(|&v| heap.push(NonNan(v)));
            let mut step = 0;
            while cur_sum > target {
                let mut val = heap.peek_mut().unwrap();
                cur_sum -= (*val).0 / 2.0;
                step += 1;
                *val = NonNan((*val).0 / 2.0);
            }
            return step;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                nums: vec![5,19,8,1],
                ans: 3,
            },
            Solution {
                nums: vec![3,8,20],
                ans: 3,
            },
            Solution {
                nums: vec![10000000; 100000],
                ans: 100000,
            }
        ];
        for i in testcases {
            let ans = Solution::halve_array(i.nums);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
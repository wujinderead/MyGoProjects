// https://leetcode.com/problems/count-the-hidden-sequences/

// You are given a 0-indexed array of n integers differences, which describes
// the differences between each pair of consecutive integers of a hidden sequence of
// length (n + 1). More formally, call the hidden sequence hidden, then we have
// that differences[i] = hidden[i + 1] - hidden[i].
// You are further given two integers lower and upper that describe the 
// inclusive range of values [lower, upper] that the hidden sequence can contain.
// For example, given differences = [1, -3, 4], lower = 1, upper = 6, the
// hidden sequence is a sequence of length 4 whose elements are in between 1 and 6 (inclusive).
//   [3, 4, 1, 5] and [4, 5, 2, 6] are possible hidden sequences.
//   [5, 6, 3, 7] is not possible since it contains an element greater than 6.
//   [1, 2, 3, 4] is not possible since the differences are not correct.
// Return the number of possible hidden sequences there are. If there are no 
// possible sequences, return 0.
// Example 1:
//   Input: differences = [1,-3,4], lower = 1, upper = 6
//   Output: 2
//   Explanation: The possible hidden sequences are:
//     - [3, 4, 1, 5]
//     - [4, 5, 2, 6]
//     Thus, we return 2.
// Example 2:
//   Input: differences = [3,-4,5,1,-2], lower = -4, upper = 5
//   Output: 4
//   Explanation: The possible hidden sequences are:
//     - [-3, 0, -4, 1, 2, 0]
//     - [-2, 1, -3, 2, 3, 1]
//     - [-1, 2, -2, 3, 4, 2]
//     - [0, 3, -1, 4, 5, 3]
//     Thus, we return 4.
// Example 3:
//   Input: differences = [4,-7,2], lower = 3, upper = 6
//   Output: 0
//   Explanation: There are no possible hidden sequences. Thus, we return 0.
// Constraints:
//   n == differences.length
//   1 <= n <= 10⁵
//   -10⁵ <= differences[i] <= 10⁵
//   -10⁵ <= lower <= upper <= 10⁵

mod _count_the_hidden_sequences {
    struct Solution{
        diff: Vec<i32>,
        low: i32,
        up: i32,
        ans: i32,
    }

    // let 0 be the first number. generate the sequence based on the differences,
    // record the maximal and minimal values, so we can compare the difference between
    // (max-min) and (upper-lower).
    impl Solution {
        pub fn number_of_arrays(differences: Vec<i32>, lower: i32, upper: i32) -> i32 {
            let (mut min, mut max, mut f) = (0i64, 0i64, 0i64);
            for &item in differences.iter() {
                f = f + item as i64;
                min = min.min(f);
                max = max.max(f);
            }
            if max - min > (upper - lower) as i64 {
                return 0;
            }
            return (upper-lower)-((max-min) as i32)+1;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                diff: vec![1,-3,4],
                low: 1,
                up: 6,
                ans: 2,
            },
            Solution {
                diff: vec![3,-4,5,1,-2],
                low: -4,
                up: 5,
                ans: 4,
            },
            Solution {
                diff: vec![4,-7,2],
                low: 3,
                up: 6,
                ans: 0,
            },
            Solution {
                diff: vec![-40],
                low: -46,
                up: 53,
                ans: 60,
            },
            Solution {
                diff: vec![100000; 100000],
                low: 100000,
                up: 100000,
                ans: 0,
            }
        ];
        for i in testcases {
            let ans = Solution::number_of_arrays(i.diff, i.low, i.up);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
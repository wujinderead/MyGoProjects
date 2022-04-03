// https://leetcode.com/problems/find-palindrome-with-fixed-length/

// Given an integer array queries and a positive integer intLength, return an array answer
// where answer[i] is either the queries[i]ᵗʰ smallest positive palindrome of length intLength
// or -1 if no such palindrome exists.
// A palindrome is a number that reads the same backwards and forwards. 
// Palindromes cannot have leading zeros.
// Example 1:
//   Input: queries = [1,2,3,4,5,90], intLength = 3
//   Output: [101,111,121,131,141,999]
//   Explanation:
//     The first few palindromes of length 3 are:
//     101, 111, 121, 131, 141, 151, 161, 171, 181, 191, 202, ...
//     The 90ᵗʰ palindrome of length 3 is 999.
// Example 2:
//   Input: queries = [2,4,6], intLength = 4
//   Output: [1111,1331,1551]
//   Explanation:
//     The first six palindromes of length 4 are:
//     1001, 1111, 1221, 1331, 1441, and 1551.
// Constraints:
//   1 <= queries.length <= 5 * 10⁴
//   1 <= queries[i] <= 10⁹
//   1 <= intLength <= 15

mod _find_palindrome_with_fixed_length {
    struct Solution{
        queries: Vec<i32>,
        int_length: i32,
        ans: Vec<i64>,
    }

    impl Solution {
        // e.g., for int_length=4,the left 2 digits can be 10 to 99,
        // for query=5, the left 2 digits will be 10+5-1=14,
        // so the 5-th palindrome of length 4 is 1441
        pub fn kth_palindrome(queries: Vec<i32>, int_length: i32) -> Vec<i64> {
            let mut max = 9;
            for i in 2..=int_length {
                max = if i%2 == 1 { max*10 } else { max };
            }
            let base = max/9;
            let mut ans = vec![0; queries.len()];
            for (i, &q) in queries.iter().enumerate() {
                if q > max {
                    ans[i] = -1;
                    continue;
                }
                let mut x = (base+q-1) as i64;
                let mut xx = x;
                if int_length % 2 == 1 {
                    xx /= 10;
                }
                for _i in 0..(int_length/2) {
                    x = x*10+(xx%10);
                    xx /= 10;
                }
                ans[i] = x;
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                queries: vec![1,2,3,4,5,90],
                int_length: 3,
                ans: vec![101,111,121,131,141,999],
            },
            Solution {
                queries: vec![2,4,6],
                int_length: 4,
                ans: vec![1111,1331,1551],
            },
            Solution {
                queries: vec![1,2,3,4,5,6,7,8,9,10],
                int_length: 1,
                ans: vec![1,2,3,4,5,6,7,8,9,-1],
            },
            Solution {
                queries: vec![1,2,3,4,5,6,7,8,9,10],
                int_length: 2,
                ans: vec![11,22,33,44,55,66,77,88,99,-1],
            },
        ];
        for i in testcases {
            let ans = Solution::kth_palindrome(i.queries, i.int_length);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
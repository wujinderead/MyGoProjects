// https://leetcode.com/problems/maximum-good-people-based-on-statements/

// There are two types of persons:
//   The good person: The person who always tells the truth.
//   The bad person: The person who might tell the truth and might lie.
// You are given a 0-indexed 2D integer array statements of size n x n that represents the
// statements made by n people about each other. More specifically, statements[i][j] could be
// one of the following:
//   0 which represents a statement made by person i that person j is a bad person.
//   1 which represents a statement made by person i that person j is a good person.
//   2 represents that no statement is made by person i about person j.
//   Additionally, no person ever makes a statement about themselves. Formally,
//     we have that statements[i][i] = 2 for all 0 <= i < n.
// Return the maximum number of people who can be good based on the statements made by the n people.
// Example 1:
//   Input: statements = [[2,1,2],[1,2,2],[2,0,2]]
//   Output: 2
// Example 2:
//   Input: statements = [[2,0],[0,2]]
//   Output: 1
// Constraints:
//   n == statements.length == statements[i].length
//   2 <= n <= 15
//   statements[i][j] is either 0, 1, or 2.
//   statements[i][i] == 2

mod _maximum_good_people_based_on_statements {
    use std::collections::HashMap;

    struct Solution{
        stat: Vec<Vec<i32>>,
        ans: i32,
    }

    // just generate 2^n bitmasks. for a certain bitmask, check those good people's statements,
    // check if the statement violate the bitmask.
    impl Solution {
        pub fn maximum_good(statements: Vec<Vec<i32>>) -> i32 {
            let n = statements.len() as i32;
            let mut map = HashMap::new();
            for i in 1..(1<<n) {
                let bits = bit_count(i);
                map.entry(bits).or_insert(Vec::new()).push(i);
            }
            let mut i = n;
            while i>0 {
                let masks = map.get(&i).unwrap();
                for &mask in masks {
                    if can_this_bits(mask, &statements) {
                        return i;
                    }
                }
                i -= 1;
            }
            return 0;
        }
    }

    fn can_this_bits(mask: i32, statements: &Vec<Vec<i32>>) -> bool {
        let n = statements.len();
        let mut i = 0;
        while i<n {
            if (1<<i) & mask > 0 {
                let stat = statements.get(i).unwrap();
                for (ind, &j) in stat.iter().enumerate() {
                    if j==0 && (1<<ind) & mask > 0 {
                        return false;
                    }
                    if j==1 && (1<<ind) & mask == 0 {
                        return false;
                    }
                }
            }
            i += 1;
        }
        return true;
    }

    fn bit_count(i: i32) -> i32 {
        let mut i = i;
        i = i - ((i >> 1) & 0x55555555);
        i = (i & 0x33333333) + ((i >> 2) & 0x33333333);
        i = (i + (i >> 4)) & 0x0f0f0f0f;
        i = i + (i >> 8);
        i = i + (i >> 16);
        return i & 0x3f;
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                stat: vec![vec![2,1,2],vec![1,2,2],vec![2,0,2]],
                ans: 2,
            },
            Solution {
                stat: vec![vec![2,2,2],vec![2,2,2],vec![2,2,2]],
                ans: 3,
            },
            Solution {
                stat: vec![vec![2,1,1],vec![1,2,1],vec![1,1,2]],
                ans: 3,
            },
            Solution {
                stat: vec![vec![2,0],vec![0,2]],
                ans: 1,
            },
            Solution {
                stat: vec![vec![2,1],vec![1,2]],
                ans: 2,
            },
            Solution {
                stat: vec![vec![2,1],vec![0,2]],
                ans: 1,
            },
            Solution {
                stat: vec![vec![2,2],vec![0,2]],
                ans: 1,
            },
        ];
        for i in testcases {
            let ans = Solution::maximum_good(i.stat);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
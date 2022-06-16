// https://leetcode.com/problems/fruit-into-baskets/

// You are visiting a farm that has a single row of fruit trees arranged from left to right.
// The trees are represented by an integer array fruits where fruits[i] is the type of fruit
// the iᵗʰ tree produces.
// You want to collect as much fruit as possible. However, the owner has some strict rules that you
// must follow:
// You only have two baskets, and each basket can only hold a single type of fruit. There is no
//   limit on the amount of fruit each basket can hold.
// Starting from any tree of your choice, you must pick exactly one fruit from every tree (including
//   the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
// Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
// Given the integer array fruits, return the maximum number of fruits you can pick.
// Example 1:
//   Input: fruits = [1,2,1]
//   Output: 3
//   Explanation: We can pick from all 3 trees.
// Example 2:
//   Input: fruits = [0,1,2,2]
//   Output: 3
//   Explanation: We can pick from trees [1,2,2].
//     If we had started at the first tree, we would only pick from trees [0,1].
// Example 3:
//   Input: fruits = [1,2,3,2,2]
//   Output: 4
//   Explanation: We can pick from trees [2,3,2,2].
//     If we had started at the first tree, we would only pick from trees [1,2].
// Constraints:
//   1 <= fruits.length <= 10⁵
//   0 <= fruits[i] < fruits.length

mod _fruit_into_baskets {
    struct Solution{
        fruits: Vec<i32>,
        ans: i32,
    }

    // can not use map but just two variables
    use std::collections::HashMap;
    impl Solution {
        pub fn total_fruit(fruits: Vec<i32>) -> i32 {
            let (mut start, mut ans) = (0, 0);
            let mut map = HashMap::with_capacity(4);
            for i in 0..fruits.len() {
                *map.entry(&fruits[i]).or_insert(0) += 1;
                // while has more than 2 types, shrink window to only two types
                while map.len() > 2 {
                    let count = map.get_mut(&fruits[start]).unwrap();
                    *count -= 1;
                    if *count == 0 {
                        map.remove(&fruits[start]);
                    }
                    start += 1;
                }
                ans = ans.max(i-start+1);
            }
            return ans as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                fruits: vec![1,2,1],
                ans: 3,
            },
            Solution {
                fruits: vec![0,1,2,2],
                ans: 3,
            },
            Solution {
                fruits: vec![1,2,3,2,2],
                ans: 4,
            }
        ];
        for i in testcases {
            let ans = Solution::total_fruit(i.fruits);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
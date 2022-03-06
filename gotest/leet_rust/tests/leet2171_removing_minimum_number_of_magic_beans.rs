// https://leetcode.com/problems/removing-minimum-number-of-magic-beans/

// You are given an array of positive integers beans, where each integer represents the number
// of magic beans found in a particular magic bag.
// Remove any number of beans (possibly none) from each bag such that the number of beans in
// each remaining non-empty bag (still containing at least one bean) is equal. Once a bean has
// been removed from a bag, you are not allowed to return it to any of the bags.
// Return the minimum number of magic beans that you have to remove.
// Example 1:
//   Input: beans = [4,1,6,5]
//   Output: 4
//   Explanation:
//     - We remove 1 bean from the bag with only 1 bean.
//       This results in the remaining bags: [4,0,6,5]
//     - Then we remove 2 beans from the bag with 6 beans.
//       This results in the remaining bags: [4,0,4,5]
//     - Then we remove 1 bean from the bag with 5 beans.
//       This results in the remaining bags: [4,0,4,4]
//     We removed a total of 1 + 2 + 1 = 4 beans to make the remaining non-empty
//     bags have an equal number of beans.
//     There are no other solutions that remove 4 beans or fewer.
// Example 2:
//   Input: beans = [2,10,3,2]
//   Output: 7
//   Explanation:
//     - We remove 2 beans from one of the bags with 2 beans.
//       This results in the remaining bags: [0,10,3,2]
//     - Then we remove 2 beans from the other bag with 2 beans.
//       This results in the remaining bags: [0,10,3,0]
//     - Then we remove 3 beans from the bag with 3 beans.
//       This results in the remaining bags: [0,10,0,0]
//     We removed a total of 2 + 2 + 3 = 7 beans to make the remaining non-empty
//     bags have an equal number of beans.
//     There are no other solutions that removes 7 beans or fewer.
// Constraints:
//   1 <= beans.length <= 10⁵
//   1 <= beans[i] <= 10⁵

mod _removing_minimum_number_of_magic_beans {
    struct Solution{
        beans: Vec<i32>,
        ans: i64,
    }

    // sort, [1,4,5,6], if choose 1, the remain is [1,1,1,1];
    // if choose 4, the remain is [0,4,4,4]; so the removed is sum_array - beans[i]*(N-i)
    impl Solution {
        pub fn minimum_removal(beans: Vec<i32>) -> i64 {
            let mut bb = beans.iter().map(|&x| x as i64).collect::<Vec<_>>();
            bb.sort();
            let sum = bb.iter().sum::<i64>();
            let mut ans = sum;
            let mut i = 1;
            while i < bb.len() {
                ans = ans.min(sum - bb[i] * ((beans.len()-i) as i64));
                i += 1;
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                beans: vec![4,1,6,5],
                ans: 4,
            },
            Solution {
                beans: vec![2,10,3,2],
                ans: 7,
            },
        ];
        for i in testcases {
            let ans = Solution::minimum_removal(i.beans);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
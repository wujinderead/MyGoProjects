// https://leetcode.com/problems/find-players-with-zero-or-one-losses/

// You are given an integer array matches where matches[i] = [winneri, loseri]
// indicates that the player winneri defeated player loseri in a match:
// Return a list answer of size 2 where:
//   answer[0] is a list of all players that have not lost any matches.
//   answer[1] is a list of all players that have lost exactly one match.
// The values in the two lists should be returned in increasing order.
// Note:
//   You should only consider the players that have played at least one match.
//   The testcases will be generated such that no two matches will have the same outcome.
// Example 1:
//   Input: matches = [[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]]
//   Output: [[1,2,10],[4,5,7,8]]
//   Explanation:
//     Players 1, 2, and 10 have not lost any matches.
//     Players 4, 5, 7, and 8 each have lost one match.
//     Players 3, 6, and 9 each have lost two matches.
//     Thus, answer[0] = [1,2,10] and answer[1] = [4,5,7,8].
// Example 2:
//   Input: matches = [[2,3],[1,3],[5,4],[6,4]]
//   Output: [[1,2,5,6],[]]
//   Explanation:
//     Players 1, 2, 5, and 6 have not lost any matches.
//     Players 3 and 4 each have lost two matches.
//     Thus, answer[0] = [1,2,5,6] and answer[1] = [].
// Constraints:
//   1 <= matches.length <= 10⁵
//   matches[i].length == 2
//   1 <= winneri, loseri <= 10⁵
//   winneri != loseri
//   All matches[i] are unique.

mod _find_players_with_zero_or_one_losses {
    struct Solution{
        matches: Vec<Vec<i32>>,
        ans: Vec<Vec<i32>>,
    }

    // just use hash table
    use std::collections::{HashMap, HashSet};
    impl Solution {
        pub fn find_winners(matches: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
            let mut wins = HashSet::new();
            let mut loses = HashMap::new();
            for v in &matches {
                let (win, lose) = (v[0], v[1]);
                wins.insert(win);
                *loses.entry(lose).or_insert(0) += 1;
            }
            let mut w = wins.iter().filter(|&w| !loses.contains_key(w)).map(|&s| s).collect::<Vec<i32>>();
            let mut l = loses.iter().filter(|&(_k, v)| *v==1).map(|(&k, v)| k).collect::<Vec<i32>>();
            w.sort();
            l.sort();
            return vec![w, l];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                matches: vec![[1,3],[2,3],[3,6],[5,6],[5,7],[4,5],[4,8],[4,9],[10,4],[10,9]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![vec![1,2,10],vec![4,5,7,8]],
            },
            Solution {
                matches: vec![[2,3],[1,3],[5,4],[6,4]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![vec![1,2,5,6],vec![]],
            },
        ];
        for i in testcases {
            let ans = Solution::find_winners(i.matches);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
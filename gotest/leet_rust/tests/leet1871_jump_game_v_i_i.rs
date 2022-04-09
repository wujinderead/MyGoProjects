// https://leetcode.com/problems/jump-game-vii/

// You are given a 0-indexed binary string s and two integers minJump and maxJump. In the beginning,
// you are standing at index 0, which is equal to '0'. You can move from index i to index j if the
// following conditions are fulfilled:
//   i + minJump <= j <= min(i + maxJump, s.length - 1), and
//   s[j] == '0'.
// Return true if you can reach index s.length - 1 in s, or false otherwise.
// Example 1:
//   Input: s = "011010", minJump = 2, maxJump = 3
//   Output: true
//   Explanation:
//     In the first step, move from index 0 to index 3.
//     In the second step, move from index 3 to index 5.
// Example 2:
//   Input: s = "01101110", minJump = 2, maxJump = 3
//   Output: false
// Constraints:
//   2 <= s.length <= 10⁵
//   s[i] is either '0' or '1'.
//   s[0] == '0'
//   1 <= minJump <= maxJump < s.length

mod _jump_game_v_i_i {
    struct Solution{
        s: String,
        min_jump: i32,
        max_jump: i32,
        ans: bool,
    }

    // another solution: bfs
    impl Solution {
        // to check if position i reachable, check if [i-max_jump, i-min_jump] reachable
        // can use prefix sum or sliding window
        pub fn can_reach(s: String, min_jump: i32, max_jump: i32) -> bool {
            let ss = s.as_bytes();
            let mut can = vec![0; s.len()];  // can[i]==1 means reachable
            can[0] = 1;
            let mut prefix = vec![1; s.len()];  // sum(can[i...j]) = prefix[j+1]-prefix[i]
            prefix[0] = 0;  // because can[0]=1, so prefix[1...min_jump-1] = 1
            for i in (min_jump as usize)..s.len() {
                prefix[i] = prefix[i-1] + can[i-1];
                if ss[i] == b'1' {  // can't reach
                    continue;
                }
                let l = (i as i32-max_jump).max(0);
                // check segment of [l...i-min_jump]
                let seg_sum = prefix[(i- min_jump as usize +1) as usize] - prefix[l as usize];
                if seg_sum > 0 {
                    can[i] = 1;
                }
            }
            return can[s.len()-1] == 1;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                s: "011010".to_string(),
                min_jump: 2,
                max_jump: 3,
                ans: true,
            },
            Solution {
                s: "01101110".to_string(),
                min_jump: 2,
                max_jump: 3,
                ans: false,
            },
        ];
        for i in testcases {
            let ans = Solution::can_reach(i.s, i.min_jump, i.max_jump);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
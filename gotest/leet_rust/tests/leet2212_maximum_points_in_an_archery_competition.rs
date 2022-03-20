// https://leetcode.com/problems/maximum-points-in-an-archery-competition/

// Alice and Bob are opponents in an archery competition. The competition has
// set the following rules:
// Alice first shoots numArrows arrows and then Bob shoots numArrows arrows. 
// The points are then calculated as follows:
//   The target has integer scoring sections ranging from 0 to 11 inclusive.
//   For each section of the target with score k (in between 0 to 11), say Alice
//     and Bob have shot ak and bk arrows on that section respectively. If ak >= bk,
//     then Alice takes k points. If ak < bk, then Bob takes k points.
//   However, if ak == bk == 0, then nobody takes k points.
// For example, if Alice and Bob both shot 2 arrows on the section with score 11, then
// Alice takes 11 points. On the other hand, if Alice shot 0 arrows on the section with
// score 11 and Bob shot 2 arrows on that same section, then Bob takes 11 points.
// You are given the integer numArrows and an integer array aliceArrows of size 12, which
// represents the number of arrows Alice shot on each scoring section from 0 to 11.
// Now, Bob wants to maximize the total number of points he can obtain.
// Return the array bobArrows which represents the number of arrows Bob shot on 
// each scoring section from 0 to 11. The sum of the values in bobArrows should
// equal numArrows.
// If there are multiple ways for Bob to earn the maximum total points, return 
// any one of them.
// Example 1:
//   Input: numArrows = 9, aliceArrows = [1,1,0,1,0,0,2,1,0,1,2,0]
//   Output: [0,0,0,0,1,1,0,0,1,2,3,1]
//   Explanation: The table above shows how the competition is scored.
//     Bob earns a total point of 4 + 5 + 8 + 9 + 10 + 11 = 47.
//     It can be shown that Bob cannot obtain a score higher than 47 points.
// Example 2:
//   Input: numArrows = 3, aliceArrows = [0,0,1,0,0,0,0,0,0,0,0,2]
//   Output: [0,0,0,0,0,0,0,0,1,1,1,0]
//   Explanation: The table above shows how the competition is scored.
//     Bob earns a total point of 8 + 9 + 10 = 27.
//     It can be shown that Bob cannot obtain a score higher than 27 points.
// Constraints:
//   1 <= numArrows <= 10⁵
//   aliceArrows.length == bobArrows.length == 12
//   0 <= aliceArrows[i], bobArrows[i] <= numArrows
//   sum(aliceArrows[i]) == numArrows

mod _maximum_points_in_an_archery_competition {
    struct Solution{
        num_arrows: i32,
        alice_arrows: Vec<i32>,
        ans: Vec<i32>,
    }

    impl Solution {
        pub fn maximum_bob_points(num_arrows: i32, alice_arrows: Vec<i32>) -> Vec<i32> {
            let (mut max_score, mut max_mask) = (0, 0);
            // bit mask for bob's shooting
            for mask in 1..2048 {
                let (mut score, mut arrow) = (0, 0);
                for i in 0..11 {
                    if mask & (1<<i) > 0 {
                        score += i+1;
                        arrow += alice_arrows[i as usize + 1] + 1; // if bob want score, he must over alice
                    }
                }
                if arrow <= num_arrows && score > max_score { // check if possible
                    max_score = score;
                    max_mask = mask;
                }
            }
            let mut ans = vec![0; 12]; // get bob's win shoot
            for i in 0..11 {
                if max_mask & (1<<i) > 0 {
                    ans[i as usize+1] = alice_arrows[i as usize + 1] + 1;
                }
            }
            ans[0] = num_arrows - ans[1..].iter().sum::<i32>(); // for unused arrows
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                num_arrows: 9,
                alice_arrows: vec![1,1,0,1,0,0,2,1,0,1,2,0],
                ans: vec![0,0,0,0,1,1,0,0,1,2,3,1],
            },
            Solution {
                num_arrows: 3,
                alice_arrows: vec![0,0,1,0,0,0,0,0,0,0,0,2],
                ans: vec![0,0,0,0,0,0,0,0,1,1,1,0],
            },
            Solution {
                num_arrows: 89,
                alice_arrows: vec![3,2,28,1,7,1,16,7,3,13,3,5],
                ans: vec![21,3,0,2,8,2,17,8,4,14,4,6],
            }
        ];
        for i in testcases {
            let ans = Solution::maximum_bob_points(i.num_arrows, i.alice_arrows);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
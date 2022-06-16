// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/

// You are given two positive integer arrays spells and potions, of length n and m respectively,
// where spells[i] represents the strength of the iᵗʰ spell and potions[j] represents the strength
// of the jᵗʰ potion.
// You are also given an integer success. A spell and potion pair is considered successful if the
// product of their strengths is at least success.
// Return an integer array pairs of length n where pairs[i] is the number of potions that will form
// a successful pair with the iᵗʰ spell.
// Example 1:
//   Input: spells = [5,1,3], potions = [1,2,3,4,5], success = 7
//   Output: [4,0,3]
//   Explanation:
//     - 0ᵗʰ spell: 5 * [1,2,3,4,5] = [5,10,15,20,25]. 4 pairs are successful.
//     - 1ˢᵗ spell: 1 * [1,2,3,4,5] = [1,2,3,4,5]. 0 pairs are successful.
//     - 2ⁿᵈ spell: 3 * [1,2,3,4,5] = [3,6,9,12,15]. 3 pairs are successful.
//     Thus, [4,0,3] is returned.
// Example 2:
//   Input: spells = [3,1,2], potions = [8,5,8], success = 16
//   Output: [2,0,2]
//   Explanation:
//     - 0ᵗʰ spell: 3 * [8,5,8] = [24,15,24]. 2 pairs are successful.
//     - 1ˢᵗ spell: 1 * [8,5,8] = [8,5,8]. 0 pairs are successful.
//     - 2ⁿᵈ spell: 2 * [8,5,8] = [16,10,16]. 2 pairs are successful.
//     Thus, [2,0,2] is returned.
// Constraints:
//   n == spells.length
//   m == potions.length
//   1 <= n, m <= 10⁵
//   1 <= spells[i], potions[i] <= 10⁵
//   1 <= success <= 10¹⁰

mod _successful_pairs_of_spells_and_potions {
    struct Solution{
        spells: Vec<i32>,
        potions: Vec<i32>,
        success: i64,
        ans: Vec<i32>,
    }

    impl Solution {
        pub fn successful_pairs(spells: Vec<i32>, mut potions: Vec<i32>, success: i64) -> Vec<i32> {
            let mut ans = vec![0; spells.len()];
            potions.sort();
            let mut sp = spells.iter().enumerate().map(|(i, &v)| (i, v)).collect::<Vec<_>>();
            sp.sort_by_key(|s| -s.1);
            let mut pi = 0;
            for i in 0..sp.len() {
                let ind = sp[i].0;
                let s = sp[i].1;
                while pi < potions.len() && (potions[pi] as i64)*(s as i64) < success {
                    pi += 1;
                }
                ans[ind] = (potions.len()-pi) as i32;
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                spells: vec![5,1,3],
                potions: vec![1,2,3,4,5],
                success: 7,
                ans: vec![4,0,3],
            },
            Solution {
                spells: vec![3,1,2],
                potions: vec![8,5,8],
                success: 16,
                ans: vec![2,0,2],
            },
        ];
        for i in testcases {
            let ans = Solution::successful_pairs(i.spells, i.potions, i.success);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
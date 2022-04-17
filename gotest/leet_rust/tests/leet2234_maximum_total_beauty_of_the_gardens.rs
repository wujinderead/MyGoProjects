// https://leetcode.com/problems/maximum-total-beauty-of-the-gardens/

// Alice is a caretaker of n gardens and she wants to plant flowers to maximize the total beauty of
// all her gardens.
// You are given a 0-indexed integer array flowers of size n, where flowers[i] is the number of
// flowers already planted in the iᵗʰ garden. Flowers that are already planted cannot be removed.
// You are then given another integer newFlowers, which is the maximum number of flowers that Alice
// can additionally plant. You are also given the integers target, full, and partial.
// A garden is considered complete if it has at least target flowers. The total beauty of the
// gardens is then determined as the sum of the following:
//   The number of complete gardens multiplied by full.
//   The minimum number of flowers in any of the incomplete gardens multiplied by partial.
//     If there are no incomplete gardens, then this value will be 0.
// Return the maximum total beauty that Alice can obtain after planting at most newFlowers flowers.
// Example 1:
//   Input: flowers = [1,3,1,1], newFlowers = 7, target = 6, full = 12, partial = 1
//   Output: 14
//   Explanation: Alice can plant
//     - 2 flowers in the 0ᵗʰ garden
//     - 3 flowers in the 1ˢᵗ garden
//     - 1 flower in the 2ⁿᵈ garden
//     - 1 flower in the 3ʳᵈ garden
//     The gardens will then be [3,6,2,2]. She planted a total of 2 + 3 + 1 + 1 = 7 flowers.
//     There is 1 garden that is complete.
//     The minimum number of flowers in the incomplete gardens is 2.
//     Thus, the total beauty is 1 * 12 + 2 * 1 = 12 + 2 = 14.
//     No other way of planting flowers can obtain a total beauty higher than 14.
// Example 2:
//   Input: flowers = [2,4,5,3], newFlowers = 10, target = 5, full = 2, partial = 6
//   Output: 30
//   Explanation: Alice can plant
//     - 3 flowers in the 0ᵗʰ garden
//     - 0 flowers in the 1ˢᵗ garden
//     - 0 flowers in the 2ⁿᵈ garden
//     - 2 flowers in the 3ʳᵈ garden
//     The gardens will then be [5,4,5,5]. She planted a total of 3 + 0 + 0 + 2 = 5 flowers.
//     There are 3 gardens that are complete.
//     The minimum number of flowers in the incomplete gardens is 4.
//     Thus, the total beauty is 3 * 2 + 4 * 6 = 6 + 24 = 30.
//     No other way of planting flowers can obtain a total beauty higher than 30.
//     Note that Alice could make all the gardens complete but in this case,
//     she would obtain a lower total beauty.
// Constraints:
//   1 <= flowers.length <= 10⁵
//   1 <= flowers[i], target <= 10⁵
//   1 <= newFlowers <= 10¹⁰
//   1 <= full, partial <= 10⁵

mod _maximum_total_beauty_of_the_gardens {
    struct Solution{
        flowers: Vec<i32>,
        new_flowers: i64,
        target: i32,
        full: i32,
        partial: i32,
        ans: i64,
    }

    impl Solution {
        pub fn maximum_beauty(flowers: Vec<i32>, new_flowers: i64, target: i32, full: i32, partial: i32) -> i64 {
            let mut fs = flowers.iter()
                .filter(|&&f| f < target)
                .map(|&s| s as i64).collect::<Vec<_>>();  // exclude >= target
            fs.sort();
            let ans = (full as i64) * ((flowers.len()-fs.len()) as i64);
            if fs.len() == 0 {   // all >= target
                return ans;
            }
            // make prefix array
            let mut prefix = vec![0; fs.len()+1];
            for i in 1..=fs.len() {
                prefix[i] = prefix[i-1] + fs[i-1];
            }
            let partial = partial as i64;
            let target = target as i64;
            let full = full as i64;
            // do not make target, all new_flowers to enlarge min
            let (mut j, mut min_val) = Solution::get_max_sum(&prefix, &fs, fs.len() as i64 - 1, new_flowers, target);
            let mut max = min_val*partial;  // initial score
            for i in (0..fs.len()).into_iter().rev() {
                // make fs[i:] as target
                if target*((fs.len()-i) as i64) - (prefix[fs.len()]-prefix[i]) > new_flowers {
                    break;  // if can't make fs[i..] as target, break
                }
                // remain flowers to enlarge min
                let remain = new_flowers - (target*((fs.len()-i) as i64) - (prefix[fs.len()]-prefix[i]));
                let x = Solution::get_max_sum(&prefix, &fs, j.min(i as i64-1), remain, target);
                j = x.0;
                min_val = x.1;
                max = max.max(full*((fs.len()-i) as i64) + min_val*partial);
            }
            return ans+max;
        }

        fn get_max_sum(prefix: &Vec<i64>, fs: &Vec<i64>, mut j: i64, mut k: i64, target: i64) -> (i64, i64) {
            if j < 0 { // edge case, make all fs as target, no need enlarge min
                return (j, 0);
            }  // else, we can always find j that fs[j]*(j+1)-sum(fs[0...j]) <= k
            while j >= 0 {  // j always go left
                let jj = j as usize;
                if fs[jj] * (j+1) - (prefix[jj+1]-prefix[0]) <= k {
                    break;
                }
                j -= 1;
            }
            // we can make fs[0...j] = fs[j]
            let jj = j as usize;
            k -= fs[jj] * (j+1) - (prefix[jj+1]-prefix[0]); // remain some k
            return (j, (target-1).min(fs[jj]+k/(j+1)));
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                flowers: vec![1,3,1,1],
                new_flowers: 7,
                target: 6,
                full: 12,
                partial: 1,
                ans: 14,
            },
            Solution {
                flowers: vec![2,4,5,3],
                new_flowers: 10,
                target: 5,
                full: 2,
                partial: 6,
                ans: 30,
            },
            Solution {
                flowers: vec![1,2,3,4,5,6],
                new_flowers: 1000,
                target: 8,
                full: 1,
                partial: 100,
                ans: 705,
            },
            Solution {
                flowers: vec![1,2,3,4,5,6],
                new_flowers: 1000,
                target: 8,
                full: 100,
                partial: 1,
                ans: 600,
            },
        ];
        for i in testcases {
            let ans = Solution::maximum_beauty(i.flowers, i.new_flowers, i.target, i.full, i.partial);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
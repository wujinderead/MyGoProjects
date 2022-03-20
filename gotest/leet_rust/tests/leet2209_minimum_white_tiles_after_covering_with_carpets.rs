// https://leetcode.com/problems/minimum-white-tiles-after-covering-with-carpets/

// You are given a 0-indexed binary string floor, which represents the colors of
// tiles on a floor:
// floor[i] = '0' denotes that the iᵗʰ tile of the floor is colored black. 
// On the other hand, floor[i] = '1' denotes that the iᵗʰ tile of the floor is colored white.
// You are also given numCarpets and carpetLen. You have numCarpets black carpets, each of
// length carpetLen tiles. Cover the tiles with the given carpets such that the number of
// white tiles still visible is minimum. Carpets may overlap one another.
// Return the minimum number of white tiles still visible.
// Example 1:
//   Input: floor = "10110101", numCarpets = 2, carpetLen = 2
//   Output: 2
//   Explanation:
//   The figure above shows one way of covering the tiles with the carpets such that only 2 white tiles are visible.
//   No other way of covering the tiles with the carpets can leave less than 2 white tiles visible.
// Example 2:
//   Input: floor = "11111", numCarpets = 2, carpetLen = 3
//   Output: 0
//   Explanation:
//   The figure above shows one way of covering the tiles with the carpets such that no white tiles are visible.
//   Note that the carpets are able to overlap one another.
// Constraints:
//   1 <= carpetLen <= floor.length <= 1000
//   floor[i] is either '0' or '1'.
//   1 <= numCarpets <= 1000

mod _minimum_white_tiles_after_covering_with_carpets {
    struct Solution(String, i32, i32, i32);

    impl Solution {
        pub fn minimum_white_tiles(floor: String, num_carpets: i32, carpet_len: i32) -> i32 {
            let floor = floor.as_bytes();
            let (num_carpets, carpet_len) = (num_carpets as usize, carpet_len as usize);
            if carpet_len*num_carpets >= floor.len() {
                return 0;
            }
            let mut dp = vec![0; floor.len()];
            let mut new = vec![0; floor.len()];
            // init state:
            // how many "1" left in floor[0..i] with 0 carpet: the number of "1" in floor[0..i]
            dp[0] = (floor[0]-b'0') as i32;
            for i in 1..floor.len() {
                dp[i] = dp[i-1] + (floor[i]-b'0') as i32;
            }
            for i in 1..=num_carpets {  // for number of n carpet
                for j in 0..i*carpet_len {
                    new[j] = 0;
                }
                for j in i*carpet_len..floor.len() {
                    // dp(floor[0...j], i), use i carpets to cover floor[0...j], it's the minimal of:
                    // - dp(floor[0...j-1], i)+floor[j], use i carpets to cover floor[0...j-1],
                    //   and expose floor[j]
                    // - dp(floor[0...j-carLen], i-1), use i-1 carpets to cover floor[0...j-carLen],
                    //   and use i-th carpet to cover floor[j-carLen+1...j]
                    new[j] = (new[j-1] + (floor[j]-b'0') as i32).min(dp[j-carpet_len]);
                }
                std::mem::swap(&mut dp, &mut new);
            }
            return dp[floor.len()-1];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution("10110101".to_string(), 2, 2, 2),
            Solution("11111".to_string(), 2, 3, 0),
            Solution("0111101".to_string(),1, 2, 3),
        ];
        for i in testcases {
            let ans = Solution::minimum_white_tiles(i.0, i.1, i.2);
            println!("{}, {}", ans, i.3);
        }
    } 
}
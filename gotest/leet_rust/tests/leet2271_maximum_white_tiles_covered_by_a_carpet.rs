// https://leetcode.com/problems/maximum-white-tiles-covered-by-a-carpet/

// You are given a 2D integer array tiles where tiles[i] = [li, ri] represents that every tile j in
// the range li <= j <= ri is colored white.
// You are also given an integer carpetLen, the length of a single carpet that can be placed anywhere.
// Return the maximum number of white tiles that can be covered by the carpet.
// Example 1:
//   Input: tiles = [[1,5],[10,11],[12,18],[20,25],[30,32]], carpetLen = 10
//   Output: 9
//   Explanation: Place the carpet starting on tile 10.
//     It covers 9 white tiles, so we return 9.
//     Note that there may be other places where the carpet covers 9 white tiles.
//     It can be shown that the carpet cannot cover more than 9 white tiles.
// Example 2:
//   Input: tiles = [[10,11],[1,1]], carpetLen = 2
//   Output: 2
//   Explanation: Place the carpet starting on tile 10. It covers 2 white tiles, so we return 2.
// Constraints:
//   1 <= tiles.length <= 5 * 10⁴
//   tiles[i].length == 2
//   1 <= li <= ri <= 10⁹
//   1 <= carpetLen <= 10⁹
//   The tiles are non-overlapping.

mod _maximum_white_tiles_covered_by_a_carpet {
    struct Solution{
        tiles: Vec<Vec<i32>>,
        carpet_len: i32,
        ans: i32,
    }

    impl Solution {
        pub fn maximum_white_tiles(mut tiles: Vec<Vec<i32>>, carpet_len: i32) -> i32 {
            let mut max = 0;
            // sort tiles
            tiles.sort_by_key(|t| t[0]);
            let mut ri = 0;
            let mut cur = 0;
            for i in 0..tiles.len() {
                // always put the carpet at the start of a group
                let right = tiles[i][0]+carpet_len-1;  // the right of the carpet
                if i > 0 {
                    cur -= tiles[i-1][1]-tiles[i-1][0]+1;  // remove prev tiles
                }
                while ri < tiles.len() && tiles[ri][1] <= right { // extend right
                    cur += tiles[ri][1]-tiles[ri][0]+1;
                    ri += 1;
                }
                let mut cur_cur = cur;
                if ri < tiles.len() && right >= tiles[ri][0] {  // right maybe in the middle of tiles[ri]
                    cur_cur += right-tiles[ri][0]+1;
                }
                max = max.max(cur_cur);
            }
            return max;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                tiles: vec![[1,5],[10,11],[12,18],[20,25],[30,32]].iter().map(|s| s.to_vec()).collect(),
                carpet_len: 10,
                ans: 9,
            },
            Solution {
                tiles: vec![[10,11],[1,1]].iter().map(|s| s.to_vec()).collect(),
                carpet_len: 2,
                ans: 2,
            },
            Solution {
                tiles: vec![[1,3]].iter().map(|s| s.to_vec()).collect(),
                carpet_len: 2,
                ans: 2,
            },
            Solution {
                tiles: vec![[1,3]].iter().map(|s| s.to_vec()).collect(),
                carpet_len: 4,
                ans: 3,
            },
        ];
        for i in testcases {
            let ans = Solution::maximum_white_tiles(i.tiles, i.carpet_len);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
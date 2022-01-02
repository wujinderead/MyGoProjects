// https://leetcode.com/problems/number-of-laser-beams-in-a-bank/

// Anti-theft security devices are activated inside a bank. You are given a 0-indexed
// binary stringarray bank representing the floor plan of the bank, which isan m x n 2D matrix.
// bank[i] represents the iᵗʰ row, consisting of '0's and '1's. '0' means the cell is empty,
// while'1' means the cell has a security device.
// There is one laser beam between any two security devices if both conditions are met:
//   The two devices are located on two different rows: r1 and r2, where r1 < r2.
//   For each row i where r1 < i < r2, there are no security devices in the iᵗʰ row.
//   Laser beams are independent, i.e., one beam does not interfere nor join with another.
// Return the total number of laser beams in the bank.
// Example 1:
//   Input: bank = ["011001","000000","010100","001000"]
//   Output: 8
//   Explanation: Between each of the following device pairs, there is one beam.
//     In total, there are 8 beams:
//      * bank[0][1] -- bank[2][1]
//      * bank[0][1] -- bank[2][3]
//      * bank[0][2] -- bank[2][1]
//      * bank[0][2] -- bank[2][3]
//      * bank[0][5] -- bank[2][1]
//      * bank[0][5] -- bank[2][3]
//      * bank[2][1] -- bank[3][2]
//      * bank[2][3] -- bank[3][2]
//     Note that there is no beam between any device on the 0ᵗʰ row with any on the 3ʳᵈ row.
//     This is because the 2ⁿᵈ row contains security devices, which breaks the second condition.
// Example 2:
//   Input: bank = ["000","111","000"]
//   Output: 0
//   Explanation: There does not exist two devices located on two different rows.
// Constraints:
//   m == bank.length
//   n == bank[i].length
//   1 <= m, n <= 500
//   bank[i][j] is either '0' or '1'.

mod _number_of_laser_beams_in_a_bank {
    struct Solution{
        bank: Vec<String>,
        ans: i32,
    }

    impl Solution {
        pub fn number_of_beams(bank: Vec<String>) -> i32 {
            let mut sum = 0;
            let mut i: usize = 0;
            let mut prev = 0;
            while i < bank.len() {
                let mut cur = 0;
                for c in bank[i].chars() {
                    if c == '1' {
                        cur += 1;
                    }
                }
                if cur > 0 {  // if current line has 1's, multiple with previous line
                    sum += prev*cur;
                    prev = cur;
                }
                i += 1;
            }
            return sum;
        }
    }
    
    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                bank: vec!["011001".to_string(),"000000".to_string(),"010100".to_string(),"001000".to_string()],
                ans: 8,
            },
            Solution {
                bank: vec!["000".to_string(),"111".to_string(),"000".to_string()],
                ans: 0,
            },
        ];
        for i in testcases {
            let ans = Solution::number_of_beams(i.bank);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
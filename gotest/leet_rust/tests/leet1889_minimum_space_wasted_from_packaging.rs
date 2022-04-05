// https://leetcode.com/problems/minimum-space-wasted-from-packaging/

// You have n packages that you are trying to place in boxes, one package in each box. There are
// m suppliers that each produce boxes of different sizes (with infinite supply). A package can
// be placed in a box if the size of the package is less than or equal to the size of the box.
// The package sizes are given as an integer array packages, where packages[i] is the size of the
// iᵗʰ package. The suppliers are given as a 2D integer array boxes, where boxes[j] is an array of
// box sizes that the jᵗʰ supplier produces.
// You want to choose a single supplier and use boxes from them such that the total wasted space is
// minimized. For each package in a box, we define the space wasted to be size of the box - size of
// the package. The total wasted space is the sum of the space wasted in all the boxes.
// For example, if you have to fit packages with sizes [2,3,5] and the supplier offers boxes of
// sizes [4,8], you can fit the packages of size-2 and size-3 into two boxes of size-4 and the
// package with size-5 into a box of size-8. This would result in a waste of (4-2) + (4-3) + (8-5) = 6.
// Return the minimum total wasted space by choosing the box supplier optimally, or -1 if it is
// impossible to fit all the packages inside boxes. Since the answer may be large, return it modulo 10⁹ + 7.
// Example 1:
//   Input: packages = [2,3,5], boxes = [[4,8],[2,8]]
//   Output: 6
//   Explanation: It is optimal to choose the first supplier, using two size-4 boxes and one size-8 box.
//     The total waste is (4-2) + (4-3) + (8-5) = 6.
// Example 2:
//   Input: packages = [2,3,5], boxes = [[1,4],[2,3],[3,4]]
//   Output: -1
//   Explanation: There is no box that the package of size 5 can fit in.
// Example 3:
//   Input: packages = [3,5,8,10,11,12], boxes = [[12],[11,9],[10,5,14]]
//   Output: 9
//   Explanation: It is optimal to choose the third supplier, using two size-5
//     boxes, two size-10 boxes, and two size-14 boxes.
//     The total waste is (5-3) + (5-5) + (10-8) + (10-10) + (14-11) + (14-12) = 9.
// Constraints:
//   n == packages.length
//   m == boxes.length
//   1 <= n <= 10⁵
//   1 <= m <= 10⁵
//   1 <= packages[i] <= 10⁵
//   1 <= boxes[j].length <= 10⁵
//   1 <= boxes[j][k] <= 10⁵
//   sum(boxes[j].length) <= 10⁵
//   The elements in boxes[j] are distinct.

mod _minimum_space_wasted_from_packaging {
    struct Solution{
        packages: Vec<i32>,
        boxes: Vec<Vec<i32>>,
        ans: i32,
    }

    impl Solution {
        pub fn min_wasted_space(mut packages: Vec<i32>, boxes: Vec<Vec<i32>>) -> i32 {
            // sort packages and all boxes
            packages.sort();
            let mut bs = Vec::new();
            for i in 0..boxes.len() {
                for &j in &boxes[i] {
                    bs.push((j as i64, i));  // box size and index
                }
            }
            bs.sort_by_key(|x| x.0);

            // find all space needed for each box supplier
            let mut sum = vec![0; boxes.len()];
            let mut ind = vec![-1; boxes.len()];
            let mut cur = -1;
            for (size, j) in bs { // a box with `size` size and belongs to supplier j
                while cur+1 < packages.len() as i64 && packages[(cur+1) as usize] as i64 <= size {
                    cur += 1;   // find the package that this box can contain
                }
                sum[j] += (cur-ind[j])*size;
                ind[j] = cur;
            }
            let mut min = 1e12 as i64;
            let mut ans = -1;
            let s = packages.iter().map(|&s| s as i64).collect::<Vec<_>>().iter().sum::<i64>();
            for i in 0..sum.len() {
                // `ind[i] == packages.len()-1` means it can cover all packages
                if ind[i] == (packages.len()-1) as i64 && sum[i] < min {
                    min = sum[i];
                    ans = min-s;
                }
            }
            return (ans%(1e9 as i64+7)) as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                packages: vec![2,3,5],
                boxes: vec![vec![4,8],vec![2,8]],
                ans: 6,
            },
            Solution {
                packages: vec![2,3,5],
                boxes: vec![vec![1,4],vec![2,3],vec![3,4]],
                ans: -1,
            },
            Solution {
                packages: vec![3,5,8,10,11,12],
                boxes: vec![vec![12],vec![11,9],vec![10,5,14]],
                ans: 9,
            },
            Solution {
                packages: vec![3,5,8,10,11,12],
                boxes: vec![vec![12],vec![2,11,9],vec![2,10,5,14,16]],
                ans: 9,
            },
            Solution {
                packages: vec![1; 40000],
                boxes: vec![vec![40000]],
                ans: 599959993,
            },
        ];
        for i in testcases {
            let ans = Solution::min_wasted_space(i.packages, i.boxes);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
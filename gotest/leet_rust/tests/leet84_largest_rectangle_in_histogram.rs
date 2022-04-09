// https://leetcode.com/problems/largest-rectangle-in-histogram/

// Given an array of integers heights representing the histogram's bar height where the width of
// each bar is 1, return the area of the largest rectangle in the histogram.
// Example 1:
//   Input: heights = [2,1,5,6,2,3]
//   Output: 10
//   Explanation: The above is a histogram where width of each bar is 1.
//     The largest rectangle is shown in the red area, which has an area = 10 units.
// Example 2:
//   Input: heights = [2,4]
//   Output: 4
// Constraints:
//   1 <= heights.length <= 10⁵
//   0 <= heights[i] <= 10⁴

mod _largest_rectangle_in_histogram {
    struct Solution{
        heights: Vec<i32>,
        ans: i32,
    }

    impl Solution {
        // for each height[i]=x, find the index of the leftmost index l and rightmost index r,
        // that height[l]<x and height[r]<x, so the max area candidate is x*(r-l-1)
        pub fn largest_rectangle_area(heights: Vec<i32>) -> i32 {
            let mut ans = 0;
            let mut l = vec![0; heights.len()];
            let mut r = vec![0; heights.len()];
            // for l
            let mut stack = vec![-1]; // initial stack [-1]
            for i in 0..heights.len() {  // find leftmost l that height[l]<height[i]
                while stack.len() > 1 && heights[*stack.last().unwrap() as usize] >= heights[i] {
                    stack.pop();
                }
                l[i] = *stack.last().unwrap();
                stack.push(i as i32);
            }

            // for r
            stack.truncate(1);
            stack[0] = heights.len() as i32;  // initial stack [height.len()]
            for i in (0..heights.len()).into_iter().rev() {  // find rightmost r that height[r]<height[i]
                while stack.len() > 1 && heights[*stack.last().unwrap() as usize] >= heights[i] {
                    stack.pop();
                }
                r[i] = *stack.last().unwrap();
                stack.push(i as i32);
                ans = ans.max(heights[i]*(r[i]-l[i]-1));  // update max area
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                heights: vec![2,1,5,6,2,3],
                ans: 10,
            },
            Solution {
                heights: vec![2,4],
                ans: 4,
            },
        ];
        for i in testcases {
            let ans = Solution::largest_rectangle_area(i.heights);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/merge-intervals/

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
// and return an array of the non-overlapping intervals that cover all the intervals in the input.
// Example 1:
//   Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
//   Output: [[1,6],[8,10],[15,18]]
//   Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
// Example 2:
//   Input: intervals = [[1,4],[4,5]]
//   Output: [[1,5]]
//   Explanation: Intervals [1,4] and [4,5] are considered overlapping.
// Constraints:
//   1 <= intervals.length <= 10⁴
//   intervals[i].length == 2
//   0 <= starti <= endi <= 10⁴

mod _merge_intervals {
    struct Solution{
        intervals: Vec<Vec<i32>>,
        ans: Vec<Vec<i32>>,
    }

    impl Solution {
        pub fn merge(mut intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
            intervals.sort_by_key(|s| s[0]);
            let mut ans = Vec::new();
            let mut cur = (intervals[0][0], intervals[0][1]);
            for i in 1..intervals.len() {
                let (l, r) = (intervals[i][0],  intervals[i][1]);
                if l > cur.1 {
                    ans.push(vec![cur.0,  cur.1]);
                    cur = (l, r);
                } else {
                    cur.1 = cur.1.max(r);
                }
            }
            ans.push(vec![cur.0, cur.1]);
            return ans;
        }

        // this is actual merge segment, this do not apply for empty intervals: like [5,5]
        pub fn merge_prefix_sum(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
            let mut arr = vec![0; 10];
            for i in intervals {
                arr[i[0] as usize] += 1;
                arr[i[1] as usize] -= 1;
            }
            let mut sum = 0;
            let mut s = -1;
            let mut ans = Vec::new();
            for i in 0..10 {
                sum += arr[i];
                if sum>0 && s==-1 {
                    s = i as i32;
                } else if sum==0 && s>=0 {
                    ans.push(vec![s, i as i32]);
                    s = -1;
                }
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                intervals: vec![[1,3],[2,6],[8,10],[15,18]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![[1,6],[8,10],[15,18]].iter().map(|s| s.to_vec()).collect(),
            },
            Solution {
                intervals: vec![[1,4],[4,5]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![[1,5]].iter().map(|s| s.to_vec()).collect(),
            },
            Solution {
                intervals: vec![[1,4],[4,5],[9995,9998],[9997,10000]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![[1,5],[9995,10000]].iter().map(|s| s.to_vec()).collect(),
            },
            Solution {
                intervals: vec![[1,4],[5,6]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![[1,4],[5,6]].iter().map(|s| s.to_vec()).collect(),
            },
            Solution {
                intervals: vec![[1,4],[3,6]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![[1,6]].iter().map(|s| s.to_vec()).collect(),
            },
            Solution {
                intervals: vec![[1,4],[4,6]].iter().map(|s| s.to_vec()).collect(),
                ans: vec![[1,6]].iter().map(|s| s.to_vec()).collect(),
            },
        ];
        for i in &testcases {
            let ans = Solution::merge(i.intervals.clone());
            println!("{:?}, {:?}", ans, i.ans);
        }
        for i in &testcases {
            let ans = Solution::merge_prefix_sum(i.intervals.clone());
            println!("{:?} {:?}", ans, i.ans);
        }
    } 
}
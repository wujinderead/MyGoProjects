// https://leetcode.com/problems/number-of-flowers-in-full-bloom/

// You are given a 0-indexed 2D integer array flowers, where flowers[i] = [starti, endi] means the
// iᵗʰ flower will be in full bloom from starti to endi (inclusive). You are also given a 0-indexed
// integer array persons of size n, where persons[i] is the time that the iᵗʰ person will arrive to
// see the flowers.
// Return an integer array answer of size n, where answer[i] is the number of flowers that are in
// full bloom when the iᵗʰ person arrives.
// Example 1:
//   Input: flowers = [[1,6],[3,7],[9,12],[4,13]], persons = [2,3,7,11]
//   Output: [1,2,2,2]
//   Explanation: The figure above shows the times when the flowers are in full bloom and when the
//     people arrive. For each person, we return the number of flowers in full bloom during their arrival.
// Example 2:
//   Input: flowers = [[1,10],[3,3]], persons = [3,3,2]
//   Output: [2,2,1]
//   Explanation: The figure above shows the times when the flowers are in full bloom and when the
//     people arrive. For each person, we return the number of flowers in full bloom during their arrival.
// Constraints:
//   1 <= flowers.length <= 5 * 10⁴
//   flowers[i].length == 2
//   1 <= starti <= endi <= 10⁹
//   1 <= persons.length <= 5 * 10⁴
//   1 <= persons[i] <= 10⁹

mod _number_of_flowers_in_full_bloom {
    struct Solution{
        flowers: Vec<Vec<i32>>,
        persons: Vec<i32>,
        ans: Vec<i32>,
    }

    impl Solution {
        pub fn full_bloom_flowers(flowers: Vec<Vec<i32>>, persons: Vec<i32>) -> Vec<i32> {
            let mut events = Vec::with_capacity(2*flowers.len()+persons.len());
            for f in flowers {
                events.push([f[0], -1]);    // event type -1: start bloom
                events.push([f[1]+1, -2]);  // event type -2: end bloom
            }
            for (i, &p) in persons.iter().enumerate() {
                events.push([p, i as i32]);  // event type >=0, watch
            }
            events.sort_by(|a, b| if a[0] != b[0] {
                a[0].cmp(&b[0])
            } else {
                a[1].cmp(&b[1])
            });
            let mut prefix = 0;
            let mut ans = vec![0; persons.len()];
            for e in events {
                if e[1] >=0 {
                    ans[e[1] as usize] = prefix;
                }
                if e[1] == -1 {
                    prefix += 1;
                }
                if e[1] == -2 {
                    prefix -= 1;
                }
            }
            return ans;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                flowers: vec![[1,6],[3,7],[9,12],[4,13]].iter().map(|x| x.to_vec()).collect(),
                persons: vec![2,3,7,11],
                ans: vec![1,2,2,2],
            },
            Solution {
                flowers: vec![[1,10],[3,3]].iter().map(|x| x.to_vec()).collect(),
                persons: vec![3,3,2],
                ans: vec![2,2,1],
            },
        ];
        for i in testcases {
            let ans = Solution::full_bloom_flowers(i.flowers, i.persons);
            println!("{:?}, {:?}", ans, i.ans);
        }
    } 
}
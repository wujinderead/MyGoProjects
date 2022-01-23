// https://leetcode.com/problems/solving-questions-with-brainpower/

// You are given a 0-indexed 2D integer array questions where questions[i] = [pointsi, brainpoweri].
// The array describes the questions of an exam, where you have to process the  questions in order
// (i.e., starting from question 0) and make a decision whether  to solve or skip each question.
// Solving question i will earn you pointsi points but you will be unable to solve each of the next
// brainpoweri questions. If you skip question i, you get to make the decision on the next question.
// For example, given questions = [[3, 2], [4, 3], [4, 4], [2, 5]]:
// If question 0 is solved, you will earn 3 points but you will be unable to 
//   solve questions 1 and 2.
// If instead, question 0 is skipped and question 1 is solved, you will earn 4 
//   points but you will be unable to solve questions 2 and 3.
// Return the maximum points you can earn for the exam.
// Example 1:
//   Input: questions = [[3,2],[4,3],[4,4],[2,5]]
//   Output: 5
//   Explanation: The maximum points can be earned by solving questions 0 and 3.
//     - Solve question 0: Earn 3 points, will be unable to solve the next 2 questions
//     - Unable to solve questions 1 and 2
//     - Solve question 3: Earn 2 points
//     Total points earned: 3 + 2 = 5. There is no other way to earn 5 or more points.
// Example 2:
//   Input: questions = [[1,1],[2,2],[3,3],[4,4],[5,5]]
//   Output: 7
//   Explanation: The maximum points can be earned by solving questions 1 and 4.
//     - Skip question 0
//     - Solve question 1: Earn 2 points, will be unable to solve the next 2 questions
//     - Unable to solve questions 2 and 3
//     - Solve question 4: Earn 5 points
//     Total points earned: 2 + 5 = 7. There is no other way to earn 7 or more points.
// Constraints:
//   1 <= questions.length <= 10⁵
//   questions[i].length == 2
//   1 <= pointsi, brainpoweri <= 10⁵

mod _solving_questions_with_brainpower {
    struct Solution{
        questions: Vec<Vec<i32>>,
        ans: i64,
    }

    // easy dp:
    // let dp[i] be the max point we can get from questions[i:], then for question i,
    // we have two choices:
    // - take question[i], can get point[i]+dp[i+brainpower[i]+1];
    // - skip question[i], can get dp[i+1];
    // choose the larger value as dp[i].
    impl Solution {
        pub fn most_points(questions: Vec<Vec<i32>>) -> i64 {
            let mut dp: Vec<i64> = vec![0; questions.len()];
            dp[questions.len()-1] = questions[questions.len()-1][0] as i64;
            let mut i = (questions.len()-2) as i32;
            while i>=0 {
                let ii = i as usize;
                let mut cur = questions[ii][0] as i64;
                if ii + (questions[ii][1] as usize) + 1 < questions.len() {
                    cur += dp[ii + (questions[ii][1] as usize) + 1];
                }
                dp[ii] = dp[ii+1].max(cur);
                i -= 1;
            }
            return dp[0];
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                questions: vec![vec![3,2],vec![4,3],vec![4,4],vec![2,5]],
                ans: 5,
            },
            Solution {
                questions: vec![vec![1,1],vec![2,2],vec![3,3],vec![4,4],vec![5,5]],
                ans: 7,
            },
        ];
        for i in testcases {
            let ans = Solution::most_points(i.questions);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
// https://leetcode.com/problems/minimum-consecutive-cards-to-pick-up/

// You are given an integer array cards where cards[i] represents the value of the iᵗʰ card.
// A pair of cards are matching if the cards have the same value.
// Return the minimum number of consecutive cards you have to pick up to have a pair of matching
// cards among the picked cards. If it is impossible to have matching cards, return -1.
// Example 1:
//   Input: cards = [3,4,2,3,4,7]
//   Output: 4
//   Explanation: We can pick up the cards [3,4,2,3] which contain a matching pair
//     of cards with value 3. Note that picking up the cards [4,2,3,4] is also optimal.
// Example 2:
//   Input: cards = [1,0,5,3]
//   Output: -1
//   Explanation: There is no way to pick up a set of consecutive cards that
//     contain a pair of matching cards.
// Constraints:
//   1 <= cards.length <= 10⁵
//   0 <= cards[i] <= 10⁶

mod _minimum_consecutive_cards_to_pick_up {
    struct Solution{
        cards: Vec<i32>,
        ans: i32,
    }

    use std::collections::HashMap;
    impl Solution {
        pub fn minimum_card_pickup(cards: Vec<i32>) -> i32 {
            let mut map = HashMap::new();
            let mut min = cards.len()+1;
            for i in 0..cards.len() {
                if let Some(&prev) = map.get(&cards[i]) {
                    min = min.min(i-prev+1);
                }
                map.insert(cards[i], i);
            }
            if min == cards.len()+1 {
                return -1;
            }
            return min as i32;
        }
    }

    #[test]
    fn test() {
        let testcases = vec![
            Solution {
                cards: vec![3,4,2,3,4,7],
                ans: 4,
            },
            Solution {
                cards: vec![1,0,5,3],
                ans: -1,
            },
        ];
        for i in testcases {
            let ans = Solution::minimum_card_pickup(i.cards);
            println!("{}, {}", ans, i.ans);
        }
    } 
}
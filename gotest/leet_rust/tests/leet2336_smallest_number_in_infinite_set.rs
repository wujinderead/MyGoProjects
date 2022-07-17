// https://leetcode.com/problems/smallest-number-in-infinite-set/

// You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].
// Implement the SmallestInfiniteSet class:
//   SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
//   int popSmallest() Removes and returns the smallest integer contained in the infinite set.
//   void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.
// Example 1:
//   Input
//     ["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest",
//       "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
//      [[], [2], [], [], [], [1], [], [], []]
//   Output
//     [null, null, 1, 2, 3, null, 1, 4, 5]
//   Explanation
//     SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
//     smallestInfiniteSet.addBack(2);    // 2 is already in the set, so no change is made.
//     smallestInfiniteSet.popSmallest(); // return 1, since 1 is the smallest number, and remove it from the set.
//     smallestInfiniteSet.popSmallest(); // return 2, and remove it from the set.
//     smallestInfiniteSet.popSmallest(); // return 3, and remove it from the set.
//     smallestInfiniteSet.addBack(1);    // 1 is added back to the set.
//     smallestInfiniteSet.popSmallest(); // return 1, since 1 was added back to the set and
//                                        // is the smallest number, and remove it from the set.
//     smallestInfiniteSet.popSmallest(); // return 4, and remove it from the set.
//     smallestInfiniteSet.popSmallest(); // return 5, and remove it from the set.
// Constraints:
//   1 <= num <= 1000
//   At most 1000 calls will be made in total to popSmallest and addBack.

mod _smallest_number_in_infinite_set {

    use std::collections::{BinaryHeap, HashSet};
    struct SmallestInfiniteSet {
        min: i32,
        back_heap: BinaryHeap<i32>,
        back_set: HashSet<i32>,
    }

    impl SmallestInfiniteSet {
        fn new() -> Self {
            Self {
                min: 1,
                back_heap: BinaryHeap::new(),
                back_set: HashSet::new(),
            }
        }

        fn pop_smallest(&mut self) -> i32 {
            if self.back_heap.len() > 0 {
                let pop = self.back_heap.pop().unwrap();
                self.back_set.remove(&-pop);
                return -pop;
            }
            self.min += 1;
            return self.min-1;
        }

        fn add_back(&mut self, num: i32) {
            if num >= self.min {
                return;
            }
            if self.back_set.contains(&num) {
                return;
            }
            self.back_set.insert(num);
            self.back_heap.push(-num);
        }
    }

    #[test]
    fn test() {
        let mut obj = SmallestInfiniteSet::new();
        let obj = &mut obj;
        println!("{}", obj.pop_smallest());
        println!("{}", obj.pop_smallest());
        obj.add_back(1);
        println!("{}", obj.pop_smallest());
        println!("{}", obj.pop_smallest());
        obj.add_back(3);
        println!("{}", obj.pop_smallest());
        println!("{}", obj.pop_smallest());
    }
}
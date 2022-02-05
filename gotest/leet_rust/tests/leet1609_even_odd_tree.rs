// https://leetcode.com/problems/even-odd-tree/

// A binary tree is named Even-Odd if it meets the following conditions:
//   The root of the binary tree is at level index 0, its children are at level
//     index 1, their children are at level index 2, etc.
//   For every even-indexed level, all nodes at the level have odd integer values
//     in strictly increasing order (from left to right).
//   For every odd-indexed level, all nodes at the level have even integer values
//     in strictly decreasing order (from left to right).
// Given the root of a binary tree, return true if the binary tree is Even-Odd, 
// otherwise return false.
// Example 1:
//   Input: root = [1,10,4,3,null,7,9,12,8,6,null,null,2]
//   Output: true
//   Explanation: The node values on each level are:
//     Level 0: [1]
//     Level 1: [10,4]
//     Level 2: [3,7,9]
//     Level 3: [12,8,6,2]
//     Since levels 0 and 2 are all odd and increasing and levels 1 and 3 are all even
//     and decreasing, the tree is Even-Odd.
//               1
//          /        \
//        10         4
//        /         / \
//       3        7   9
//     /  \      /     \
//   12   8     6       2
// Example 2:
//   Input: root = [5,4,2,3,3,7]
//   Output: false
//   Explanation: The node values on each level are:
//     Level 0: [5]
//     Level 1: [4,2]
//     Level 2: [3,3,7]
//     Node values in level 2 must be in strictly increasing order, so the tree is not Even-Odd.
// Example 3:
//   Input: root = [5,9,1,3,5,7]
//   Output: false
//   Explanation: Node values in the level 1 should be even integers.
// Constraints:
//   The number of nodes in the tree is in the range [1, 10⁵].
//   1 <= Node.val <= 10⁶

mod _even_odd_tree {
    struct Solution;

    #[derive(Debug, PartialEq, Eq)]
    pub struct TreeNode {
      pub val: i32,
      pub left: Option<Rc<RefCell<TreeNode>>>,
      pub right: Option<Rc<RefCell<TreeNode>>>,
    }

    use std::rc::Rc;
    use std::cell::RefCell;
    use std::collections::VecDeque;

    // bfs to get each level
    impl Solution {
        pub fn is_even_odd_tree(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
            let mut queue = VecDeque::new();
            // as_ref() makes Option<T> become Option<&T> so we can get
            // T's reference, other than take the ownership of T.
            // here clone the reference of Rc<RefCell<TreeNode>>, and push to queue.
            queue.push_back(Rc::clone(root.as_ref().unwrap()));
            let mut check_odd = true;
            while queue.len() > 0 {
                let l = queue.len();
                let mut i = 0;
                let mut prev = match check_odd {
                    true => -1,
                    false => 1000001,
                };
                while i < l {
                    // pop the queue, got Rc<RefCell<TreeNode>>, and its ownership
                    let op = queue.pop_front().unwrap();
                    // borrow to get TreeNode's reference in RefCell, we can visit reference
                    let cur = op.borrow();
                    if (check_odd && (cur.val%2 == 0 || prev >= cur.val)) ||
                        (!check_odd && (cur.val%2 == 1 || prev <= cur.val)) {
                        return false;
                    }
                    if cur.left.is_some() {
                        // as_ref().unwrap() to get &Rc<RefCell<TreeNode>> then clone it
                        queue.push_back(Rc::clone(&cur.left.as_ref().unwrap()));
                    }
                    if cur.right.is_some() {
                        queue.push_back(Rc::clone(&cur.right.as_ref().unwrap()));
                    }
                    prev = cur.val;
                    i += 1;
                }
                check_odd = !check_odd;
            }
            return true;
        }
    }

    #[test]
    fn test() {
//               1
//          /        \
//        10         4
//        /         / \
//       3        7   9
//     /  \      /     \
//   12   8     6       2
        {
            let a12 = Rc::new(RefCell::new(TreeNode{val: 12, left: None, right: None }));
            let a8 = Rc::new(RefCell::new(TreeNode{val: 8, left: None, right: None }));
            let a3 = Rc::new(RefCell::new(TreeNode{val: 3, left: Some(Rc::clone(&a12)), right: Some(Rc::clone(&a8)) }));
            let a6 = Rc::new(RefCell::new(TreeNode{val: 6, left: None, right: None }));
            let a2 = Rc::new(RefCell::new(TreeNode{val: 2, left: None, right: None }));
            let a7 = Rc::new(RefCell::new(TreeNode{val: 7, left: Some(Rc::clone(&a6)), right: None }));
            let a9 = Rc::new(RefCell::new(TreeNode{val: 9, left: None, right: Some(Rc::clone(&a2)) }));
            let a4 = Rc::new(RefCell::new(TreeNode{val: 4, left: Some(Rc::clone(&a7)), right: Some(Rc::clone(&a9)) }));
            let a10 = Rc::new(RefCell::new(TreeNode{val: 10, left: Some(Rc::clone(&a3)), right: None }));
            let a1 = Rc::new(RefCell::new(TreeNode{val: 1, left: Some(Rc::clone(&a10)), right: Some(Rc::clone(&a4)) }));
            let ans = Solution::is_even_odd_tree(Some(a1));
            println!("ans: {}, expected: {}", ans, true);
        }
//        5
//       / \
//     4   2
//    / \   \
//   3  3   7
        {
            let a5 = Some(Rc::new(RefCell::new(TreeNode{
                val: 5,
                left: Some(Rc::new(RefCell::new(TreeNode{
                    val: 4,
                    left: Some(Rc::new(RefCell::new(TreeNode{val: 3, left: None, right: None }))),
                    right: Some(Rc::new(RefCell::new(TreeNode{val: 3, left: None, right: None }))),
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode{
                    val: 2,
                    left: None,
                    right: Some(Rc::new(RefCell::new(TreeNode{val: 7, left: None, right: None }))),
                }))),
            })));
            let ans = Solution::is_even_odd_tree(a5);
            println!("ans: {}, expected: {}", ans, false);
        }
//        5
//       / \
//     9   1
//    / \   \
//   3  5   7
        {
            let a5 = Some(Rc::new(RefCell::new(TreeNode{
                val: 5,
                left: Some(Rc::new(RefCell::new(TreeNode{
                    val: 9,
                    left: Some(Rc::new(RefCell::new(TreeNode{val: 3, left: None, right: None }))),
                    right: Some(Rc::new(RefCell::new(TreeNode{val: 5, left: None, right: None }))),
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode{
                    val: 1,
                    left: Some(Rc::new(RefCell::new(TreeNode{val: 7, left: None, right: None }))),
                    right: None,
                }))),
            })));
            let ans = Solution::is_even_odd_tree(a5);
            println!("ans: {}, expected: {}", ans, false);
        }
    } 
}
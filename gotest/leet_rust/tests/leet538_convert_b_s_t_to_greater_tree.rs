// https://leetcode.com/problems/convert-bst-to-greater-tree/

// Given the root of a Binary Search Tree (BST), convert it to a Greater Tree such that every key
// of the original BST is changed to the original key plus the sum of all keys greater than the
// original key in BST.
// As a reminder, a binary search tree is a tree that satisfies these constraints:
//   The left subtree of a node contains only nodes with keys less than the node's key.
//   The right subtree of a node contains only nodes with keys greater than the node's key.
//   Both the left and right subtrees must also be binary search trees.
// Example 1:
//   Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
//   Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
// Example 2:
//   Input: root = [0,null,1]
//   Output: [1,null,1]
// Constraints:
//   The number of nodes in the tree is in the range [0, 10⁴].
//   -10⁴ <= Node.val <= 10⁴
//   All the values in the tree are unique.
//   root is guaranteed to be a valid binary search tree.
// Note: This question is the same as 1038: https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/

mod _convert_b_s_t_to_greater_tree {
    struct Solution;

    //Definition for a binary tree node.
    #[derive(Debug, PartialEq, Eq)]
    pub struct TreeNode {
      pub val: i32,
      pub left: Option<Rc<RefCell<TreeNode>>>,
      pub right: Option<Rc<RefCell<TreeNode>>>,
    }

    impl TreeNode {
      #[inline]
      pub fn new(val: i32) -> Self {
        TreeNode {
          val,
          left: None,
          right: None
        }
      }
    }

    use std::rc::Rc;
    use std::cell::RefCell;
    use std::collections::VecDeque;

    // just traverse the tree in mid-order
    impl Solution {
        pub fn convert_bst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
            let mut queue = VecDeque::new();
            let mut prev = 0;
            let mut cur = root.clone();
            while cur.is_some() || queue.len() > 0 {
                if cur.is_some() {
                    let x = cur.unwrap();
                    queue.push_back(Rc::clone(&x));
                    cur = x.borrow().right.clone();
                } else {
                    let x = queue.pop_back().unwrap();
                    let mut xx = x.borrow_mut();
                    xx.val += prev;
                    prev = xx.val;
                    cur = xx.left.clone();
                }
            }
            return root;
        }

        fn pre_order(node: Option<Rc<RefCell<TreeNode>>>) {
            let mut queue = VecDeque::new();
            let mut cur = node;
            while cur.is_some() || queue.len() > 0 {
                if cur.is_some() {
                    let x = cur.unwrap();
                    println!("{}", x.borrow().val);
                    queue.push_back(Rc::clone(&x));
                    cur = x.borrow().left.clone();
                } else {
                    let x = queue.pop_back().unwrap();
                    cur = x.borrow().right.clone();
                }
            }
        }

        fn mid_order(node: Option<Rc<RefCell<TreeNode>>>) {
            let mut queue = VecDeque::new();
            let mut cur = node;
            while cur.is_some() || queue.len() > 0 {
                if cur.is_some() {
                    let x = cur.unwrap();
                    queue.push_back(Rc::clone(&x));
                    cur = x.borrow().left.clone();
                } else {
                    let x = queue.pop_back().unwrap();
                    println!("{}", x.borrow().val);
                    cur = x.borrow().right.clone();
                }
            }
        }
    }

    #[test]
    fn test() {
//      4                       30
//   1    6                 36     21
//  0 2  5 7              36  35 26  15
//    3     8                  33     8
        {
            let r = Some(Rc::new(RefCell::new(TreeNode::new(4))));
            r.as_ref().unwrap().borrow_mut().left = Some(Rc::new(RefCell::new(TreeNode::new(1))));
            r.as_ref().unwrap().borrow_mut().right = Some(Rc::new(RefCell::new(TreeNode::new(6))));
            r.as_ref().unwrap().borrow().left.as_ref().unwrap().borrow_mut().left = Some(Rc::new(RefCell::new(TreeNode::new(0))));
            r.as_ref().unwrap().borrow().left.as_ref().unwrap().borrow_mut().right = Some(Rc::new(RefCell::new(TreeNode::new(2))));
            r.as_ref().unwrap().borrow().left.as_ref().unwrap().borrow().right.as_ref().unwrap().borrow_mut().right = Some(Rc::new(RefCell::new(TreeNode::new(3))));
            r.as_ref().unwrap().borrow().right.as_ref().unwrap().borrow_mut().left = Some(Rc::new(RefCell::new(TreeNode::new(5))));
            r.as_ref().unwrap().borrow().right.as_ref().unwrap().borrow_mut().right = Some(Rc::new(RefCell::new(TreeNode::new(7))));
            r.as_ref().unwrap().borrow().right.as_ref().unwrap().borrow().right.as_ref().unwrap().borrow_mut().right = Some(Rc::new(RefCell::new(TreeNode::new(8))));
            Solution::pre_order(r.clone());
            Solution::mid_order(r.clone());
            let ans = Solution::convert_bst(r);
            Solution::mid_order(ans);
        }
    } 
}
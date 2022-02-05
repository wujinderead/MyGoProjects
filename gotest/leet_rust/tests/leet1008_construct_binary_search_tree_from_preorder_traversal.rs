// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/

// Given an array of integers preorder, which represents the preorder traversal of a BST
// (i.e., binary search tree), construct the tree and return its root.
// It is guaranteed that there is always possible to find a binary search tree with the
// given requirements for the given test cases.
// A binary search tree is a binary tree where for every node, any descendant of Node.left
// has a value strictly less than Node.val, and any descendant of Node.right has a value
// strictly greater than Node.val.
// A preorder traversal of a binary tree displays the value of the node first, then traverses
// Node.left, then traverses Node.right.
// Example 1:
//   Input: preorder = [8,5,1,7,10,12]
//   Output: [8,5,10,1,7,null,12]
//        8
//       / \
//      5  10
//     / \  \
//    1  7  12
// Example 2:
//   Input: preorder = [1,3]
//   Output: [1,null,3]
// Constraints:
//   1 <= preorder.length <= 100
//   1 <= preorder[i] <= 1000
//   All the values of preorder are unique.

mod _construct_binary_search_tree_from_preorder_traversal {
    struct Solution;

    // Definition for a binary tree node.
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
    impl Solution {
        pub fn bst_from_preorder(preorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
            return Self::recur(&preorder[..]);
        }

        fn recur(preorder: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
            if preorder.is_empty() {
                return None;
            }
            let mut node = TreeNode::new(preorder[0]);
            let mut i = 0;
            while i+1 < preorder.len() && preorder[i+1] < preorder[0] {
                i += 1;
            }
            if i > 0 {
                let l = Self::recur(&preorder[1..i+1]);
                node.left = Some(Rc::clone(l.as_ref().unwrap()));
            }
            if i+1 < preorder.len() {
                let r = Self::recur(&preorder[i+1..]);
                node.right = Some(Rc::clone(r.as_ref().unwrap()));
            }
            println!("val{} left{} right{}",
                     node.val,
                     RefCell::borrow(node.left.as_ref().unwrap_or(&Rc::new(RefCell::new(TreeNode::new(-1))))).val,
                     RefCell::borrow(node.right.as_ref().unwrap_or(&Rc::new(RefCell::new(TreeNode::new(-1))))).val,
            );
            return Some(Rc::new(RefCell::new(node)));
        }

        fn pre_order(node: &TreeNode) {
            println!("{}", node.val);
            if let Some(x) = node.left.as_ref() {
                Self::pre_order(&RefCell::borrow(x));
            }
            if let Some(x) = node.right.as_ref() {
                Self::pre_order(&RefCell::borrow(x));
            }
        }

        fn pre_order_1(node: Option<&Rc<RefCell<TreeNode>>>) {
            if let Some(&x) = node.as_ref() {
                let xx = RefCell::borrow(x);
                println!("{}", xx.val);
                if xx.left.is_some() {
                    Self::pre_order_1(xx.left.as_ref());
                }
                if xx.right.is_some() {
                    Self::pre_order_1(xx.right.as_ref());
                }
            }
        }

        fn pre_order_2(node: &Option<Rc<RefCell<TreeNode>>>) {
            if let Some(x) = node {
                let xx = RefCell::borrow(x);
                println!("{}", xx.val);
                if xx.left.is_some() {
                    Self::pre_order_2(&xx.left);
                }
                if xx.right.is_some() {
                    Self::pre_order_2(&xx.right);
                }
            }
        }
    }

    #[test]
    fn test() {
        let ans = Solution::bst_from_preorder(vec![8,5,1,7,10,12]);
        Solution::pre_order_1(ans.as_ref());
        Solution::pre_order_2(&ans);
        Solution::pre_order(&RefCell::borrow(ans.as_ref().unwrap()));
    }
}
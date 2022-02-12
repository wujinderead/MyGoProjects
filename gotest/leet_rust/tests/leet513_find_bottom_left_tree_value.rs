// https://leetcode.com/problems/find-bottom-left-tree-value/

// Given the root of a binary tree, return the leftmost value in the last row of the tree.
// Example 1:
//   Input: root = [2,1,3]
//   Output: 1
// Example 2:
//   Input: root = [1,2,3,4,null,5,6,null,null,7]
//   Output: 7
// Constraints:
//   The number of nodes in the tree is in the range [1, 10⁴].
//   -2³¹ <= Node.val <= 2³¹ - 1

mod _find_bottom_left_tree_value {
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
        pub fn find_bottom_left_value(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
            let mut most = (0, 0);
            Solution::dfs(&root, 1, &mut most);
            return most.1;
        }

        fn dfs(cur: &Option<Rc<RefCell<TreeNode>>>, level: i32, most: &mut (i32, i32)) {
            let c = cur.as_ref().unwrap().borrow();
            if level > most.0 {  // if current node is visited the first time in a level
                most.0 = level;
                most.1 = c.val;
            }
            if c.left.is_some() {
                Solution::dfs(&c.left, level+1, most);
            }
            if c.right.is_some() {
                Solution::dfs(&c.right, level+1, most);
            }
        }
    }

    #[test]
    fn test() {
        //      2
        //     / \
        //    1  3
        {
            let root = Some(Rc::new(RefCell::new(TreeNode{
                val: 2,
                left: Some(Rc::new(RefCell::new(TreeNode::new(1)))),
                right: Some(Rc::new(RefCell::new(TreeNode::new(3)))),
            })));
            let ans = Solution::find_bottom_left_value(root);
            println!("{}", ans);
        }
        //       1
        //      / \
        //     2   3
        //    /    /\
        //   4    5  6
        //       /
        //      7
        {
            let root = Some(Rc::new(RefCell::new(TreeNode{
                val: 1,
                left: Some(Rc::new(RefCell::new(TreeNode {
                    val: 2,
                    left: Some(Rc::new(RefCell::new(TreeNode::new(4)))),
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode {
                    val: 3,
                    left: Some(Rc::new(RefCell::new(TreeNode {
                        val: 5,
                        left: Some(Rc::new(RefCell::new(TreeNode::new(7)))),
                        right: None,
                    }))),
                    right: Some(Rc::new(RefCell::new(TreeNode::new(6)))),
                }))),
            })));
            let ans = Solution::find_bottom_left_value(root);
            println!("{}", ans);
        }
    } 
}
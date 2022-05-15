// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/

// Given the root of a binary tree, return the number of nodes where the value of the node is
// equal to the average of the values in its subtree.
// Note:
// The average of n elements is the sum of the n elements divided by n and rounded down to the
// nearest integer.
// A subtree of root is a tree consisting of root and all of its descendants.
// Example 1:
//   Input: root = [4,8,5,0,1,null,6]
//   Output: 5
//   Explanation:
//     For the node with value 4: The average of its subtree is (4 + 8 + 5 + 0 + 1 +6) / 6 = 24 / 6 = 4.
//     For the node with value 5: The average of its subtree is (5 + 6) / 2 = 11 / 2 = 5.
//     For the node with value 0: The average of its subtree is 0 / 1 = 0.
//     For the node with value 1: The average of its subtree is 1 / 1 = 1.
//     For the node with value 6: The average of its subtree is 6 / 1 = 6.
// Example 2:
//   Input: root = [1]
//   Output: 1
//   Explanation: For the node with value 1: The average of its subtree is 1 / 1 = 1.
// Constraints:
//   The number of nodes in the tree is in the range [1, 1000].
//   0 <= Node.val <= 1000

mod _count_nodes_equal_to_average_of_subtree {
    struct Solution;
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
        pub fn average_of_subtree(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
            let mut ans = 0;
            Solution::helper(root.as_ref(), &mut ans);
            return ans;
        }

        pub fn helper(root: Option<&Rc<RefCell<TreeNode>>>, count: &mut i32) -> (i32, i32) {
            if root.is_none() {
                return (0, 0);
            }
            let cur = root.as_ref().unwrap().borrow();
            let (l_sum, l_count) = Solution::helper(cur.left.as_ref(), count);
            let (r_sum, r_count) = Solution::helper(cur.right.as_ref(), count);
            if (l_sum+r_sum+cur.val)/(l_count+r_count+1) == cur.val {
                *count += 1;
            }
            return (l_sum+r_sum+cur.val, l_count+r_count+1)
        }
    }

    #[test]
    fn test() {
        let root = Some(Rc::new(RefCell::new(TreeNode{
            left: Some(Rc::new(RefCell::new(TreeNode{
                left: Some(Rc::new(RefCell::new(TreeNode::new(0)))),
                right: Some(Rc::new(RefCell::new(TreeNode::new(1)))),
                val: 8,
            }))),
            right: Some(Rc::new(RefCell::new(TreeNode{
                left: Some(Rc::new(RefCell::new(TreeNode::new(6)))),
                right: None,
                val: 5,
            }))),
            val: 4,
        })));
        println!("{}", Solution::average_of_subtree(root));
    } 
}
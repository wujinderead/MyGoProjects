// https://leetcode.com/problems/most-frequent-subtree-sum/

// Given the root of a binary tree, return the most frequent subtree sum. If there is a tie,
// return all the values with the highest frequency in any order.
// The subtree sum of a node is defined as the sum of all the node values formed by the
// subtree rooted at that node (including the node itself).
// Example 1:
//   Input: root = [5,2,-3]
//   Output: [2,-3,4]
// Example 2:
//   Input: root = [5,2,-5]
//   Output: [2]
// Constraints:
//   The number of nodes in the tree is in the range [1, 10⁴].
//   -10⁵ <= Node.val <= 10⁵

mod _most_frequent_subtree_sum {
    struct Solution;

    //Definition for a binary tree node.
    #[derive(Debug, PartialEq, Eq)]
    pub struct TreeNode {
      pub val: i32,
      pub left: Option<Rc<RefCell<TreeNode>>>,
      pub right: Option<Rc<RefCell<TreeNode>>>,
    }

    use std::rc::Rc;
    use std::cell::RefCell;
    use std::collections::hash_map::HashMap;
    impl Solution {
        pub fn find_frequent_tree_sum(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
            let mut map = HashMap::new();
            let _ = Self::recur(root.as_ref().unwrap(), &mut map);
            let mut max = 0;
            let mut ans = Vec::new();
            for (&_k, &v) in map.iter() {
                max = max.max(v);
            }
            for (&k, &v) in map.iter() {
                if v == max {
                    ans.push(k);
                }
            }
            return ans;
        }

        fn recur(root: &Rc<RefCell<TreeNode>>, map: &mut HashMap<i32, i32>) -> i32 {
            let node = root.borrow();
            let mut sum = node.val;
            if node.left.is_some() {
                sum += Self::recur(node.left.as_ref().unwrap(), map);
            }
            if node.right.is_some() {
                sum += Self::recur(node.right.as_ref().unwrap(), map);
            }
            *map.entry(sum).or_insert(0) += 1;
            return sum;
        }
    }

    #[test]
    fn test() {
        {
            let root = Some(Rc::new(RefCell::new(TreeNode{
                val: 5,
                left: Some(Rc::new(RefCell::new(TreeNode{
                    val: 2,
                    left: None,
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode{
                    val: -3,
                    left: None,
                    right: None,
                }))),
            })));
            let ans = Solution::find_frequent_tree_sum(root);
            println!("{:?}", ans);
        }
        {
            let root = Some(Rc::new(RefCell::new(TreeNode{
                val: 5,
                left: Some(Rc::new(RefCell::new(TreeNode{
                    val: 2,
                    left: None,
                    right: None,
                }))),
                right: Some(Rc::new(RefCell::new(TreeNode{
                    val: -5,
                    left: None,
                    right: None,
                }))),
            })));
            let ans = Solution::find_frequent_tree_sum(root);
            println!("{:?}", ans);
        }
    } 
}
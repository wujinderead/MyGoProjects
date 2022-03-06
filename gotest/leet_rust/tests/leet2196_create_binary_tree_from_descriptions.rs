// https://leetcode.com/problems/create-binary-tree-from-descriptions/

// You are given a 2D integer array descriptions where descriptions[i] = [parenti, childi, isLefti]
// indicates that parenti is the parent of childi in a binary tree of unique values. Furthermore,
//   If isLefti == 1, then childi is the left child of parenti.
//   If isLefti == 0, then childi is the right child of parenti.
// Construct the binary tree described by descriptions and return its root.
// The test cases will be generated such that the binary tree is valid.
// Example 1:
//   Input: descriptions = [[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
//   Output: [50,20,80,15,17,19]
//             50
//           /   \
//          20   80
//         / \   /
//        15 17 19
//   Explanation: The root node is the node with value 50 since it has no parent.
//     The resulting binary tree is shown in the diagram.
// Example 2:
//   Input: descriptions = [[1,2,1],[2,3,0],[3,4,1]]
//   Output: [1,2,null,null,3,4]
//         1
//        /
//       2
//        \
//        3
//       /
//      4
//   Explanation: The root node is the node with value 1 since it has no parent.
//     The resulting binary tree is shown in the diagram.
// Constraints:
//   1 <= descriptions.length <= 10⁴
//   descriptions[i].length == 3
//   1 <= parenti, childi <= 10⁵
//   0 <= isLefti <= 1
//   The binary tree described by descriptions is valid.

mod _create_binary_tree_from_descriptions {
    struct Solution;

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

    use std::cell::RefCell;
    use std::collections::{HashMap, HashSet};
    use std::rc::Rc;
    impl Solution {
        pub fn create_binary_tree(descriptions: Vec<Vec<i32>>) -> Option<Rc<RefCell<TreeNode>>> {
            let mut parent = HashSet::new();
            let mut nodes = HashMap::new();
            for v in descriptions {
                // for each description
                let (p, c, l) = (v[0], v[1], v[2]);
                let pn = match nodes.get(&p) {
                    Some(node) => Rc::clone(node),
                    None => {
                        let n = Rc::new(RefCell::new(TreeNode::new(p)));
                        nodes.insert(p, Rc::clone(&n));
                        Rc::clone(&n)
                    },
                };
                let cn = match nodes.get(&c) {
                    Some(node) => Rc::clone(node),
                    None => {
                        // create reference count pointer of TreeNode
                        let n = Rc::new(RefCell::new(TreeNode::new(c)));
                        nodes.insert(c, Rc::clone(&n));
                        Rc::clone(&n)
                    },
                };
                parent.insert(c);
                if l == 1 {
                    pn.borrow_mut().left = Some(Rc::clone(&cn));
                } else {
                    pn.borrow_mut().right = Some(Rc::clone(&cn));
                }
                // for example, if we process first edge,
                // parent's strong_count is 2 (one in map, one is pn),
                // child's string_count is 3 (on in map, one is cn, one referred by parent)
                println!("{} {} {} {}", p, Rc::strong_count(&pn), c, Rc::strong_count(&cn))
                // when the scope end, pn and cn are dropped.
            }
            for (k, _) in &nodes {
                if !parent.contains(k) {
                    return Some(Rc::clone(nodes.get(k).unwrap()));
                }
            }
            // when the function return, the cloned Rc in map will drop
            return None;
        }
    }

    #[test]
    fn test() {
        {
            let v = vec![vec![20,15,1],vec![20,17,0],vec![50,20,1],vec![50,80,0],vec![80,19,1]];
            let r = Solution::create_binary_tree(v);
            println!("r: {}", r.as_ref().unwrap().borrow().val);
            println!("rl: {}", r.as_ref().unwrap().borrow().left.as_ref().unwrap().borrow().val);
            println!("rll: {}", r.as_ref().unwrap().borrow().left.as_ref().unwrap().borrow().left.as_ref().unwrap().borrow().val);
            println!("rlr: {}", r.as_ref().unwrap().borrow().left.as_ref().unwrap().borrow().right.as_ref().unwrap().borrow().val);
            println!("rr: {}", r.as_ref().unwrap().borrow().right.as_ref().unwrap().borrow().val);
            println!("rrl: {}", r.as_ref().unwrap().borrow().right.as_ref().unwrap().borrow().left.as_ref().unwrap().borrow().val);
        }
        {
            let v = vec![vec![1,2,1],vec![2,3,0],vec![3,4,1]];
            let r = Solution::create_binary_tree(v);
            println!("r: {}", r.as_ref().unwrap().borrow().val);
            println!("rl: {}", r.as_ref().unwrap().borrow().
                left.as_ref().unwrap().borrow().val);
            println!("rlr: {}", r.as_ref().unwrap().borrow().
                left.as_ref().unwrap().borrow().
                right.as_ref().unwrap().borrow().val);
            println!("rlrl: {}", r.as_ref().unwrap().borrow().
                left.as_ref().unwrap().borrow().
                right.as_ref().unwrap().borrow().
                left.as_ref().unwrap().borrow().val);
        }
    } 
}
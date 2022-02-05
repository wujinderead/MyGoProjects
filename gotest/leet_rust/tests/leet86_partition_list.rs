// https://leetcode.com/problems/partition-list/

// Given the head of a linked list and a value x, partition it such that all nodes less than
// x come before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in each of the two partitions.
// Example 1:
//   Input: head = [1,4,3,2,5,2], x = 3
//   Output: [1,2,2,4,3,5]
// Example 2:
//   Input: head = [2,1], x = 2
//   Output: [1,2]
// Constraints:
//   The number of nodes in the list is in the range [0, 200].
//   -100 <= Node.val <= 100
//   -200 <= x <= 200

mod _partition_list {
    struct Solution;

    #[derive(PartialEq, Eq, Clone, Debug)]
    pub struct ListNode {
      pub val: i32,
      pub next: Option<Box<ListNode>>
    }

    impl ListNode {
      #[inline]
      fn new(val: i32) -> Self {
        ListNode {
          next: None,
          val
        }
      }
    }

    impl Solution {
        pub fn partition(head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
            if head.is_none() || head.as_ref().unwrap().next.is_none() {
                return head;
            }
            let mut l_head = Box::new(ListNode::new(-1));
            let mut g_head = Box::new(ListNode::new(-1));
            let mut cur = head.as_ref();
            let mut l_tail = &mut l_head;
            let mut g_tail = &mut g_head;
            let (mut l_empty, mut g_empty) = (true, true);
            while cur.is_some() {
                let cur_node = cur.unwrap();
                if cur_node.val < x {
                    l_empty = false;
                    l_tail.next = Some(Box::new(ListNode::new(cur_node.val)));
                    l_tail = l_tail.next.as_mut().unwrap();
                } else {
                    g_empty = false;
                    g_tail.next = Some(Box::new(ListNode::new(cur_node.val)));
                    g_tail = g_tail.next.as_mut().unwrap();
                }
                cur = cur_node.next.as_ref();
            }
            return if l_empty {
                g_head.next
            } else if g_empty {
                l_head.next
            } else {
                l_tail.next = g_head.next;
                l_head.next
            }
        }

        pub fn partition1(head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
            let (mut lhead, mut hhead) = (None, None);
            let (mut low, mut high) = (&mut lhead, &mut hhead);
            let mut cur = head.as_ref();
            while let Some(node) = cur {
                cur = node.next.as_ref();
                if node.val < x {
                    *low = Some(Box::new(ListNode::new(node.val)));
                    low = &mut low.as_deref_mut().unwrap().next;
                } else {
                    *high = Some(Box::new(ListNode::new(node.val)));
                    high = &mut high.as_deref_mut().unwrap().next;
                }
            }
            *low = hhead;
            lhead
        }

        fn print(head: Option<&Box<ListNode>>, a: String) {
            let mut cur = head;
            while cur.is_some() {
                println!("{}", cur.unwrap().val);
                cur = cur.unwrap().next.as_ref();
            }
            println!("end {}", a);
        }
    }

    #[test]
    fn test() {
        {
            let head = Some(Box::new(ListNode{
                val: 2,
                next: Some(Box::new(ListNode{
                    val: 1,
                    next: None,
                })),
            }));
            Solution::print(head.as_ref(), "a1".to_string());
            let ans = Solution::partition(head, 2);
            Solution::print(ans.as_ref(), "a2".to_string());
        }
        {
            let head = Some(Box::new(ListNode{
                val: 2,
                next: Some(Box::new(ListNode{
                    val: 1,
                    next: None,
                })),
            }));
            Solution::print(head.as_ref(), "a1".to_string());
            let ans = Solution::partition1(head, 2);
            Solution::print(ans.as_ref(), "a2".to_string());
        }
        {
            // 1,4,3,2,5,2
            let a6 = Some(Box::new(ListNode{val: 2, next: None}));
            let a5 = Some(Box::new(ListNode{val: 5, next: a6}));
            let a4 = Some(Box::new(ListNode{val: 2, next: a5}));
            let a3 = Some(Box::new(ListNode{val: 3, next: a4}));
            let a2 = Some(Box::new(ListNode{val: 4, next: a3}));
            let a1 = Some(Box::new(ListNode{val: 1, next: a2}));
            Solution::print(a1.as_ref(), "a1".to_string());
            let ans = Solution::partition(a1, 3);
            Solution::print(ans.as_ref(), "a2".to_string());
        }
        {
            // 1,4,3,2,5,2
            let mut a1 = ListNode::new(1);
            let mut a2 = ListNode::new(4);
            let mut a3 = ListNode::new(3);
            let mut a4 = ListNode::new(2);
            let mut a5 = ListNode::new(5);
            let a6 = ListNode::new(2);
            a5.next = Some(Box::new(a6));
            a4.next = Some(Box::new(a5));
            a3.next = Some(Box::new(a4));
            a2.next = Some(Box::new(a3));
            a1.next = Some(Box::new(a2));
            let head = Some(Box::new(a1));
            Solution::print(head.as_ref(), "a1".to_string());
            let ans = Solution::partition(head, 3);
            Solution::print(ans.as_ref(), "a2".to_string());
        }
    } 
}

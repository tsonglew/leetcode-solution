// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
// 
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if head.is_none() {
            return head;
        }
        let mut head_node = head.unwrap();
        let mut nex: Option<Box<ListNode>> = head_node.next.take();
        let mut cur: Option<Box<ListNode>> = Some(head_node);

        while !nex.is_none() {
            let mut nex_node = nex.unwrap();
            let nn = nex_node.next.take();
            nex_node.next = cur;
            cur = Some(nex_node);
            nex = nn;
        }
        return cur;
    }
}

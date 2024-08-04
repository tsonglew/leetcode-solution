// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::cell::RefCell;
use std::rc::Rc;

fn tree2String(root: Option<Rc<RefCell<TreeNode>>>) -> String {
    match root {
        Some(n) => {
            let mut s = String::from(n.borrow().val.to_string());
            return format!(
                "|{}|{}{}",
                s,
                tree2String(n.borrow().left.clone()),
                tree2String(n.borrow().right.clone())
            );
        }
        None => {
            return String::from("|n|");
        }
    }
}

impl Solution {
    pub fn is_subtree(
        root: Option<Rc<RefCell<TreeNode>>>,
        sub_root: Option<Rc<RefCell<TreeNode>>>,
    ) -> bool {
        let root_string = tree2String(root.clone());
        let sub_root_string = tree2String(sub_root.clone());
        // println!("root_str: {}, sub_root_str {}", root_string, sub_root_string);
        return tree2String(root).contains(&tree2String(sub_root));
    }
}

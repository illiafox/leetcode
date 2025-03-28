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
            right: None,
        }
    }
}
use std::cell::RefCell;
use std::rc::Rc;

struct Solution;

impl Solution {
    fn count_good_nodes(node: Option<Rc<RefCell<TreeNode>>>, max_val: i32) -> i32 {
        match node {
            Some(rc_node) => {
                let node_ref = rc_node.borrow();
                let val = node_ref.val;
                let new_max = max_val.max(val);

                (val >= max_val) as i32
                    + Solution::count_good_nodes(node_ref.left.clone(), new_max)
                    + Solution::count_good_nodes(node_ref.right.clone(), new_max)
            }
            None => 0,
        }
    }

    pub fn good_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if let Some(rc_root) = &root {
            let root_val = rc_root.borrow().val;
            Solution::count_good_nodes(root, root_val)
        } else {
            0
        }
    }
}

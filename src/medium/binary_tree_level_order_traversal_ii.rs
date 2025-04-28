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
            right: None,
        }
    }
}

use std::cell::RefCell;
use std::collections::VecDeque;
use std::rc::Rc;

impl Solution {
    pub fn level_order_bottom(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        if root.is_none() {
            return vec![];
        }

        let mut queue = VecDeque::from([root.unwrap()]);
        let mut result = vec![];

        while !queue.is_empty() {
            let size = queue.len();

            let mut layer = Vec::with_capacity(size);

            for _ in 0..size {
                let node = queue.pop_front().unwrap();
                let node_ref = node.borrow();
                layer.push(node_ref.val);

                if let Some(n) = &node_ref.left {
                    queue.push_back(Rc::clone(n))
                }
                if let Some(n) = &node_ref.right {
                    queue.push_back(Rc::clone(n))
                }
            }

            result.push(layer)
        }

        result.reverse();
        result
    }
}

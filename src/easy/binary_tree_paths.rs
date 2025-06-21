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
use std::rc::Rc;
impl Solution {
    pub fn traverse(
        node: Rc<RefCell<TreeNode>>,
        all_paths: &mut Vec<Vec<i32>>,
        current_path: &mut Vec<i32>,
    ) {
        let node = node.borrow();
        current_path.push(node.val);

        if node.left.is_none() && node.right.is_none() {
            all_paths.push(current_path.clone());
        } else {
            if let Some(left) = &node.left {
                Self::traverse(left.clone(), all_paths, current_path);
            }
            if let Some(right) = &node.right {
                Self::traverse(right.clone(), all_paths, current_path);
            }
        }

        current_path.pop();
    }

    pub fn binary_tree_paths(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<String> {
        let mut all_paths = vec![];

        if let Some(node) = root {
            Self::traverse(node, &mut all_paths, &mut Vec::with_capacity(1));
        }

        all_paths
            .iter()
            .map(|path| {
                path.iter()
                    .map(i32::to_string)
                    .collect::<Vec<_>>()
                    .join("->")
            })
            .collect()
    }
}

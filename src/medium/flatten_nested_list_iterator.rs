use std::collections::VecDeque;

#[derive(Debug, PartialEq, Eq)]
pub enum NestedInteger {
    Int(i32),
    List(Vec<NestedInteger>),
}

struct NestedIterator {
    stack: VecDeque<NestedInteger>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NestedIterator {
    fn new(nested_list: Vec<NestedInteger>) -> Self {
        let mut stack = VecDeque::new();
        for l in nested_list.into_iter().rev() {
            stack.push_front(l);
        }

        NestedIterator { stack }
    }

    fn next(&mut self) -> i32 {
        if let Some(NestedInteger::Int(n)) = self.stack.pop_front() {
            n
        } else {
            panic!("next() called without checking has_next()");
        }
    }

    fn has_next(&mut self) -> bool {
        while let Some(top) = self.stack.front() {
            match top {
                NestedInteger::Int(_) => return true,
                NestedInteger::List(_) => {
                    let NestedInteger::List(nested) = self.stack.pop_front().unwrap() else {
                        unreachable!();
                    };
                    for ni in nested.into_iter().rev() {
                        self.stack.push_front(ni);
                    }
                }
            }
        }
        false
    }
}

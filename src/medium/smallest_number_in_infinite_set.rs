use std::cmp::Reverse;
use std::collections::{BinaryHeap, HashSet};

struct SmallestInfiniteSet {
    heap: BinaryHeap<Reverse<i32>>,
    added: HashSet<i32>,
    last_inserted: i32,
}

impl SmallestInfiniteSet {
    fn new() -> Self {
        SmallestInfiniteSet {
            heap: BinaryHeap::new(),
            added: HashSet::new(),
            last_inserted: 1,
        }
    }

    fn pop_smallest(&mut self) -> i32 {
        if let Some(Reverse(n)) = self.heap.pop() {
            self.added.remove(&n);
            return n;
        }

        let result = self.last_inserted;
        self.last_inserted += 1;
        result
    }

    fn add_back(&mut self, num: i32) {
        if num < self.last_inserted && !self.added.contains(&num) {
            self.heap.push(Reverse(num));
            self.added.insert(num);
        }
    }
}

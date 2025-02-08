use std::collections::{BTreeSet, HashMap};

struct NumberContainers {
    all: HashMap<i32, BTreeSet<i32>>, // number -> indexes
    current: HashMap<i32, i32>,       // index -> number
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NumberContainers {
    fn new() -> Self {
        Self {
            all: HashMap::new(),
            current: HashMap::new(),
        }
    }

    fn change(&mut self, index: i32, number: i32) {
        if let Some(&cur_list) = self.current.get(&index) {
            if let Some(s) = self.all.get_mut(&cur_list) {
                s.remove(&index);
            }
        }

        self.all.entry(number).or_default().insert(index);
        self.current.insert(index, number);
    }

    fn find(&self, number: i32) -> i32 {
        if let Some(t) = self.all.get(&number) {
            return t.first().copied().unwrap_or(-1);
        }

        -1
    }
}

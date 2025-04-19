use std::collections::VecDeque;

struct RecentCounter {
    queue: VecDeque<i32>,
}

impl RecentCounter {
    fn new() -> Self {
        RecentCounter {
            queue: VecDeque::new(),
        }
    }

    fn ping(&mut self, t: i32) -> i32 {
        self.queue.push_back(t);

        while let Some(v) = self.queue.front() {
            if *v >= (t - 3000) {
                break;
            }
            _ = self.queue.pop_front()
        }

        self.queue.len() as i32
    }
}

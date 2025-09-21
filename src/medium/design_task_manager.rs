use std::cmp::Ordering;
use std::collections::{BinaryHeap, HashMap};

// https://leetcode.com/problems/design-task-manager

struct TaskManager {
    task_to_user: HashMap<i32, (i32, i32)>,
    heap: BinaryHeap<Task>,
}

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
struct Task {
    priority: i32,
    task_id: i32,
}

impl Ord for Task {
    fn cmp(&self, other: &Self) -> Ordering {
        self.priority
            .cmp(&other.priority)
            .then_with(|| self.task_id.cmp(&other.task_id))
    }
}
impl PartialOrd for Task {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl TaskManager {
    fn new(tasks: Vec<Vec<i32>>) -> Self {
        let mut tm = TaskManager {
            task_to_user: HashMap::with_capacity(tasks.len()),
            heap: BinaryHeap::new(),
        };
        for t in tasks {
            tm.add(t[0], t[1], t[2]);
        }
        tm
    }

    fn add(&mut self, user_id: i32, task_id: i32, priority: i32) {
        self.task_to_user.insert(task_id, (user_id, priority));
        self.heap.push(Task { priority, task_id });
    }

    fn edit(&mut self, task_id: i32, new_priority: i32) {
        if let Some((_, p)) = self.task_to_user.get_mut(&task_id) {
            *p = new_priority;
            self.heap.push(Task {
                priority: new_priority,
                task_id,
            });
        }
    }

    fn rmv(&mut self, task_id: i32) {
        self.task_to_user.remove(&task_id);
    }

    fn exec_top(&mut self) -> i32 {
        while let Some(top) = self.heap.peek().copied() {
            match self.task_to_user.get(&top.task_id) {
                Some(&(user_id, cur_priority)) if cur_priority == top.priority => {
                    self.heap.pop();
                    self.task_to_user.remove(&top.task_id);
                    return user_id;
                }
                _ => {
                    self.heap.pop();
                }
            }
        }
        -1
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn v(u: i32, t: i32, p: i32) -> Vec<i32> {
        vec![u, t, p]
    }

    #[test]
    fn init_and_exec_top_basic() {
        // [[user, task, priority]]
        let tasks = vec![v(1, 101, 8), v(2, 102, 20), v(3, 103, 5)];
        let mut tm = TaskManager::new(tasks);
        // Highest priority is (20, task 102) => user 2
        assert_eq!(tm.exec_top(), 2);
        // Next is (8, task 101) => user 1
        assert_eq!(tm.exec_top(), 1);
        // Then (5, task 103) => user 3
        assert_eq!(tm.exec_top(), 3);
        // Nothing left
        assert_eq!(tm.exec_top(), -1);
    }

    #[test]
    fn tie_breaker_uses_higher_task_id() {
        let tasks = vec![v(1, 10, 5), v(2, 11, 5), v(3, 9, 4)];
        let mut tm = TaskManager::new(tasks);
        // Same priority 5, higher task_id=11 should win (user 2)
        assert_eq!(tm.exec_top(), 2);
        // Then task_id=10 (user 1)
        assert_eq!(tm.exec_top(), 1);
        // Then the remaining one
        assert_eq!(tm.exec_top(), 3);
        assert_eq!(tm.exec_top(), -1);
    }

    #[test]
    fn edit_updates_priority_and_skips_stale_heap_entries() {
        let tasks = vec![v(1, 101, 8), v(2, 102, 20)];
        let mut tm = TaskManager::new(tasks);
        // Boost task 101 above 102
        tm.edit(101, 25);
        // Top should now be user 1 (task 101)
        assert_eq!(tm.exec_top(), 1);
        // Then user 2 (task 102)
        assert_eq!(tm.exec_top(), 2);
        assert_eq!(tm.exec_top(), -1);
    }

    #[test]
    fn multiple_edits_only_last_priority_counts() {
        let tasks = vec![v(1, 101, 8), v(2, 102, 7)];
        let mut tm = TaskManager::new(tasks);
        // Push a couple of edits to the same task
        tm.edit(102, 9);
        tm.edit(102, 30); // final intended state
                          // Top must be the latest (30, task 102) -> user 2
        assert_eq!(tm.exec_top(), 2);
        // Then (8, task 101) -> user 1
        assert_eq!(tm.exec_top(), 1);
        assert_eq!(tm.exec_top(), -1);
    }

    #[test]
    fn rmv_removes_and_lazy_deletes_heap_entries() {
        let tasks = vec![v(1, 101, 8), v(2, 102, 20), v(3, 103, 5)];
        let mut tm = TaskManager::new(tasks);
        // Remove the current top (task 102)
        tm.rmv(102);
        // Next valid should be task 101 -> user 1
        assert_eq!(tm.exec_top(), 1);
        // Then 103 -> user 3
        assert_eq!(tm.exec_top(), 3);
        assert_eq!(tm.exec_top(), -1);
    }

    #[test]
    fn remove_then_readd_same_taskid_different_user() {
        // Reproduce the sequence from your example
        let tasks = vec![v(1, 101, 8), v(2, 102, 20), v(3, 103, 5)];
        let mut tm = TaskManager::new(tasks);

        tm.add(4, 104, 5); // ["add",[4,104,5]]
        tm.edit(102, 9); // ["edit",[102,9]]

        // execTop -> highest is (9,102) -> user 2
        assert_eq!(tm.exec_top(), 2);

        // rmv(101), then add same taskId=101 for a different user
        tm.rmv(101);
        tm.add(50, 101, 8);

        // Now top among remaining is (8,101)-> user 50 (tie-breaker vs 104/103 with prio 5)
        assert_eq!(tm.exec_top(), 50);
        // Next the 5's: highest taskId is 104 -> user 4
        assert_eq!(tm.exec_top(), 4);
        // Then 103 -> user 3
        assert_eq!(tm.exec_top(), 3);
        // No more
        assert_eq!(tm.exec_top(), -1);
    }

    #[test]
    fn empty_returns_minus_one() {
        let mut tm = TaskManager::new(vec![]);
        assert_eq!(tm.exec_top(), -1);
        // Add and remove then ensure -1 again
        tm.add(7, 200, 1);
        tm.rmv(200);
        assert_eq!(tm.exec_top(), -1);
    }
}

struct Solution;

impl Solution {
    pub fn can_visit_all_rooms(rooms: Vec<Vec<i32>>) -> bool {
        let mut keys = vec![false; rooms.len()];
        keys[0] = true;

        let mut stack = vec![0];

        while let Some(room) = stack.pop() {
            for &key in &rooms[room] {
                let key = key as usize;
                if !keys[key] {
                    keys[key] = true;
                    stack.push(key);
                }
            }
        }

        !keys.contains(&false)
    }
}

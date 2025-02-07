use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn query_results(_limit: i32, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut current_colors: HashMap<i32, i32> = HashMap::new();

        let mut m: HashMap<i32, i32> = HashMap::new();

        let mut out = Vec::new();

        for n in queries {
            if let Some(c) = current_colors.get(&n[0]) {
                let s = m.get_mut(c).unwrap();
                *s -= 1;
                if *s == 0 {
                    m.remove(c);
                }
            }

            let new_color = n[1];
            *m.entry(new_color).or_default() += 1;

            out.push(m.len() as i32);

            current_colors.insert(n[0], new_color);
        }

        out
    }
}

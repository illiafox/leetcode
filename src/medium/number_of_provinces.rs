struct Solution;

impl Solution {
    fn dfs(city: usize, is_connected: &Vec<Vec<i32>>, visited: &mut Vec<bool>) {
        visited[city] = true;
        for (next_city, &connected) in is_connected[city].iter().enumerate() {
            if connected == 1 && !visited[next_city] {
                Solution::dfs(next_city, is_connected, visited)
            }
        }
    }

    pub fn find_circle_num(is_connected: Vec<Vec<i32>>) -> i32 {
        let mut visited = vec![false; is_connected.len()];
        let mut provinces = 0;

        for city in 0..is_connected.len() {
            if !visited[city] {
                Solution::dfs(city, &is_connected, &mut visited);
                provinces += 1;
            }
        }

        provinces
    }
}

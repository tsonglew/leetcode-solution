impl Solution {
    pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
        let mut m = vec![vec![0; obstacle_grid[0].len()]; obstacle_grid.len()];
        if let Some(last_row) = m.last_mut() {
            if let Some(last_el) = last_row.last_mut() {
                *last_el = 1;
            }
        }

        for j in (0..obstacle_grid[0].len()).rev() {
            for i in (0..obstacle_grid.len()).rev() {
                if obstacle_grid[i][j] == 1 {
                    m[i][j] = 0;
                    continue;
                }
                if i < obstacle_grid.len() - 1 {
                    m[i][j] += m[i+1][j];
                }
                if j < obstacle_grid[0].len() - 1 {
                    m[i][j] += m[i][j+1];
                }
            }
        }
        m[0][0]
    }
}

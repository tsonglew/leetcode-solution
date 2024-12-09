use std::cmp::min;

impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        let mut m = vec![vec![0; grid[0].len()]; grid.len()];
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                if i == 0 && j == 0 {
                    m[i][j] = grid[i][j];
                    continue;
                }
                if i == 0 {
                    m[i][j] = m[i][j-1] + grid[i][j];
                } else if j == 0{
                    m[i][j] = m[i-1][j] + grid[i][j];
                } else {
                    m[i][j] = min(m[i][j-1], m[i-1][j]) + grid[i][j];
                }
            }
        }
        m[grid.len()-1][grid[0].len()-1]
    }
}

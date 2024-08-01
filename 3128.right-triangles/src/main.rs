impl Solution {
    pub fn number_of_right_triangles(grid: Vec<Vec<i32>>) -> i64 {
        let mut r = 0 as i64;
        let mut row_sum = vec![0; grid.len()];
        let mut col_sum = vec![0; grid[0].len()];
        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                row_sum[i] += grid[i][j] as i64;
                col_sum[j] += grid[i][j] as i64;
            }
        }
        for i in 0..row_sum.len() {
            for j in 0..col_sum.len() {
                if grid[i][j] == 0 {
                    continue;
                }
                r += (&row_sum[i] - 1) * (&col_sum[j] - 1);
            }
        }
        r
    }
}

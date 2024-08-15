use std::cmp;

impl Solution {
    pub fn max_score(grid: Vec<Vec<i32>>) -> i32 {
        let mut res = -100000;
        let mut mem = grid.clone();
        for i in (0..mem.len()).rev() {
            for j in (0..mem[0].len()).rev() {
                let mut r_val = -100000;
                let mut d_val = -100000;
                if i < mem.len() - 1 {
                    d_val = grid[i+1][j] + mem[i+1][j] - grid[i][j];
                }
                if j < mem[0].len() - 1 {
                    r_val = grid[i][j+1] + mem[i][j+1] - grid[i][j];
                }
                mem[i][j] = cmp::max(r_val, cmp::max(0, d_val));
            }
        }
        println!("{:?}", mem);

        for i in (0..mem.len()).rev() {
            for j in (0..mem[0].len()).rev() {
                let mut r_val = -100000;
                let mut d_val = -100000;
                if i < mem.len() - 1 {
                    d_val = mem[i+1][j] + (grid[i+1][j]-grid[i][j]);
                }
                if j < mem[0].len() - 1 {
                    r_val = mem[i][j+1] + (grid[i][j+1]-grid[i][j]);
                }
                res = cmp::max(res, cmp::max(r_val, d_val));
            }
        }

        res
    }
}

use std::cmp;

impl Solution {
    pub fn min_falling_path_sum(matrix: Vec<Vec<i32>>) -> i32 {
        if matrix.len() == 1 {
            return matrix[0][0];
        }
        let mut m = matrix.clone();
        for i in 1..matrix.len() {
            for j in 0..matrix.len() {
                let mut upper_left = 10000;
                let mut upper_right = 10000;
                let upper = m[i - 1][j];
                if j > 0 {
                    upper_left = m[i - 1][j - 1];
                }
                if j < matrix.len() - 1 {
                    upper_right = m[i - 1][j + 1];
                }
                m[i][j] += cmp::min(cmp::min(upper_left, upper_right), upper);
            }
        }
        *m[matrix.len() - 1].iter().min().unwrap()
    }
}

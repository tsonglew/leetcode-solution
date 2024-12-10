pub struct Solution {}

impl Solution {
    pub fn minimum_total(mut triangle: Vec<Vec<i32>>) -> i32 {
        let depth = triangle.len();
        for i in (0..depth-1).rev() {
            for j in (0..triangle[i].len()) {
                triangle[i][j] += std::cmp::min(triangle[i+1][j], triangle[i+1][j+1]);
            }
        }
        triangle[0][0]
    }
}

struct Solution;

impl Solution {
    pub fn oranges_rotting(grid: Vec<Vec<i32>>) -> i32 {
        let mut m: Vec<Vec<i32>> = grid.clone();
        for i in 0..m.len() {
            for j in 0..m[i].len() {
                match grid[i][j] {
                    0 => m[i][j] = -2,
                    1 => m[i][j] = -1,
                    _ => m[i][j] = 0,
                }
            }
        }
        println!("{:?}", m);
        for i in 0..m.len() {
            for j in 0..m[i].len() {
                if m[i][j] == 2 {
                    Self::rot(&mut m, i as i32, j as i32, 0);
                }
            }
        }
        0
    }

    fn deepsearch(m: &mut Vec<Vec<i32>>, i: i32, j: i32, depth: i32) {}
}

fn main() {
    assert!(Solution::oranges_rotting(vec![vec![2, 1, 1], vec![1, 1, 0], vec![0, 1, 1]]) == 4);
}

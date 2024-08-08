use std::cmp;

impl Solution {
    pub fn maximal_square(matrix: Vec<Vec<char>>) -> i32 {
        let mut continue_1_l = vec![vec![0; matrix[0].len()]; matrix.len()];
        let mut continue_1_u = vec![vec![0; matrix[0].len()]; matrix.len()];
        continue_1_l[0][0] = matrix[0][0].to_digit(10).unwrap() as i32;
        continue_1_u[0][0] = matrix[0][0].to_digit(10).unwrap() as i32;
        let mut mx = matrix[0][0].to_digit(10).unwrap() as i32;

        for i in 0..matrix.len() {
            for j in 0..matrix[0].len() {
                let num = matrix[i][j].to_digit(10).unwrap();
                if num == 0 {
                    continue_1_l[i][j] = 0;
                    continue_1_u[i][j] = 0;
                    continue;
                }
                if mx == 0 {
                    mx = 1;
                }

                let mut prev_l_mx = 0;
                let mut prev_u_mx = 0;
                if i > 0 {
                    prev_u_mx = continue_1_u[i - 1][j];
                }
                if j > 0 {
                    prev_l_mx = continue_1_l[i][j - 1];
                }
                continue_1_u[i][j] = prev_u_mx + 1;
                continue_1_l[i][j] = prev_l_mx + 1;

                if i > 0 && j > 0 {
                    let r = cmp::min(
                        cmp::min(1 + continue_1_l[i - 1][j - 1], continue_1_l[i][j]),
                        cmp::min(1 + continue_1_u[i - 1][j - 1], continue_1_u[i][j]),
                    );
                    mx = cmp::max(r.pow(2), mx);
                    continue_1_l[i][j] = r;
                    continue_1_u[i][j] = r;
                }
            }
        }
        // println!("continue_1_l: {:?}", continue_1_l);
        // println!("continue_1_u: {:?}", continue_1_u);

        mx
    }
}

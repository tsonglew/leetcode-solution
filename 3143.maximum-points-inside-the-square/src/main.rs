impl Solution {
    pub fn max_points_inside_square(points: Vec<Vec<i32>>, s: String) -> i32 {
        let mut max_abs_val = 0;
        for p in points.iter() {
            let p0_abs = p[0].abs();
            let p1_abs = p[1].abs();
            if p0_abs > max_abs_val {
                max_abs_val = p0_abs;
            }
            if p1_abs > max_abs_val {
                max_abs_val = p1_abs;
            }
        }

        let mut p_min_border = vec![max_abs_val + 1; 26];
        let mut border_val = max_abs_val;
        for (i, c) in s.chars().enumerate() {
            let p_dis = std::cmp::max(points[i][0].abs(), points[i][1].abs());
            let ci = (c as u8 - b'a') as usize;
            if p_min_border[ci] == p_dis {
                border_val = std::cmp::min(p_dis - 1, border_val);
            } else {
                border_val = std::cmp::min(border_val, std::cmp::max(p_min_border[ci], p_dis) - 1);
            }
            p_min_border[ci] = std::cmp::min(p_min_border[ci], p_dis);
        }
        let mut r = 0;
        for i in p_min_border.iter() {
            if (*i <= border_val) {
                r += 1;
            }
        }
        r
    }
}


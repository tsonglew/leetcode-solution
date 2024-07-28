impl Solution {
    pub fn cal_points(operations: Vec<String>) -> i32 {
        let mut points: Vec<i32> = Vec::new();
        for i in &operations {
            match i.parse::<i32>() {
                Ok(n) => points.push(n),
                Err(e) => {
                    if "C" == i {
                        points.pop();
                    } else if "D" == i {
                        match points.last() {
                            Some(&last_value) => points.push(last_value * 2),
                            None => {}
                        }
                    } else if "+" == i {
                        points.push(&points[points.len() - 2] + &points[points.len() - 1])
                    }
                }
            }
        }
        return points.iter().sum();
    }
}

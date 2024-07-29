fn pow_mod(n: i32, i: i32, md: i32) -> i32 {
    if i == 1 {
        return n;
    } else if i % 2 == 0 {
        let m = pow_mod(n, i >> 1, md);
        return m * m % md;
    } else {
        let m = pow_mod(n, i >> 1, md);
        return m * m * n % md;
    }
}

impl Solution {
    pub fn get_good_indices(variables: Vec<Vec<i32>>, target: i32) -> Vec<i32> {
        let mut r: Vec<i32> = Vec::new();
        for (i, sv) in variables.iter().enumerate() {
            if pow_mod(pow_mod(sv[0], sv[1], 10) % 10, sv[2], sv[3]) % sv[3] == target {
                r.push(i as i32);
            }
        }
        return r;
    }
}


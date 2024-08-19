impl Solution {
    pub fn max_energy_boost(energy_drink_a: Vec<i32>, energy_drink_b: Vec<i32>) -> i64 {
        let n = energy_drink_a.len();
        let mut drink_a = vec![0 as i64; n];
        let mut drink_b = vec![0 as i64; n];
        drink_a[0] = energy_drink_a[0] as i64;
        drink_a[1] = energy_drink_a[1] as i64 + drink_a[0];
        drink_b[0] = energy_drink_b[0] as i64;
        drink_b[1] = energy_drink_b[1] as i64 + drink_b[0];

        for i in 2..n {
            drink_a[i] = energy_drink_a[i] as i64
                + if drink_a[i - 1] > drink_b[i - 2] {
                    drink_a[i - 1]
                } else {
                    drink_b[i - 2]
                };
            drink_b[i] = energy_drink_b[i] as i64
                + if drink_b[i - 1] > drink_a[i - 2] {
                    drink_b[i - 1]
                } else {
                    drink_a[i - 2]
                };
        }
        // println!("a: {:?}", drink_a);
        // println!("b: {:?}", drink_b);
        if drink_a[n - 1] > drink_b[n - 1] {
            drink_a[n - 1]
        } else {
            drink_b[n - 1]
        }
    }
}

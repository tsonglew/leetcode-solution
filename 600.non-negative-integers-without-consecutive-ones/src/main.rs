impl Solution {
    pub fn find_integers(n: i32) -> i32 {
        if n == 1 {
            return 2;
        }
        let mut n_bits = vec![];
        let mut n_copy = n;
        let mut r = 0;

        while n_copy > 0 {
            n_bits.push(n_copy % 2);
            n_copy >>= 1;
        }

        let mut cache = vec![0; n_bits.len()];
        for i in 0..cache.len() {
            if i == 0 {
                cache[i] = 1;
            } else if i == 1 {
                cache[i] = 2;
            } else if i == 2 {
                cache[i] = 3;
            } else {
                cache[i] = cache[i - 1] + cache[i - 2];
            }
        }

        for i in (0..=(n_bits.len() - 1)).rev() {
            if n_bits[i] == 0 {
                continue;
            }
            r += cache[i];

            if (i < n_bits.len() - 1 && n_bits[i + 1] == 1) {
                return r;
            }
        }
        return r + 1;
    }
}

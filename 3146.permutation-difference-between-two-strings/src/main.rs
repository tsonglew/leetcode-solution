use std::collections::HashMap;

impl Solution {
    pub fn find_permutation_difference(s: String, t: String) -> i32 {
        let mut s_idx = HashMap::new();
        let mut res: i32 = 0;

        for (i, c) in s.char_indices() {
            s_idx.insert(c, i as i32);
        }
        // println!("{:?}", s_idx);
        for (i, c) in t.char_indices() {
            res += (s_idx.get(&c).unwrap() - (i as i32)).abs();
        }

        res
    }
}


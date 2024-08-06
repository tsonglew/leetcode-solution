use std::cmp;
use std::collections::HashMap;
use std::collections::HashSet;

impl Solution {
    pub fn delete_and_earn(nums: Vec<i32>) -> i32 {
        let mut dedup = vec![];
        let mut h = HashMap::new();
        for n in nums.iter() {
            if !h.contains_key(n) {
                dedup.push(n);
            }
            let e = h.entry(n).or_insert(0);
            *e += 1;
        }

        dedup.sort_by(|a, b| a.cmp(&b));

        let mut s = vec![0; dedup.len()];
        for (i, n) in dedup.iter().enumerate() {
            if let Some(v) = h.get(n) {
                if i == 0 {
                    s[i] = dedup[i] * v;
                } else {
                    if *dedup[i] == dedup[i - 1] + 1 {
                        let l1 = if i > 1 { s[i - 2] } else { 0 };
                        let l2 = if i > 2 { s[i - 3] } else { 0 };
                        s[i] = dedup[i] * v + cmp::max(l1, l2);
                    } else {
                        let l1 = if i > 0 { s[i - 1] } else { 0 };
                        let l2 = if i > 1 { s[i - 2] } else { 0 };
                        s[i] = dedup[i] * v + cmp::max(l1, l2);
                    }
                }
            }
        }
        // println!("dedup: {:?}, h: {:?}, s: {:?}", dedup, h, s);
        match s.len() {
            1 => s[0],
            _ => *(s.iter().rev().take(2).max().unwrap()),
        }
    }
}

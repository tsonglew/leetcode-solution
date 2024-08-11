use std::cmp;

impl Solution {
    pub fn longest_palindrome_subseq(s: String) -> i32 {
        let mut st = vec![vec![0; s.len()]; s.len()];
        let mut mx = 0;
        let chars: Vec<char> = s.chars().collect();
        for j in 0..s.len() {
            for i in (0..=j).rev() {
                if i == j {
                    st[i][j] = 1;
                } else if j == i + 1 {
                    st[i][j] = if chars[i] == chars[j] { 2 } else { 1 };
                } else {
                    st[i][j] = if chars[i] == chars[j] {
                        2 + st[i + 1][j - 1]
                    } else {
                        cmp::max(st[i][j - 1], st[i + 1][j])
                    };
                }
                if st[i][j] > mx {
                    mx = st[i][j];
                }
            }
        }
        mx
    }
}

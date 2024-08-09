use std::cmp;

impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        let mut st = vec![vec![0; s.len()]; s.len()];
        let mut mx = 1;
        let mut l = 0;
        let mut r = 0;
        let chars: Vec<char> = s.chars().collect();
        for i in 0..st.len() {
            for j in 0..=i {
                if i == j {
                    st[i][j] = 1;
                } else if i == j + 1 {
                    st[i][j] = if chars[i] == chars[j] { 1 } else { 0 };
                } else {
                    st[i][j] = if chars[i] == chars[j] && st[i - 1][j + 1] == 1 {
                        1
                    } else {
                        0
                    };
                }

                // println!("{} {}: {}", i, j, st[i][j]);
                if st[i][j] == 1 && i + 1 - j > mx {
                    mx = i + 1 - j;
                    l = i;
                    r = j;
                }
            }
        }
        // println!("{:?}, {} {}", st, r, l);
        s[r..=l].to_string()
    }
}

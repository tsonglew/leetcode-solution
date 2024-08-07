impl Solution {
    pub fn unique_paths(m: i32, n: i32) -> i32 {
        let mut s = vec![vec![0; n as usize]; m as usize];
        s[0][0] = 1;
        for i in 0..m as usize {
            for j in 0..n as usize {
                if i > 0 {
                    s[i][j] += s[i - 1][j]
                }
                if j > 0 {
                    s[i][j] += s[i][j - 1]
                }
            }
        }
        s[m as usize - 1][n as usize - 1]
    }
}

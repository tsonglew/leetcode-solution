impl Solution {
    pub fn tribonacci(n: i32) -> i32 {
        let mut tribs = vec![0; 38];
        tribs[0] = 0;
        tribs[1] = 1;
        tribs[2] = 1;
        for i in 3..tribs.len() {
            tribs[i] = tribs[i - 3] + tribs[i - 2] + tribs[i - 1];
        }
        return tribs[n as usize];
    }
}

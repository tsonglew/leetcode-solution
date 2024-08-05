impl Solution {
    pub fn fib(n: i32) -> i32 {
        let mut fibs = vec![0; 31];
        fibs[0] = 0;
        fibs[1] = 1;
        for i in 2..fibs.len() {
            fibs[i] = fibs[i - 1] + fibs[i - 2];
        }
        return fibs[n as usize];
    }
}

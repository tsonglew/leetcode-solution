impl Solution {
    pub fn is_array_special(nums: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<bool> {
        let mut mem = vec![0; nums.len()];
        for i in 1..mem.len() {
            if (nums[i] + nums[i - 1]) % 2 == 1 {
                mem[i] = mem[i - 1] + 1;
            } else {
                mem[i] = 0;
            }
        }
        let mut res = vec![];
        for q in queries.iter() {
            if q[1] - mem[q[1] as usize] <= q[0] {
                res.push(true);
            } else {
                res.push(false);
            }
        }
        res
    }
}

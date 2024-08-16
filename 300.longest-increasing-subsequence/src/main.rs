impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let mut mem = vec![0; nums.len()];
        mem[0] = 1;
        let mut res = 1;
        for i in 1..nums.len() {
            let mut x = 1;
            for j in 0..i {
                if nums[j] < nums[i] && mem[j] + 1 > x {
                    x = mem[j] + 1;
                }
            }
            mem[i] = x;
            if mem[i] > res {
                res = mem[i];
            }
        }
        println!("{:?}", mem);
        res
    }
}

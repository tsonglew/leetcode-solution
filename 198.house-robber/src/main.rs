use std::cmp;

impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        if nums.len() < 3 {
            return *nums.iter().max().unwrap();
        }
        let mut s = vec![0; nums.len()];
        s[0] = nums[0];
        s[1] = nums[1];
        s[2] = nums[0] + nums[2];
        for i in 3..s.len() {
            s[i] = cmp::max(s[i - 2], s[i - 3]) + nums[i];
        }
        cmp::max(s[s.len() - 1], s[s.len() - 2])
    }
}

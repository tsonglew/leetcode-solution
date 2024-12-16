impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let mut dp = vec![0; nums.len()];
        dp[0] = nums[0];
        for i in 1..dp.len() {
            dp[i] = std::cmp::max(dp[i-1]+nums[i], nums[i]);
        }        
        *dp.iter().max().unwrap()
    }
}

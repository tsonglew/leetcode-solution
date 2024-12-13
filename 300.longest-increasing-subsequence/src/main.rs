impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let mut dp = vec![1; nums.len()];
        let mut max_v = 0;
        for i in 0..nums.len() {
            let mut max_len = 1;
            for j in 0..i {
                if nums[i] > nums[j] {
                    max_len = std::cmp::max(max_len, dp[j] + 1);
                }
            }
            dp[i] = max_len;
            max_v = std::cmp::max(max_v, max_len);
        }
        max_v
    }
}

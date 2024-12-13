impl Solution {
    pub fn find_number_of_lis(nums: Vec<i32>) -> i32 {
        let mut dp = vec![vec![0; nums.len()]];
        for i in 0..dp.len() {
            let mut mx = nums[i];
            for j in i..dp.len() {
                if nums[j] >= mx {
                    
                }
            }
        }
    }
}

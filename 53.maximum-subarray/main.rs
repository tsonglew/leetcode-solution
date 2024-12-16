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

use std::cmp::max;

#[derive(Debug)]
struct Result {
    left_max: i32,
    right_max: i32,
    inner_max: i32,
    sum: i32
}
impl PrefixTreeSolution {

    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let r = Solution::f(&nums, 0, 0+nums.len()-1);
        max(max(r.left_max, r.right_max), r.inner_max)
    }

    fn f(nums: &Vec<i32>, i: usize, j: usize) -> Result {
        if i == j {
            return Result{left_max: nums[i], right_max: nums[i], inner_max: nums[i], sum: nums[i]}
        }
        let left_r = Solution::f(nums, i, (i+j)/2);
        let right_r = Solution::f(nums, (i+j)/2+1, j);
        Result{
            left_max: max(left_r.left_max, left_r.sum + right_r.left_max),
            right_max: max(right_r.right_max, right_r.sum + left_r.right_max),
            inner_max: max(max(max(max(left_r.right_max, right_r.left_max), left_r.right_max + right_r.left_max), left_r.inner_max), right_r.inner_max),
            sum: left_r.sum + right_r.sum
        }
    }
}

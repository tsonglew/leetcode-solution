impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        let mut i = 0;
        let mut j = 0;
        while j < nums.len() {
            if nums[i] == nums[j] {
                j+= 1;
            } else {
                i+= 1;
                nums[i] = nums[j];
            }
        }
        (i+1) as i32
    }
}

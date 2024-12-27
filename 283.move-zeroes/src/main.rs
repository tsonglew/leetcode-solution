impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let mut i = 0;
        let mut j = 0;
        while i < nums.len() && nums[i] != 0 {
            i += 1;
        }
        if i >= nums.len() {
            return;
        }
        j = i;
        while j < nums.len() && i < nums.len() {
            while i < nums.len() && nums[i] != 0 {
                i += 1;
            }
            if i >= nums.len() {
                return;
            }
            while j < nums.len() && nums[j] == 0  {
                j += 1;
            }
            if j >= nums.len() {
                return;
            }
            nums[i] = nums[j];
            nums[j] = 0;
            i+=1
        }
    }
}

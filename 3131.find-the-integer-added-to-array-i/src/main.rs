impl Solution {
    pub fn added_integer(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        let sum2: i32 = nums2.iter().sum();
        let sum1: i32 = nums1.iter().sum();
        (sum2 - sum1) / nums1.len() as i32
    }
}

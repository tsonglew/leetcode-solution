impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let mut j = m - 1;
        let mut k = n - 1;
        for i in (0..nums1.len()).rev() {
            if j >= 0 && (k < 0 || nums1[j as usize] > nums2[k as usize]) {
                nums1[i] = nums1[j as usize];
                j -= 1;
            } else {
                nums1[i] = nums2[k as usize];
                k -= 1;
            }
        }
    }
}

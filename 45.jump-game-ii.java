class Solution {
    public int jump(int[] nums) {
        int p = 0;
        int ans = 0;
        while (p < nums.length-1) {
            int next_p = 0;
            int next_p_next = 0;
            for (int i = 1; i <= nums[p] ; i++) {
                if (p + i >= nums.length-1){
                    return ans + 1;
                }
                if (nums[p + i] + p + i > next_p_next) {
                    next_p = p + i;
                    next_p_next = nums[p+i] + p + i;
                }
            }
            ans ++;
            p = next_p;
        }
        return ans;
    }
}

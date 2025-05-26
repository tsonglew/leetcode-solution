class Solution {
    public int maxSubArray(int[] nums) {
        int[] ans = new int[nums.length];
        int maxval = nums[0];
        for (int i = 0; i < ans.length; i++) {
            if (i == 0) {
                ans[i] = nums[i];
            } else {
                ans[i] = Math.max(nums[i], ans[i-1] + nums[i]);
            }
            maxval = Math.max(ans[i], maxval);
        }
        return maxval;
    }
}

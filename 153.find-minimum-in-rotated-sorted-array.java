class Solution {
    public int findMin(int[] nums) {
        // 3 4 5 1 2
        // nums[i] <= 2
        // 0 0 0 1 1
        //
        // 0 1 2 3 4
        // nums[i] <= 4
        // 1 1 1 1 1
        int l = 0, r = nums.length - 1;
        int lastVal = nums[r];
        int ans = nums.length;
        while (l <= r) {
            int mid = (int) Math.floor((l + r) / 2);
            int val = this.getVal(nums, mid, lastVal);
            if (val == 1) {
                ans = Math.min(ans, mid);
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }
        return nums[ans];
    }

    public int getVal(int[] nums, int idx, int val) {
        return nums[idx] <= val ? 1 : 0;
    }
}

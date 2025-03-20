class Solution {
    public int findMin(int[] nums) {
        // 3 4 5 1 2
        // nums[i] <= 2
        // 0 0 0 1 1
        //
        // 0 1 2 3 4
        // nums[i] <= 4
        // 1 1 1 1 1
        //
        // 3 1 3
        // 1 1 1 

        // 找最小的大于 0 的元素
        int l = 0, r = nums.length - 1;
        int lastVal = nums[r];
        int ans = nums.length;
        int ansVal = nums[0];
        while (l <= r) {
            lastVal = nums[r];
            int mid = (int) Math.floor((l + r) / 2);
            int val = this.getVal(nums, mid, lastVal);
            if (val == 1) {
                ansVal = Math.min(ansVal, nums[mid]);
                r = mid - 1;
            } else if (val == 0) {
                l = mid + 1;
            } else {
                ansVal = Math.min(ansVal, nums[r]);
                r -= 1;
            }
        }
        return ansVal;
    }

    public int getVal(int[] nums, int idx, int val) {
        if (nums[idx] < val) {
            return 1;
        } else if (nums[idx] > val) {
            return 0;
        }
        return -1;
        // return nums[idx] <= val ? 1 : 0;
    }
}


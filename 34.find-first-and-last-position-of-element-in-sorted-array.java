class Solution {
    public int[] searchRange(int[] nums, int target) {
        return new int[]{this.searchFirst(nums, target), this.searchLast(nums, target)};
    }

    public int searchLast(int[] nums, int target) {
        // 5 7 7 8 8 10
        // < 9
        int i = 0, j = nums.length - 1;
        int ans = -1;
        while (i <= j) {
            int mid = (int) Math.ceil((i + j) / 2);
            if (nums[mid] > target) {
                j = mid - 1;
            } else if (nums[mid] == target) {
                ans = Math.max(ans, mid);
                i = mid + 1;
            } else {
                i = mid + 1;
            }
        }
        return ans;
    }

    public int searchFirst(int[] nums, int target) {
        // 5 7 7 8 8 10
        int i = 0, j = nums.length - 1;
        int ans = nums.length;
        while (i <= j) {
            int mid = (int) Math.ceil((i + j) / 2);
            if (nums[mid] > target) {
                j = mid - 1;
            } else if (nums[mid] == target) {
                ans = Math.min(ans, mid);
                j = mid - 1;
            } else {
                i = mid + 1;
            }
        }
        if (ans > nums.length - 1) {
            return -1;
        }
        return ans;
    }
}

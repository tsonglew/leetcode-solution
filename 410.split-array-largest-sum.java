class Solution {
    public int splitArray(int[] nums, int k) {
        int sum = 0;
        for (int i = 0; i < nums.length; i++) {
            sum += nums[i];
        }

        int l = 0, r = sum;
        int minmax = sum;
        while (l <= r) {
            int mid = (int)Math.floor((l+r)/2);
            if (this.fit(nums, k, mid)) {
                minmax = Math.min(minmax, mid);
                r = mid -1;
            } else {
                l = mid + 1;
            }
        }

        int maxsum = 0;
        sum = 0;
        for (int i = 0; i < nums.length; i++) {
            if (sum + nums[i] > minmax) {
                maxsum = Math.max(sum, maxsum);
                sum = nums[i];
            } else {
                sum += nums[i];
            }   
        }
        maxsum = Math.max(sum, maxsum);
        return maxsum;
    }

    public boolean fit(int[] nums, int k, int max) {
        int sum = 0;
        int subcnt = 0;
        for (int i = 0; i < nums.length; i++) {
            if (sum + nums[i] > max) {
                sum = nums[i];
                subcnt ++;
            } else {
                sum += nums[i];
            }
            if (subcnt > k) {
                return false;
            }
        }
        if (sum > 0) {
            return sum <= max && subcnt <= k-1;
        }
        return subcnt <= k;

    }
}

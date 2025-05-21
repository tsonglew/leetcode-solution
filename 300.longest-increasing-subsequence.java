class Solution {
    int[] nums;

    public int lengthOfLIS(int[] nums) {
        this.nums = nums;
        int[] ans = new int[nums.length];
        ans[0] = 1;
        int maxans = 1;
        for (int i = 1; i < nums.length; i++) {
            int maxval = 1;
            int maxvalidx = -1;

            for (int j = 0; j < i; j++) {
                if (nums[j] >= nums[i]) {
                    continue;
                }
                if (ans[j] + 1 > maxval) {
                    maxval = ans[j] + 1;
                    maxvalidx = j;
                }
            }

            ans[i] = maxval;
            if (ans[i] > maxans) {
                maxans = ans[i];
            }
        }

        return maxans;
    }

    private void iter(int[] nums) {
        for (int num : nums) {
            System.out.printf("%d ", num);
        }
        System.out.println();
    }
}

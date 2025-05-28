class Solution {
    public int maxProduct(int[] nums) {
        int ans = nums[0];
        int n = nums.length;
        int[] minNums = new int[n];
        int[] maxNums = new int[n];
        for (int i = 0; i < nums.length; i++) {
            minNums[i] = nums[i];
            maxNums[i] = nums[i];
            if (i == 0) {
                continue;
            } 
            
            minNums[i] = Math.min(Math.min(minNums[i], nums[i] * minNums[i-1]), nums[i] * maxNums[i-1]);
            maxNums[i] = Math.max(Math.max(maxNums[i], nums[i] * minNums[i-1]), nums[i] * maxNums[i-1]);
    
            ans = Math.max(maxNums[i], ans);
        }

        // iterNums(minNums);
        // iterNums(maxNums);
        return ans;
    }

    private void iterNums(int[] nums) {
        for (int num : nums) {
            System.out.printf("%d ", num);
        }
        System.out.println();
    }
}

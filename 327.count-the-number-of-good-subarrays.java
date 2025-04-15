class Solution {
    public int countRangeSum(int[] nums, int lower, int upper) {
        long sum = 0, ans = 0;
        long[] sumArr = new long[nums.length];
        for (int i = 0; i < nums.length; i++) {
            sum += nums[i];
            sumArr[i] = sum;
        }
        // this.iterArray(sumArr);
        
        for (int i = 0; i < nums.length; i++) {
            for (int j = i; j < nums.length; j++) {

                sum = this.sum(i, j, sumArr);
                if (sum >= lower && sum <= upper) {
                    // System.out.printf("%d %d\n", i, j);
                    ans ++;
                }
            }
        }
        return (int)ans;
    }

    private long sum(int i, int j, long[] sumArr) {
        if (i > 0) {
            return sumArr[j] - sumArr[i-1];
        }
        return sumArr[j];
    }

    private void iterArray(long[] nums){ 
        for (long i : nums) {
            System.out.printf("%d ", i);
        }
        System.out.println();
    }
}

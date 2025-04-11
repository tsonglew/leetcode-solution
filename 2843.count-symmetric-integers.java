class Solution {
    public int countSymmetricIntegers(int low, int high) {
        int ans = 0;
        for (int i = low; i <= high; i++) {
            if (this.isSymmetric(i)) {
                ans++;
            }
        }
        return ans;
    }


    private boolean isSymmetric(int num) {
        int sum1 = 0;
        int sum2 = 0;
        int num2 = num;
        while (num != 0) {
            for (int i = 0; i < 2; i++) {
                sum1 += num % 10;
                num = (int) Math.floor(num / 10);
                if (i == 0 && num == 0) {
                    return false;
                }
            }
            sum2 += num2 % 10;
            num2 = (int) Math.floor(num2 / 10);
        }
        return sum1 == sum2 * 2;
    }
}

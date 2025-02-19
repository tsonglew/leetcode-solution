class Solution {
    public int[] evenOddBit(int n) {
        int[] res = new int[2];
        boolean isEven = true;
        while (n > 0) {
            if (n % 2 == 1) {
                if (isEven) {
                    res[0]++;
                } else {
                    res[1]++;
                }
            }
            isEven = !isEven;
            n >>= 1;
        }
        return res;
    }
}
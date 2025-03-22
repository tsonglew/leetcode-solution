class Solution {
    public int mySqrt(int x) {
        long ans = 0;
        long i = 0, j = x;
        while (i <= j) {
            long mid = (long) Math.ceil((i + j) / 2);
            long expMid = mid * mid;
            if (i * i > x) {
                break;
            }
            if (expMid == x) {
                return (int)mid;
            } else if (expMid > x) {
                // 0 1 2 3 4 5 6 7 8
                j = mid - 1;
            } else {
                ans = Math.max(mid, ans);
                i = mid + 1;
            }
        }
        return (int)ans;
    }
}

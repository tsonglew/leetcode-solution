class Solution {
    public int shipWithinDays(int[] weights, int days) {
        int total = Arrays.stream(weights).sum();
        int l = 0, r = total, ans = total;
        while (l <= r) {
            int mid = (int) Math.floor((l + r) / 2);
            if (this.valid(weights, days, mid)) {
                ans = Math.min(ans, mid);
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }
        return ans;

    }

    public boolean valid(int[] weights, int days, int capacity) {
        int usedCap = 0;
        int usedDay = 0;
        for (int w : weights) {
            if (w > capacity) {
                return false;
            }
            if (w + usedCap > capacity) {
                usedCap = w;
                usedDay++;
                if (usedDay > days) {
                    return false;
                }
            } else {
                usedCap += w;
            }
        }
        return usedDay + 1 <= days;
    }
}

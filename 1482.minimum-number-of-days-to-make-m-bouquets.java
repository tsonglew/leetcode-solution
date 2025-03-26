class Solution {
    public int minDays(int[] bloomDay, int m, int k) {
        int inf = 1000000001;
        int l = 0, r = 0, ans = inf;
        for (int i = 0; i < bloomDay.length; i++) {
            r = Math.max(r, bloomDay[i]);
        }
        while (l <= r) {
            int mid = (int) Math.floor((l+r)/2);
            if (this.valid(bloomDay, m, k, mid)) {
                ans = Math.min(ans, mid);
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }
        return ans < inf ? ans : -1;
    }

    public boolean valid(int[] bloomDay, int m, int k, int day) {
        int bloomed_sum = 0;
        int baquest_sum = 0;

        for (int i = 0; i < bloomDay.length; i++) {
            if (bloomDay[i] <= day) {
                bloomed_sum ++;
                if (bloomed_sum == k) {
                    baquest_sum ++;
                    bloomed_sum = 0;
                    if (baquest_sum == m) {
                        return true;
                    }
                }
            } else {
                bloomed_sum = 0;
            }
        }
        return false;
    }
}

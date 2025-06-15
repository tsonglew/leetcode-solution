class Solution {
    public int climbStairs(int n) {
        int step1prev = 1, step2prev = 1;
        int now = 1;
        int ans = 1;
        while (now < n) {
            ans = step1prev + step2prev;
            step1prev = step2prev;
            step2prev = ans;

            // System.out.printf("now: %d, s1: %d, s2: %d, ans: %d\n", now, step1prev, step2prev, ans);
            now ++;
        }
        return ans;
    }
}

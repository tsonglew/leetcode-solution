class Solution:
    def nthUglyNumber(self, n: int) -> int:
        dp = [1] + [0] * (n-1)
        p2, p3, p5 = 0, 0, 0
        i = 0
        while i < n-1:
            dp[i+1] = min(dp[p2]*2, dp[p3]*3, dp[p5]*5)
            if dp[i+1] == dp[p2]*2:
                p2 += 1
            if dp[i+1] == dp[p3]*3:
                p3 += 1
            if dp[i+1] == dp[p5]*5:
                p5 += 1
            i += 1
        return dp[-1]

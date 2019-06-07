class Solution:
    dp = [0, 1, 2, 3, 1]
        
    @classmethod
    def numSquares(cls, n: int) -> int:
        import math
        while len(cls.dp) < n+1:
            current = len(cls.dp)
            if pow(int(math.sqrt(current)), 2) == current:
                current_sq = 1
            else:
                i = 1
                current_sq = 5
                while i <= int(math.sqrt(current)):
                    current_sq = min(current_sq, cls.dp[current-i*i]+1)
                    i += 1
            cls.dp.append(current_sq)
        return cls.dp[n]
class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 1:
            return x
        if n == 0:
            return 1
        if n < 0:
            return 1 / self.myPow(x, -n)
        ans = 1
        if n % 2 == 1:
            n -= 1
            ans *= x
        v = self.myPow(x, n // 2)
        return v * v * ans

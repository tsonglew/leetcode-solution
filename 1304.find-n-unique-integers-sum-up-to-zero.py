class Solution:
    def sumZero(self, n: int) -> List[int]:
        for i in range(1, n//2+1):
            yield i
            yield -i
        if n % 2 == 1:
            yield 0

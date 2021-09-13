class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        bits = [0]*32
        for n in nums:
            for i in self.bit(n):
                print(i)
                bits[i] += 1
        for i, _ in enumerate(bits):
            bits[i] %= 3
        i = 0
        r = 0
        while i < 32:
            r += bits[i] * pow(2, i)
            i += 1
        return r

    def bit(self, n):
        r = 0
        while n:
            if n % 2 == 1:
                yield r
            r += 1
            n >>= 1

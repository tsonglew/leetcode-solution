class Solution:
    def digitsNum(self, n):
        cnt = 0
        while n:
            n //= 10
            cnt += 1
        return cnt

    def seqDigits(self, n, min_value, max_value):
        s = []
        for i in range(1, 11-n):
            num = 0
            b = i
            for j in range(n):
                num = 10*num + b
                if num > max_value:
                    return s
                b += 1
            if num < min_value:
                continue
            s.append(num)
        return s
                
    def sequentialDigits(self, low: int, high: int) -> List[int]:
        m, n = self.digitsNum(low), self.digitsNum(high)
        res = []
        for i in range(m, n+1):
            res.extend(self.seqDigits(i, low, high))
        return res

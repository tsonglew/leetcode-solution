class Solution:
    def seqDigits(self, n, min_value, max_value):
        s = []
        for i in range(1, 11-n):
            num, b = 0, i
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
        return reduce(lambda x, y: x+y, map(lambda x: self.seqDigits(x, low, high), range(len(str(low)), len(str(high))+1)))


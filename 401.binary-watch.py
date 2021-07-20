class Solution:
    def __init__(self):
        super().__init__()
        self.min_table = [[] for i in range(4)]
        self.sec_table = [[] for i in range(6)]
        for i in range(12):
            self.min_table[self.cntOnes(i)].append(i)
        for i in range(60):
            self.sec_table[self.cntOnes(i)].append(i)

    @staticmethod
    def cntOnes(num):
        cnt = 0
        while num > 0:
            cnt += num%2
            num >>= 1
        return cnt

    def readBinaryWatch(self, turnedOn: int) -> List[str]:
        res = []
        for i in range(max(0, turnedOn - 5), min(4,turnedOn+1)):
            for m in self.min_table[i]:
                for n in self.sec_table[turnedOn-i]:
                    res.append("%d:%02d" % (m, n))
        return res

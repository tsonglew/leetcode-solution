class Solution:
    def __init__(self):
        super().__init__()
        self.mul = [0] * 59
        self.mul[1] = 1
        self.mul[2] = 1
        for i in range(3, 59):
            for j in range(1, i):
                self.mul[i] = max(self.mul[i], max(self.mul[j], j) * max(i-j, self.mul[i-j]))

    def cuttingRope(self, n: int) -> int:
        return self.mul[n]

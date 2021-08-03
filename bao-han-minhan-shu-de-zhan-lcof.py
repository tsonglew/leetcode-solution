class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.s = []
        self.ms = []
        self.m = None


    def push(self, x: int) -> None:
        if self.m is None or self.m >= x:
            self.m = x
            self.ms.append(x)
        self.s.append(x)

    def pop(self) -> None:
        r = self.s.pop()
        if r == self.m:
            self.ms.pop()
            self.m = self.ms[-1] if len(self.ms) > 0 else None
        return r


    def top(self) -> int:
        return self.s[-1]


    def min(self) -> int:
        return self.m



# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.min()

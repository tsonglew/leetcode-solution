class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.stack = []
        self.minval = None
        

    def push(self, x: int) -> None:
        if self.minval == None:
            self.minval = x
            self.stack.append(None)
            self.stack.append(x)
        else:
            if x <= self.minval:
                self.stack.append(self.minval)
                self.stack.append(x)
                self.minval = x
            else:
                self.stack.append(x)

    def pop(self) -> None:
        if len(self.stack) == 1:
            return
        if self.stack.pop() == self.minval:
            self.minval = self.stack.pop()
        if len(self.stack) == 1:
            self.minval = 0
        

    def top(self) -> int:
        return self.stack[-1]
        

    def getMin(self) -> int:
        return self.minval
        


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
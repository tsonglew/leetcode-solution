class MaxQueue:

    def __init__(self):
        self._max_value = -1
        self.q = []

    def max_value(self) -> int:
        if self._max_value == -1 and len(self.q) > 0:
            self._max_value = heapq.nlargest(1, self.q)[0]
        return self._max_value

    def push_back(self, value: int) -> None:
        self.q.append(value)
        if self._max_value != -1 and value > self._max_value:
            self._max_value = value

    def pop_front(self) -> int:
        if len(self.q) == 0:
            return -1
        r = self.q.pop(0)
        if r == self._max_value:
            self._max_value = -1
        return r

import heapq


class SolutionMonotonicQueue:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        q = []
        ans = []
        for i in range(len(nums)):
            while q and nums[q[-1]] <= nums[i]:
                q.pop()
            q.append(i)
            if i < k - 1:
                continue
            ans.append(nums[q[0]])
            if i + 1 >= q[0] + k:
                q.pop(0)
        return ans


class Item:
    def __init__(self, idx, val):
        self.idx = idx
        self.val = val

    def __lt__(self, other):
        return self.val > other.val


class SolutionMaxHeap:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        h = [Item(i, nums[i]) for i in range(k - 1)]
        ans = []
        heapq.heapify(h)
        for i in range(k - 1, len(nums)):
            heapq.heappush(h, Item(i, nums[i]))
            while h[0].idx <= i - k:
                heapq.heappop(h)
            ans.append(h[0].val)

        return ans

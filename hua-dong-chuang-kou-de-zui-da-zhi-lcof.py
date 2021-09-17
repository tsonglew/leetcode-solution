class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        if len(nums) == 0:
            return []
        window = list(map(lambda x: -x, nums[:k]))
        heapq.heapify(window)
        window = list(map(lambda x: -x, window))
        res = [window[0]]
        i = k
        while i < len(nums):
            idx = window.index(nums[i-k])
            if window[idx] <= nums[i]:
                window[idx] = nums[i]
                while idx > 0:
                    if window[idx//2] >= window[idx]:
                        break
                    window[idx//2], window[idx] = window[idx], window[idx//2]
                    idx //= 2
            else:
                window[idx] = nums[i]
                while idx < len(window):
                    next_idx = idx
                    next_val = window[idx]
                    if idx*2 < len(window) and window[idx*2] > next_val:
                        next_idx = idx*2
                        next_val = window[idx*2]
                    if idx*2+1 < len(window) and window[idx*2+1] > next_val:
                        next_idx = idx*2+1
                    if next_idx == idx:
                        break
                    window[idx], window[next_idx] = window[next_idx], window[idx]
                    idx = next_idx
            res.append(window[0])
            i += 1
        return res


class Solution:
    def trap(self, height: List[int]) -> int:
        left_top = [0] * len(height)
        right_top = [0] * len(height)
        for i in range(len(height)):
            left_top[i] = 0 if i == 0 else max(left_top[i-1], height[i-1])
            ri = len(height) - i - 1
            right_top[ri] = 0 if i == 0 else max(right_top[ri+1], height[ri+1])
        ans = 0
        for i in range(len(height)):
            ans += max(0, min(left_top[i], right_top[i]) - height[i])
        return ans

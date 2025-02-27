class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        preCnt = {0: 1}
        res = 0
        pre = [0] * len(nums)
        for i in range(len(nums)):
            pre[i] = nums[i]

            if i > 0:
                pre[i] += pre[i - 1]
            # print('preCnt: ', preCnt, ' pre[i]', pre[i])
            res += preCnt.get(pre[i] - k, 0)
            preCnt[pre[i]] = preCnt.get(pre[i], 0) + 1
        return res

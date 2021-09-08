class Solution:
    def singleNumbers(self, nums: List[int]) -> List[int]:
        xor = reduce(lambda x, y: x^y, nums)
        lowbit = 0
        n = xor
        while n % 2 == 0:
            n >>= 1
            lowbit += 1
        dispatcher = 1 << lowbit
        l1 = []
        l2 = []
        for i in nums:
            if i | dispatcher == i:
                l1.append(i)
            else:
                l2.append(i)
        return [reduce(lambda x, y: x^y, l1), reduce(lambda x, y: x^y, l2)]

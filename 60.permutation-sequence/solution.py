class Solution:
    def getPermutation(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """
        nums = list(range(1, n+1))
        factorials = [1]
        for i in range(1, n):
            factorials.append(i*factorials[-1])
        cnt = n
        k -= 1
        result = []
        while cnt > 0:
            quotient = k // factorials[cnt-1]
            reminder = k % factorials[cnt-1]
            result.append(str(nums[quotient]))
            k = reminder
            cnt -= 1
            nums.pop(quotient)
        return ''.join(result)
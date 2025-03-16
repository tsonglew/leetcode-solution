class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        if amount == 0:
            return 0
        mem = [-1] * (amount + 1)
        for i in range(len(mem)):
            if i <= 0:
                i = -1
            elif i in coins:
                mem[i] = 1
            else:
                for j in coins:
                    if i - j >= 0 and mem[i - j] > 0:
                        mem[i] = min(mem[i] if mem[i] > 0 else 10001, mem[i - j] + 1)
        # print(mem)
        return mem[-1]

class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        cash = {5: 0, 10: 0, 20: 0}
        for b in bills:
            if b == 5:
                pass
            elif b == 10:
                if cash[5] == 0:
                    return False
                cash[5] -= 1
            else:
                if cash[10] > 0 and cash[5] > 0:
                    cash[10] -= 1
                    cash[5] -= 1
                elif cash[10] > 0:
                    return False
                elif cash[5] < 3:
                    return False
                else:
                    cash[5] -= 3
            cash[b] += 1
        return True

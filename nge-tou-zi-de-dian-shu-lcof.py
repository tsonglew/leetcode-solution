class Solution:
    def dicesProbability(self, n: int) -> List[float]:
        r = [1/6]*6
        for i in range(n-1):
            nr = [0]*(len(r)+5)
            for k in range(len(nr)):
                for j in range(1, 7):
                    if k+2 <= j or k+2-j-1 >= len(r):
                        continue
                    nr[k] += r[k+2-j-1]/6
            r = nr
        return r

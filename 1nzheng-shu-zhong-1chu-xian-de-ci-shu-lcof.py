class Solution:
    def countDigitOne(self, n: int) -> int:
        n_s = str(n)
        res = 0
        for i in range(len(n_s)):
            num = int(n_s[i])
            left, right = self.surround(n_s, i)
            if num == 0:
                res += left*pow(10, len(n_s)-1-i)
            elif num == 1:
                res += left*pow(10, len(n_s)-1-i) + right + 1
            else:
                res += (left+1)*pow(10, len(n_s)-1-i)
        return res
            


    def surround(self, n_s, i):
        l_s = n_s[:i]
        r_s = n_s[i+1:]
        l = 0
        r = 0
        if len(l_s) > 0:
            l = int(l_s)
        if len(r_s) > 0:
            r = int(r_s)
        return l, r

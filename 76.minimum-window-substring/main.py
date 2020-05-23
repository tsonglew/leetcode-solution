class Solution:
    def minWindow(self, s: str, t: str) -> str:
        from collections import defaultdict
        s += "0"
        self.tgtm = defaultdict(lambda: 0)
        self.curm = defaultdict(lambda: 0)
        for i in t:
            self.tgtm[i] += 1
        i, j = 0, 0
        minLen = len(s)+1
        minS = ""
        while j < len(s):
            print(1, i, j, self.curm)
            while self.check() and i < j:
                print(2, i, j, self.curm)
                if j-i < minLen:
                    minLen = j-i
                    minS = s[i:j]
                self.curm[s[i]] -= 1
                i+=1
            self.curm[s[j]] += 1
            j += 1
        return minS
    def check(self):
        for k, v in self.tgtm.items():
            if self.curm[k] < v:
                return False
        return True

print(Solution().minWindow("ADOBECODEBANC", "ABC"))

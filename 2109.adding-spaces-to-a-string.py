class Solution:
    def addSpaces(self, s: str, spaces: List[int]) -> str:
        start = 0
        j = 0
        end = spaces[j]
        ans = ""
        while True:
            ans += s[start:end] + " "
            j += 1
            start = end
            if j >= len(spaces):
                break
            end = spaces[j]
        if len(s) + len(spaces) > len(ans):
            ans += s[start:]
        return ans

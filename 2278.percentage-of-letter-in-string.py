class Solution:
    def percentageLetter(self, s: str, letter: str) -> int:
        return int(sum(1 for i in s if i == letter) * 100 / len(s))
        

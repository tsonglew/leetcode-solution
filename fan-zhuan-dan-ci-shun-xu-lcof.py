class Solution:
    def reverseWords(self, s: str) -> str:
        return " ".join(filter(lambda x: x!="", reversed(s.split(" "))))


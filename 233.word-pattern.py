class Solution(object):
    def wordPattern(self, pattern, s):
        """
        :type pattern: str
        :type s: str
        :rtype: bool
        """
        chars = list(pattern)
        words = s.split(" ")
        if len(chars) != len(words):
            return False
        cwMap = {}
        for i, c in enumerate(chars):
            if c not in cwMap:
                if words[i] in cwMap.values():
                    return False
                cwMap[c] = words[i]
            if cwMap[c] != words[i]:
                return False
        return True

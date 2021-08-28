class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) == 0:
            return 0
        char_set = {s[0]: 0}
        res = 1
        start = 0
        for i in range(1, len(s)):
            if s[i] in char_set and char_set[s[i]] >= start:
                start = char_set[s[i]] + 1
            char_set[s[i]] = i
            cur_len = i - start + 1
            res = max(res, cur_len)
        return res

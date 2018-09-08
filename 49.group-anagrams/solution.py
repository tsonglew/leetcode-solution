#!/usr/bin/python3

class Solution:
    def groupAnagrams(self, strs):
        """
        :type strs: List[str]
        :rtype: List[List[str]]
        """
        m = dict()
        for str in strs:
            sorted_str = ''.join(sorted(str))
            if sorted_str not in m.keys():
                m[sorted_str] = list()
            m[sorted_str].append(str)
        return list(m.values())
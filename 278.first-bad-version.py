# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):

class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        return self.find(1, n-1)
    
    def find(self, lo, hi):
        mid = (lo+hi)//2
        if not isBadVersion(mid) and isBadVersion(mid+1):
            return mid+1
        if isBadVersion(mid):
            return self.find(lo, mid-1)
        return self.find(mid+1, hi)

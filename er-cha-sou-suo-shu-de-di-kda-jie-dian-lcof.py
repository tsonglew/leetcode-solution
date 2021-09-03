# https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/submissions/
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def kthLargest(self, root: TreeNode, k: int) -> int:
        p = root
        cache = []
        while p:
            cache.append(p)
            p = p.right
        while len(cache):
            n = cache.pop()
            k -= 1
            if k == 0:
                return n.val
            p = n.left
            while p:
                cache.append(p)
                p = p.right

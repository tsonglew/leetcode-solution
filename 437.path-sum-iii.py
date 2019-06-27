# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
import copy

class Solution:
    def pathSum(self, root: TreeNode, sum: int) -> int:
        self.cnt = 0
        def canSum(root, num_set, total):
            if root is None:
                return
            set_copy = []
            for n in num_set:
                if n == root.val:
                    self.cnt += 1
                set_copy.append(n - root.val)
            set_copy.append(total)
            canSum(root.left, set_copy, total)
            canSum(root.right, set_copy, total)
        canSum(root, [sum], sum)
        return self.cnt
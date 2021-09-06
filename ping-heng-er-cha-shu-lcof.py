# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def isBalanced(self, root: TreeNode) -> bool:
        _, r = self.f(root)
        return r
    
    def f(self, root):
        if root is None:
            return 0, True
        left_depth, left_balanced = self.f(root.left)
        right_depth, right_balanced = self.f(root.right)
        if not all([left_balanced, right_balanced]):
            return 0, False
        return max(left_depth, right_depth)+1, abs(left_depth-right_depth) <= 1

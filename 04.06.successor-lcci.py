# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def inorderSuccessor(self, root: TreeNode, p: TreeNode) -> TreeNode:
        self.min = None
        return self.f(root, p)

    def f(self, root: TreeNode, p: TreeNode) -> TreeNode:
        if root is None:
            return None
        if root is p:
            if p.right is not None:
                return self.mostLeft(p.right)
            return self.min
        if (self.min is None or root.val < self.min.val) and root.val >= p.val:
            self.min = root
        if root.val > p.val:
            return self.f(root.left, p)
        return self.f(root.right, p)

    def mostLeft(self, root: TreeNode) -> TreeNode:
        if root.left is None:
            return root
        return self.mostLeft(root.left)
        

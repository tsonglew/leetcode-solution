# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def sumNumbers(self, root: TreeNode) -> int:
        if root == None:
            return 0
        return self.sumFromRoot(root, 0)
        
        
    def sumFromRoot(self, root, parent):
        sm = 0
        if root.left == None and root.right == None:
            return root.val + parent
        if root.left != None:
            sm += self.sumFromRoot(root.left, (parent+root.val) * 10)
        if root.right != None:
            sm += self.sumFromRoot(root.right, (parent+root.val) * 10)
        return sm
        
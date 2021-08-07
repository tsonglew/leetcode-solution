# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[int]:
        stack = [root]
        res = []
        while len(stack):
            l = []
            while len(stack):
                n = stack.pop(0)
                if n is None:
                    continue
                res.append(n.val)
                l.append(n.left)
                l.append(n.right)
            stack = l
        return res
        

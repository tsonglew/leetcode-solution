# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrderBottom(self, root: TreeNode) -> List[List[int]]:
        self.result = []
        self.wft([root])
        return self.result
    
    def wft(self, nodes):
        newNodes = []
        values = []
        for n in nodes:
            if n is not None:
                values.append(n.val)
                newNodes += [n.left, n.right]
        if len(values) != 0:
            self.result = [values] + self.result
            self.wft(newNodes)
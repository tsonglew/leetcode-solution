class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def generateTrees(self, n: int) -> List[TreeNode]:
        if n<1:
            return []
        return self.buildTree(1, n+1)
        
    def buildTree(self, l, r):
        if l >= r:
            return [None]
        result = []
        for i in range(l, r):
            for ln in self.buildTree(l, i):
                for rn in self.buildTree(i+1, r):
                    root_node = TreeNode(i)
                    root_node.left = ln
                    root_node.right = rn
                    result.append(root_node)
        return result
"""
# Definition for a Node.
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
"""
class Solution:
    def treeToDoublyList(self, root: 'Node') -> 'Node':
        if root is None:
            return root
        ll, rr = self.tree(root)
        ll.left = rr
        rr.right = ll
        return ll


    def tree(self, root):
        if root is None:
            return None, None
        ll, lr = self.tree(root.left)
        rl, rr = self.tree(root.right)

        ll = ll or root
        lr = lr or root
        rl = rl or root
        rr = rr or root

        rl.left = root
        root.left = lr
        lr.right = root
        root.right = rl
        return ll, rr

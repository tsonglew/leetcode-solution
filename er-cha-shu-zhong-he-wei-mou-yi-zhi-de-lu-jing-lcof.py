# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pathSum(self, root: TreeNode, target: int) -> List[List[int]]:
        res = []
        self.sum(root, target, 0, [], res)
        return res

    def sum(self, root, target, sm, buf, res):
        if root is None:
            return
        sm += root.val
        buf.append(root.val)
        if sm == target and root.left is None and root.right is None:
            res.append(copy.copy(buf))
        self.sum(root.left, target, sm, buf, res)
        self.sum(root.right, target, sm, buf, res)
        buf.pop()

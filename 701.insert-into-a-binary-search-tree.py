# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def insertIntoBST(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        node = TreeNode(val)
        if not root:
            return node
        parent = None
        ptr = root
        while ptr != None:
            parent = ptr
            if ptr.val > node.val:
                ptr = ptr.left
            else:
                ptr = ptr.right
        if parent.val > node.val:
            parent.left = node
        else:
            parent.right = node
        return root

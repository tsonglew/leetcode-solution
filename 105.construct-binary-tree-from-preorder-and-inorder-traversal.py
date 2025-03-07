# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if not preorder:
            return None
        root = TreeNode(preorder[0])
        rootInOrderIdx = inorder.index(root.val)
        leftInOrder = inorder[:rootInOrderIdx]
        rightInOrder = inorder[rootInOrderIdx + 1 :]
        leftPreOrder = preorder[1 : len(leftInOrder) + 1]
        rightPreOrder = preorder[len(leftInOrder) + 1 :]
        root.left = self.buildTree(leftPreOrder, leftInOrder)
        root.right = self.buildTree(rightPreOrder, rightInOrder)
        return root

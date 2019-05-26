# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
        
    def recoverTree(self, root: TreeNode) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        nums = []
        inorder_travel(root, nums)
        inorder_set(root, sorted(nums))
        
        
def inorder_travel(root: TreeNode, nums) -> None:
    if not root:
        return
    inorder_travel(root.left, nums)
    nums.append(root.val)
    inorder_travel(root.right, nums)
        
def inorder_set(root: TreeNode, nums) -> None:
    if not root:
        return
    inorder_set(root.left, nums)
    root.val = nums.pop(0)
    inorder_set(root.right, nums)
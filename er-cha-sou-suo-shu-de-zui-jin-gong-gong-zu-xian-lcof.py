# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        _, _, res = self.f(root, p, q)
        return res

    def f(self, root, p, q):
        left_contain_p, left_contain_q, left_parent = False, False, None
        right_contain_p, right_contain_q, right_parent = False, False, None
        if root.left is not None:
            left_contain_p, left_contain_q, left_parent = self.f(root.left, p, q)
        if left_parent:
            return True, True, left_parent
        if root.right is not None:
            right_contain_p, right_contain_q, right_parent = self.f(root.right, p, q)
        if right_parent:
            return True, True, right_parent
        self_contain_p = root.val == p.val or left_contain_p or right_contain_p
        self_contain_q = root.val == q.val or left_contain_q or right_contain_q
        return self_contain_p, self_contain_q, root if self_contain_p and self_contain_q else None
        

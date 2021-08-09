# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        l = [root]
        seq = True
        res = []
        while len(l):
            l_buf = []
            res_buf = []
            for i in l:
                if i is None:
                    continue
                res_buf.append(i.val)
                l_buf.append(i.left)
                l_buf.append(i.right)
            if len(res_buf):
                res.append(res_buf[::1 if seq else -1])
            l = l_buf
            seq = not seq
        return res

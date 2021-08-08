# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        l = [root]
        res = []
        while len(l):
            l_buf = []
            res_buf = []
            for i in l:
                if i:
                    l_buf.append(i.left)
                    l_buf.append(i.right)
                    res_buf.append(i.val)
            if len(res_buf):
                res.append(res_buf)
            l = l_buf
        return res


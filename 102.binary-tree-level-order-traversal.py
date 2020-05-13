# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        q1 = [root]
        q2 = []
        result = []
        while len(q1) > 0:
            cur = []
            for t in q1:
                if t is None:
                    continue
                q2.append(t.left)
                q2.append(t.right)
                cur.append(t.val)
            if len(cur) > 0:
                result.append(cur)
            q1 = q2
            q2 = []
        return result

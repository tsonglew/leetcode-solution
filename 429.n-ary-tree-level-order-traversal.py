"""
# Definition for a Node.
class Node:
    def __init__(self, val: Optional[int] = None, children: Optional[List['Node']] = None):
        self.val = val
        self.children = children
"""


class Solution:
    def levelOrder(self, root: "Node") -> List[List[int]]:
        q = [root] if root else []
        ans = []
        while q:
            next_level = []
            level = []
            while q:
                e = q.pop(0)
                for c in e.children:
                    next_level.append(c)
                level.append(e.val)
            q = next_level
            ans.append(level)

        return ans

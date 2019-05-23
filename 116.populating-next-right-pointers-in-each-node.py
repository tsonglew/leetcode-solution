"""
# Definition for a Node.
class Node:
    def __init__(self, val, left, right, next):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""
class Solution:
    def connect(self, root: 'Node') -> 'Node':
        upperLeftMost = root
        while upperLeftMost is not None:
            p = upperLeftMost
            nextLeftMost = p.left
            if nextLeftMost is None:
                break
            while p is not None:
                p.left.next = p.right
                if p.next is not None:
                    p.right.next = p.next.left
                p = p.next
            upperLeftMost = nextLeftMost
        return root
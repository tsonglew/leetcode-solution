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
        def add_children(queue, node):
            if node.left != None:
                queue.append(node.left)
            if node.right != None:
                queue.append(node.right)
        if root == None:
            return root
        queue = [root]
        while len(queue) != 0:
            new_queue = list()
            prev = queue.pop(0)
            add_children(new_queue, prev)
            while len(queue) != 0:
                new_node = queue.pop(0)
                prev.next = new_node
                add_children(new_queue, new_node)
                prev = new_node
            queue = new_queue
        
        return root
            
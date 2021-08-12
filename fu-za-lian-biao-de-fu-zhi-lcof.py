"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""
class Solution:
    def copyRandomList(self, head: 'Node') -> 'Node':
        new_head = Node(0)

        origin_p = head
        new_p = new_head

        while origin_p:
            origin_p_next = origin_p.next
            new_p.next = Node(origin_p.val)

            new_p.next.random = origin_p.random
            origin_p.next = new_p.next

            new_p = new_p.next
            origin_p = origin_p_next
            
        new_p = new_head.next
        while new_p:
            if new_p.random:
                new_p.random = new_p.random.next
            new_p = new_p.next
        return new_head.next

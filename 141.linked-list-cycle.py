# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        p1, p2 = head, head
        while p1 and p2 and p2.next:
            p1, p2 = p1.next, p2.next.next
            if p1 == p2:
                return True
        return False

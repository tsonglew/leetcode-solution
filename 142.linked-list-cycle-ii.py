# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None or head.next.next is None:
            return None
        slow = head.next
        fast = head.next.next
        while slow is not None and fast is not None and fast.next is not None:
            if slow == fast:
                s1 = slow
                h1 = head
                while s1 != h1:
                    s1 = s1.next
                    h1 = h1.next
                return s1
            slow = slow.next
            fast = fast.next.next
        return None
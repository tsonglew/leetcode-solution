# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        if not all([headA, headB]):
            return None
        pa, pb, a_end, b_end = headA, headB, False, False
        while 1:
            if (pa is None and a_end) or (pb is None and b_end):
                return None
            if pa is None:
                pa = headB
                a_end = True
            if pb is None:
                pb = headA
                b_end = True
            if pa is pb:
                return pa
            pa = pa.next
            pb = pb.next

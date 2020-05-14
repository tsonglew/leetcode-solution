# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        cf = 0
        result = h = ListNode(0)
        while l1 is not None or l2 is not None:
            l1n = l2n = 0
            if l1 is not None:
                l1n = l1.val
                l1 = l1.next
            if l2 is not None:
                l2n = l2.val
                l2 = l2.next
            sm = l1n + l2n + cf
            h.next = ListNode(sm % 10)
            cf = sm // 10
            h = h.next
        if cf > 0:
            h.next = ListNode(cf)
        return result.next

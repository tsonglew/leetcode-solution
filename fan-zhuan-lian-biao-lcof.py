# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        p1, p2 = None, head
        while p2:
            p3 = p2.next
            p2.next, p1, p2 = p1, p2, p3
        return p1

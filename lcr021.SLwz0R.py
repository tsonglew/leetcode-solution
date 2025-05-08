# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        vn = ListNode()
        vn.next = head

        n += 1
        p1 = p2 = vn
        while p2:
            if n == 0:
                p1 = p1.next
            else:
                n -= 1

            p2 = p2.next
        if p1.next:
            p1.next = p1.next.next
        return vn.next

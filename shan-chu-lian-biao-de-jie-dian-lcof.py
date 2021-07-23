class Solution:
    def deleteNode(self, head: ListNode, val: int) -> ListNode:
        fakeHead = ListNode(0)
        fakeHead.next = head
        p = fakeHead
        while p.next is not None and p.next.val != val:
            p = p.next
        if p.next is not None:
            p.next = p.next.next
        return fakeHead.next

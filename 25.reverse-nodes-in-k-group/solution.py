# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        new_head = None
        new_end = None

        num = k-1
        h = head
        e = head

        while h:
            while num and e:
                e = e.next
                num -= 1
            if num != 0:
                new_end.next = h
                return new_head
            if not e:
                new_end.next = h
                return new_head

            next_head = e.next
            e.next = None
            cur_head, cur_end = self.reverseGroup(h)
            if new_head is None:
                new_head = cur_head
            else:
                new_end.next = cur_head
            new_end = cur_end
            h = next_head
            e = next_head
            num = k-1

        return new_head
        
    def reverseGroup(self, head: Optional[ListNode]) -> Tuple[Optional[ListNode], Optional[ListNode]]:
        pre = None
        cur = head
        new_end = None
        while cur:
            if new_end is None:
                new_end = cur
            n = cur.next
            cur.next = pre
            pre = cur
            cur = n
        return pre, new_end


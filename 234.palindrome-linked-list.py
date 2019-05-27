# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        if head is None or head.next is None:
            return True
        slow_pointer = fast_pointer = head
        even = False
        while fast_pointer is not None:
            if fast_pointer.next is None or fast_pointer.next.next is None:
                break
            slow_pointer = slow_pointer.next
            fast_pointer = fast_pointer.next.next
        if fast_pointer.next is not None:
            even = True
        second_half_head = slow_pointer.next
        first_half_head = slow_pointer
        first_half_head.next = None
        self.reverseLinkedList(head)
        if not even:
            first_half_head = first_half_head.next
        p1 = first_half_head
        p2 = second_half_head
        while p1 is not None and p2 is not None:
            if p1.val != p2.val:
                return False
            p1 = p1.next
            p2 = p2.next
        if p1 is not None or p2 is not None:
            return False
        return True
            
    def reverseLinkedList(self, head: ListNode) -> None:
        pre_head = None
        while head is not None:
            head_next = head.next
            head.next = pre_head
            pre_head = head
            head = head_next
        
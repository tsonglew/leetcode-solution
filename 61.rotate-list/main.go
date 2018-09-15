package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	headHolder := &ListNode{
		Val:  0,
		Next: head,
	}
	length := 0
	h := headHolder
	for h.Next != nil {
		length++
		h = h.Next
	}
	tail := h
	rotateTimes := k % length
	cnt := length - rotateTimes
	h = headHolder
	for cnt > 0 {
		h = h.Next
		cnt--
	}
	tail.Next = headHolder.Next
	head = h.Next
	h.Next = nil
	return head
}

func main() {}

package main

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	cnt := 1
	p := head
	if p == nil {
		return
	}
	for p.Next != nil {
		cnt++
		p = p.Next
	}
	if cnt < 3 {
		return
	}
	last := p
	p = head
	for i := 0; i < (cnt+1)/2-1; i++ {
		p = p.Next
	}
	next := p.Next
	p.Next = nil
	p = next
	var pre *ListNode = nil
	for p != nil {
		next = p.Next
		p.Next = pre
		pre = p
		p = next
	}
	p1, p2 := head, last
	for p1 != nil {
		var p2Next *ListNode
		p1Next := p1.Next
		p1.Next = p2
		if p2 != nil {
			p2Next = p2.Next
			p2.Next = p1Next
		}
		p1, p2 = p1Next, p2Next
	}
}

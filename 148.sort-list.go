package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	preHf := head
	nexHf := splitList(head)
	st1 := sortList(preHf)
	st2 := sortList(nexHf)
	r := mergeList(st1, st2)
	return r
}

func splitList(head *ListNode) *ListNode {
	slowHead := head
	fastHead := head
	slowPre := slowHead
	for fastHead != nil && fastHead.Next != nil {
		slowPre = slowHead
		slowHead = slowHead.Next
		fastHead = fastHead.Next.Next
	}
	slowPre.Next = nil
	return slowHead
}

func mergeList(h1, h2 *ListNode) *ListNode {
	if h1 == h2 {
		return h1
	}
	virtualHead := &ListNode{}
	p := virtualHead
	for h1 != nil || h2 != nil {
		var tmp *ListNode
		if h1 == nil {
			p.Next = h2
			break
		}
		if h2 == nil {
			p.Next = h1
			break
		}
		if h1.Val < h2.Val {
			tmp = h1
			h1 = h1.Next
		} else {
			tmp = h2
			h2 = h2.Next
		}
		tmp.Next = nil
		p.Next = tmp
		p = p.Next
	}
	return virtualHead.Next
}

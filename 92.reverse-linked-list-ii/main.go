package main

import "fmt"

// ListNode is node of the list
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	p := reverseBetween(head, 2, 4)
	cnt := 10
	for p != nil && cnt > 0 {
		fmt.Println(p.Val)
		p = p.Next
		cnt--
	}
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	newHead := &ListNode{Next: head}
	p := newHead
	var pre, a, apre *ListNode
	for i := 0; p != nil; i++ {
		pNext := p.Next
		if i == m {
			a = p
			apre = pre
		}
		if i > m && i <= n {
			p.Next = pre
		}
		if i == n {
			apre.Next = p
			a.Next = pNext
			break
		}
		pre = p
		p = pNext
	}
	return newHead.Next
}

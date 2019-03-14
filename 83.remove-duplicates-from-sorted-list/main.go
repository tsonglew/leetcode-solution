package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1}
	initList(head, []int{2, 3, 3, 4, 4, 5})
	printList(head)
	newHead := deleteDuplicates(head)
	printList(newHead)
}

func initList(head *ListNode, nums []int) {
	for i := range nums {
		head.Next = &ListNode{Val: nums[i]}
		head = head.Next
	}
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	p := newHead
	for p != nil {
		if p.Next != nil && p.Next.Val == p.Val {
			for p.Next != nil && p.Next.Val == p.Val {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}
	if head != nil && head.Val == 0 {
		return newHead
	}
	return newHead.Next
}

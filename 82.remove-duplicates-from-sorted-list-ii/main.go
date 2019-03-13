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
	head = newHead
	for head != nil {
		if head.Next != nil && head.Next.Next != nil {
			if head.Next.Val == head.Next.Next.Val {
				currentv := head.Next.Val
				for head.Next != nil && head.Next.Val == currentv {
					head.Next = head.Next.Next
				}
			} else {
				head = head.Next
			}
		} else {
			head = head.Next
		}
	}
	return p.Next
}

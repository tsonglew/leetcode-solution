package main

import "fmt"

func main() {
	head := &ListNode{-1, &ListNode{5, &ListNode{4, &ListNode{3, nil}}}}
	result := insertionSortList(head)
	printHead(result)
}

func printHead(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fakeHead := &ListNode{Val: 0, Next: head}
	currentSorting := head.Next
	preCurrentSorting := head
	for currentSorting != nil {
		nextSorting := currentSorting.Next
		preP := fakeHead
		p := preP.Next
		for p != currentSorting && p.Val <= currentSorting.Val {
			preP = p
			p = p.Next
		}
		if p != currentSorting {
			preCurrentSorting.Next = currentSorting.Next
			preP.Next = currentSorting
			currentSorting.Next = p
		} else {
			preCurrentSorting = currentSorting
		}
		currentSorting = nextSorting
	}
	return fakeHead.Next
}

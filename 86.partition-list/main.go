package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func partition(head *ListNode, x int) *ListNode {
	leNos, geNos := &ListNode{}, &ListNode{}
	lep, gep := leNos, geNos
	for head != nil {
		headNext := head.Next
		if head.Val < x {
			lep.Next = head
			lep = lep.Next
			lep.Next = nil
		} else {
			gep.Next = head
			gep = gep.Next
			gep.Next = nil
		}
		head = headNext
	}
	lep.Next = geNos.Next
	return leNos.Next
}

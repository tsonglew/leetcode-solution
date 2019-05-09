/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	nodeAddrs := make(map[*ListNode]bool)
	p := head
	for p != nil {
		if _, ok := nodeAddrs[p]; ok {
			return p
		}
		nodeAddrs[p] = true
		p = p.Next
	}
	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	// Using Floyd Detection
	if head == nil {
		return nil
	}
	p1, p2 := head, head
	if p1.Next == nil || p2.Next == nil || p2.Next.Next == nil {
		return nil
	}
	p1 = p1.Next
	p2 = p2.Next.Next
	for *p1 != *p2 {
		if p1.Next == nil || p2.Next == nil || p2.Next.Next == nil {
			return nil
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	p1 = head
	for *p1 != *p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

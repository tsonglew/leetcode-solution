/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	p := head
	var pp, pn *ListNode
	for p != nil {
		pn = p.Next
		p.Next = pp
		pp = p
		p = pn
	}
	return pp
}
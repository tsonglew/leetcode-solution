/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	sp, fp := head, head
	for sp != nil && fp != nil {
		if fp.Next != nil {
			fp = fp.Next.Next
		} else {
			return false
		}
		sp = sp.Next
		if fp == sp {
			return true
		}
	}
	return false
}
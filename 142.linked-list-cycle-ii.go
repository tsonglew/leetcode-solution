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
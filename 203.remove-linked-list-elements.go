/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeElements(head *ListNode, val int) *ListNode {
    fakeHead := &ListNode{0, head}
    for p1, p2 := fakeHead, head; p2 != nil; {
        if p2.Val == val {
            p1.Next = p2.Next
            p2 = p2.Next
        } else {
            p1, p2 = p2, p2.Next 
        }
    }
    return fakeHead.Next
}
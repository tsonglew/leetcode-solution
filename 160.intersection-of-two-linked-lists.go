/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ppa := headA
	pa := ppa.Next
	ppb := headB
	pb := ppb.Next
	for pa != nil && pb != nil {
		ppa = pa
		pa = pa.Next
		ppb = pb
		pb = pb.Next
	}
	lenMinusA := 0
	lenMinusB := 0
	for pa != nil {
		ppa = pa
		pa = pa.Next
		lenMinusA++
	}
	for pb != nil {
		ppb = pb
		pb = pb.Next
		lenMinusB++
	}
	pa = headA
	pb = headB
	for lenMinusA > 0 {
		pa = pa.Next
		lenMinusA--
	}
	for lenMinusB > 0 {
		pb = pb.Next
		lenMinusB--
	}
	var dis *ListNode
	for pa != nil && pb != nil {
		if pa == pb {
			if dis == nil {
				dis = pa
			}
		} else {
			dis = nil
		}
		pa = pa.Next
		pb = pb.Next
	}
	return dis
}

// Without knowing the difference of lens
func getIntersectionNodeWIthoutLens(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa := headA
	pb := headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
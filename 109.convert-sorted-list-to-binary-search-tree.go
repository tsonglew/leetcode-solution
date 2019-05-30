/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func sortedListToBST(head *ListNode) *TreeNode {
    nums := make([]int, 0)
    for head != nil {
        nums = append(nums, head.Val)
        head = head.Next
    }
    return sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return &TreeNode{
        Val: nums[len(nums)/2],
        Left: sortedArrayToBST(nums[0:len(nums)/2]),
        Right: sortedArrayToBST(nums[len(nums)/2+1:len(nums)]),
    }
}
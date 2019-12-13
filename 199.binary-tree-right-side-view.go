/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := rightSideView(root.Left)
	right := rightSideView(root.Right)
	var toAppend []int
	if len(left) <= len(right) {
		toAppend = right
	} else {
	 	toAppend = append(right, left[len(right):]...)
	}
	return append([]int{root.Val}, toAppend...)
}
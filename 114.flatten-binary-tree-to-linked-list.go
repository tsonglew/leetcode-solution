/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		p := root.Left
		for p.Right != nil {
			p = p.Right
		}
		p.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	flatten(root.Right)
}
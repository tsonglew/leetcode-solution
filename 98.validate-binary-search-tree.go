/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValidBSTSub(root, -1, -1)
}

func isValidBSTSub(root *TreeNode, lower, upper int) bool {
	if lower != -1 && root.Val <= lower {
		return false
	}
	if upper != -1 && root.Val >= upper {
		return false
	}
	lValid, rValid := true, true
	if root.Left != nil {
		lValid = isValidBSTSub(root.Left, lower, root.Val)
	}
	if root.Right != nil {
		rValid = isValidBSTSub(root.Right, root.Val, upper)
	}
	return lValid && rValid
}
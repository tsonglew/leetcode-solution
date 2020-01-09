/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	path(root, "", &result)
	return result
}

func path(root *TreeNode, current string, result *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*result = append(*result, fmt.Sprintf("%s%d", current, root.Val))
	}
	path(root.Left, fmt.Sprintf("%s%d->", current, root.Val), result)
	path(root.Right, fmt.Sprintf("%s%d->", current, root.Val), result)
}

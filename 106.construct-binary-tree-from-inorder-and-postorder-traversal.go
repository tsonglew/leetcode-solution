/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	rootIndexOfInorder := getIndex(inorder, rootVal)
	leftLen := rootIndexOfInorder
	root.Left = buildTree(inorder[0:leftLen], postorder[0:leftLen])
	root.Right = buildTree(inorder[rootIndexOfInorder+1:len(inorder)], postorder[leftLen:len(postorder)-1])
	return root
}

func getIndex(s []int, v int) int {
	for i := range s {
		if s[i] == v {
			return i
		}
	}
	return -1
}
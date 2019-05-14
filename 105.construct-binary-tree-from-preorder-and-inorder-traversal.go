/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	currentNode := &TreeNode{Val: preorder[0]}
	currentNodeIndex := 0
	for i := range inorder {
		if inorder[i] == currentNode.Val {
			currentNodeIndex = i
			break
		}
	}

	leftPreorder, rightPreorder := []int{}, []int{}

	leftPreorder = preorder[1 : currentNodeIndex+1]
	rightPreorder = preorder[currentNodeIndex+1 : len(preorder)]

	leftInorder := inorder[0:currentNodeIndex]
	rightInorder := inorder[currentNodeIndex+1 : len(inorder)]

	currentNode.Left = buildTree(leftPreorder, leftInorder)
	currentNode.Right = buildTree(rightPreorder, rightInorder)
	return currentNode
}
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

	currentNode.Left = buildTree(preorder[1:currentNodeIndex+1], inorder[0:currentNodeIndex])
	currentNode.Right = buildTree(preorder[currentNodeIndex+1:len(preorder)], inorder[currentNodeIndex+1:len(inorder)])
	return currentNode
}
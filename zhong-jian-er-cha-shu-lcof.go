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
    pi, ii := 0, 0
    for preorder[pi] != inorder[ii] {
        ii++
    }
    return &TreeNode{
        Val: preorder[0],
        Left: buildTree(preorder[1:1+ii], inorder[0:ii]),
        Right: buildTree(preorder[1+ii:], inorder[ii+1:]),
    }
}

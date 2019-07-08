/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func diameterOfBinaryTree(root *TreeNode) int {
    max := 0 
    diameterSub(root, &max)
    return max
}

func diameterSub(root *TreeNode, max *int) int {
    if root == nil {
        return 0
    }
    left := diameterSub(root.Left, max)
    right := diameterSub(root.Right, max)
    if left + right > *max {
        *max = left+right
    }
    if left > right {
        return left + 1
    }
    return right + 1
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    ld, rd := minDepth(root.Left), minDepth(root.Right)
    if ld == 0 {
        return rd + 1
    }
    if rd == 0 {
        return ld + 1
    }
    return min(rd, ld) + 1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
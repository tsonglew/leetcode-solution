/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    lHeight, lBalanced := heightOf(root.Left)
    rHeight, rBalanced := heightOf(root.Right)
    if lBalanced && rBalanced && abs(lHeight-rHeight) <= 1 {
        return true
    }
    return false
}

func heightOf(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }
    lHeight, lBalanced := heightOf(root.Left)
    rHeight, rBalanced := heightOf(root.Right)
    if lBalanced && rBalanced && abs(lHeight-rHeight) <= 1 {
        return max(lHeight, rHeight) + 1, true
    }
    return 0, false
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
                       
func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}
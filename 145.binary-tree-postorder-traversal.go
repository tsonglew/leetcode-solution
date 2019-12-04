/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func postorderTraversal(root *TreeNode) []int {
    result := []int{}
    travel(root, &result)
    return result
}

func travel(root *TreeNode, result *[]int) {
    if root == nil {
        return
    }
    travel(root.Left, result)
    travel(root.Right, result)
    *result = append(*result, root.Val)
}
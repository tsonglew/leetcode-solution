/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    _, kth := smallest(root, k)
    return kth
}

func smallest(root *TreeNode, k int) (int, int) {
    if root == nil {
        return 0, -1
    }
    leftCnt, kth := smallest(root.Left, k)
    if kth != -1 {
        return 0, kth
    }
    if leftCnt == k-1 {
        return 0, root.Val
    }
    rightCnt, kth := smallest(root.Right, k-1-leftCnt)
    if kth != -1 {
        return 0, kth
    }
    return leftCnt+rightCnt+1, -1
}

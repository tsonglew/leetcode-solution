/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    ancestor := make(chan *TreeNode, 1)
    go findAncestor(root, p, q, ancestor)
    return <- ancestor
}

func findAncestor(root, p, q *TreeNode, ancestor chan *TreeNode) {
    if root == nil {
        return
    }
    findAncestor(root.Left, p, q, ancestor)
    findAncestor(root.Right, p, q, ancestor)
    if isAncestor(root.Left, p) && isAncestor(root.Left, q) {
        ancestor <- root.Left
    }
    if isAncestor(root.Right, p) && isAncestor(root.Right, q) {
        ancestor <- root.Right
    }
    if isAncestor(root, p) && isAncestor(root, q) {
        ancestor <- root
    }
}

func isAncestor(root, p *TreeNode) bool {
    if root == nil {
        return false
    }
    if root.Val == p.Val {
        return true
    }
    return isAncestor(root.Left, p) || isAncestor(root.Right, p)
}
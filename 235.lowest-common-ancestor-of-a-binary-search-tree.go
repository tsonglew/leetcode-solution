func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, result := hasLeaf(root, p, q)
	return result
}

func hasLeaf(root, p, q *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	lc, left := hasLeaf(root.Left, p, q)
	if lc == 2 {
		return 2, left
	}
	rc, right := hasLeaf(root.Right, p, q)
	if rc == 2 {
		return 2, right
	}
	cnt := lc + rc
	if root.Val == p.Val || root.Val == q.Val {
		cnt++
	}
	if cnt == 2 {
		return 2, root
	}

	return cnt, nil
}

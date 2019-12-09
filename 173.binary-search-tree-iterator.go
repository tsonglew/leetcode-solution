type BSTIterator struct {
    Root *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    return BSTIterator{root}
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	var prev, p *TreeNode
	for p = this.Root; p.Left != nil ; p = p.Left {
		prev = p
	}
	if prev != nil {
		prev.Left = p.Right
	} else {
		this.Root = p.Right
	}
	return p.Val
}
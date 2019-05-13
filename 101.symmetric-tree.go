package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	nodes := []*TreeNode{root}
	newNodes := []*TreeNode{}
	for len(nodes) != 0 {
		for i := range nodes {
			if nodes[i] != nil {
				newNodes = append(newNodes, nodes[i].Left, nodes[i].Right)
			}
		}
		for i, j := 0, len(newNodes)-1; i < j; {
			if newNodes[i] == nil || newNodes[j] == nil {
				if newNodes[i] != newNodes[j] {
					return false
				}

			} else {
				if newNodes[i].Val != newNodes[j].Val {
					return false
				}
			}
			i++
			j--
		}
		nodes = newNodes
		newNodes = []*TreeNode{}

	}
	return true
}

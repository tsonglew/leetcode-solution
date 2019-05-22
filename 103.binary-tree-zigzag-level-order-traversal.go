/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	travel([]*TreeNode{root}, &result, false)
	return result
}

func travel(nodes []*TreeNode, result *[][]int, forward bool) {
	if len(nodes) == 0 {
		return
	}
	newNodes := []*TreeNode{}
	nodeVals := []int{}
	for i := len(nodes) - 1; i >= 0; i-- {
		if nodes[i] != nil {
			nodeVals = append(nodeVals, nodes[i].Val)
			if forward {
				newNodes = append(newNodes, nodes[i].Right, nodes[i].Left)
			} else {
				newNodes = append(newNodes, nodes[i].Left, nodes[i].Right)
			}
		}
	}
	if len(nodeVals) == 0 {
		return
	}
	*result = append(*result, append([]int{}, nodeVals...))
	travel(newNodes, result, !forward)
}
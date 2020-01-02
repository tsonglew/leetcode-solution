package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func inorderTraversal(root *TreeNode) []int {
	result := &[]int{}

	if root != nil {
		if root.Left != nil {
			travelSub(root.Left, result)
		}
		(*result) = append(*result, root.Val)
		if root.Right != nil {
			travelSub(root.Right, result)
		}
	}
	return *result
}

func travelSub(root *TreeNode, result *[]int) {
	if root.Left != nil {
		travelSub(root.Left, result)
	}
	(*result) = append(*result, root.Val)
	if root.Right != nil {
		travelSub(root.Right, result)
	}
}

// Morris Traversal
func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    cur := root
	for cur != nil {
		if cur.Left == nil {
			result = append(result, cur.Val)
			cur = cur.Right
		} else {
			p := cur.Left
			for p.Right != nil {
				p = p.Right
			}
			p.Right = cur
			next := cur.Left
			cur.Left = nil
			cur = next
		}
    }
    return result
}

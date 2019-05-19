package main

import "fmt"

// TreeNode is a representation of a tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	sum := 0
	switch {
	case root.Val < L:
		sum = rangeSumBST(root.Right, L, R)
	case root.Val > R:
		sum = rangeSumBST(root.Left, L, R)
	default:
		sum = root.Val
		sum += rangeSumBST(root.Left, L, R)
		sum += rangeSumBST(root.Right, L, R)
	}
	return sum
}
func main() {
	a := TreeNode{10, nil, nil}
	a.Left = &TreeNode{5, nil, nil}
	a.Right = &TreeNode{15, nil, nil}
	a.Left.Left = &TreeNode{3, nil, nil}
	a.Left.Right = &TreeNode{7, nil, nil}
	a.Right.Right = &TreeNode{18, nil, nil}
	res := rangeSumBST(&a, 7, 15)
	fmt.Println(res)
}

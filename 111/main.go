package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	depth := tree(root, 1)
	return depth
}

func main() {
	// root := &TreeNode{
	// 	Val:  2,
	// 	Left: nil,
	// 	Right: &TreeNode{
	// 		Val:  3,
	// 		Left: nil,
	// 		Right: &TreeNode{
	// 			Val:  4,
	// 			Left: nil,
	// 			Right: &TreeNode{
	// 				Val:  5,
	// 				Left: nil,
	// 				Right: &TreeNode{
	// 					Val:   6,
	// 					Left:  nil,
	// 					Right: nil,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	depth := tree(root, 1)

	log.Printf("depth : %+v", depth)
	log.Printf("minDepth : %+v", minDepth)
}
func tree(root *TreeNode, depth int) int {

	// log.Printf("depth : %+v", depth)

	if root.Left == nil && root.Right == nil {
		return depth
	}

	depth = depth + 1
	leftDepth := 100000
	rightDepth := 100000
	if root.Left != nil {
		leftDepth = tree(root.Left, depth)
	}

	if root.Right != nil {
		rightDepth = tree(root.Right, depth)
	}

	// log.Printf("leftDepth : %+v", leftDepth)
	// log.Printf("rightDepth : %+v", rightDepth)

	if leftDepth < rightDepth {
		return leftDepth
	}
	return rightDepth
}

package main

import "log"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	treeAry := []int{0, 1, 5, 3, 2, 15, 6, 4, 9, 8, -1, -1, -1, 11, -1, -1, 7, 18, 17, 13, -1, 12, 14, 10, -1, -1, -1, 19, -1, -1, -1, -1, -1, -1, -1, -1, 16}
	root := &TreeNode{
		Val: treeAry[0],
	}
	targetVal := 16
	targetNode := makeTree(treeAry[1:], []*TreeNode{root}, targetVal)
	log.Printf("root : %+v", root)
	log.Printf("targetNode : %+v", targetNode)
	// log.Printf("Val : %+v", targetNode.Val)
	ary := distanceK(root, targetNode, 6)
	log.Printf("ary : %+v", ary)
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {

	_, _, ary := rootTree(root, target, k)

	return ary
}

func rootTree(root *TreeNode, target *TreeNode, k int) (bool, int, []int) {

	ary := []int{}
	if root == nil {
		return false, k, ary
	}
	if root.Val == target.Val {
		if k == 0 {
			ary = append(ary, root.Val)
			return false, k, ary
		}
		if target.Left != nil {
			ary = append(ary, subTree(target.Left, k-1)...)
		}
		if target.Right != nil {
			ary = append(ary, subTree(target.Right, k-1)...)
		}
		return true, k - 1, ary
	}

	leftFather, leftK, aryLeft := rootTree(root.Left, target, k)
	ary = append(ary, aryLeft...)
	if leftFather && leftK == 0 {
		ary = append(ary, root.Val)
	} else if leftFather && leftK > 0 {
		ary = append(ary, subTree(root.Right, leftK-1)...)
		return true, leftK - 1, ary
	}

	rightFather, rightK, aryRight := rootTree(root.Right, target, k)
	ary = append(ary, aryRight...)
	if rightFather && rightK == 0 {
		ary = append(ary, root.Val)
	} else if rightFather && rightK > 0 {
		ary = append(ary, subTree(root.Left, rightK-1)...)
		return true, rightK - 1, ary
	}

	return false, k, ary
}

func subTree(target *TreeNode, k int) []int {

	if target == nil {
		return []int{}
	} else if k == 0 {
		return []int{target.Val}
	}

	ary := []int{}
	ary = append(ary, subTree(target.Left, k-1)...)
	ary = append(ary, subTree(target.Right, k-1)...)

	return ary

}

func makeTree(ary []int, nodeAry []*TreeNode, targetVal int) *TreeNode {

	var targetNode *TreeNode
	if len(ary) == 0 || len(nodeAry) == 0 {
		return nil
	}
	var newNodeAry []*TreeNode
	for _, v := range nodeAry {
		if ary[0] != -1 {
			v.Left = &TreeNode{
				Val: ary[0],
			}
			if v.Left.Val == targetVal {
				targetNode = v.Left
			}
			newNodeAry = append(newNodeAry, v.Left)
		}
		if ary[1] != -1 {
			v.Right = &TreeNode{
				Val: ary[1],
			}
			if v.Right.Val == targetVal {
				targetNode = v.Right
			}
			newNodeAry = append(newNodeAry, v.Right)
		}
		if len(ary) == 2 {
			return targetNode
		} else {
			ary = ary[2:]
		}
	}
	if targetNode != nil {
		makeTree(ary, newNodeAry, targetVal)
	} else {
		targetNode = makeTree(ary, newNodeAry, targetVal)
	}
	return targetNode
}

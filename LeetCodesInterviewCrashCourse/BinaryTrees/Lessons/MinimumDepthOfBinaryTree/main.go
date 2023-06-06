package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	left1 := &TreeNode{
		Val: 15,
	}
	right2 := &TreeNode{
		Val: 7,
	}

	//left := &TreeNode{
	//	Val: 9,
	//}
	right := &TreeNode{
		Val:   20,
		Left:  left1,
		Right: right2,
	}
	root := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: right,
	}
	fmt.Println(minDepth(root))
	//minDepth(root, 1)

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	min := math.MaxInt
	if root.Left != nil {
		ans := minDepth(root.Left)
		if ans < min {
			min = ans
		}
	}
	if root.Right != nil {
		ans := minDepth(root.Right)
		if ans < min {
			min = ans
		}
	}

	return min + 1

}

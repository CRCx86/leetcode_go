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

	/* left of root*/
	left3 := &TreeNode{
		Val: 4,
	}
	right3 := &TreeNode{
		Val: 7,
	}

	left2 := &TreeNode{
		Val: 1,
	}
	right2 := &TreeNode{
		Val:   6,
		Left:  left3,
		Right: right3,
	}

	left1 := &TreeNode{
		Val:   3,
		Left:  left2,
		Right: right2,
	}

	/* right of root*/
	right5 := &TreeNode{
		Val: 13,
	}

	right4 := &TreeNode{
		Val:   14,
		Right: right5,
	}

	right1 := &TreeNode{
		Val:   10,
		Right: right4,
	}

	root := &TreeNode{
		Val:   8,
		Left:  left1,
		Right: right1,
	}

	fmt.Println(maxAncestorDiff(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return dfs(root, root.Val, root.Val)
}

func dfs(node *TreeNode, minimum int, maximum int) int {

	if node == nil {
		return maximum - minimum
	}

	maximum = int(math.Max(float64(maximum), float64(node.Val)))
	minimum = int(math.Min(float64(minimum), float64(node.Val)))
	left := dfs(node.Left, minimum, maximum)
	right := dfs(node.Right, minimum, maximum)

	return int(math.Max(float64(left), float64(right)))
}

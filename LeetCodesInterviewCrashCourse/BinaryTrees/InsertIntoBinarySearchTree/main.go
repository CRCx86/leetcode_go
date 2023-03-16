package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	/* left of root*/
	left2 := &TreeNode{
		Val: 1,
	}
	right2 := &TreeNode{
		Val: 3,
	}
	left1 := &TreeNode{
		Val:   2,
		Left:  left2,
		Right: right2,
	}

	/* right of root*/
	right3 := &TreeNode{
		Val: 7,
	}

	right1 := &TreeNode{
		Val:   5,
		Right: right3,
	}

	/* root*/
	root := &TreeNode{
		Val:   4,
		Left:  left1,
		Right: right1,
	}

	fmt.Println(insertIntoBST(root, 6))

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if val < root.Val {
		left := insertIntoBST(root.Left, val)
		root.Left = left
	}
	if val > root.Val {
		right := insertIntoBST(root.Right, val)
		root.Right = right
	}

	return root
}

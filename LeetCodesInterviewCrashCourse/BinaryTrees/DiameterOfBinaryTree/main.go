package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	/* left of root*/

	left2 := &TreeNode{
		Val: 4,
	}
	//right2 := &TreeNode{
	//	Val: 5,
	//}
	left1 := &TreeNode{
		Val:  2,
		Left: left2,
		//Right: right2,
	}

	/* right of root*/

	right1 := &TreeNode{
		Val: 3,
	}

	/* root*/
	root := &TreeNode{
		Val:   1,
		Left:  left1,
		Right: right1,
	}

	fmt.Println(diameterOfBinaryTree(root))

	left1 = &TreeNode{
		Val: 2,
	}

	root = &TreeNode{
		Val:  1,
		Left: left1,
	}

	fmt.Println(diameterOfBinaryTree(root))
}

var diameter int

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {

	diameter = 0
	dfs(root)

	return diameter

}

func dfs(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	if left+right > diameter {
		diameter = left + right
	}

	ans := left
	if right > ans {
		ans = right
	}

	return ans + 1

}

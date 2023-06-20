package main

import (
	"fmt"
)

/*
Дан корень root двоичного дерева и целочисленную target, удалите все конечные узлы с target.

Обратите внимание, что после удаления конечного узла с target, если его родительский узел
становится конечным узлом и имеет target, он также должен быть удален
(вы должны продолжать делать это, пока не сможете).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//q := generateTree1()
	//fmt.Println(removeLeafNodes(q, 2))

	//m := generateTree2()
	//fmt.Println(removeLeafNodes(m, 3))

	m := generateTree3()
	fmt.Println(removeLeafNodes(m, 1))
}

func generateTree1() *TreeNode {

	left2 := &TreeNode{
		Val: 2,
	}

	right2 := &TreeNode{
		Val: 2,
	}

	right1 := &TreeNode{
		Val:  3,
		Left: right2,
		//Right: two,
	}

	left1 := &TreeNode{
		Val:  2,
		Left: left2,
		//Right: right2,
	}

	q := &TreeNode{
		Val:   1,
		Left:  left1,
		Right: right1,
	}
	return q
}

func generateTree2() *TreeNode {

	left2 := &TreeNode{
		Val: 3,
	}

	right2 := &TreeNode{
		Val: 2,
	}

	right1 := &TreeNode{
		Val: 3,
		//Left: right2,
		//Right: two,
	}

	left1 := &TreeNode{
		Val:   3,
		Left:  left2,
		Right: right2,
	}

	q := &TreeNode{
		Val:   1,
		Left:  left1,
		Right: right1,
	}
	return q
}

func generateTree3() *TreeNode {

	right1 := &TreeNode{
		Val: 1,
	}

	left1 := &TreeNode{
		Val: 1,
	}

	q := &TreeNode{
		Val:   1,
		Left:  left1,
		Right: right1,
	}

	return q
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {

	dfs(root, target)

	if root.Left == nil && root.Right == nil && root.Val == target {
		root = nil
	}

	return root

	// leetcode solution
	//if root.Left != nil {
	//	root.Left = removeLeafNodes(root.Left, target)
	//}
	//if root.Right != nil {
	//	root.Right = removeLeafNodes(root.Right, target)
	//}
	//if root.Left == nil && root.Right == nil && root.Val == target {
	//	return nil
	//}
	//return root
}

func dfs(root *TreeNode, target int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == target {
		return true
	}

	left := dfs(root.Left, target)
	if left {
		root.Left = nil
	}
	right := dfs(root.Right, target)
	if right {
		root.Right = nil
	}

	return root.Left == nil && root.Right == nil && root.Val == target
}

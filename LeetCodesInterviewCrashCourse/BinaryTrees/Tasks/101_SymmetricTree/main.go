package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	three := &TreeNode{
		Val: 2,
	}

	four := &TreeNode{
		Val: 2,
	}

	one := &TreeNode{
		Val:  2,
		Left: three,
	}

	two := &TreeNode{
		Val:  2,
		Left: four,
	}

	q := &TreeNode{
		Val:   1,
		Left:  one,
		Right: two,
	}

	//fmt.Println(isSymmetric(q))

	fmt.Println(isSymmetric2(q))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {

		currLen := queue.Len()
		buff := make([]int, currLen*2, currLen*2)
		for i := 0; i < currLen; i++ {

			node := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())

			if node.Left != nil {
				queue.PushBack(node.Left)
				buff[i*2] = node.Left.Val
			} else {
				buff[i*2] = -101
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
				buff[i*2+1] = node.Right.Val
			} else {
				buff[i*2+1] = -101
			}

		}

		left := 0
		right := currLen*2 - 1
		for left < right {

			if buff[left] != buff[right] {
				return false
			}

			left++
			right--
		}

	}

	return true

}

func isSymmetric2(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return dfs(root.Left, root.Right)

}

func dfs(left *TreeNode, right *TreeNode) bool {

	if left == nil && right == nil {
		return left.Val == right.Val
	}

	if left == nil || right == nil {
		return false
	}

	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

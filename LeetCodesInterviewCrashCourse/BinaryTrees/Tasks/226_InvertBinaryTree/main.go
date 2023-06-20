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

	qLeft := &TreeNode{
		Val: 2,
	}

	//qRight := &TreeNode{
	//	Val: 3,
	//}

	q := &TreeNode{
		Val:  1,
		Left: qLeft,
		//Right: qRight,
	}

	fmt.Println(invertTree(q))

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		currLen := queue.Len()

		//buff := make([]*TreeNode, 0, currLen)
		for i := 0; i < currLen; i++ {
			node := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())

			//buff = append(buff, node)

			node.Left, node.Right = node.Right, node.Left

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		//left := 0
		//right := currLen - 1
		//for left < right {
		//	buff[left], buff[right] = buff[right], buff[left]
		//	left++
		//	right--
		//}

	}

	return root

}

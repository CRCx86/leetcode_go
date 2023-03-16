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
	/* left of root*/

	left3 := &TreeNode{
		Val: 7,
	}

	left2 := &TreeNode{
		Val: 4,
	}
	right2 := &TreeNode{
		Val:  5,
		Left: left3,
	}
	left1 := &TreeNode{
		Val:   2,
		Left:  left2,
		Right: right2,
	}

	/* right of root*/

	right3 := &TreeNode{
		Val: 6,
	}

	right1 := &TreeNode{
		Val:   3,
		Right: right3,
	}

	/* root*/
	root := &TreeNode{
		Val:   1,
		Left:  left1,
		Right: right1,
	}

	fmt.Println(zigzagLevelOrder(root))

	left1 = &TreeNode{
		Val: 2,
	}

	root = &TreeNode{
		Val:  1,
		Left: left1,
	}

	fmt.Println(zigzagLevelOrder(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {

	ans := make([][]int, 0)

	if root == nil {
		return ans
	}

	queue := list.New()
	queue.PushBack(root)
	zigzag := true

	for queue.Len() > 0 {

		values := make([]int, 0)

		length := queue.Len()
		buff := list.New()
		for i := 0; i < length; i++ {
			node := queue.Front()
			queue.Remove(node)

			buff.PushBack(node.Value)

			if node.Value.(*TreeNode).Left != nil {
				queue.PushBack(node.Value.(*TreeNode).Left)
			}

			if node.Value.(*TreeNode).Right != nil {
				queue.PushBack(node.Value.(*TreeNode).Right)
			}
		}

		for buff.Len() > 0 {
			if !zigzag {
				node := buff.Back()
				values = append(values, node.Value.(*TreeNode).Val)
				buff.Remove(node)
			} else {
				node := buff.Front()
				values = append(values, node.Value.(*TreeNode).Val)
				buff.Remove(node)
			}
		}

		ans = append(ans, values)
		zigzag = !zigzag

	}

	return ans

}

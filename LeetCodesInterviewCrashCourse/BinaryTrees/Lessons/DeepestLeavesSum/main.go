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

	fmt.Println(deepestLeavesSum(root))

	left1 = &TreeNode{
		Val: 2,
	}

	root = &TreeNode{
		Val:  1,
		Left: left1,
	}

	fmt.Println(deepestLeavesSum(root))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {

	if root == nil {
		return 0
	}

	ans := 0
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {

		length := queue.Len()

		lowLevel := false
		leftLevel := false
		rightLevel := false
		for i := 0; i < length; i++ {
			node := queue.Front()
			ans += node.Value.(*TreeNode).Val
			queue.Remove(node)

			if node.Value.(*TreeNode).Left != nil {
				queue.PushBack(node.Value.(*TreeNode).Left)
				leftLevel = true
			}

			if node.Value.(*TreeNode).Right != nil {
				queue.PushBack(node.Value.(*TreeNode).Right)
				rightLevel = true
			}

			lowLevel = !leftLevel && !rightLevel
		}

		if queue.Len() != 0 && !lowLevel {
			ans = 0
		}

	}

	return ans

}

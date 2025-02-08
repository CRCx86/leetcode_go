package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(levelOrder(generateTree1()))
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	queue := list.New()
	queue.PushBack(root)

	ans := make([][]int, 0)
	i := 0
	for queue.Len() > 0 {

		length := queue.Len()

		values := make([]int, 0)
		for i := 0; i < length; i++ {

			node := queue.Front()
			queue.Remove(node)

			values = append(values, node.Value.(*TreeNode).Val)

			if node.Value.(*TreeNode).Left != nil {
				queue.PushBack(node.Value.(*TreeNode).Left)
			}

			if node.Value.(*TreeNode).Right != nil {
				queue.PushBack(node.Value.(*TreeNode).Right)
			}
		}
		ans = append(ans, values)
		i++
	}

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTree1() *TreeNode {

	left2 := &TreeNode{
		Val: 15,
	}

	right2 := &TreeNode{
		Val: 7,
	}

	left1 := &TreeNode{
		Val: 9,
	}

	right1 := &TreeNode{
		Val:   20,
		Left:  left2,
		Right: right2,
	}

	root := &TreeNode{
		Val:   3,
		Left:  left1,
		Right: right1,
	}
	return root
}

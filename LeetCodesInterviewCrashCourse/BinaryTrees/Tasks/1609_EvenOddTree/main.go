package main

import "fmt"

func main() {
	fmt.Println(isEvenOddTree(generateTree1()))
}

func isEvenOddTree(root *TreeNode) bool {

	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {

		length := len(queue)
		var prev *TreeNode
		for i := 0; i < length; i++ {

			node := queue[0]
			queue = queue[1:]

			state := true
			if level%2 == 0 {
				if node.Val%2 == 0 {
					return false
				}
				state = isInc(prev, node)
			} else {
				if node.Val%2 != 0 {
					return false
				}
				state = isDecr(prev, node)
			}

			if !state {
				return false
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			prev = node

		}
		level++

	}

	return true

}

func isInc(prev *TreeNode, curr *TreeNode) bool {

	if prev == nil {
		return true
	}

	return prev.Val < curr.Val

}

func isDecr(prev *TreeNode, curr *TreeNode) bool {

	if prev == nil {
		return true
	}

	return prev.Val > curr.Val

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTree1() *TreeNode {

	left2 := &TreeNode{
		Val: 7,
	}

	right2 := &TreeNode{
		Val: 15,
	}

	left1 := &TreeNode{
		Val: 20,
	}

	right1 := &TreeNode{
		Val:   10,
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

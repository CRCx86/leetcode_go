package main

import "fmt"

func main() {
	fmt.Println(averageOfLevels(generateTree1()))
}

func averageOfLevels(root *TreeNode) []float64 {

	if root == nil {
		return []float64{}
	}

	ans := make([]float64, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {

		length := len(queue)
		sum := 0
		for i := 0; i < length; i++ {

			node := queue[0]
			queue = queue[1:]

			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		ans = append(ans, float64(sum)/float64(length))

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

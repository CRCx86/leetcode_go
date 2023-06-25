package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree1 := generateTree1()
	fmt.Println(longestZigZag(tree1)) // 2

	tree2 := generateTree2()
	fmt.Println(longestZigZag(tree2)) // 4

	tree3 := generateTree3()
	fmt.Println(longestZigZag(tree3)) // 3
}

func generateTree1() *TreeNode {

	right5 := &TreeNode{
		Val: -2,
	}

	left3 := &TreeNode{
		Val: 3,
	}

	right4 := &TreeNode{
		Val: 1,
	}

	right3 := &TreeNode{
		Val: 11,
	}

	left2 := &TreeNode{
		Val:   3,
		Left:  left3,
		Right: right5,
	}

	right2 := &TreeNode{
		Val:   2,
		Right: right4,
	}

	left1 := &TreeNode{
		Val:   5,
		Left:  left2,
		Right: right2,
	}

	right1 := &TreeNode{
		Val:   -3,
		Right: right3,
	}

	root := &TreeNode{
		Val:   10,
		Left:  left1,
		Right: right1,
	}
	return root
}

func generateTree2() *TreeNode {

	right5 := &TreeNode{
		Val: 1,
	}

	//left3 := &TreeNode{
	//	Val: 3,
	//}
	//
	//right3 := &TreeNode{
	//	Val: 11,
	//}

	right4 := &TreeNode{
		Val: 1,
	}

	left2 := &TreeNode{
		Val:   1,
		Right: right5,
	}

	right2 := &TreeNode{
		Val:   1,
		Left:  left2,
		Right: right4,
	}

	left1 := &TreeNode{
		Val:   1,
		Right: right2,
	}

	right1 := &TreeNode{
		Val: 1,
	}

	root := &TreeNode{
		Val:   1,
		Left:  left1,
		Right: right1,
	}
	return root
}

func generateTree3() *TreeNode {

	right3 := &TreeNode{
		Val: 8,
	}

	right5 := &TreeNode{
		Val:   7,
		Right: right3,
	}

	right4 := &TreeNode{
		Val: 6,
	}

	left2 := &TreeNode{
		Val:   5,
		Right: right5,
	}

	right2 := &TreeNode{
		Val:   4,
		Left:  left2,
		Right: right4,
	}

	left1 := &TreeNode{
		Val: 3,
	}

	right1 := &TreeNode{
		Val:   2,
		Left:  left1,
		Right: right2,
	}

	root := &TreeNode{
		Val:   1,
		Right: right1,
	}
	return root
}

func longestZigZag(root *TreeNode) int {
	return max(maxZigZag2(root.Left, "right", 0), maxZigZag2(root.Right, "left", 0))
}

//func maxZigZag(root *TreeNode, direction int) int {
//
//	if root == nil {
//		return 0
//	}
//
//	left := 0
//	if direction == 0 {
//		left += maxZigZag(root.Right, 1) + root.Val
//	}
//
//	right := 0
//	if direction == 1 {
//		right += maxZigZag(root.Left, 0) + root.Val
//	}
//
//	return max(left, right)
//}

func maxZigZag2(root *TreeNode, direct string, ans int) int {

	if root == nil {
		return ans
	}

	if direct == "left" {
		return max(maxZigZag2(root.Left, "right", ans+1), maxZigZag2(root.Right, "left", 0))
	} else {
		return max(maxZigZag2(root.Right, "left", ans+1), maxZigZag2(root.Left, "right", 0))
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

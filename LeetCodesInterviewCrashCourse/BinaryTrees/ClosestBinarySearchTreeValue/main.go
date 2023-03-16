package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	/* left of root*/
	left2 := &TreeNode{
		Val: 1,
	}
	right2 := &TreeNode{
		Val: 3,
	}
	left1 := &TreeNode{
		Val:   2,
		Left:  left2,
		Right: right2,
	}

	/* right of root*/
	right1 := &TreeNode{
		Val: 5,
	}

	/* root*/
	root := &TreeNode{
		Val:   4,
		Left:  left1,
		Right: right1,
	}

	fmt.Println(closestValue(root, 3.714286))
	fmt.Println(closestValue(root, 3.114286))
	fmt.Println(closestValue(root, 1.114286))
	fmt.Println(closestValue(root, 1.914286))
	fmt.Println(closestValue(root, 4.014286))
	fmt.Println(closestValue(root, 4.614286))
	fmt.Println(closestValue(root, 5.214286))

	root = &TreeNode{
		Val: 1,
	}

	fmt.Println(closestValue(root, 4.428571))

	left1 = &TreeNode{
		Val: 1,
	}

	root = &TreeNode{
		Val:  8,
		Left: left1,
	}

	fmt.Println(closestValue(root, 6.0))

	right2 = &TreeNode{
		Val: 2,
	}

	left1 = &TreeNode{
		Val:   1,
		Right: right2,
	}

	right1 = &TreeNode{
		Val: 4,
	}

	root = &TreeNode{
		Val:   3,
		Left:  left1,
		Right: right1,
	}

	fmt.Println(closestValue(root, 0.428571))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	ans := root.Val
	if root.Left != nil && target < float64(root.Val) {
		left := closestValue(root.Left, target)
		a := math.Abs(float64(root.Val) - target)
		b := math.Abs(float64(left) - target)
		if a < b {
			ans = root.Val
		} else {
			ans = left
		}
	}

	if root.Right != nil && target > float64(root.Val) {
		right := closestValue(root.Right, target)
		a := math.Abs(float64(root.Val) - target)
		b := math.Abs(float64(right) - target)
		if a < b {
			ans = root.Val
		} else {
			ans = right
		}
	}

	return ans
}

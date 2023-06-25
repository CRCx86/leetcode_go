package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	tree1 := generateTree1()
	//fmt.Println(pathSum(tree1, 8))

	fmt.Println(pathSum2(tree1, 8))

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

var ans int

func pathSum(root *TreeNode, targetSum int) int {

	ans = 0
	dfs(root, targetSum, []int{})
	return ans
}

func dfs(root *TreeNode, target int, stack []int) {

	if root == nil {
		return
	}

	sum := 0
	stack = append(stack, root.Val)
	for i := len(stack) - 1; i > 0; i-- {
		sum += stack[i]
		if sum == target {
			ans++
		}
	}

	dfs(root.Left, target, stack)
	dfs(root.Right, target, stack)

}

func pathSum2(root *TreeNode, targetSum int) int {

	if root == nil {
		return 0
	}

	return dfs2(root, targetSum) + pathSum2(root.Left, targetSum) + pathSum2(root.Right, targetSum)
}

func dfs2(root *TreeNode, target int) int {

	if root == nil {
		return 0
	}

	result := 0

	result += dfs2(root.Left, target-root.Val)
	result += dfs2(root.Right, target-root.Val)
	if target == root.Val {
		result += 1
	}

	return result

}

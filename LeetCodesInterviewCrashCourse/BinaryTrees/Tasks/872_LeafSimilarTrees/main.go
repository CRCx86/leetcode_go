package main

import (
	"fmt"
)

func main() {

	pLeft := &TreeNode{
		Val: 2,
	}

	//pRight := &TreeNode{
	//	Val: 3,
	//}

	p := &TreeNode{
		Val:  1,
		Left: pLeft,
		//Right: pRight,
	}

	qLeft := &TreeNode{
		Val: 2,
	}

	//qRight := &TreeNode{
	//	Val: 3,
	//}

	q := &TreeNode{
		Val:  2,
		Left: qLeft,
		//Right: qRight,
	}

	fmt.Println(leafSimilar(p, q))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	left := make([]int, 0)
	left = dfs(root1, left)

	right := make([]int, 0)
	right = dfs(root2, right)

	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if right[i] != left[i] {
			return false
		}
	}

	//reflect.DeepEqual(left, right)

	return true
}

func dfs(root *TreeNode, leafs []int) []int {

	if root == nil {
		return leafs
	}

	if root.Left == nil && root.Right == nil {
		leafs = append(leafs, root.Val)
		return leafs
	}

	leafs = dfs(root.Left, leafs)
	leafs = dfs(root.Right, leafs)

	return leafs
}

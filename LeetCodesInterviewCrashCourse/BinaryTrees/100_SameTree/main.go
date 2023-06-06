package main

import "fmt"

func main() {

	pLeft := &TreeNode{
		Val: 2,
	}

	pRight := &TreeNode{
		Val: 3,
	}

	p := &TreeNode{
		Val:   1,
		Left:  pLeft,
		Right: pRight,
	}

	qLeft := &TreeNode{
		Val: 2,
	}

	qRight := &TreeNode{
		Val: 3,
	}

	q := &TreeNode{
		Val:   1,
		Left:  qLeft,
		Right: qRight,
	}

	fmt.Println(isSameTree(p, q))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	return check(p, q)

}

func check(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return false
	}

	if p.Left == nil && q.Left == nil && p.Right == nil && q.Right == nil {
		return p.Val == q.Val
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

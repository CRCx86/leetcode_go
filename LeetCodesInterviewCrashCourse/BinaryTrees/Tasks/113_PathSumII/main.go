package main

import "fmt"

/**
  Дан корень двоичного дерева и целое число targetSum, вернуть все пути от корня к листу,
  где сумма значений узлов в пути равна targetSum. Каждый путь должен быть возвращен в виде списка значений узлов,
  а не ссылок на узлы.

  Путь от корня к листу — это путь, начинающийся от корня и заканчивающийся любым конечным узлом.
  Лист — это узел без потомков.

  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	q := generateTree1()
	fmt.Println(pathSum(q, 22))

	m := generateTree2()
	fmt.Println(pathSum(m, 5))

	c := generateTree3()
	fmt.Println(pathSum(c, 22))
}

func generateTree1() *TreeNode {
	two := &TreeNode{
		Val: 2,
	}

	seven := &TreeNode{
		Val: 7,
	}

	eleven := &TreeNode{
		Val:   11,
		Left:  seven,
		Right: two,
	}

	four := &TreeNode{
		Val:  4,
		Left: eleven,
	}

	q := &TreeNode{
		Val:  5,
		Left: four,
	}
	return q
}

func generateTree3() *TreeNode {

	one := &TreeNode{
		Val: 1,
	}

	five := &TreeNode{
		Val: 5,
	}

	four := &TreeNode{
		Val:   4,
		Left:  five,
		Right: one,
	}

	thirteen := &TreeNode{
		Val: 13,
	}

	eight := &TreeNode{
		Val:   8,
		Left:  thirteen,
		Right: four,
	}

	q := &TreeNode{
		Val:   5,
		Right: eight,
	}
	return q
}

func generateTree2() *TreeNode {
	two := &TreeNode{
		Val: 2,
	}

	three := &TreeNode{
		Val: 3,
	}

	q := &TreeNode{
		Val:   1,
		Left:  two,
		Right: three,
	}
	return q
}

func pathSum(root *TreeNode, targetSum int) [][]int {

	ints := make([][]int, 0)

	dfs(root, 0, targetSum, []int{}, &ints)

	return ints
}

func dfs(root *TreeNode, current int, target int, in []int, ints *[][]int) {

	if root == nil {
		return
	}

	in = append(in, root.Val)
	if current+root.Val == target && root.Left == nil && root.Right == nil {
		buff := make([]int, len(in))
		copy(buff, in)
		*(ints) = append(*(ints), buff)
	}

	dfs(root.Left, current+root.Val, target, in, ints)

	dfs(root.Right, current+root.Val, target, in, ints)

}

package BinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New() *TreeNode {
	return &TreeNode{}
}

func (t *TreeNode) Fill(nodes ...int) {

	for i := range nodes {
		slide := i*2 + 1
		if slide < len(nodes) {

		}
	}

}

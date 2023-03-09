package SingleLinkedList

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
}

func NewSingleLinkedList() *List {
	return &List{}
}

func (list *List) Insert(val int) {

	node := &ListNode{
		Val:  val,
		Next: nil,
	}

	if list.Head == nil {
		list.Head = node
	} else {
		prev := list.Head
		for prev.Next != nil {
			prev = prev.Next
		}
		prev.Next = node
	}
}

func (list *List) Fill(border int, max int) {
	for i := 0; i < border; i++ {
		list.Insert(rand.Intn(max))
	}
}

func (list *List) FillNums(nums ...int) {
	for _, e := range nums {
		list.Insert(e)
	}
}

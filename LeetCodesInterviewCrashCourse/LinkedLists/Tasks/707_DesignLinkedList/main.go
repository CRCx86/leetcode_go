package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	param_1 := obj.Get(1)
	fmt.Println(param_1)
	obj.DeleteAtIndex(1)
	param_1 = obj.Get(1)
	fmt.Println(param_1)

	obj = Constructor()
	obj.AddAtHead(1)
	obj.DeleteAtIndex(0)

	obj = Constructor()
	obj.AddAtHead(7)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtIndex(3, 0)
	obj.DeleteAtIndex(2)
	obj.AddAtHead(6)
	obj.AddAtTail(4)
	param_1 = obj.Get(4)
	fmt.Println(param_1)
	obj.AddAtHead(4)
	obj.AddAtIndex(5, 0)
	obj.AddAtHead(6)

	obj = Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	param_1 = obj.Get(1)
	fmt.Println(param_1)
	obj.DeleteAtIndex(0)
	param_1 = obj.Get(0)
	fmt.Println(param_1)

	fmt.Println("=======")
	obj = Constructor()
	obj.AddAtIndex(0, 10)
	obj.AddAtIndex(0, 20)
	obj.AddAtIndex(1, 30)
	param_1 = obj.Get(0)
	fmt.Println(param_1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	Head *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (t *MyLinkedList) Get(index int) int {

	curr := t.Head
	for curr != nil && index > 0 {
		curr = curr.Next
		index--
	}

	if index == 0 && curr != nil {
		return curr.Val
	}

	return -1
}

func (t *MyLinkedList) AddAtHead(val int) {

	node := &ListNode{
		Val:  val,
		Next: nil,
	}

	if t.Head == nil {
		t.Head = node
	} else {
		node.Next = t.Head
		t.Head = node
	}
}

func (t *MyLinkedList) AddAtTail(val int) {

	node := &ListNode{
		Val:  val,
		Next: nil,
	}

	if t.Head == nil {
		t.Head = node
	} else {
		curr := t.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
	}

}

func (t *MyLinkedList) AddAtIndex(index int, val int) {

	var prev *ListNode
	curr := t.Head
	for index > 0 && curr != nil {
		prev = curr
		curr = curr.Next
		index--
	}

	if index == 0 {
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		node.Next = curr
		if prev != nil {
			prev.Next = node
		} else {
			t.Head = node
		}
	}

}

func (t *MyLinkedList) DeleteAtIndex(index int) {

	if t.Head == nil {
		return
	}

	var prev *ListNode
	curr := t.Head
	for index > 0 && curr != nil {
		prev = curr
		curr = curr.Next
		index--
	}

	if index == 0 && curr != nil {
		next := curr.Next
		if prev == nil {
			t.Head = next
		} else {
			prev.Next = next
		}
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

package main

import (
	"leetcode_go/utils/SingleLinkedList"
)

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	//singleLinkedList.FillNums(1, 2, 3, 4, 5)
	//oddEvenList(singleLinkedList.Head)
	//singleLinkedList.FillNums(2, 1, 3, 5, 6, 4, 7)
	//oddEvenList(singleLinkedList.Head)
	singleLinkedList.FillNums(1, 2, 3, 4, 5, 6, 7, 8)
	oddEvenList(singleLinkedList.Head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *SingleLinkedList.ListNode) *SingleLinkedList.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	o := head
	e := head.Next
	for iter, i := head.Next, 1; iter != nil; iter = iter.Next {
		if i%2 == 0 {
			e.Next = iter.Next
			e = e.Next
		} else {
			o.Next = iter.Next
			if o.Next != nil {
				o = o.Next
			}
		}
		i++
	}

	o.Next = even

	return odd
}

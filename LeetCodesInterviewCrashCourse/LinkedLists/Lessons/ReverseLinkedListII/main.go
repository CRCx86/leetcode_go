package main

import (
	"leetcode_go/utils/SingleLinkedList"
)

func main() {

	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 2, 3, 4, 5)
	reverseBetween(singleLinkedList.Head, 2, 4)

	singleLinkedList = SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 2, 3)
	reverseBetween(singleLinkedList.Head, 2, 3)

	singleLinkedList = SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 2, 3, 4, 5)
	reverseBetween(singleLinkedList.Head, 1, 4)

	singleLinkedList = SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(3, 5)
	reverseBetween(singleLinkedList.Head, 1, 2)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *SingleLinkedList.ListNode, left int, right int) *SingleLinkedList.ListNode {

	if left >= right {
		return head
	}

	start := head
	i := 1
	var prev *SingleLinkedList.ListNode
	var begin *SingleLinkedList.ListNode
	var tail *SingleLinkedList.ListNode
	for head != nil {

		next := head.Next

		if i == left {
			begin = prev
			tail = head
		}

		if begin != nil && i <= right {
			begin.Next = head
		}
		if i >= left && i <= right {
			head.Next = prev
		}

		prev = head
		if i == right {
			tail.Next = next
			break
		}

		head = next
		i++
	}
	if left == 1 {
		start = prev
	}

	return start

}

package main

import "leetcode_go/utils/SingleLinkedList"

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 1, 2)
	deleteDuplicates(singleLinkedList.Head)

	singleLinkedList = SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 1, 2, 3, 3)
	deleteDuplicates(singleLinkedList.Head)

	singleLinkedList = SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 1, 1)
	deleteDuplicates(singleLinkedList.Head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *SingleLinkedList.ListNode) *SingleLinkedList.ListNode {

	next := head
	for next != nil && next.Next != nil {
		if next.Val == next.Next.Val {
			next.Next = next.Next.Next
		} else {
			next = next.Next
		}
	}
	return head
}

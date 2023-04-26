package main

import "leetcode_go/utils/SingleLinkedList"

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	//singleLinkedList.FillNums(1, 3, 4, 7, 1, 2, 6)
	singleLinkedList.FillNums(1)
	deleteMiddle(singleLinkedList.Head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *SingleLinkedList.ListNode) *SingleLinkedList.ListNode {

	if head.Next == nil {
		return nil
	}

	fast := head
	slow := head
	preSlow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		preSlow = slow
		slow = slow.Next
	}

	preSlow.Next = slow.Next

	return head
}

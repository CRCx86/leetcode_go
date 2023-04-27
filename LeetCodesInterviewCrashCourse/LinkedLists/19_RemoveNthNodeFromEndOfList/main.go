package main

import "leetcode_go/utils/SingleLinkedList"

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	//singleLinkedList.FillNums(1, 2, 3, 4, 5)
	//removeNthFromEnd(singleLinkedList.Head, 2)
	//singleLinkedList.FillNums(1)
	//removeNthFromEnd(singleLinkedList.Head, 1)
	singleLinkedList.FillNums(1, 2)
	removeNthFromEnd(singleLinkedList.Head, 2)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *SingleLinkedList.ListNode, n int) *SingleLinkedList.ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	fast := head
	slow := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	var pre *SingleLinkedList.ListNode
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}

	if pre == nil {
		head = slow.Next
	} else {
		pre.Next = slow.Next
	}

	return head
}

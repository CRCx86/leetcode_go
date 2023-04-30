package main

import "leetcode_go/utils/SingleLinkedList"

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 2, 6, 3, 4, 5, 6)
	removeElements(singleLinkedList.Head, 6)
	//singleLinkedList.FillNums(7, 7, 7, 7)
	//removeElements(singleLinkedList.Head, 7)
	//singleLinkedList.FillNums(1, 2, 2, 1)
	//removeElements(singleLinkedList.Head, 2)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *SingleLinkedList.ListNode, n int) *SingleLinkedList.ListNode {

	if head == nil {
		return head
	}

	var prev *SingleLinkedList.ListNode
	curr := head
	for curr != nil {
		if curr.Val == n {
			if prev == nil {
				curr = curr.Next
				head = head.Next
			} else {
				prev.Next = curr.Next
				curr = curr.Next
			}
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return head
}

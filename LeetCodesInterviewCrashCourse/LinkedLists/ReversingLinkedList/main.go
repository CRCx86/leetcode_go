package main

import "leetcode_go/utils/SingleLinkedList"

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 2, 3, 4, 5)
	reverseList(singleLinkedList.Head)
}

func reverseList(head *SingleLinkedList.ListNode) *SingleLinkedList.ListNode {

	curr := head
	var prev *SingleLinkedList.ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

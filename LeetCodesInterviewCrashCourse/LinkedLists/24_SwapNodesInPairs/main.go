package main

import "leetcode_go/utils/SingleLinkedList"

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 2, 3, 4, 5)
	swapPairs(singleLinkedList.Head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *SingleLinkedList.ListNode) *SingleLinkedList.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	dummy := head.Next
	var prev *SingleLinkedList.ListNode
	for head != nil && head.Next != nil {
		if prev != nil {
			prev.Next = head.Next
		}

		prev = head

		nextNode := head.Next.Next
		head.Next.Next = head

		head.Next = nextNode
		head = nextNode
	}

	return dummy
}

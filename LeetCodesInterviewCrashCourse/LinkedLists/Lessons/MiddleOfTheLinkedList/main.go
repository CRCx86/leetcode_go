package main

import (
	"fmt"
	"leetcode_go/utils/SingleLinkedList"
)

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.Fill(10, 100)
	result := middleNode(singleLinkedList.Head)
	fmt.Println(result)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *SingleLinkedList.ListNode) *SingleLinkedList.ListNode {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

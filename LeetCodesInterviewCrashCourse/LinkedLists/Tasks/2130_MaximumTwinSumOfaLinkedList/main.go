package main

import (
	"fmt"
	"leetcode_go/utils/SingleLinkedList"
)

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(5, 4, 2, 1)
	fmt.Println(pairSum(singleLinkedList.Head))

	singleLinkedList = SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(4, 2, 2, 3)
	fmt.Println(pairSum(singleLinkedList.Head))

	singleLinkedList = SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 100000)
	fmt.Println(pairSum(singleLinkedList.Head))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *SingleLinkedList.ListNode) int {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	curr := slow
	var prev *SingleLinkedList.ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	max := 0
	start := head
	for prev != nil {

		n := start.Val + prev.Val
		if n > max {
			max = n
		}
		start = start.Next
		prev = prev.Next
	}

	return max
}

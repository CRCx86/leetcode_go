package main

import (
	"fmt"
	"leetcode_go/utils/SingleLinkedList"
)

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	//singleLinkedList.FillNums(1, 2, 3, 4, 5)
	//fmt.Println(isPalindrome(singleLinkedList.Head))
	singleLinkedList.FillNums(1, 2, 2, 1)
	fmt.Println(isPalindrome(singleLinkedList.Head))
	//singleLinkedList.FillNums(1, 0, 1)
	//fmt.Println(isPalindrome(singleLinkedList.Head))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *SingleLinkedList.ListNode) bool {

	if head.Next == nil {
		return true
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	curr := slow
	var prev *SingleLinkedList.ListNode
	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}

	for prev != nil {
		if head.Val != prev.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}

	return true
}

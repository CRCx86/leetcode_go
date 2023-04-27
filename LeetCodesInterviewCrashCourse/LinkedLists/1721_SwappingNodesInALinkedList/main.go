package main

import "leetcode_go/utils/SingleLinkedList"

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	//singleLinkedList.FillNums(1, 2, 3, 4, 5)
	//swapNodes1(singleLinkedList.Head, 2)
	//singleLinkedList.FillNums(7, 9, 6, 6, 7, 8, 3, 0, 9, 5)
	//swapNodes(singleLinkedList.Head, 5)
	singleLinkedList.FillNums(100, 90)
	swapNodes2(singleLinkedList.Head, 2)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *SingleLinkedList.ListNode, k int) *SingleLinkedList.ListNode {

	left := head
	right := head
	fast := head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		right = right.Next
		fast = fast.Next
	}

	for i := 1; i < k; i++ {
		left = left.Next
	}

	left.Val, right.Val = right.Val, left.Val

	return head

}

func swapNodes1(head *SingleLinkedList.ListNode, k int) *SingleLinkedList.ListNode {

	left, right := head, head

	for iter := head; iter != nil; iter = iter.Next {
		k--
		if k == 0 {
			left = iter
		}
		if k < 0 {
			right = right.Next
		}
	}

	left.Val, right.Val = right.Val, left.Val

	return head

}

func swapNodes2(head *SingleLinkedList.ListNode, k int) *SingleLinkedList.ListNode {

	left, fast := head, head

	for k-1 > 0 {
		left = left.Next
		fast = fast.Next
		k--
	}

	right := head
	for fast = fast.Next; fast != nil; fast = fast.Next {
		right = right.Next
	}

	left.Val, right.Val = right.Val, left.Val

	return head

}

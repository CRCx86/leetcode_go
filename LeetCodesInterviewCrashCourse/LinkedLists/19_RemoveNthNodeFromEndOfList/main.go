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
	pre := head

	m := 0
	k := 0

	for slow != nil {

		if fast == nil || fast.Next == nil {
			if m-k == n {
				if slow.Next == nil {
					pre.Next = nil
				} else {
					slow.Next = slow.Next.Next
				}
				break
			} else {
				pre = slow
				slow = slow.Next
				k += 1
			}
		} else {
			fast = fast.Next.Next
			m += 2
		}
	}

	return head
}

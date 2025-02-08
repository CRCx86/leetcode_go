package main

import "leetcode_go/utils/SingleLinkedList"

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	//singleLinkedList.FillNums(1, 2, 3, 3, 4, 4, 5)
	//deleteDuplicates(singleLinkedList.Head)
	//singleLinkedList.FillNums(1, 1, 1, 2, 3)
	//deleteDuplicates(singleLinkedList.Head)
	singleLinkedList.FillNums(1, 1)
	deleteDuplicates(singleLinkedList.Head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *SingleLinkedList.ListNode) *SingleLinkedList.ListNode {

	fast := head.Next
	slow := head
	var pre *SingleLinkedList.ListNode
	flag := false

	for fast != nil {
		if fast.Val == slow.Val {
			fast = fast.Next
			flag = true
			continue
		}
		if flag {
			if pre == nil {
				head = fast
				slow = fast
				fast = fast.Next
			} else {
				pre.Next = fast
				slow = pre
			}
			flag = false
		} else {
			pre = slow
			slow = slow.Next
			fast = fast.Next
		}
	}

	if flag {
		if pre == nil {
			head = fast
			slow = fast
		} else {
			pre.Next = fast
			slow = pre
		}
		flag = false
	}

	return head
}

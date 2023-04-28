package main

import (
	"leetcode_go/utils/SingleLinkedList"
)

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	//singleLinkedList.FillNums(5, 2, 6, 3, 9, 1, 7, 3, 8, 4)
	//reverseEvenLengthGroups(singleLinkedList.Head)
	singleLinkedList.FillNums(1, 1, 0, 6)
	reverseEvenLengthGroups(singleLinkedList.Head)
	//singleLinkedList.FillNums(1, 1, 0, 6, 5)
	//reverseEvenLengthGroups(singleLinkedList.Head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseEvenLengthGroups(head *SingleLinkedList.ListNode) *SingleLinkedList.ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	i, j, k := 1, 1, 1
	k = i
	l := 0
	var A *SingleLinkedList.ListNode
	B, C, D := head, head, head.Next
	for D != nil {

		if j == 1 {
			j++
			i++
			k = i + j
			A = head
			B = D
			continue
		}

		if i != k {
			C = C.Next
			D = D.Next
			l++
			i++
		} else {

			if l%2 == 0 {

				var prev *SingleLinkedList.ListNode
				curr := B
				for curr != D {
					nextNode := curr.Next
					curr.Next = prev
					prev = curr
					curr = nextNode
				}
				A.Next = prev
				B.Next = D
				//A = C
				A = B
				B = D
				C = D
				D = D.Next

			} else {
				A = C
				B = D
				C = D
				D = D.Next
			}
			l = 1
			j++
			k = i + j
			i++
		}

	}

	if l%2 == 0 {
		var prev *SingleLinkedList.ListNode
		curr := B
		for curr != D {
			nextNode := curr.Next
			curr.Next = prev
			prev = curr
			curr = nextNode
		}
		A.Next = prev
		B.Next = D
	}

	return head
}

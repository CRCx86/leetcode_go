package main

import (
	"fmt"
	"leetcode_go/utils/SingleLinkedList"
)

func main() {
	singleLinkedList := SingleLinkedList.NewSingleLinkedList()
	singleLinkedList.FillNums(1, 0, 1)
	fmt.Println(getDecimalValue(singleLinkedList.Head))
	//singleLinkedList.FillNums(0)
	//fmt.Println(getDecimalValue(singleLinkedList.Head))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *SingleLinkedList.ListNode) int {

	//n := -1
	//
	//fast := head
	//slow := head
	//
	//for fast != nil {
	//	n++
	//	fast = fast.Next
	//}
	//
	//ans := 0
	//for slow != nil {
	//	ans += slow.Val << n
	//	n--
	//	slow = slow.Next
	//}
	//
	//return ans

	var res int
	for head != nil {
		i := res << 1 // двигаясь к младшему биту имеем ввиду, что предыдущее число - это на разряд больше
		i2 := i | head.Val
		res = i2
		fmt.Printf("i: %d, i2: %d", i, i2)
		head = head.Next
	}
	return res

}

package main

import "fmt"

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	param := obj.SumRange(0, 2)
	fmt.Println(param)
	param = obj.SumRange(2, 5)
	fmt.Println(param)
	param = obj.SumRange(0, 5)
	fmt.Println(param)
}

type NumArray struct {
	obj []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		obj: nums,
	}
}

func (n *NumArray) SumRange(left int, right int) int {

	prefixes := make([]int, 0)
	prefixes = append(prefixes, n.obj[0])
	for i, e := range n.obj[1:] {
		prefixes = append(prefixes, prefixes[i]+e)
	}

	if left == 0 {
		return prefixes[right]
	}

	return prefixes[right] - prefixes[left] + n.obj[left]

}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}

func pivotIndex(nums []int) int {

	sum := 0
	for _, e := range nums {
		sum += e
	}

	left := 0
	for i, e := range nums {
		if left == (sum - nums[i] - left) {
			return i
		}
		left += e
	}

	return -1
}

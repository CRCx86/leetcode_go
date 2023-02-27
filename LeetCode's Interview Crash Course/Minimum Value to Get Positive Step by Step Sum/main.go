package main

import (
	"fmt"
)

func main() {

	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))
	fmt.Println(minStartValue([]int{1, 2}))
	fmt.Println(minStartValue([]int{1, -2, -3}))

}

func minStartValue(nums []int) int {

	total := 0
	min := 0

	for i := range nums {
		total += nums[i]
		if total < min {
			min = total
		}
	}

	return -min + 1

}

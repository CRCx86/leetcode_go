package main

import (
	"fmt"
	"math"
)

// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}

func sortedSquares(nums []int) []int {

	l := len(nums) - 1

	left := 0
	right := l

	result := make([]int, l+1)

	for i := l; i >= 0; i-- {
		buff := 0
		if math.Abs(float64(nums[left])) > math.Abs(float64(nums[right])) {
			buff = nums[left]
			left++
		} else {
			buff = nums[right]
			right--
		}
		result[i] = buff * buff
	}

	return result
}

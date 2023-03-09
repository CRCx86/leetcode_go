package main

import (
	"fmt"
	"math"
)

//	Given an array of positive integers nums and a positive integer target, return the minimal length of a
//	subarray
//	whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

func main() {

	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(7, []int{1, 2, 4}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
	fmt.Println(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
	fmt.Println(minSubArrayLen(15, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}))

}

func minSubArrayLen(target int, nums []int) int {

	min := 0
	max := len(nums)
	left := 0
	curr := 0
	for right, e := range nums {
		curr += e
		if curr >= target {
			for curr >= target {
				curr -= nums[left]
				left++
			}
			min = int(math.Min(float64(max), float64(right-left+2)))
			max = min
		}
	}

	return min
}

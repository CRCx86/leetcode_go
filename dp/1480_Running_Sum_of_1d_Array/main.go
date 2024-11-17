package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}

func runningSum(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return nums
}

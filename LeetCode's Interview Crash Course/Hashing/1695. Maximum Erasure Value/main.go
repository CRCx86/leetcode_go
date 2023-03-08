package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}) == 17)
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}) == 8)
	fmt.Println(maximumUniqueSubarray([]int{10000, 1, 10000, 1, 1, 1, 1, 1, 1}) == 10001)
	fmt.Println(maximumUniqueSubarray([]int{187, 470, 25, 436, 538, 809, 441, 167, 477, 110, 275, 133, 666, 345, 411, 459, 490, 266, 987, 965, 429, 166, 809, 340, 467, 318, 125, 165, 809, 610, 31, 585, 970, 306, 42, 189, 169, 743, 78, 810, 70, 382, 367, 490, 787, 670, 476, 278, 775, 673, 299, 19, 893, 817, 971, 458, 409, 886, 434}) == 16911)
}

func maximumUniqueSubarray(nums []int) int {

	max := 0
	sum := 0
	counts := make(map[int]int)
	left := 0
	for right := 0; right < len(nums); right++ {

		counts[nums[right]]++

		for counts[nums[right]] > 1 {
			sum -= nums[left]
			counts[nums[left]]--
			left++
		}

		sum += nums[right]
		if sum > max {
			max = sum
		}

	}

	return max
}

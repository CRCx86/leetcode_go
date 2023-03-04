package main

import "fmt"

func main() {
	fmt.Println(largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1}))
	fmt.Println(largestUniqueNumber([]int{9, 9, 8, 8}))
}

func largestUniqueNumber(nums []int) int {

	max := -1
	counts := make(map[int]int)
	for k := range nums {
		counts[nums[k]]++
	}

	for k := range counts {
		if v := counts[k]; v == 1 {
			if k > max {
				max = k
			}
		}
	}

	return max
}

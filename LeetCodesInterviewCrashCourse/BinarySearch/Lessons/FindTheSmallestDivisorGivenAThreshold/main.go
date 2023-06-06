package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
	fmt.Println(smallestDivisor([]int{44, 22, 33, 11, 1}, 5))
}

func smallestDivisor(nums []int, threshold int) int {

	left := 1
	right := 0
	for _, n := range nums {
		if n > right {
			right = n
		}
	}

	for left <= right {
		mid := left + (right-left)/2
		if check(mid, nums, threshold) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func check(mid int, nums []int, threshold int) bool {
	sum := 0
	for _, n := range nums {
		sum += int(math.Ceil(float64(n) / float64(mid)))
	}

	return sum <= threshold
}

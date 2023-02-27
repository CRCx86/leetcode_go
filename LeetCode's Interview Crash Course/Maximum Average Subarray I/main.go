package main

import (
	"fmt"
	"math"
)

// You are given an integer array nums consisting of n elements, and an integer k.

// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
// Any answer with a calculation error less than 10-5 will be accepted.

func main() {

	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))

}

func findMaxAverage(nums []int, k int) float64 {

	curr := 0.0
	for i := 0; i < k; i++ {
		curr += float64(nums[i])
	}
	curr /= float64(k)

	ans := curr
	for i := k; i < len(nums); i++ {
		curr += float64(nums[i]-nums[i-k]) / float64(k)
		ans = math.Max(ans, curr)
	}

	return ans
}

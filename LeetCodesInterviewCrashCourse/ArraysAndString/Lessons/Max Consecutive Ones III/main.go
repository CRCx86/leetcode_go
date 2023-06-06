package main

import (
	"fmt"
)

// Given a binary array nums and an integer k,
// return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

func main() {

	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))

}

func longestOnes(nums []int, k int) int {

	left := 0
	curr := 0
	ans := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			curr++
		}

		for curr > k {
			if nums[left] == 0 {
				curr--
			}
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}

func max(a int, b int) int {

	if a > b {
		return a
	}

	return b
}

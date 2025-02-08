package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Println(longestSubarray([]int{1, 1, 1}))
	fmt.Println(longestSubarray([]int{1, 1, 0, 0, 1, 1, 1, 0, 1}))
}

func longestSubarray(nums []int) int {

	left := 0
	curr := 0
	ans := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			curr++
		}

		for curr > 1 {
			if nums[left] == 0 {
				curr--
			}
			left++
		}

		ans = max(ans, right-left+1)

	}

	return ans - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

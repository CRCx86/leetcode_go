package main

import "fmt"

// Разбор задачи

func main() {
	fmt.Println(lengthOfLIS([]int{1, 2, 5, 3, 4}))
	//fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	//fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	//fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

var array []int
var memo []int

func lengthOfLIS(nums []int) int {

	return bottomUp(nums)

	//array = nums
	//memo = make([]int, len(nums))
	//for i := range memo {
	//	memo[i] = -1
	//}
	//
	//ans := 0
	//for i := 0; i < len(array); i++ {
	//	ref := topDown(i)
	//	if ref > ans {
	//		ans = ref
	//	}
	//}
	//
	//return ans
}

func topDown(i int) int {

	if memo[i] != -1 {
		return memo[i]
	}

	ans := 1 // base case
	for j := 0; j < i; j++ {
		if array[i] > array[j] {
			ref := topDown(j) + 1
			if ref > ans {
				ans = ref
			}
		}
	}

	memo[i] = ans
	return memo[i]
}

func bottomUp(nums []int) int {

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	ans := 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				ref := dp[j] + 1
				if ref > dp[i] {
					dp[i] = ref
				}
				if dp[i] > ans {
					ans = dp[i]
				}
			}
		}
	}

	return ans
}

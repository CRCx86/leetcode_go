package main

import "fmt"

// разбор задачи
func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

var array []int
var memo []int

func rob(nums []int) int {

	return bottomUp(nums)

	//array = nums
	//memo = make([]int, len(array))
	//for i := range memo {
	//	memo[i] = -1
	//}
	//
	//return topDown(len(array) - 1)
}

func topDown(i int) int {

	if i == 0 {
		return array[0]
	}

	if i == 1 {
		ref := array[0]
		if array[1] > ref {
			ref = array[1]
		}
		return ref
	}

	if memo[i] != -1 {
		return memo[i]
	}

	left := topDown(i - 1)
	right := topDown(i-2) + array[i]

	if right > left {
		left = right
	}

	memo[i] = left
	return memo[i]
}

func bottomUp(nums []int) int {

	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = nums[0]
	if nums[1] > dp[1] {
		dp[1] = nums[1]
	}

	for i := 2; i < n; i++ {
		left := dp[i-1]
		right := dp[i-2] + nums[i]
		if right > left {
			left = right
		}
		dp[i] = left
	}

	return dp[n-1]

}

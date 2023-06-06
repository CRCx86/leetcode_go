package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
	fmt.Println(answerQueries([]int{2, 3, 4, 5}, []int{1}))
}

func answerQueries(nums []int, queries []int) []int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := make([]int, 0)

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	for _, query := range queries {
		i := bs(nums, query)
		ans = append(ans, i)
	}

	return ans

}

func bs(arr []int, target int) int {

	left := 0
	right := len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

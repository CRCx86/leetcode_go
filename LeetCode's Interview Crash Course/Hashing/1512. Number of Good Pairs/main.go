package main

import (
	"fmt"
)

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}))
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
}

func numIdenticalPairs(nums []int) int {

	ans := 0
	counts := make(map[int][]int)
	for i := range nums {
		counts[nums[i]] = append(counts[nums[i]], i)
	}

	for _, v := range counts {
		right := len(v) - 1
		if right > 0 {
			left := 0
			for left < right {
				ans += right - left
				left++
			}
		}
	}

	//counts := make(map[int]int)
	//for i := range nums {
	//	if counts[nums[i]] > 0 {
	//		ans += counts[nums[i]]
	//	}
	//	counts[nums[i]]++
	//}

	return ans
}

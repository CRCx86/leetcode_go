package main

import "fmt"

func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(sumOfUnique([]int{1, 1, 1, 1, 1}))
	fmt.Println(sumOfUnique([]int{1, 2, 3, 4, 5}))
}

func sumOfUnique(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	counts := make(map[int]int)
	for n := range nums {
		counts[nums[n]]++
	}

	ans := 0
	for k, v := range counts {
		if v == 1 {
			ans += k
		}
	}

	return ans
}

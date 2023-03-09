package main

import "fmt"

func main() {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}

func numSubarraysWithSum(nums []int, goal int) int {

	curr, ans := 0, 0
	counts := make(map[int]int)
	counts[0] = 1
	for i := range nums {
		curr += nums[i]
		ans += counts[curr-goal]
		counts[curr]++
	}

	return ans
}

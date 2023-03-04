package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, -1, 1, -1}, 3))
}

func subarraySum(nums []int, k int) int {

	counts := make(map[int]int)
	counts[0] = 1

	ans, curr := 0, 0

	for _, i := range nums {
		curr += i
		ans += counts[curr-k]
		counts[curr]++
	}

	return ans

}

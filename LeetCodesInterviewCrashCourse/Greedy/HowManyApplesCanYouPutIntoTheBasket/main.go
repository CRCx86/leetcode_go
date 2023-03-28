package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxNumberOfApples([]int{100, 200, 150, 1000}))
	fmt.Println(maxNumberOfApples([]int{900, 950, 800, 1000, 700, 800}))
}

func maxNumberOfApples(weight []int) int {

	sort.Slice(weight, func(i, j int) bool {
		return weight[i] < weight[j]
	})

	ans := 0
	pivot := 5000
	for i := 0; i < len(weight); i++ {
		pivot -= weight[i]
		if pivot < 0 {
			return ans
		}
		ans++
	}

	return ans
}

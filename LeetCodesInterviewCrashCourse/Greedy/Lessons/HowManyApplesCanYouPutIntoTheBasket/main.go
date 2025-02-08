package main

import (
	"sort"
)

func MaxNumberOfApples(weight []int) int {
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

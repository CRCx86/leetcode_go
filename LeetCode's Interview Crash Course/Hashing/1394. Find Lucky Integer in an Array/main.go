package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLucky([]int{2, 2, 3, 4}))
	fmt.Println(findLucky([]int{1, 2, 2, 3, 3, 3}))
	fmt.Println(findLucky([]int{2, 2, 2, 3, 3}))
}

func findLucky(arr []int) int {

	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {
		if arr[0] == 1 {
			return arr[0]
		} else {
			return -1
		}
	}

	counts := make(map[int]int)
	for i := range arr {
		counts[arr[i]]++
	}

	keys := make([]int, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[j] < keys[i]
	})

	for _, v := range keys {
		if v == counts[v] {
			return v
		}
	}

	return -1
}

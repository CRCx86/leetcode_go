package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
	fmt.Println(uniqueOccurrences([]int{3, 5, -2, -3, -6, -6}))
}

func uniqueOccurrences(arr []int) bool {

	numbers := make(map[int]int)
	for i := range arr {
		numbers[arr[i]]++
	}

	counts := make(map[int]int)
	for _, v := range numbers {
		counts[v]++
		if counts[v] > 1 {
			return false
		}
	}

	return true
}

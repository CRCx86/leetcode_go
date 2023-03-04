package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
	fmt.Println(findWinners([][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}))
}

func findWinners(matches [][]int) [][]int {

	winners := make(map[int]int)
	losers := make(map[int]int)

	for i := range matches {
		winners[matches[i][0]]++
		losers[matches[i][1]]++
	}

	answer0 := make([]int, 0)
	answer1 := make([]int, 0)
	for k := range winners {
		if value := losers[k]; value == 0 {
			answer0 = append(answer0, k)
		}
	}

	for k, v := range losers {
		if v == 1 {
			answer1 = append(answer1, k)
		}
	}

	sort.Ints(answer0)
	sort.Ints(answer1)

	return [][]int{answer0, answer1}
}

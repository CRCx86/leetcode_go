package main

import "fmt"

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
	fmt.Println(maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}

func maximumWealth(accounts [][]int) int {

	max := 0
	//for i := 0; i < len(accounts); i++ {
	//	buff := 0
	//	for j := 0; j < len(accounts[i]); j++ {
	//		buff += accounts[i][j]
	//	}
	//	if buff > max {
	//		max = buff
	//	}
	//}

	for _, i := range accounts {
		buff := 0
		for _, j := range i {
			buff += j
		}
		if buff > max {
			max = buff
		}
	}

	return max
}

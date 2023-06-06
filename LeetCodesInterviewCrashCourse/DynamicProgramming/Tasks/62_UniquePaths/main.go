package main

import "fmt"

// Разбор задачи

func main() {
	//fmt.Println(uniquePaths(3, 5))
	fmt.Println(uniquePaths(1, 1))
	fmt.Println(uniquePaths(2, 3))
}

var memo [][]int

func uniquePaths(m int, n int) int {

	return bu(m, n)

	//memo = make([][]int, m)
	//for i := range memo {
	//	memo[i] = make([]int, n)
	//	for j := range memo[i] {
	//		memo[i][j] = -1
	//	}
	//}
	//
	//return td(m-1, n-1)
}

func td(row int, col int) int {

	if row+col == 0 {
		return 1
	}

	if memo[row][col] != -1 {
		return memo[row][col]
	}

	ways := 0
	if row > 0 {
		ways += td(row-1, col)
	}

	if col > 0 {
		ways += td(row, col-1)
	}

	memo[row][col] = ways
	return ways
}

func bu(m int, n int) int {

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row > 0 {
				dp[row][col] += dp[row-1][col]
			}
			if col > 0 {
				dp[row][col] += dp[row][col-1]
			}
		}
	}

	return dp[m-1][n-1]
}

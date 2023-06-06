package main

import (
	"fmt"
	"math"
)

// Разбор задачи

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

var matrix [][]int
var memo [][]int

func minPathSum(grid [][]int) int {

	return bu(grid)

	//m := len(grid)
	//n := len(grid[0])
	//matrix = grid
	//
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

func td(row, col int) int {

	if row+col == 0 {
		return matrix[row][col]
	}

	if memo[row][col] != -1 {
		return memo[row][col]
	}

	ans := math.MaxInt
	if row > 0 {
		ref := td(row-1, col)
		if ref < ans {
			ans = ref
		}
	}

	if col > 0 {
		ref := td(row, col-1)
		if ref < ans {
			ans = ref
		}
	}

	memo[row][col] = matrix[row][col] + ans
	return memo[row][col]
}

func bu(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row+col == 0 {
				continue
			}

			ans := math.MaxInt
			if row > 0 {
				ref := dp[row-1][col]
				if ref < ans {
					ans = ref
				}
			}
			if col > 0 {
				ref := dp[row][col-1]
				if ref < ans {
					ans = ref
				}
			}

			dp[row][col] = grid[row][col] + ans
		}
	}

	return dp[m-1][n-1]
}

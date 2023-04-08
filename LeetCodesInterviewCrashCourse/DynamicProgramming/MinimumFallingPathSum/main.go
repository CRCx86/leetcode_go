package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
	fmt.Println(minFallingPathSum([][]int{{-19, 57}, {-40, -5}}))
}

func minFallingPathSum(matrix [][]int) int {

	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = matrix[0][i]
	}

	for row := 1; row < m; row++ {
		for col := 0; col < n; col++ {

			ans := math.MaxInt
			if col > 0 {
				left := dp[row-1][col-1] + matrix[row][col]
				if left < ans {
					ans = left
				}
			}
			center := dp[row-1][col] + matrix[row][col]
			if center < ans {
				ans = center
			}
			if col < n-1 {
				right := dp[row-1][col+1] + matrix[row][col]
				if right < ans {
					ans = right
				}
			}

			dp[row][col] = ans
		}
	}

	min := math.MaxInt
	for i := 0; i < n; i++ {
		if dp[m-1][i] < min {
			min = dp[m-1][i]
		}
	}

	return min
}

package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{1}}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {

			if obstacleGrid[row][col] == 1 {
				dp[row][col] = 0
				continue
			}

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

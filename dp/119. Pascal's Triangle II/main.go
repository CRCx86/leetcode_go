package main

import "fmt"

func main() {

	fmt.Println(getRow(3))

}

func getRow(rowIndex int) []int {

	dp := make([][]int, rowIndex+1)

	for i := 1; i <= rowIndex+1; i++ {
		dp[i-1] = make([]int, i)
		dp[i-1][0] = 1
		dp[i-1][i-1] = 1
		for j := 2; j < i; j++ {
			dp[i-1][j-1] = dp[i-2][j-1] + dp[i-2][j-2]
		}
	}

	return dp[rowIndex]

}

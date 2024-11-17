package main

import "fmt"

func main() {
	fmt.Println(generate(1))
	fmt.Println(generate(2))
	fmt.Println(generate(3))
	fmt.Println(generate(4))
	fmt.Println(generate(5))
	fmt.Println(generate(30))
}

func generate(numRows int) [][]int {

	dp := make([][]int, numRows)

	for i := 1; i <= numRows; i++ {
		dp[i-1] = make([]int, i)
		dp[i-1][0] = 1
		dp[i-1][i-1] = 1
		for j := 2; j < i; j++ {
			dp[i-1][j-1] = dp[i-2][j-1] + dp[i-2][j-2]
		}
	}

	return dp
}

package main

import "fmt"

func main() {
	fmt.Println(mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
	fmt.Println(mostPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}))
}

var array [][]int

func mostPoints(questions [][]int) int64 {

	return bottomUp(questions)

	//array = questions
	//memo := make([]int64, len(questions))
	//
	//return topDown(0, memo)
}

func topDown(i int, memo []int64) int64 {

	if i >= len(array) {
		return 0
	}

	if memo[i] != 0 {
		return memo[i]
	}

	j := i + array[i][1] + 1
	left := int64(array[i][0]) + topDown(j, memo)
	right := topDown(i+1, memo)
	if right > left {
		left = right
	}

	memo[i] = left
	return memo[i]
}

func bottomUp(questions [][]int) int64 {

	n := len(questions)
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		j := i + questions[i][1] + 1

		if j > n {
			j = n
		}

		dp[i] = questions[i][0] + dp[j]
		if dp[i+1] > dp[i] {
			dp[i] = dp[i+1]
		}
	}

	return int64(dp[0])
}

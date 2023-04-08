package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbingStairs(cost []int) int {

	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i <= n; i++ {
		c := 0
		if i < n {
			c = cost[i]
		}
		dp[i] = c + min(dp, i)
	}

	return dp[n]

}

func min(dp []int, i int) int {
	min := dp[i-1]
	if dp[i-2] < dp[i-1] {
		min = dp[i-2]
	}
	return min
}

package main

import "fmt"

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {

	if n < 2 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	//for stair := 2; stair < n; stair++ {
	//	for step := 1; step <= 2; step++ {
	//		dp[stair] += dp[stair-step]
	//	}
	//}

	for stair := 2; stair < n; stair++ {
		dp[stair] = dp[stair-1] + dp[stair-2]
	}

	return dp[n-1]
}

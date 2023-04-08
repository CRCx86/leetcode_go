package main

import "fmt"

func main() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(23))
}

func climbStairs(n int) int {

	if n <= 1 {
		return 1
	}

	dp := make([]int, n)

	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]

}
